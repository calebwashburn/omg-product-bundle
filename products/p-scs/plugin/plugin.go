package pscs

import (
	"fmt"

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
		pcli.CreateStringFlag("deployment-name", "the name bosh will use for the deployment", "p-rabbitmq"),
		pcli.CreateStringFlag("system-domain", "the system domain"),
		pcli.CreateStringFlag("network", "the name of the network to use"),
		pcli.CreateStringFlag("stemcell-ver", "the version number of the stemcell you wish to use", StemcellVersion),
		pcli.CreateBoolFlag("skip-ssl-verify", "skip SSL verification"),
		pcli.CreateStringFlag("broker-username", "the service broker username", generatePassword),
		pcli.CreateStringFlag("broker-password", "the service broker password", generatePassword),
		pcli.CreateStringFlag("worker-client-secret", "client secret for worker", generatePassword),
		pcli.CreateStringFlag("worker-password", "worker password", generatePassword),
		pcli.CreateStringFlag("instances-password", "instances password", generatePassword),
		pcli.CreateStringFlag("broker-dashboard-secret", "broker dashboard secret", generatePassword),
		pcli.CreateStringFlag("encryption-key", "encryption key", generatePassword),
		pcli.CreateStringFlag("cf-admin-password", "CF admin password"),
		pcli.CreateStringFlag("uaa-admin-client-secret", "UAA client secret for admin account"),
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

	cfg, err := configFromContext(c)
	if err != nil {
		lo.G.Error(err)
	}

	dm := new(enaml.DeploymentManifest)
	dm.SetName(cfg.DeploymentName)
	dm.AddRelease(enaml.Release{Name: SpringCloudBrokerReleaseName, Version: SpringCloudBrokerReleaseVersion})
	dm.AddStemcell(enaml.Stemcell{OS: StemcellName, Version: cfg.StemcellVersion, Alias: StemcellAlias})
	dm.AddInstanceGroup(NewDeployServiceBroker(cfg))
	return dm.Bytes()
}
