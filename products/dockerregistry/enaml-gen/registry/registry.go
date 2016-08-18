package registry 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Registry struct {

	/*Ssl - Descr: PEM-encoded private key for HTTPS registry operation Default: <nil>
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

	/*Cookie - Descr: A secret cookie used to sign state against tampering. Default: <nil>
*/
	Cookie interface{} `yaml:"cookie,omitempty"`

	/*Root - Descr: Path (on-disk, locally) where the Docker registry should store its data Default: /var/vcap/store/registry
*/
	Root interface{} `yaml:"root,omitempty"`

	/*Port - Descr: What port to run the Docker Registry v2 API on Default: 5000
*/
	Port interface{} `yaml:"port,omitempty"`

}