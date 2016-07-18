package cloud_controller_ng 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Doppler struct {

	/*Enabled - Descr: Whether to expose the doppler_logging_endpoint listed at /v2/info Default: true
*/
	Enabled interface{} `yaml:"enabled,omitempty"`

	/*UseSsl - Descr: Whether to use ssl for the doppler_logging_endpoint listed at /v2/info Default: true
*/
	UseSsl interface{} `yaml:"use_ssl,omitempty"`

	/*Port - Descr: Port for doppler_logging_endpoint listed at /v2/info Default: 443
*/
	Port interface{} `yaml:"port,omitempty"`

}