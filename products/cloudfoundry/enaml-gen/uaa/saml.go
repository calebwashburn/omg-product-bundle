package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Saml struct {

	/*Entityid - Descr: This is used as the SAML Service Provider Entity ID. Each  Default: <nil>
*/
	Entityid interface{} `yaml:"entityid,omitempty"`

	/*ServiceProviderKey - Descr: Private key for the service provider certificate. Default: <nil>
*/
	ServiceProviderKey interface{} `yaml:"serviceProviderKey,omitempty"`

	/*Providers - Descr: Contains a hash of SAML Identity Providers, the key is the IDP Alias, followed by key/value pairs. To learn more about how to setup a saml identity provider go to https://simplesamlphp.org Default: <nil>
*/
	Providers interface{} `yaml:"providers,omitempty"`

	/*EntityBaseUrl - Descr: The URL for which SAML identity providers will post assertions to.
If set it overrides the default.
This URL should NOT have the schema (http:// or https:// prefix in it) instead just the hostname.
The schema is derived by #{login.protocol} property.
The default value is #{uaa.url}.replaceFirst('uaa','login'), typically login.example.com
The UAA will display this link in the cf --sso call if there is a SAML provider enabled.
 Default: <nil>
*/
	EntityBaseUrl interface{} `yaml:"entity_base_url,omitempty"`

	/*ServiceProviderCertificate - Descr: Service provider certificate. Default: <nil>
*/
	ServiceProviderCertificate interface{} `yaml:"serviceProviderCertificate,omitempty"`

	/*SignRequest - Descr: Global property to sign Local/SP requests Default: true
*/
	SignRequest interface{} `yaml:"signRequest,omitempty"`

	/*IdpMetadataFile - Descr: Deprecated. Use login.saml.providers list objects Default: <nil>
*/
	IdpMetadataFile interface{} `yaml:"idp_metadata_file,omitempty"`

	/*SignMetaData - Descr: Global property to sign Local/SP metadata Default: true
*/
	SignMetaData interface{} `yaml:"signMetaData,omitempty"`

	/*AssertionConsumerIndex - Descr: Deprecated. Use login.saml.providers list objects Default: 1
*/
	AssertionConsumerIndex interface{} `yaml:"assertionConsumerIndex,omitempty"`

	/*IdpEntityAlias - Descr: Deprecated. Use login.saml.providers list objects Default: <nil>
*/
	IdpEntityAlias interface{} `yaml:"idpEntityAlias,omitempty"`

	/*Socket - Descr: Read timeout in milliseconds for SAML metadata HTTP requests Default: 10000
*/
	Socket *Socket `yaml:"socket,omitempty"`

	/*NameidFormat - Descr: Deprecated. Use login.saml.providers list objects Default: urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress
*/
	NameidFormat interface{} `yaml:"nameidFormat,omitempty"`

	/*IdpMetadataURL - Descr: Deprecated. Use login.saml.providers list objects Default: <nil>
*/
	IdpMetadataURL interface{} `yaml:"idpMetadataURL,omitempty"`

	/*ServiceProviderKeyPassword - Descr: Password to protect the service provider private key. Default: <nil>
*/
	ServiceProviderKeyPassword interface{} `yaml:"serviceProviderKeyPassword,omitempty"`

	/*WantAssertionSigned - Descr: Global property to request that external IDPs sign their SAML assertion before sending them to the UAA Default: false
*/
	WantAssertionSigned interface{} `yaml:"wantAssertionSigned,omitempty"`

	/*MetadataTrustCheck - Descr: Deprecated. Use login.saml.providers list objects Default: true
*/
	MetadataTrustCheck interface{} `yaml:"metadataTrustCheck,omitempty"`

}