package route_emitter 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Nats struct {

	/*Password - Descr: Password for server authentication. Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*Port - Descr: The port for the NATS server to listen on. Default: 4222
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Machines - Descr: IP of each NATS cluster member. Default: <nil>
*/
	Machines interface{} `yaml:"machines,omitempty"`

	/*User - Descr: Username for server authentication. Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

}