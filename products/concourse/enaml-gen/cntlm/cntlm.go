package cntlm 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Cntlm struct {

	/*Domain - Descr: Domain to use for NTLM proxy Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*Proxy - Descr: Proxy server and port example 10.0.0.0:8080 Default: <nil>
*/
	Proxy interface{} `yaml:"proxy,omitempty"`

	/*Username - Descr: Username to use for NTLM proxy Default: <nil>
*/
	Username interface{} `yaml:"username,omitempty"`

	/*Password - Descr: Password to use for NTLM proxy Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

}