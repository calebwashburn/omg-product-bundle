package hm9000 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Hm9000 struct {

	/*ServerKey - Descr: PEM-encoded server key Default: <nil>
*/
	ServerKey interface{} `yaml:"server_key,omitempty"`

	/*CaCert - Descr: PEM-encoded CA certificate Default: <nil>
*/
	CaCert interface{} `yaml:"ca_cert,omitempty"`

	/*SenderMessageLimit - Descr: The maximum number of messages the sender should send per invocation. Default: 60
*/
	SenderMessageLimit interface{} `yaml:"sender_message_limit,omitempty"`

	/*Port - Descr: The port to serve API requests Default: <nil>
*/
	Port interface{} `yaml:"port,omitempty"`

	/*FetcherNetworkTimeoutInSeconds - Descr: Each API call to the CC must succeed within this timeout. Default: 30
*/
	FetcherNetworkTimeoutInSeconds interface{} `yaml:"fetcher_network_timeout_in_seconds,omitempty"`

	/*ServerCert - Descr: PEM-encoded server certificate Default: <nil>
*/
	ServerCert interface{} `yaml:"server_cert,omitempty"`

	/*DesiredStateBatchSize - Descr: The batch size when fetching desired state information from the CC. Default: 5000
*/
	DesiredStateBatchSize interface{} `yaml:"desired_state_batch_size,omitempty"`

}