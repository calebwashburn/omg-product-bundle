package cloudfoundry

import (
	"github.com/codegangsta/cli"
	"github.com/enaml-ops/enaml"
	consullib "github.com/enaml-ops/omg-product-bundle/products/cloudfoundry/enaml-gen/consul_agent"
	"github.com/enaml-ops/pluginlib/util"
	"github.com/xchapter7x/lo"
)

//NewConsulAgent -
func NewConsulAgent(c *cli.Context, services []string) *ConsulAgent {
	ca := &ConsulAgent{
		EncryptKeys: c.StringSlice("consul-encryption-key"),
		NetworkIPs:  c.StringSlice("consul-ip"),
		Services:    services,
	}
	ca.loadSSL(c)
	return ca
}

//NewConsulAgentServer -
func NewConsulAgentServer(c *cli.Context) *ConsulAgent {
	ca := &ConsulAgent{
		EncryptKeys: c.StringSlice("consul-encryption-key"),
		NetworkIPs:  c.StringSlice("consul-ip"),
		Mode:        "server",
	}
	ca.loadSSL(c)
	return ca
}

func (ca *ConsulAgent) loadSSL(c *cli.Context) {
	caCert, err := pluginutil.LoadResourceFromContext(c, "consul-ca-cert")
	if err != nil {
		lo.G.Fatalf("consul ca cert: %s\n", err.Error())
	}
	agentCert, err := pluginutil.LoadResourceFromContext(c, "consul-agent-cert")
	if err != nil {
		lo.G.Fatalf("consul agent cert: %s\n", err.Error())
	}
	agentKey, err := pluginutil.LoadResourceFromContext(c, "consul-agent-key")
	if err != nil {
		lo.G.Fatalf("consul agent key: %s\n", err.Error())
	}
	serverCert, err := pluginutil.LoadResourceFromContext(c, "consul-server-cert")
	if err != nil {
		lo.G.Fatalf("consul server cert: %s\n", err.Error())
	}
	serverKey, err := pluginutil.LoadResourceFromContext(c, "consul-server-key")
	if err != nil {
		lo.G.Fatalf("consul server key: %s\n", err.Error())
	}

	ca.CaCert = caCert
	ca.AgentCert = agentCert
	ca.ServerCert = serverCert
	ca.AgentKey = agentKey
	ca.ServerKey = serverKey
}

//CreateJob - Create the yaml job structure
func (s *ConsulAgent) CreateJob() enaml.InstanceJob {

	serviceMap := make(map[string]map[string]string)
	for _, serviceName := range s.Services {
		serviceMap[serviceName] = make(map[string]string)
	}

	return enaml.InstanceJob{
		Name:    "consul_agent",
		Release: "cf",
		Properties: &consullib.ConsulAgentJob{
			Consul: &consullib.Consul{
				EncryptKeys: s.EncryptKeys,
				CaCert:      s.CaCert,
				AgentCert:   s.AgentCert,
				AgentKey:    s.AgentKey,
				ServerCert:  s.ServerCert,
				ServerKey:   s.ServerKey,
				Agent: &consullib.Agent{
					Domain: "cf.internal",
					Mode:   s.getMode(),
					Servers: &consullib.Servers{
						Lan: s.NetworkIPs,
					},
					Services: serviceMap,
				},
			},
		},
	}
}

func (s *ConsulAgent) getMode() interface{} {
	if s.Mode != "" {
		return s.Mode
	}
	return nil
}

//HasValidValues -
func (s *ConsulAgent) HasValidValues() bool {
	return len(s.NetworkIPs) > 0 &&
		len(s.EncryptKeys) > 0 &&
		s.CaCert != "" &&
		s.AgentCert != "" &&
		s.AgentKey != "" &&
		s.ServerCert != "" &&
		s.ServerKey != ""
}
