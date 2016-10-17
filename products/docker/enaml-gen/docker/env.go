package docker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Env struct {

	/*HttpsProxy - Descr: HTTPS proxy that Docker should use Default: <nil>
*/
	HttpsProxy interface{} `yaml:"https_proxy,omitempty"`

	/*NoProxy - Descr: List of comma-separated hosts that Docker should skip connecting to the proxy Default: <nil>
*/
	NoProxy interface{} `yaml:"no_proxy,omitempty"`

	/*HttpProxy - Descr: HTTP proxy that Docker should use Default: <nil>
*/
	HttpProxy interface{} `yaml:"http_proxy,omitempty"`

}