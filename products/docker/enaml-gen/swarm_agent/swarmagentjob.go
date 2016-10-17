package swarm_agent 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type SwarmAgentJob struct {

	/*Env - Descr: HTTPS proxy that Docker should use Default: <nil>
*/
	Env *Env `yaml:"env,omitempty"`

	/*SwarmAgent - Descr: Time in second between each heartbeat Default: 20s
*/
	SwarmAgent *SwarmAgent `yaml:"swarm_agent,omitempty"`

	/*Swarm - Descr: Swarm discovery options Default: []
*/
	Swarm *Swarm `yaml:"swarm,omitempty"`

	/*Docker - Descr: TCP port where Docker daemon will listen to (if not set, TCP will not be available) Default: 4243
*/
	Docker *Docker `yaml:"docker,omitempty"`

}