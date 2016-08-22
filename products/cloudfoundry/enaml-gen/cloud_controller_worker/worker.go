package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Worker struct {

	/*RestartIfConsistentlyAboveMb - Descr: The cc will restart if memory remains above this threshold for 15 monit cycles Default: 384
*/
	RestartIfConsistentlyAboveMb interface{} `yaml:"restart_if_consistently_above_mb,omitempty"`

	/*AlertIfAboveMb - Descr: The cc will alert if memory remains above this threshold for 3 monit cycles Default: 384
*/
	AlertIfAboveMb interface{} `yaml:"alert_if_above_mb,omitempty"`

	/*RestartIfAboveMb - Descr: The cc will restart if memory remains above this threshold for 3 monit cycles Default: 512
*/
	RestartIfAboveMb interface{} `yaml:"restart_if_above_mb,omitempty"`

}