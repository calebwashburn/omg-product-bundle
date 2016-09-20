package pscs

import (
	"fmt"

	cli "gopkg.in/urfave/cli.v2"

	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/pluginlib/pcli"
	"github.com/enaml-ops/pluginlib/product"
	"github.com/enaml-ops/pluginlib/util"
	"github.com/xchapter7x/lo"
)

// Plugin is an omg product plugin for depoying p-spring-cloud-services.
type Plugin struct {
	Version string
}

// generatePassword is the default for password flags that should be generated by
// the plugin if not specified by the user
const generatePassword = "[autogenerated]"

func (p *Plugin) GetFlags() []pcli.Flag {
	return []pcli.Flag{
		pcli.CreateStringFlag("deployment-name", "the name bosh will use for the deployment", "p-scs"),
		pcli.CreateStringFlag("system-domain", "the system domain"),
		pcli.CreateStringSliceFlag("app-domain", "Applications Domains"),
		pcli.CreateStringSliceFlag("az", "the AZs to use"),
		pcli.CreateStringFlag("network", "the name of the network to use"),
		pcli.CreateStringFlag("stemcell-ver", "the version number of the stemcell you wish to use", StemcellVersion),
		pcli.CreateStringFlag("vm-type", "VM type to use for SCS instance groups"),
		pcli.CreateBoolFlag("skip-ssl-verify", "skip SSL verification"),
		pcli.CreateStringFlag("broker-username", "the service broker username", generatePassword),
		pcli.CreateStringFlag("broker-password", "the service broker password", generatePassword),
		pcli.CreateStringFlag("worker-client-secret", "client secret for worker", generatePassword),
		pcli.CreateStringFlag("worker-password", "worker password", generatePassword),
		pcli.CreateStringFlag("instances-password", "instances password", generatePassword),
		pcli.CreateStringFlag("broker-dashboard-secret", "broker dashboard secret", generatePassword),
		pcli.CreateStringFlag("encryption-key", "encryption key", generatePassword),
		pcli.CreateStringFlag("admin-password", "CF admin password"),
		pcli.CreateStringFlag("uaa-admin-secret", "UAA client secret for admin account"),
		pcli.CreateBoolFlag("infer-from-cloud", "attempt to pull defaults from your targetted bosh"),

		pcli.CreateStringFlag("vault-domain", "the location of your vault server (ie. http://10.0.0.1:8200)"),
		pcli.CreateStringFlag("vault-token", "the token to make connections to your vault"),
		pcli.CreateStringSliceFlag("vault-hash", "a list of vault hashes to pull values from"),
	}
}

func (p *Plugin) GetMeta() product.Meta {
	return product.Meta{
		Name: "p-spring-cloud-services",
		Properties: map[string]interface{}{
			"version":                      p.Version,
			"stemcell":                     StemcellVersion,
			"pivotal-spring-cloud-servies": fmt.Sprintf("%s / %s", "pivotal-spring-cloud-servies", ProductVersion),
			"spring-cloud-broker-release":  fmt.Sprintf("%s / %s", SpringCloudBrokerReleaseName, SpringCloudBrokerReleaseVersion),
		},
	}
}

func (p *Plugin) GetProduct(args []string, cloudConfig []byte) []byte {
	flags := p.GetFlags()
	c := pluginutil.NewContext(args, pluginutil.ToCliFlagArray(flags))

	if c.Bool("infer-from-cloud") {
		inferFromCloud(cloudConfig, flags, c)
		c = pluginutil.NewContext(args, pluginutil.ToCliFlagArray(flags))
	}

	domain := c.String("vault-domain")
	tok := c.String("vault-token")
	hashes := c.StringSlice("vault-hash")
	if domain != "" && tok != "" && len(hashes) > 0 {
		lo.G.Debug("connecting to vault at", domain)
		v := pluginutil.NewVaultUnmarshal(domain, tok)
		for _, hash := range hashes {
			err := v.UnmarshalFlags(hash, flags)
			if err != nil {
				lo.G.Errorf("error reading vault hash %s: %s", hash, err.Error())
			}
		}
		c = pluginutil.NewContext(args, pluginutil.ToCliFlagArray(flags))
	}

	cfg, err := configFromContext(c)
	if err != nil {
		lo.G.Error(err)
	}

	dm := new(enaml.DeploymentManifest)
	dm.SetName(cfg.DeploymentName)
	dm.AddRelease(enaml.Release{Name: SpringCloudBrokerReleaseName, Version: SpringCloudBrokerReleaseVersion})
	dm.AddStemcell(enaml.Stemcell{OS: StemcellName, Version: cfg.StemcellVersion, Alias: StemcellAlias})
	dm.AddInstanceGroup(NewDeployServiceBroker(cfg))
	dm.AddInstanceGroup(NewRegisterBroker(cfg))
	dm.AddInstanceGroup(NewDestroyServiceBroker(cfg))

	dm.Update = enaml.Update{
		Canaries:        1,
		CanaryWatchTime: "30000-300000",
		UpdateWatchTime: "30000-300000",
		MaxInFlight:     1,
		Serial:          true,
	}

	return dm.Bytes()
}

func inferFromCloud(cloudConfig []byte, flags []pcli.Flag, c *cli.Context) {
	inferer := pluginutil.NewCloudConfigInferFromBytes(cloudConfig)

	vm := inferer.InferDefaultVMType()
	network := inferer.InferDefaultNetwork()
	az := inferer.InferDefaultAZ()

	for i := range flags {
		name := flags[i].Name
		if !c.IsSet(name) {
			switch name {
			case "network":
				lo.G.Debugf("got network '%s' from cloud config", network)
				flags[i].Value = network
			case "az":
				lo.G.Debugf("got azs %v from cloud config", az)
				flags[i].Value = az
			case "vm-type":
				lo.G.Debugf("got vm type %s from cloud config", vm)
				flags[i].Value = vm
			}
		}
	}
}
