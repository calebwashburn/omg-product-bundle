package plugin

import (
	"strings"

	"gopkg.in/urfave/cli.v2"
	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/omg-cli/utils"
	"github.com/enaml-ops/omg-product-bundle/products/dockerregistry"
	"github.com/enaml-ops/pluginlib/pcli"
	"github.com/enaml-ops/pluginlib/product"
	"github.com/enaml-ops/pluginlib/util"
	"github.com/xchapter7x/lo"
)

const (
	defaultRegistryReleaseURL string = "https://bosh.io/d/github.com/cloudfoundry-community/docker-registry-boshrelease?v=3"
	defaultRegistryReleaseSHA string = "834f8ca9fd8f5280d94007b724a3b710739619db"
	defaultRegistryReleaseVer string = "3"

	registryReleaseURL string = "docker-registry-release-url"
	registryReleaseSHA string = "docker-registry-release-sha"
	registryReleaseVer string = "docker-registry-release-version"
	stemcellAlias      string = "stemcell-alias"
	stemcellOS         string = "stemcell-os"
	stemcellVersion    string = "stemcell-version"
	az                 string = "az"
	deploymentName     string = "deployment-name"
	networkName        string = "network-name"
	registryIP         string = "registry-ip"
	registryVMType     string = "registry-vm-type"
	proxyVMType        string = "proxy-vm-type"
	proxyIP            string = "proxy-ip"
	nfsVMType          string = "nfs-server-vm-type"
	nfsDiskType        string = "nfs-server-disk-type"
	nfsIP              string = "nfs-server-ip"
	publicHostName     string = "public-host-name"
)

type Plugin struct {
	PluginVersion string
}

func (p *Plugin) GetFlags() (flags []pcli.Flag) {
	flags = []pcli.Flag{
		pcli.Flag{FlagType: pcli.StringFlag, Name: deploymentName, Value: "docker-registry", Usage: "deployment name"},
		pcli.Flag{FlagType: pcli.StringSliceFlag, Name: az, Usage: "list of AZ names to use"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: networkName, Usage: "the name of the network to use"},
		pcli.Flag{FlagType: pcli.StringSliceFlag, Name: registryIP, Usage: "ip for registry job"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: registryVMType, Usage: "vm type for registry"},
		pcli.Flag{FlagType: pcli.StringSliceFlag, Name: proxyIP, Usage: "ip for proxy job"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: proxyVMType, Usage: "vm type for proxy job"},

		pcli.Flag{FlagType: pcli.StringFlag, Name: nfsIP, Usage: "ip for nfs server job"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: nfsVMType, Usage: "vm type for nfs server job"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: nfsDiskType, Usage: "disk type for nfs server job"},
		pcli.Flag{FlagType: pcli.StringSliceFlag, Name: publicHostName, Usage: "host name/ip for proxy"},

		pcli.Flag{FlagType: pcli.StringFlag, Name: registryReleaseURL, Value: defaultRegistryReleaseURL, Usage: "release url for docker-registry bosh release"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: registryReleaseSHA, Value: defaultRegistryReleaseSHA, Usage: "release sha for docker-registry bosh release"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: registryReleaseVer, Value: defaultRegistryReleaseVer, Usage: "release version for docker-registry bosh release"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: stemcellAlias, Value: "trusty", Usage: "alias of stemcell"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: stemcellOS, Value: "ubuntu-trusty", Usage: "os of stemcell"},
		pcli.Flag{FlagType: pcli.StringFlag, Name: stemcellVersion, Value: "latest", Usage: "version of stemcell"},
	}
	return
}

func (p *Plugin) GetMeta() product.Meta {
	return product.Meta{
		Name: "docker-registry",
		Properties: map[string]interface{}{
			"version":                 p.PluginVersion,
			"docker-registry-release": strings.Join([]string{defaultRegistryReleaseURL, defaultRegistryReleaseVer, defaultRegistryReleaseSHA}, " / "),
		},
	}
}

func (p *Plugin) GetProduct(args []string, cloudConfig []byte) (b []byte) {
	if len(cloudConfig) == 0 {
		lo.G.Debug("plugin: empty cloud config")
		panic("cloud config cannot be empty")
	}
	c := pluginutil.NewContext(args, pluginutil.ToCliFlagArray(p.GetFlags()))
	dm := NewDeploymentManifest(c, cloudConfig)
	return dm.Bytes()
}

func NewDeploymentManifest(c *cli.Context, cloudConfig []byte) *enaml.DeploymentManifest {
	deployment := dockerregistry.DockerRegistry{
		DeploymentName:           c.String(deploymentName),
		DockerRegistryReleaseVer: c.String(registryReleaseVer),
		DockerRegistryReleaseURL: c.String(registryReleaseURL),
		DockerRegistryReleaseSHA: c.String(registryReleaseSHA),
		StemcellVersion:          c.String(stemcellVersion),
		StemcellAlias:            c.String(stemcellAlias),
		StemcellOS:               c.String(stemcellOS),
		AZs:                      c.StringSlice(az),
		NetworkName:              c.String(networkName),
		RegistryVMType:           c.String(registryVMType),
		RegistryIPs:              c.StringSlice(registryIP),
		ProxyVMType:              c.String(proxyVMType),
		ProxyIPs:                 c.StringSlice(proxyIP),
		PublicIP:                 c.StringSlice(publicHostName),
		NFSServerVMType:          c.String(nfsVMType),
		NFSDiskType:              c.String(nfsDiskType),
		NFSIP:                    c.String(nfsIP),
		Secret:                   utils.NewPassword(20),
	}

	certIPs := append(deployment.ProxyIPs, deployment.PublicIP...)
	if _, cert, key, err := utils.GenerateCert(certIPs); err != nil {
		lo.G.Error(err.Error())
		panic(err.Error())
	} else {
		deployment.ProxyCert = cert
		deployment.ProxyCertKey = key
	}

	if manifest, err := deployment.CreateDeploymentManifest(cloudConfig); err != nil {
		lo.G.Error(err.Error())
		panic(err)
	} else {
		return manifest
	}

}
