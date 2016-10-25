package docker

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/omg-product-bundle/products/docker/enaml-gen/containers"
	"github.com/enaml-ops/omg-product-bundle/products/docker/enaml-gen/docker"
	"github.com/enaml-ops/omg-product-bundle/products/docker/enaml-gen/docker-volume-netshare"
	"github.com/enaml-ops/pluginlib/pcli"
	"github.com/enaml-ops/pluginlib/pluginutil"
	"github.com/enaml-ops/pluginlib/product"
	"github.com/xchapter7x/lo"
)

const (
	BoshDockerReleaseURL = "https://github.com/calebwashburn/docker-boshrelease/releases/tag/v28.0.1-dev-14"
	BoshDockerReleaseVer = "28.0.1-dev-14"
	BoshDockerReleaseSHA = "debaf48c7e7b8fbb4ac385f5c41fc26dcbdd8018"
	defaultReleaseName   = "docker"
	defaultStemcellName  = "trusty"
)

type jobBucket struct {
	JobName   string
	JobType   int
	Instances int
}
type Plugin struct {
	PluginVersion      string
	DeploymentName     string
	Containers         interface{}
	NetworkName        string
	IPs                []string
	VMTypeName         string
	DiskTypeName       string
	AZs                []string
	StemcellName       string
	StemcellURL        string
	StemcellVersion    string
	StemcellSHA        string
	RegistryMirrors    []string
	InsecureRegistries []string
	EnableNFS          bool
	NFSVersion         int
	NFSOptions         string
}

func (s *Plugin) GetFlags() (flags []pcli.Flag) {
	return []pcli.Flag{
		pcli.Flag{FlagType: pcli.StringFlag, Name: "deployment-name", Value: "docker", Usage: "the name bosh will use for this deployment"},
		pcli.Flag{FlagType: pcli.BoolFlag, Name: "infer-from-cloud", Usage: "setting this flag will attempt to pull as many defaults from your targetted bosh's cloud config as it can (vmtype, network, disk, etc)."},
		pcli.Flag{FlagType: pcli.StringSliceFlag, Name: "ip", Usage: "multiple static ips for each redis leader vm"},
		pcli.Flag{FlagType: pcli.StringSliceFlag, Name: "az", Usage: "list of AZ names to use"},
		pcli.Flag{FlagType: pcli.StringSliceFlag, Name: "insecure-registry", Usage: "Array of insecure registries (no certificate verification for HTTPS and enable HTTP fallback)"},
		pcli.Flag{FlagType: pcli.StringSliceFlag, Name: "registry-mirror", Usage: "Array of preferred Docker registry mirrors"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: "network", Usage: "the name of the network to use"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: "vm-type", Usage: "name of your desired vm type"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: "disk-type", Usage: "name of your desired disk type"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: "stemcell-url", Usage: "the url of the stemcell you wish to use"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: "stemcell-ver", Usage: "the version number of the stemcell you wish to use"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: "stemcell-sha", Usage: "the sha of the stemcell you will use"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: "stemcell-name", Value: s.GetMeta().Stemcell.Name, Usage: "the name of the stemcell you will use"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: "container-definition", Usage: "filepath to the container definition for your docker containers"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: "docker-release-url", Value: BoshDockerReleaseURL, Usage: "the url of the docker release you wish to use"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: "docker-release-ver", Value: BoshDockerReleaseVer, Usage: "the version number of the docker release you wish to use"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: "docker-release-sha", Value: BoshDockerReleaseSHA, Usage: "the sha of the docker release you will use"},

		pcli.Flag{FlagType: pcli.BoolFlag, Name: "use-nfs-volume", Usage: "enable nfs volume driver"},
		pcli.Flag{FlagType: pcli.IntFlag, Name: "nfs-version", Value: "3", Usage: "nfs version"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: "nfs-options", Value: "port=2049,nolock", Usage: "options for nfs mount"},
	}
}

func (s *Plugin) GetMeta() product.Meta {
	return product.Meta{
		Name: "docker",
		Stemcell: enaml.Stemcell{
			Name: defaultStemcellName,
		},
		Releases: []enaml.Release{
			enaml.Release{
				Name:    defaultReleaseName,
				Version: BoshDockerReleaseVer,
				URL:     BoshDockerReleaseURL,
				SHA1:    BoshDockerReleaseSHA,
			},
		},
		Properties: map[string]interface{}{
			"version":        s.PluginVersion,
			"docker-release": strings.Join([]string{BoshDockerReleaseURL, BoshDockerReleaseVer, BoshDockerReleaseSHA}, " / "),
		},
	}
}

func (s *Plugin) setContainerDefinitionFromFile(filename string) interface{} {
	var res []interface{}
	if b, e := ioutil.ReadFile(filename); e == nil {
		yaml.Unmarshal(b, &res)

	} else {
		lo.G.Fatalf("you have not given a valid path to a container definition file: %v", filename)
	}
	return res
}

func (s *Plugin) GetProduct(args []string, cloudConfig []byte) (b []byte, err error) {
	c := pluginutil.NewContext(args, pluginutil.ToCliFlagArray(s.GetFlags()))
	flgs := s.GetFlags()
	InferFromCloudDecorate(flagsToInferFromCloudConfig, cloudConfig, args, flgs)
	s.Containers = s.setContainerDefinitionFromFile(c.String("container-definition"))
	s.IPs = c.StringSlice("ip")
	s.AZs = c.StringSlice("az")
	s.InsecureRegistries = c.StringSlice("insecure-registry")
	s.RegistryMirrors = c.StringSlice("registry-mirror")
	s.DeploymentName = c.String("deployment-name")
	s.NetworkName = c.String("network")
	s.StemcellName = c.String("stemcell-name")
	s.StemcellVersion = c.String("stemcell-ver")
	s.StemcellSHA = c.String("stemcell-sha")
	s.StemcellURL = c.String("stemcell-url")
	s.VMTypeName = c.String("vm-type")
	s.DiskTypeName = c.String("disk-type")
	s.EnableNFS = c.Bool("use-nfs-volume")
	s.NFSVersion = c.Int("nfs-version")
	s.NFSOptions = c.String("nfs-options")

	if err = s.flagValidation(); err != nil {
		lo.G.Error("invalid arguments: ", err)
		return nil, err
	}

	if err = s.cloudconfigValidation(enaml.NewCloudConfigManifest(cloudConfig)); err != nil {
		lo.G.Error("invalid settings for cloud config on target bosh: ", err)
		return nil, err
	}
	lo.G.Debug("context", c)
	var dm = new(enaml.DeploymentManifest)
	dm.SetName(s.DeploymentName)
	dm.AddRemoteRelease(defaultReleaseName, c.String("docker-release-ver"), c.String("docker-release-url"), c.String("docker-release-sha"))
	dm.AddRemoteStemcell(s.StemcellName, s.StemcellName, s.StemcellVersion, s.StemcellURL, s.StemcellSHA)

	dm.AddInstanceGroup(s.NewDockerInstanceGroup())
	dm.Update = enaml.Update{
		MaxInFlight:     1,
		UpdateWatchTime: "30000-300000",
		CanaryWatchTime: "30000-300000",
		Serial:          false,
		Canaries:        1,
	}
	return dm.Bytes(), err
}

func (s *Plugin) NewDockerInstanceGroup() (ig *enaml.InstanceGroup) {
	ig = &enaml.InstanceGroup{
		Name:               s.DeploymentName,
		Instances:          len(s.IPs),
		VMType:             s.VMTypeName,
		AZs:                s.AZs,
		Stemcell:           s.StemcellName,
		PersistentDiskType: s.DiskTypeName,
		Jobs: []enaml.InstanceJob{
			s.createDockerJob(),
			s.createContainersJob(),
		},
		Networks: []enaml.Network{
			enaml.Network{Name: s.NetworkName, StaticIPs: s.IPs},
		},
		Update: enaml.Update{
			MaxInFlight: 1,
		},
	}

	if s.EnableNFS {
		ig.Jobs = append(ig.Jobs, s.createNFSVolumeShareJob())
	}
	return
}
func (s *Plugin) createNFSVolumeShareJob() enaml.InstanceJob {
	return enaml.InstanceJob{
		Name:    "docker-volume-netshare",
		Release: "docker",
		Properties: &docker_volume_netshare.DockerVolumeNetshareJob{
			Nfs: &docker_volume_netshare.Nfs{
				Version: s.NFSVersion,
				Options: s.NFSOptions,
			},
		},
	}
}
func (s *Plugin) createDockerJob() enaml.InstanceJob {
	return enaml.InstanceJob{
		Name:    "docker",
		Release: "docker",
		Properties: &docker.DockerJob{
			Docker: &docker.Docker{
				RegistryMirrors:    s.RegistryMirrors,
				InsecureRegistries: s.InsecureRegistries,
			},
		},
	}
}

func (s *Plugin) createContainersJob() enaml.InstanceJob {
	return enaml.InstanceJob{
		Name:    "containers",
		Release: "docker",
		Properties: &containers.ContainersJob{
			Containers: s.Containers,
		},
	}
}

func (s *Plugin) cloudconfigValidation(cloudConfig *enaml.CloudConfigManifest) (err error) {
	lo.G.Debug("running cloud config validation")

	for _, vmtype := range cloudConfig.VMTypes {
		err = fmt.Errorf("vm size %s does not exist in cloud config. options are: %v", s.VMTypeName, cloudConfig.VMTypes)
		if vmtype.Name == s.VMTypeName {
			err = nil
			break
		}
	}

	for _, disktype := range cloudConfig.DiskTypes {
		err = fmt.Errorf("disk size %s does not exist in cloud config. options are: %v", s.DiskTypeName, cloudConfig.DiskTypes)
		if disktype.Name == s.DiskTypeName {
			err = nil
			break
		}
	}

	for _, net := range cloudConfig.Networks {
		err = fmt.Errorf("network %s does not exist in cloud config. options are: %v", s.NetworkName, cloudConfig.Networks)
		if net.(map[interface{}]interface{})["name"] == s.NetworkName {
			err = nil
			break
		}
	}

	if len(cloudConfig.VMTypes) == 0 {
		err = fmt.Errorf("no vm sizes found in cloud config")
	}

	if len(cloudConfig.DiskTypes) == 0 {
		err = fmt.Errorf("no disk sizes found in cloud config")
	}

	if len(cloudConfig.Networks) == 0 {
		err = fmt.Errorf("no networks found in cloud config")
	}
	return
}

func (s *Plugin) flagValidation() (err error) {
	lo.G.Debug("validating given flags")

	if len(s.IPs) <= 0 {
		err = fmt.Errorf("no `ip` given")
	}

	if s.Containers == nil {
		err = fmt.Errorf("no valid container definition in given file")
	}

	if len(s.AZs) <= 0 {
		err = fmt.Errorf("no `az` given")
	}

	if s.NetworkName == "" {
		err = fmt.Errorf("no `network-name` given")
	}

	if s.VMTypeName == "" {
		err = fmt.Errorf("no `vm-type` given")
	}
	if s.DiskTypeName == "" {
		err = fmt.Errorf("no `disk-type` given")
	}

	if s.StemcellVersion == "" {
		err = fmt.Errorf("no `stemcell-ver` given")
	}
	return
}

func InferFromCloudDecorate(inferFlagMap map[string][]string, cloudConfig []byte, args []string, flgs []pcli.Flag) {
	c := pluginutil.NewContext(args, pluginutil.ToCliFlagArray(flgs))

	if c.Bool("infer-from-cloud") {
		ccinf := pluginutil.NewCloudConfigInferFromBytes(cloudConfig)
		setAllInferredFlagDefaults(inferFlagMap["disktype"], ccinf.InferDefaultDiskType(), flgs)
		setAllInferredFlagDefaults(inferFlagMap["vmtype"], ccinf.InferDefaultVMType(), flgs)
		setAllInferredFlagDefaults(inferFlagMap["az"], ccinf.InferDefaultAZ(), flgs)
		setAllInferredFlagDefaults(inferFlagMap["network"], ccinf.InferDefaultNetwork(), flgs)
	}
}

func setAllInferredFlagDefaults(matchlist []string, defaultvalue string, flgs []pcli.Flag) {

	for _, match := range matchlist {
		setFlagDefault(match, defaultvalue, flgs)
	}
}

func setFlagDefault(flagname, defaultvalue string, flgs []pcli.Flag) {
	for idx, flg := range flgs {

		if flg.Name == flagname {
			flgs[idx].Value = defaultvalue
		}
	}
}

var flagsToInferFromCloudConfig = map[string][]string{
	"disktype": []string{
		"disk-type",
	},
	"vmtype": []string{
		"vm-type",
	},
	"az":      []string{"az"},
	"network": []string{"network"},
}
