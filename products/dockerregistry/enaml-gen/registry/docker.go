package registry 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Docker struct {

	/*Registry - Descr: What port to run the Docker Registry v2 API on Default: 5000
*/
	Registry *Registry `yaml:"registry,omitempty"`

	/*Cache - Descr: IP address of the cache host Default: 
*/
	Cache *Cache `yaml:"cache,omitempty"`

}