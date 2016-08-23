package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type PasswordPolicy struct {

	/*MaxLength - Descr: Maximum number of characters required for password to be considered valid Default: 255
*/
	MaxLength interface{} `yaml:"maxLength,omitempty"`

	/*Global - Descr: Minimum number of special characters required for password to be considered valid Default: 0
*/
	Global *Global `yaml:"global,omitempty"`

	/*ExpirePasswordInMonths - Descr: Number of months after which current password expires Default: 0
*/
	ExpirePasswordInMonths interface{} `yaml:"expirePasswordInMonths,omitempty"`

	/*RequireSpecialCharacter - Descr: Minimum number of special characters required for password to be considered valid Default: 0
*/
	RequireSpecialCharacter interface{} `yaml:"requireSpecialCharacter,omitempty"`

	/*RequireUpperCaseCharacter - Descr: Minimum number of uppercase characters required for password to be considered valid Default: 0
*/
	RequireUpperCaseCharacter interface{} `yaml:"requireUpperCaseCharacter,omitempty"`

	/*MinLength - Descr: Minimum number of characters required for password to be considered valid Default: 0
*/
	MinLength interface{} `yaml:"minLength,omitempty"`

	/*RequireLowerCaseCharacter - Descr: Minimum number of lowercase characters required for password to be considered valid Default: 0
*/
	RequireLowerCaseCharacter interface{} `yaml:"requireLowerCaseCharacter,omitempty"`

	/*RequireDigit - Descr: Minimum number of digits required for password to be considered valid Default: 0
*/
	RequireDigit interface{} `yaml:"requireDigit,omitempty"`

}