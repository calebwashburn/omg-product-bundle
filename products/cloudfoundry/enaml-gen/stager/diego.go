package stager 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Diego struct {

	/*Stager - Descr: External port to access the Cloud Controller Default: 9022
*/
	Stager *Stager `yaml:"stager,omitempty"`

	/*Ssl - Descr: when connecting over https, ignore bad ssl certificates Default: false
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

}