package cc_uploader 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Diego struct {

	/*CcUploader - Descr: Address of interface on which to serve files Default: 0.0.0.0:9090
*/
	CcUploader *CcUploader `yaml:"cc_uploader,omitempty"`

	/*Ssl - Descr: when connecting over https, ignore bad ssl certificates Default: false
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

}