package cloud_controller_ng 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type DropletsCdn struct {

	/*KeyPairId - Descr: Key pair name for signed download URIs Default: 
*/
	KeyPairId interface{} `yaml:"key_pair_id,omitempty"`

	/*Uri - Descr: URI for a CDN to used for droplet downloads Default: 
*/
	Uri interface{} `yaml:"uri,omitempty"`

	/*PrivateKey - Descr: Private key for signing download URIs Default: 
*/
	PrivateKey interface{} `yaml:"private_key,omitempty"`

}