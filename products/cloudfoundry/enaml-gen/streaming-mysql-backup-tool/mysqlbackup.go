package streaming_mysql_backup_tool 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type MysqlBackup struct {

	/*Command - Descr: Command which is executed in order to take a backup Default: <nil>
*/
	Command interface{} `yaml:"command,omitempty"`

	/*Port - Descr: Port number used for listening for backup requests Default: 8081
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Credentials - Descr: Username used by backup client to stream a backup from the mysql node Default: <nil>
*/
	Credentials *Credentials `yaml:"credentials,omitempty"`

}