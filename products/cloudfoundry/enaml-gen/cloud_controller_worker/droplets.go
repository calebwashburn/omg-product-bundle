package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Droplets struct {

	/*WebdavConfig - Descr: The ca cert to use when communicating with webdav Default: 
*/
	WebdavConfig *DropletsWebdavConfig `yaml:"webdav_config,omitempty"`

	/*Cdn - Descr: Private key for signing download URIs Default: 
*/
	Cdn *DropletsCdn `yaml:"cdn,omitempty"`

	/*FogConnection - Descr: Fog connection hash Default: <nil>
*/
	FogConnection interface{} `yaml:"fog_connection,omitempty"`

	/*DropletDirectoryKey - Descr: Directory (bucket) used store droplets.  It does not have be pre-created. Default: cc-droplets
*/
	DropletDirectoryKey interface{} `yaml:"droplet_directory_key,omitempty"`

	/*BlobstoreType - Descr: The type of blobstore backing to use. Valid values: ['fog', 'webdav'] Default: fog
*/
	BlobstoreType interface{} `yaml:"blobstore_type,omitempty"`

}