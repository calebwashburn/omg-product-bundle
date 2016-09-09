package smoke_tests 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type SmokeTests struct {

	/*Org - Descr: Pre-existing CF organization to be used by the smoke tests errand Default: 
*/
	Org interface{} `yaml:"org,omitempty"`

	/*TimeoutScale - Descr: Timeout scale to be used by the smoke tests errand Default: 3
*/
	TimeoutScale interface{} `yaml:"timeout_scale,omitempty"`

	/*AppsDomain - Descr: App domain for Cloud Foundry. Defaults to cf.domain. Default: 
*/
	AppsDomain interface{} `yaml:"apps_domain,omitempty"`

}