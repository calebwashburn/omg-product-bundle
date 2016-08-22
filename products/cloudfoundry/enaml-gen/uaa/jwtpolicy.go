package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type JwtPolicy struct {

	/*Global - Descr: The global refresh token validity for all zones if nothing is configured on the client Default: 2592000
*/
	Global *Global `yaml:"global,omitempty"`

	/*AccessTokenValiditySeconds - Descr: The access token validity for the default zone if nothing is configured on the client. Will override global validity policies for the default zone only. Default: 43200
*/
	AccessTokenValiditySeconds interface{} `yaml:"accessTokenValiditySeconds,omitempty"`

	/*RefreshTokenValiditySeconds - Descr: The refresh token validity for the default zone if nothing is configured on the client. Will override global validity policies for the default zone only. Default: 2592000
*/
	RefreshTokenValiditySeconds interface{} `yaml:"refreshTokenValiditySeconds,omitempty"`

	/*Keys - Descr: Map of key ids to key pairs (signing and verification keys) Default: <nil>
*/
	Keys interface{} `yaml:"keys,omitempty"`

}