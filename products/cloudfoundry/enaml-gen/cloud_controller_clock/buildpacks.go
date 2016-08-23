package cloud_controller_clock 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Buildpacks struct {

	/*WebdavConfig - Descr: The basic auth user that CC uses to connect to the admin endpoint on webdav Default: 
*/
	WebdavConfig *BuildpacksWebdavConfig `yaml:"webdav_config,omitempty"`

	/*BlobstoreType - Descr: The type of blobstore backing to use. Valid values: ['fog', 'webdav'] Default: fog
*/
	BlobstoreType interface{} `yaml:"blobstore_type,omitempty"`

	/*Cdn - Descr: Private key for signing download URIs Default: 
*/
	Cdn *BuildpacksCdn `yaml:"cdn,omitempty"`

	/*BuildpackDirectoryKey - Descr: Directory (bucket) used store buildpacks.  It does not have be pre-created. Default: cc-buildpacks
*/
	BuildpackDirectoryKey interface{} `yaml:"buildpack_directory_key,omitempty"`

	/*FogConnection - Descr: Fog connection hash Default: <nil>
*/
	FogConnection interface{} `yaml:"fog_connection,omitempty"`

}