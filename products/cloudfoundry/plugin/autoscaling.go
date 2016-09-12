package cloudfoundry

import (
	"fmt"

	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/omg-cli/utils"
	das "github.com/enaml-ops/omg-product-bundle/products/cloudfoundry/enaml-gen/deploy-autoscaling"
	db "github.com/enaml-ops/omg-product-bundle/products/cloudfoundry/enaml-gen/destroy-broker"
	rb "github.com/enaml-ops/omg-product-bundle/products/cloudfoundry/enaml-gen/register-broker"
	ta "github.com/enaml-ops/omg-product-bundle/products/cloudfoundry/enaml-gen/test-autoscaling"

	"github.com/enaml-ops/omg-product-bundle/products/cloudfoundry/plugin/config"
)

type (
	deployAutoscaling       struct{ *config.Config }
	registerAutoscaleBroker struct{ *config.Config }
	destroyAutoscaleBroker  struct{ *config.Config }
	autoscalingTests        struct{ *config.Config }
)

func NewDeployAutoscaling(c *config.Config) InstanceGroupCreator {
	return deployAutoscaling{c}
}

func (a deployAutoscaling) ToInstanceGroup() *enaml.InstanceGroup {
	return &enaml.InstanceGroup{
		Name:      "autoscaling",
		Instances: 1,
		VMType:    a.ErrandVMType,
		Lifecycle: "errand",
		AZs:       a.AZs,
		Stemcell:  a.StemcellName,
		Networks: []enaml.Network{
			{Name: a.NetworkName},
		},
		Update: enaml.Update{
			MaxInFlight: 1,
		},
		Jobs: []enaml.InstanceJob{
			{
				Name:    "deploy-autoscaling",
				Release: CFAutoscalingReleaseName,
				Properties: &das.DeployAutoscalingJob{
					AppDomains: a.AppDomains,
					Autoscale: &das.Autoscale{
						Broker: &das.Broker{
							User:     a.AutoscaleBrokerUser,
							Password: a.AutoscaleBrokerPassword,
						},
						Cf: &das.Cf{
							AdminUser:     "admin",
							AdminPassword: a.AdminPassword,
						},
						InstanceCount: 1,
						Database: &das.Database{
							Url: fmt.Sprintf("mysql://%s:%s@%s:3306/autoscale", a.AutoscaleDBUser, a.AutoscaleDBPassword, a.MySQLProxyHost()),
						},
						EncryptionKey:     utils.NewPassword(16),
						EnableDiego:       true,
						NotificationsHost: fmt.Sprintf("https://notifications.%s", a.SystemDomain),
						Organization:      "system",
						Space:             "autoscaling",
						MarketplaceCompanyName:      "Pivotal",
						MarketplaceImageUrl:         "data:image/png;base64,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",
						MarketplaceDocumentationUrl: "http://docs.gopivotal.com/pivotalcf/",
					},
					Domain: a.SystemDomain,
					Ssl: &das.Ssl{
						SkipCertVerify: a.SkipSSLCertVerify,
					},
					Uaa: &das.Uaa{
						Clients: &das.Clients{
							AutoscalingService: &das.AutoscalingService{
								Secret: a.AutoScalingServiceClientSecret,
							},
						},
					},
				},
			},
		},
	}
}

func NewAutoscaleRegisterBroker(c *config.Config) InstanceGroupCreator {
	return registerAutoscaleBroker{c}
}

func (a registerAutoscaleBroker) ToInstanceGroup() *enaml.InstanceGroup {
	return &enaml.InstanceGroup{
		Name:      "autoscaling-register-broker",
		Instances: 1,
		VMType:    a.ErrandVMType,
		Lifecycle: "errand",
		AZs:       a.AZs,
		Stemcell:  a.StemcellName,
		Networks: []enaml.Network{
			{Name: a.NetworkName},
		},
		Update: enaml.Update{
			MaxInFlight: 1,
		},
		Jobs: []enaml.InstanceJob{
			{
				Name:    "register-broker",
				Release: CFAutoscalingReleaseName,
				Properties: &rb.RegisterBrokerJob{
					AppDomains: a.AppDomains,
					Autoscale: &rb.Autoscale{
						Broker: &rb.Broker{
							User:     a.AutoscaleBrokerUser,
							Password: a.AutoscaleBrokerPassword,
						},
						Cf: &rb.Cf{
							AdminUser:     "admin",
							AdminPassword: a.AdminPassword,
						},
						Organization: "system",
						Space:        "autoscaling",
					},
					Domain: a.SystemDomain,
					Ssl: &rb.Ssl{
						SkipCertVerify: a.SkipSSLCertVerify,
					},
				},
			},
		},
	}
}

func NewAutoscaleDestroyBroker(c *config.Config) InstanceGroupCreator {
	return destroyAutoscaleBroker{c}
}

func (d destroyAutoscaleBroker) ToInstanceGroup() *enaml.InstanceGroup {
	return &enaml.InstanceGroup{
		Name:      "autoscaling-destroy-broker",
		Instances: 1,
		VMType:    d.ErrandVMType,
		Lifecycle: "errand",
		AZs:       d.AZs,
		Stemcell:  d.StemcellName,
		Networks: []enaml.Network{
			{Name: d.NetworkName},
		},
		Update: enaml.Update{
			MaxInFlight: 1,
		},
		Jobs: []enaml.InstanceJob{
			{
				Name:    "destroy-broker",
				Release: CFAutoscalingReleaseName,
				Properties: &db.DestroyBrokerJob{
					Autoscale: &db.Autoscale{
						Broker: &db.Broker{
							User:     d.AutoscaleBrokerUser,
							Password: d.AutoscaleBrokerPassword,
						},
						Cf: &db.Cf{
							AdminUser:     "admin",
							AdminPassword: d.AdminPassword,
						},
						Organization: "system",
						Space:        "autoscaling",
					},
					Domain: d.SystemDomain,
					Ssl: &db.Ssl{
						SkipCertVerify: d.SkipSSLCertVerify,
					},
				},
			},
		},
	}
}

func NewAutoscalingTests(c *config.Config) InstanceGroupCreator {
	return autoscalingTests{c}
}

func (a autoscalingTests) ToInstanceGroup() *enaml.InstanceGroup {
	return &enaml.InstanceGroup{
		Name:      "autoscaling-tests",
		Instances: 1,
		VMType:    a.ErrandVMType,
		Lifecycle: "errand",
		AZs:       a.AZs,
		Stemcell:  a.StemcellName,
		Networks: []enaml.Network{
			{Name: a.NetworkName},
		},
		Update: enaml.Update{
			MaxInFlight: 1,
		},
		Jobs: []enaml.InstanceJob{
			{
				Name:    "test-autoscaling",
				Release: CFAutoscalingReleaseName,
				Properties: &ta.TestAutoscalingJob{
					Autoscale: &ta.Autoscale{
						Cf: &ta.Cf{
							AdminUser:     "admin",
							AdminPassword: a.AdminPassword,
						},
					},
					Domain: a.SystemDomain,
					Ssl: &ta.Ssl{
						SkipCertVerify: a.SkipSSLCertVerify,
					},
				},
			},
		},
	}
}
