package cloud_controller_clock 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Ccdb struct {

	/*Databases - Descr: Contains the name of the database on the database server Default: <nil>
*/
	Databases interface{} `yaml:"databases,omitempty"`

	/*MaxConnections - Descr: Maximum connections for Sequel Default: 25
*/
	MaxConnections interface{} `yaml:"max_connections,omitempty"`

	/*Address - Descr: The address of the database server Default: <nil>
*/
	Address interface{} `yaml:"address,omitempty"`

	/*Roles - Descr: Users to create on the database when seeding Default: <nil>
*/
	Roles interface{} `yaml:"roles,omitempty"`

	/*PoolTimeout - Descr: The timeout for Sequel pooled connections Default: 10
*/
	PoolTimeout interface{} `yaml:"pool_timeout,omitempty"`

	/*Port - Descr: The port of the database server Default: <nil>
*/
	Port interface{} `yaml:"port,omitempty"`

	/*DbScheme - Descr: The type of database being used. mysql or postgres Default: postgres
*/
	DbScheme interface{} `yaml:"db_scheme,omitempty"`

}