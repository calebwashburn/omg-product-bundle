package locator 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Locator struct {

	/*VmMemory - Descr: RAM allocated to the locator VM in MB Default: <nil>
*/
	VmMemory interface{} `yaml:"vm_memory,omitempty"`

	/*Addresses - Descr: List of GemFire Locator addresses of the form X.X.X.X Default: <nil>
*/
	Addresses interface{} `yaml:"addresses,omitempty"`

	/*Port - Descr: Port the Locator will listen on Default: 55221
*/
	Port interface{} `yaml:"port,omitempty"`

	/*RestPort - Descr: Port the Locator will listen on for REST API Default: 8080
*/
	RestPort interface{} `yaml:"rest_port,omitempty"`

}