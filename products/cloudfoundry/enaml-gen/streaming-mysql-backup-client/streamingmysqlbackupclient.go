package streaming_mysql_backup_client 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type StreamingMysqlBackupClient struct {

	/*TmpFolder - Descr: Folder to download / prepare backups Default: /var/vcap/data/mysql-backups-tmp
*/
	TmpFolder interface{} `yaml:"tmp_folder,omitempty"`

	/*SymmetricKey - Descr: Symmetric Key used to encrypt backups Default: <nil>
*/
	SymmetricKey interface{} `yaml:"symmetric_key,omitempty"`

	/*BackupGenerator - Descr: List of IP address of servers which can generate backups Default: <nil>
*/
	BackupGenerator *BackupGenerator `yaml:"backup-generator,omitempty"`

	/*OutputFolder - Descr: Folder to place the prepared backups Default: <nil>
*/
	OutputFolder interface{} `yaml:"output_folder,omitempty"`

}