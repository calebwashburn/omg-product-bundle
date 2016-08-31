package pmysql

import (
	"strings"

	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/omg-product-bundle/products/p-mysql/enaml-gen/mysql"
	"github.com/enaml-ops/omg-product-bundle/products/p-mysql/enaml-gen/send-email"
)

const (
	innodbBufferPoolSize       int = 2147483648
	maxConnections             int = 1500
	databaseStartupTimeout     int = 600
	wsrepDebug                 int = 1
	seededDBUser                   = "repcanary"
	seededDBName                   = "canary_db"
	adminUsername                  = "root"
	notificationClientUsername     = "mysql-monitoring"
)

func NewMysqlPartition(plgn *Plugin) *enaml.InstanceGroup {
	return &enaml.InstanceGroup{
		Name:               "mysql-partition",
		Lifecycle:          "service",
		Instances:          len(plgn.IPs),
		VMType:             plgn.VMTypeName,
		AZs:                plgn.AZs,
		Stemcell:           plgn.StemcellName,
		PersistentDiskType: plgn.DiskTypeName,
		Jobs: []enaml.InstanceJob{
			newCFMySQLJob(plgn),
			newSendEmailJob(plgn),
			newStreamingMysqlBackupToolJob(plgn),
		},
		Networks: []enaml.Network{
			enaml.Network{Name: plgn.NetworkName, StaticIPs: plgn.IPs},
		},
		Update: enaml.Update{
			MaxInFlight: 1,
		},
	}
}

func newSendEmailJob(plgn *Plugin) enaml.InstanceJob {
	return enaml.InstanceJob{
		Name:    "send-email",
		Release: "mysql-monitoring",
		Properties: &send_email.SendEmailJob{
			Ssl: &send_email.Ssl{
				SkipCertVerify: true,
			},
			Domain: strings.Join([]string{"sys", plgn.BaseDomain}, "."),
			MysqlMonitoring: &send_email.MysqlMonitoring{
				RecipientEmail: plgn.NotificationRecipientEmail,
				AdminClient: &send_email.AdminClient{
					Secret: plgn.UaaAdminClientSecret,
				},
				Client: &send_email.Client{
					Username: notificationClientUsername,
					Secret:   plgn.NotificationClientSecret,
				},
			},
		},
	}
}

func newStreamingMysqlBackupToolJob(plgn *Plugin) enaml.InstanceJob {
	return enaml.InstanceJob{
		Name:    "streaming-mysql-backup-tool",
		Release: "mysql-backup",
	}
}

func newCFMySQLJob(plgn *Plugin) enaml.InstanceJob {
	return enaml.InstanceJob{
		Name:    "mysql",
		Release: "cf-mysql",
		Properties: &mysql.MysqlJob{
			AdminUsername: adminUsername,
			AdminPassword: plgn.AdminPassword,
			CfMysql: &mysql.CfMysql{
				&mysql.Mysql{
					DisableAutoSst:     true,
					InterruptNotifyCmd: "/var/vcap/jobs/send-email/bin/run",
					ClusterHealth: &mysql.ClusterHealth{
						Password: plgn.ClusterHealthPassword,
					},
					GaleraHealthcheck: &mysql.GaleraHealthcheck{
						DbPassword:       plgn.GaleraHealthcheckDBPassword,
						EndpointPassword: plgn.GaleraHealthcheckPassword,
						EndpointUsername: plgn.GaleraHealthcheckUsername,
					},
				},
			},
			ClusterIps:             plgn.IPs,
			DatabaseStartupTimeout: databaseStartupTimeout,
			InnodbBufferPoolSize:   innodbBufferPoolSize,
			MaxConnections:         maxConnections,
			WsrepDebug:             wsrepDebug,
			SeededDatabases: []map[string]string{
				map[string]string{
					"name":     seededDBUser,
					"username": seededDBName,
					"password": plgn.SeededDBPassword,
				},
			},
			SyslogAggregator: &mysql.SyslogAggregator{
				Address:   plgn.SyslogAddress,
				Port:      plgn.SyslogPort,
				Transport: plgn.SyslogTransport,
			},
		},
	}
}
