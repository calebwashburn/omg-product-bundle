package deploy_service_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Broker struct {

	/*Password - Descr: Broker basic auth password Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*OrgName - Descr: Org that will host the Service Broker Default: system
*/
	OrgName interface{} `yaml:"org_name,omitempty"`

	/*MaxInstances - Descr: Maximum number of instances Default: 100
*/
	MaxInstances interface{} `yaml:"max_instances,omitempty"`

	/*SpaceName - Descr: Space that will host the Service Broker Default: p-spring-cloud-services
*/
	SpaceName interface{} `yaml:"space_name,omitempty"`

	/*User - Descr: Broker basic auth user Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

}