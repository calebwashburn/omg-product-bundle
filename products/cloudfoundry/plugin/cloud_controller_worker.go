package cloudfoundry

import (
	"github.com/codegangsta/cli"
	"github.com/enaml-ops/enaml"
	ccworkerlib "github.com/enaml-ops/omg-cli/plugins/products/cloudfoundry/enaml-gen/cloud_controller_worker"
)

//NewCloudControllerWorkerPartition - Creating a New Cloud Controller Partition
func NewCloudControllerWorkerPartition(c *cli.Context) InstanceGrouper {
	return &CloudControllerWorkerPartition{
		AZs:                   c.StringSlice("az"),
		VMTypeName:            c.String("cc-worker-vm-type"),
		StemcellName:          c.String("stemcell-name"),
		NetworkName:           c.String("network"),
		SystemDomain:          c.String("system-domain"),
		AppDomains:            c.StringSlice("app-domain"),
		AllowAppSSHAccess:     c.Bool("allow-app-ssh-access"),
		Metron:                NewMetron(c),
		ConsulAgent:           NewConsulAgent(c, []string{}),
		NFSMounter:            NewNFSMounter(c),
		StatsdInjector:        NewStatsdInjector(c),
		StagingUploadUser:     c.String("cc-staging-upload-user"),
		StagingUploadPassword: c.String("cc-staging-upload-password"),
		BulkAPIUser:           c.String("cc-bulk-api-user"),
		BulkAPIPassword:       c.String("cc-bulk-api-password"),
		InternalAPIUser:       c.String("cc-internal-api-user"),
		InternalAPIPassword:   c.String("cc-internal-api-password"),
		DbEncryptionKey:       c.String("cc-db-encryption-key"),
	}
}

//ToInstanceGroup - Convert CLoud Controller Partition to an Instance Group
func (s *CloudControllerWorkerPartition) ToInstanceGroup() (ig *enaml.InstanceGroup) {
	ig = &enaml.InstanceGroup{
		Name:      "cloud_controller_worker-partition",
		AZs:       s.AZs,
		Instances: 2, //Not sure where this number should be coming from!
		VMType:    s.VMTypeName,
		Stemcell:  s.StemcellName,
		Networks: []enaml.Network{
			enaml.Network{Name: s.NetworkName},
		},
		Jobs: []enaml.InstanceJob{
			newCloudControllerWorkerJob(s),
			s.ConsulAgent.CreateJob(),
			s.NFSMounter.CreateJob(),
			s.Metron.CreateJob(),
			s.StatsdInjector.CreateJob(),
		},
		Update: enaml.Update{
			MaxInFlight: 1,
			Serial:      true,
		},
	}
	return
}

func newCloudControllerWorkerJob(c *CloudControllerWorkerPartition) enaml.InstanceJob {
	return enaml.InstanceJob{
		Name:    "cloud_controller_worker",
		Release: "cf",
		Properties: &ccworkerlib.CloudControllerWorkerJob{
			Domain:                   c.SystemDomain,
			SystemDomain:             c.SystemDomain,
			AppDomains:               c.AppDomains,
			SystemDomainOrganization: "system",
			Cc: &ccworkerlib.Cc{
				AllowAppSshAccess: c.AllowAppSSHAccess,
				Buildpacks: &ccworkerlib.Buildpacks{
					BlobstoreType: "fog",
					FogConnection: &ccworkerlib.DefaultFogConnection{
						Provider:  "Local",
						LocalRoot: "/var/vcap/nfs/shared",
					},
				},
				Droplets: &ccworkerlib.Droplets{
					BlobstoreType: "fog",
					FogConnection: &ccworkerlib.DefaultFogConnection{
						Provider:  "Local",
						LocalRoot: "/var/vcap/nfs/shared",
					},
				},
				Packages: &ccworkerlib.Packages{
					BlobstoreType: "fog",
					FogConnection: &ccworkerlib.DefaultFogConnection{
						Provider:  "Local",
						LocalRoot: "/var/vcap/nfs/shared",
					},
				},
				ResourcePool: &ccworkerlib.ResourcePool{
					BlobstoreType: "fog",
					FogConnection: &ccworkerlib.DefaultFogConnection{
						Provider:  "Local",
						LocalRoot: "/var/vcap/nfs/shared",
					},
				},
				LoggingLevel:              "debug",
				MaximumHealthCheckTimeout: "600",
				StagingUploadUser:         c.StagingUploadUser,
				StagingUploadPassword:     c.StagingUploadPassword,
				BulkApiUser:               c.BulkAPIUser,
				BulkApiPassword:           c.BulkAPIPassword,
				InternalApiUser:           c.InternalAPIUser,
				InternalApiPassword:       c.InternalAPIPassword,
				DbEncryptionKey:           c.DbEncryptionKey,
			},
		},
	}
}

//HasValidValues - Check if valid values has been populated
func (s *CloudControllerWorkerPartition) HasValidValues() bool {
	return (len(s.AZs) > 0 &&
		s.StemcellName != "" &&
		s.VMTypeName != "" &&
		s.Metron.Zone != "" &&
		s.Metron.Secret != "" &&
		s.NetworkName != "" &&
		s.NFSMounter.hasValidValues() &&
		s.ConsulAgent.HasValidValues())

}