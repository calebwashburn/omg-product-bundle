package vault 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Tcp struct {

	/*Address - Descr: Address for TCP connection Default: 0.0.0.0
*/
	Address interface{} `yaml:"address,omitempty"`

	/*Tls - Descr: Contents of the PEM-encoded TLS server private key Default: <nil>
*/
	Tls *Tls `yaml:"tls,omitempty"`

	/*Port - Descr: Port for TCP connection Default: 8200
*/
	Port interface{} `yaml:"port,omitempty"`

}