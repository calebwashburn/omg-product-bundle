package cloudfoundry

import (
	"strings"

	"github.com/codegangsta/cli"
	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/omg-cli/pluginlib/product"
	"github.com/enaml-ops/omg-cli/pluginlib/util"
	"github.com/xchapter7x/lo"
	"gopkg.in/yaml.v2"
)

func init() {
	RegisterInstanceGrouperFactory(NewConsulPartition)
	RegisterInstanceGrouperFactory(NewNatsPartition)
	RegisterInstanceGrouperFactory(NewEtcdPartition)
	RegisterInstanceGrouperFactory(NewNFSPartition)
	RegisterInstanceGrouperFactory(NewGoRouterPartition)
	RegisterInstanceGrouperFactory(NewMySQLProxyPartition)
	RegisterInstanceGrouperFactory(NewMySQLPartition)
	RegisterInstanceGrouperFactory(NewCloudControllerPartition)
	RegisterInstanceGrouperFactory(NewClockGlobalPartition)
	RegisterInstanceGrouperFactory(NewCloudControllerWorkerPartition)
	RegisterInstanceGrouperFactory(NewDiegoBrainPartition)
	RegisterInstanceGrouperFactory(NewBootstrapPartition)
	RegisterInstanceGrouperFactory(NewDiegoDatabasePartition)
	RegisterInstanceGrouperFactory(NewDiegoCellPartition)
	RegisterInstanceGrouperFactory(NewDopplerPartition)
	RegisterInstanceGrouperFactory(NewLoggregatorTrafficController)
	RegisterInstanceGrouperFactory(NewUAAPartition)
	RegisterInstanceGrouperFactory(NewSmokeErrand)
	RegisterInstanceGrouperFactory(NewHaProxyPartition)

	acceptanceTests := func(c *cli.Context) InstanceGrouper {
		return NewAcceptanceTestsPartition(c, true)
	}
	internetLessAcceptanceTests := func(c *cli.Context) InstanceGrouper {
		return NewAcceptanceTestsPartition(c, false)
	}
	RegisterInstanceGrouperFactory(acceptanceTests)
	RegisterInstanceGrouperFactory(internetLessAcceptanceTests)
}

//GetFlags -
func (s *Plugin) GetFlags() (flags []cli.Flag) {
	return []cli.Flag{
		// shared for all instance groups:
		createStringFlag("stemcell-name", "the name of your desired stemcell"),
		createStringSliceFlag("az", "list of AZ names to use"),
		createStringFlag("network", "the name of the network to use"),
		createStringFlag("system-domain", "System Domain"),
		createStringSliceFlag("app-domain", "Applications Domain"),
		createBoolFlag("allow-app-ssh-access", "Allow SSH access for CF applications"),

		createStringSliceFlag("router-ip", "a list of the router ips you wish to use"),
		createStringFlag("router-vm-type", "the name of your desired vm size"),
		createStringFlag("router-ssl-cert", "the go router ssl cert, or a filename preceded by '@'"),
		createStringFlag("router-ssl-key", "the go router ssl key, or a filename preceded by '@'"),
		createStringFlag("router-user", "the username of the go-routers", "router_status"),
		createStringFlag("router-pass", "the password of the go-routers"),
		createBoolFlag("router-enable-ssl", "enable or disable ssl on your routers"),

		createStringSliceFlag("haproxy-ip", "a list of the haproxy ips you wish to use"),
		createStringFlag("haproxy-vm-type", "the name of your desired vm size"),

		createStringFlag("nats-vm-type", "the name of your desired vm size for NATS"),
		createStringFlag("nats-user", "username for your nats pool", "nats"),
		createStringFlag("nats-pass", "password for your nats pool", "nats-password"),
		createIntFlag("nats-port", "the port for the NATS server to listen on"),
		createStringSliceFlag("nats-machine-ip", "ip of a nats node vm"),

		createStringFlag("metron-zone", "zone guid for the metron agent"),
		createStringFlag("metron-secret", "shared secret for the metron agent endpoint"),
		createIntFlag("metron-port", "local metron agent's port"),

		createStringSliceFlag("consul-ip", "a list of the consul ips you wish to use"),
		createStringFlag("consul-vm-type", "the name of your desired vm size for consul"),
		createStringSliceFlag("consul-encryption-key", "encryption key for consul"),
		createStringFlag("consul-ca-cert", "ca cert for consul, or a filename preceded by '@'"),
		createStringFlag("consul-agent-cert", "agent cert for consul, or a filename preceded by '@'"),
		createStringFlag("consul-agent-key", "agent key for consul, or a filename preceded by '@'"),
		createStringFlag("consul-server-cert", "server cert for consul, or a filename preceded by '@'"),
		createStringFlag("consul-server-key", "server key for consul, or a filename preceded by '@'"),

		createStringFlag("syslog-address", "address of syslog server"),
		createIntFlag("syslog-port", "port of syslog server"),
		createStringFlag("syslog-transport", "transport to syslog server"),

		createStringSliceFlag("etcd-machine-ip", "ip of a etcd node vm"),
		createStringFlag("etcd-vm-type", "the name of your desired vm size for etcd"),
		createStringFlag("etcd-disk-type", "the name of your desired persistent disk type for etcd"),

		createStringSliceFlag("nfs-ip", "a list of the nfs ips you wish to use"),
		createStringFlag("nfs-vm-type", "the name of your desired vm size for nfs"),
		createStringFlag("nfs-disk-type", "the name of your desired persistent disk type for nfs"),
		createStringFlag("nfs-server-address", "NFS Server address"),
		createStringFlag("nfs-share-path", "NFS Share Path"),
		createStringSliceFlag("nfs-allow-from-network-cidr", "the network cidr you wish to allow connections to nfs from"),

		//Mysql Flags
		createStringSliceFlag("mysql-ip", "a list of the mysql ips you wish to use"),
		createStringFlag("mysql-vm-type", "the name of your desired vm size for mysql"),
		createStringFlag("mysql-disk-type", "the name of your desired persistent disk type for mysql"),
		createStringFlag("mysql-admin-password", "admin password for mysql"),
		createStringFlag("mysql-bootstrap-username", "bootstrap username for mysql"),
		createStringFlag("mysql-bootstrap-password", "bootstrap password for mysql"),

		//MySQL proxy flags
		createStringSliceFlag("mysql-proxy-ip", "a list of -mysql proxy ips you wish to use"),
		createStringFlag("mysql-proxy-vm-type", "the name of your desired vm size for mysql proxy"),
		createStringFlag("mysql-proxy-external-host", "Host name of MySQL proxy"),
		createStringFlag("mysql-proxy-api-username", "Proxy API user name"),
		createStringFlag("mysql-proxy-api-password", "Proxy API password"),

		//CC Worker Partition Flags
		createStringFlag("cc-worker-vm-type", "the name of the desired vm type for cc worker"),
		createStringFlag("cc-staging-upload-user", "user name for staging upload"),
		createStringFlag("cc-staging-upload-password", "password for staging upload"),
		createStringFlag("cc-bulk-api-user", "user name for bulk api calls"),
		createStringFlag("cc-bulk-api-password", "password for bulk api calls"),
		createIntFlag("cc-bulk-batch-size", "number of apps to fetch at once from bulk API"),
		createStringFlag("cc-db-encryption-key", "Cloud Controller DB encryption key"),
		createStringFlag("cc-internal-api-user", "user name for Internal API calls"),
		createStringFlag("cc-internal-api-password", "password for Internal API calls"),
		createIntFlag("cc-uploader-poll-interval", "CC uploader job polling interval, in seconds"),
		createIntFlag("cc-fetch-timeout", "how long to wait for completion of requests to CC, in seconds"),
		createStringFlag("cc-vm-type", "Cloud Controller VM Type"),
		createStringFlag("host-key-fingerprint", "Host Key Fingerprint"),
		createStringFlag("support-address", "Support URL"),
		createStringFlag("min-cli-version", "Min CF CLI Version supported"),

		createStringFlag("db-uaa-username", "uaa db username"),
		createStringFlag("db-uaa-password", "uaa db password"),
		createStringFlag("db-ccdb-username", "ccdb db username"),
		createStringFlag("db-ccdb-password", "ccdb db password"),
		createStringFlag("db-console-username", "console db username"),
		createStringFlag("db-console-password", "console db password"),

		//Diego Database
		createStringSliceFlag("diego-db-ip", "a list of static IPs for the diego database partitions"),
		createStringFlag("diego-db-vm-type", "the name of the desired vm type for the diego db"),
		createStringFlag("diego-db-disk-type", "the name of your desired persistent disk type for the diego db"),
		createStringFlag("diego-db-passphrase", "the passphrase for your database"),
		createStringFlag("bbs-server-cert", "BBS server SSL cert (or a file containing it: file format `@filepath`)"),
		createStringFlag("bbs-server-key", "BBS server SSL key (or a file containing it: file format `@filepath`)"),
		createStringFlag("etcd-server-key", "etcd server SSL key (or a file containing it: file format `@filepath`)"),
		createStringFlag("etcd-server-cert", "etcd server cert  (or a file containing it: file format `@filepath`)"),
		createStringFlag("etcd-client-key", "etcd client SSL key (or a file containing it: file format `@filepath`)"),
		createStringFlag("etcd-client-cert", "etcd client SSL cert (or a file containing it: file format `@filepath`)"),
		createStringFlag("etcd-peer-key", "etcd peer SSL key (or a file containing it: file format `@filepath`)"),
		createStringFlag("etcd-peer-cert", "etcd peer SSL cert (or a file containing it: file format `@filepath`)"),

		// Diego Cell
		createStringSliceFlag("diego-cell-ip", "a list of static IPs for the diego cell"),
		createStringFlag("diego-cell-vm-type", "the name of the desired vm type for the diego cell"),
		createStringFlag("diego-cell-disk-type", "the name of your desired persistent disk type for the diego cell"),

		// Diego Brain
		createStringSliceFlag("diego-brain-ip", "a list of static IPs for the diego brain"),
		createStringFlag("diego-brain-vm-type", "the name of the desired vm type for the diego brain"),
		createStringFlag("diego-brain-disk-type", "the name of your desired persistent disk type for the diego brain"),

		createStringFlag("bbs-ca-cert", "BBS CA SSL cert (or a file containing it: file format `@filepath`)"),
		createStringFlag("bbs-client-cert", "BBS client SSL cert (or a file containing it: file format `@filepath`)"),
		createStringFlag("bbs-client-key", "BBS client SSL key (or a file containing it: file format `@filepath`)"),
		createStringFlag("bbs-api", "location of the bbs api"),
		createBoolTFlag("bbs-require-ssl", "enable SSL for all communications with the BBS"),

		createBoolTFlag("skip-cert-verify", "ignore bad SSL certificates when connecting over HTTPS"),

		createStringFlag("fs-listen-addr", "address of interface on which to serve files"),
		createStringFlag("fs-static-dir", "fully qualified path to the doc root for the file server's static files"),
		createStringFlag("fs-debug-addr", "address at which to serve debug info"),
		createStringFlag("fs-log-level", "file server log level"),

		createIntFlag("cc-external-port", "external port of the Cloud Controller API"),
		createStringFlag("ssh-proxy-uaa-secret", "the OAuth client secret used to authenticate the SSH proxy"),
		createStringFlag("traffic-controller-url", "the URL of the traffic controller"),
		createStringFlag("clock-global-vm-type", "the name of the desired vm type for the clock global partition"),

		//Doppler
		createStringSliceFlag("doppler-ip", "a list of the doppler ips you wish to use"),
		createStringFlag("doppler-vm-type", "the name of your desired vm size for doppler"),
		createStringFlag("doppler-zone", "the name zone for doppler"),
		createIntFlag("doppler-drain-buffer-size", "message drain buffer size"),
		createStringFlag("doppler-shared-secret", "doppler shared secret"),

		//Loggregator Traffic Controller
		cli.StringSliceFlag{Name: "loggregator-traffic-controller-ip", Usage: "a list of loggregator traffic controller IPs"},
		cli.StringFlag{Name: "loggregator-traffic-controller-vmtype", Usage: "the name of your desired vm size for the loggregator traffic controller"},

		//UAA
		createStringFlag("uaa-vm-type", "the name of your desired vm size for uaa"),
		createIntFlag("uaa-instances", "the number of your desired vms for uaa"),

		createStringFlag("uaa-company-name", "name of company for UAA branding"),
		createStringFlag("uaa-product-logo", "product logo for UAA branding"),
		createStringFlag("uaa-square-logo", "square logo for UAA branding"),
		createStringFlag("uaa-footer-legal-txt", "legal text for UAA branding"),
		createBoolTFlag("uaa-enable-selfservice-links", "enable self service links"),
		createBoolTFlag("uaa-signups-enabled", "enable signups"),
		createStringFlag("uaa-login-protocol", "uaa login protocol, default https"),
		createStringFlag("uaa-saml-service-provider-key", "saml service provider key for uaa"),
		createStringFlag("uaa-saml-service-provider-certificate", "saml service provider certificate for uaa"),
		createStringFlag("uaa-jwt-signing-key", "signing key for jwt used by UAA"),
		createStringFlag("uaa-jwt-verification-key", "verification key for jwt used by UAA"),
		createBoolFlag("uaa-ldap-enabled", "is ldap enabled for UAA"),
		createStringFlag("uaa-ldap-url", "url for ldap server"),
		createStringFlag("uaa-ldap-user-dn", "userDN to bind to ldap with"),
		createStringFlag("uaa-ldap-user-password", "bind password for ldap user"),
		createStringFlag("uaa-ldap-search-filter", "search filter for users"),
		createStringFlag("uaa-ldap-search-base", "search base for users"),
		createStringFlag("uaa-ldap-mail-attributename", "attribute name for mail"),
		createStringFlag("uaa-admin-secret", "admin account client secret"),

		//User accounts
		createStringFlag("admin-password", "password for admin account"),
		createStringFlag("push-apps-manager-password", "password for push_apps_manager account"),
		createStringFlag("smoke-tests-password", "password for smoke_tests account"),
		createStringFlag("system-services-password", "password for system_services account"),
		createStringFlag("system-verification-password", "password for system_verification account"),

		//Client secrets
		createStringFlag("opentsdb-firehose-nozzle-client-secret", "client-secret for opentsdb firehose nozzle"),
		createStringFlag("identity-client-secret", "client-secret for identity"),
		createStringFlag("login-client-secret", "client-secret for login"),
		createStringFlag("portal-client-secret", "client-secret for portal"),
		createStringFlag("autoscaling-service-client-secret", "client-secret for autoscaling service"),
		createStringFlag("system-passwords-client-secret", "client-secret for system-passwords"),
		createStringFlag("cc-service-dashboards-client-secret", "client-secret for cc-service-dashboards"),
		createStringFlag("doppler-client-secret", "client-secret for doppler"),
		createStringFlag("gorouter-client-secret", "client-secret for gorouter"),
		createStringFlag("notifications-client-secret", "client-secret for notifications"),
		createStringFlag("notifications-ui-client-secret", "client-secret for notification-ui"),
		createStringFlag("cloud-controller-username-lookup-client-secret", "client-secret for cloud controller username lookup"),
		createStringFlag("cc-routing-client-secret", "client-secret for cc routing"),
		createStringFlag("ssh-proxy-client-secret", "client-secret for ssh proxy"),
		createStringFlag("apps-metrics-client-secret", "client-secret for apps metrics "),
		createStringFlag("apps-metrics-processing-client-secret", "client-secret for apps metrics processing"),

		createStringFlag("errand-vm-type", "vm type to be used for running errands"),
		createStringFlag("haproxy-sslpem", "SSL pem for HAProxy"),

		//Vault stuff
		createStringFlag("vault-domain", "the location of your vault server (ie. http://10.0.0.1:8200)"),
		createStringFlag("vault-hash-password", "the hashname of your secret (ie. secret/pcf-1-passwords"),
		createStringFlag("vault-hash-keycert", "the hashname of your secret (ie. secret/pcf-1-keycert"),
		createStringFlag("vault-hash-ip", "the hashname of your secret (ie. secret/pcf-1-ips"),
		createStringFlag("vault-hash-host", "the hashname of your secret (ie. secret/pcf-1-hosts"),
		createStringFlag("vault-token", "the token to make connections to your vault"),
		createBoolTFlag("vault-active", "use the data which is stored in vault for the flag values it contains"),
	}
}

func createStringFlag(name, usage string, value ...string) cli.StringFlag {
	res := cli.StringFlag{Name: name, Usage: usage, EnvVar: strings.ToUpper(name)}

	if len(value) > 0 {
		res.Value = value[0]
	}
	return res
}

func createBoolFlag(name, usage string) cli.BoolFlag {
	return cli.BoolFlag{Name: name, Usage: usage, EnvVar: strings.ToUpper(name)}
}

func createIntFlag(name, usage string) cli.IntFlag {
	return cli.IntFlag{Name: name, Usage: usage, EnvVar: strings.ToUpper(name)}
}

func createBoolTFlag(name, usage string) cli.BoolTFlag {
	return cli.BoolTFlag{Name: name, Usage: usage, EnvVar: strings.ToUpper(name)}
}

func createStringSliceFlag(name, usage string, value ...string) cli.StringSliceFlag {
	res := cli.StringSliceFlag{Name: name, Usage: usage, EnvVar: strings.ToUpper(name)}

	if len(value) > 0 {
		res.Value = &cli.StringSlice{}
		res.Value.Set(strings.Join(value, ","))
	}
	return res
}

//GetMeta -
func (s *Plugin) GetMeta() product.Meta {
	return product.Meta{
		Name: "cloudfoundry",
	}
}

//GetProduct -
func (s *Plugin) GetProduct(args []string, cloudConfig []byte) (b []byte) {
	flgs := s.GetFlags()
	s.vaultDecorate(args, flgs)
	c := pluginutil.NewContext(args, flgs)
	dm := enaml.NewDeploymentManifest([]byte(``))
	dm.SetName(DeploymentName)

	dm.AddRelease(enaml.Release{Name: CFReleaseName, Version: CFReleaseVersion})
	dm.AddRelease(enaml.Release{Name: CFMysqlReleaseName, Version: CFMysqlReleaseVersion})
	dm.AddRelease(enaml.Release{Name: DiegoReleaseName, Version: DiegoReleaseVersion})
	dm.AddRelease(enaml.Release{Name: GardenReleaseName, Version: GardenReleaseVersion})
	dm.AddRelease(enaml.Release{Name: CFLinuxReleaseName, Version: CFLinuxReleaseVersion})
	dm.AddRelease(enaml.Release{Name: EtcdReleaseName, Version: EtcdReleaseVersion})
	dm.AddRelease(enaml.Release{Name: PushAppsReleaseName, Version: PushAppsReleaseVersion})
	dm.AddRelease(enaml.Release{Name: NotificationsReleaseName, Version: NotificationsReleaseVersion})
	dm.AddRelease(enaml.Release{Name: NotificationsUIReleaseName, Version: NotificationsUIReleaseVersion})
	dm.AddRelease(enaml.Release{Name: CFAutoscalingReleaseName, Version: CFAutoscalingReleaseVersion})

	dm.AddStemcell(enaml.Stemcell{OS: StemcellName, Version: StemcellVersion, Alias: StemcellAlias})

	for _, factory := range factories {
		grouper := factory(c)

		if grouper.HasValidValues() {
			ig := grouper.ToInstanceGroup()
			lo.G.Debug("instance-group: ", ig)
			dm.AddInstanceGroup(ig)

		} else {
			b, _ := yaml.Marshal(grouper)
			lo.G.Panic("invalid values in instance group: ", string(b))
		}
	}

	return dm.Bytes()
}

func (s *Plugin) vaultDecorate(args []string, flgs []cli.Flag) {
	c := pluginutil.NewContext(args, flgs)

	if s.hasValidVaultFlags(c) {
		vault := pluginutil.NewVaultUnmarshal(c.String("vault-domain"), c.String("vault-token"), pluginutil.DefaultClient())
		hashes := []string{
			c.String("vault-hash-password"),
			c.String("vault-hash-keycert"),
			c.String("vault-hash-ip"),
			c.String("vault-hash-host"),
		}

		for _, hash := range hashes {

			if hash != "" {
				vault.UnmarshalFlags(hash, flgs)
			}
		}

	} else {
		lo.G.Debug("complete vault flagset not found:",
			"active: ", c.BoolT("vault-active"),
			"domain: ", c.String("vault-domain"),
			"passhash: ", c.String("vault-hash-password"),
			"keycerthash: ", c.String("vault-hash-keycert"),
			"iphash: ", c.String("vault-hash-ip"),
			"hosthash: ", c.String("vault-hash-host"),
			"vaulttoken: ", c.String("vault-token"),
		)

		if c.BoolT("vault-active") {
			lo.G.Panic("you've activated vault, but have not provided a complete set of values... exiting program now")
		}
	}
}

func (s *Plugin) hasValidVaultFlags(c *cli.Context) bool {
	return c.BoolT("vault-active") &&
		c.String("vault-domain") != "" &&
		c.String("vault-token") != ""
}

//GetContext -
func (s *Plugin) GetContext(args []string) (c *cli.Context) {
	c = pluginutil.NewContext(args, s.GetFlags())
	return
}