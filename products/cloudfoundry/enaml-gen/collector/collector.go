package collector 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Collector struct {

	/*MemoryThreshold - Descr: Memory threshold for collector restart (Mb) Default: 800
*/
	MemoryThreshold interface{} `yaml:"memory_threshold,omitempty"`

	/*Datadog - Descr: Datadog application key Default: <nil>
*/
	Datadog *Datadog `yaml:"datadog,omitempty"`

	/*LoggingLevel - Descr: the logging level for the collector Default: info
*/
	LoggingLevel interface{} `yaml:"logging_level,omitempty"`

	/*UseAwsCloudwatch - Descr: enable CloudWatch plugin Default: false
*/
	UseAwsCloudwatch interface{} `yaml:"use_aws_cloudwatch,omitempty"`

	/*Intervals - Descr: the interval in seconds that varz is checked Default: 30
*/
	Intervals *Intervals `yaml:"intervals,omitempty"`

	/*UseTsdb - Descr: enable OpenTsdb plugin Default: false
*/
	UseTsdb interface{} `yaml:"use_tsdb,omitempty"`

	/*DeploymentName - Descr: name for this bosh deployment. All metrics will be tagged with deployment:XXX when sending them to CloudWatch, Datadog and Graphite Default: <nil>
*/
	DeploymentName interface{} `yaml:"deployment_name,omitempty"`

	/*Aws - Descr: AWS secret for CloudWatch access Default: <nil>
*/
	Aws *Aws `yaml:"aws,omitempty"`

	/*Graphite - Descr: IP address of Graphite Default: <nil>
*/
	Graphite *Graphite `yaml:"graphite,omitempty"`

	/*UseGraphite - Descr: enable Graphite plugin Default: false
*/
	UseGraphite interface{} `yaml:"use_graphite,omitempty"`

	/*UseDatadog - Descr: enable Datadog plugin Default: false
*/
	UseDatadog interface{} `yaml:"use_datadog,omitempty"`

	/*Opentsdb - Descr: IP address of OpenTsdb Default: <nil>
*/
	Opentsdb *Opentsdb `yaml:"opentsdb,omitempty"`

}