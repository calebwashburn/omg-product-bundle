package deploy_autoscaling 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type DeployAutoscalingJob struct {

	/*Ssl - Descr: Whether to verify SSL certs when making web requests Default: <nil>
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

	/*Domain - Descr: CloudFoundry system domain Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*AppDomains - Descr: CloudFoundry application domains Default: <nil>
*/
	AppDomains interface{} `yaml:"app_domains,omitempty"`

	/*Autoscale - Descr: Enable diego deployment of autoscaling Default: <nil>
*/
	Autoscale *Autoscale `yaml:"autoscale,omitempty"`

	/*Uaa - Descr: UAA Client Secret Default: <nil>
*/
	Uaa *Uaa `yaml:"uaa,omitempty"`

}