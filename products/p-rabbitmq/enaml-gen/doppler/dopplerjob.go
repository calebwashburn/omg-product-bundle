package doppler 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type DopplerJob struct {

	/*Doppler - Descr: boolean value to turn on verbose logging for doppler system (dea agent & doppler server) Default: false
*/
	Doppler *Doppler `yaml:"doppler,omitempty"`

	/*DopplerEndpoint - Descr: Shared secret used to verify cryptographically signed dropsonde messages Default: <nil>
*/
	DopplerEndpoint *DopplerEndpoint `yaml:"doppler_endpoint,omitempty"`

	/*Loggregator - Descr: IPs pointing to the ETCD cluster Default: <nil>
*/
	Loggregator *Loggregator `yaml:"loggregator,omitempty"`

	/*MetronEndpoint - Descr: The host used to emit messages to the Metron agent Default: 127.0.0.1
*/
	MetronEndpoint *MetronEndpoint `yaml:"metron_endpoint,omitempty"`

}