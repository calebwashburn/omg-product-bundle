package broker_registrar 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type BrokerRegistrarJob struct {

	/*Cf - Descr: Username of the admin user Default: <nil>
*/
	Cf *Cf `yaml:"cf,omitempty"`

	/*Broker - Descr: Name of the service broker Default: <nil>
*/
	Broker *Broker `yaml:"broker,omitempty"`

}