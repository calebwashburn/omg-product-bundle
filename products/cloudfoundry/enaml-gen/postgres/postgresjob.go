package postgres 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type PostgresJob struct {

	/*Databases - Descr: The postgres `printf` style string that is output at the beginning of each log line Default: %m: 
*/
	Databases *Databases `yaml:"databases,omitempty"`

}