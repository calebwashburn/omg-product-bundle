package deploy_service_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type DeployServiceBrokerJob struct {

	/*SpringCloudBroker - Descr: App Push Memory limit for the Service Broker Application Default: 1024
*/
	SpringCloudBroker *SpringCloudBroker `yaml:"spring_cloud_broker,omitempty"`

	/*AppDomains - Descr: Cloud Foundry application domains Default: <nil>
*/
	AppDomains interface{} `yaml:"app_domains,omitempty"`

	/*Ssl - Descr: Whether to verify SSL certs when making web requests Default: <nil>
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

	/*Domain - Descr: Cloud Foundry system domain Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*Uaa - Descr: UAA Client Secret Default: <nil>
*/
	Uaa *Uaa `yaml:"uaa,omitempty"`

}