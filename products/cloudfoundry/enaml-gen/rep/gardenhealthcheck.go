package rep 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type GardenHealthcheck struct {

	/*CommandRetryPause - Descr: Time to wait between retrying garden commands Default: 5s
*/
	CommandRetryPause interface{} `yaml:"command_retry_pause,omitempty"`

	/*Interval - Descr: Frequency for healtchecking garden Default: 10m
*/
	Interval interface{} `yaml:"interval,omitempty"`

	/*Timeout - Descr: Maximum allowed time for garden healthcheck Default: 10m
*/
	Timeout interface{} `yaml:"timeout,omitempty"`

	/*Process - Descr: List of command line args to pass to the garden health check process Default: -c, ls > /tmp/test
*/
	Process *Process `yaml:"process,omitempty"`

}