package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Login struct {

	/*Saml - Descr: The URL for which SAML identity providers will post assertions to.
If set it overrides the default.
This URL should NOT have the schema (http:// or https:// prefix in it) instead just the hostname.
The schema is derived by #{login.protocol} property.
The default value is #{uaa.url}.replaceFirst('uaa','login'), typically login.example.com
The UAA will display this link in the cf --sso call if there is a SAML provider enabled.
 Default: <nil>
*/
	Saml *Saml `yaml:"saml,omitempty"`

	/*UaaBase - Descr: Deprecated. Use uaa.url for setting the location of UAA. Default: <nil>
*/
	UaaBase interface{} `yaml:"uaa_base,omitempty"`

	/*Notifications - Descr: The url for the notifications service (configure to use Notifications Service instead of SMTP server) Default: <nil>
*/
	Notifications *Notifications `yaml:"notifications,omitempty"`

	/*AssetBaseUrl - Descr: Base url for static assets, allows custom styling of the login server.  Use '/resources/pivotal' for Pivotal style. Default: /resources/oss
*/
	AssetBaseUrl interface{} `yaml:"asset_base_url,omitempty"`

	/*Analytics - Descr: Google analytics code. If Google Analytics is desired set both login.analytics.code and login.analytics.domain Default: <nil>
*/
	Analytics *Analytics `yaml:"analytics,omitempty"`

	/*SpringProfiles - Descr: Deprecated. Use uaa.ldap.enabled - login.spring_profiles is used for backwards compatibility to enable ldap from login config Default: <nil>
*/
	SpringProfiles interface{} `yaml:"spring_profiles,omitempty"`

	/*Ldap - Descr: Deprecated. Use uaa.ldap.localPasswordCompare - login.ldap prefix is used for backwards compatibility to enable ldap from login config Default: true
*/
	Ldap *LoginLdap `yaml:"ldap,omitempty"`

	/*Smtp - Descr: SMTP server username Default: <nil>
*/
	Smtp *Smtp `yaml:"smtp,omitempty"`

	/*Messages - Descr: A nested or flat hash of messages that the login server uses to display UI message
This will be flattened into a java.util.Properties file. The example below will lead
to four properties, where the key is the concatenated value delimited by dot, for example scope.tokens.read=message
 Default: <nil>
*/
	Messages interface{} `yaml:"messages,omitempty"`

	/*Links - Descr: URL for requesting to signup/register for an account Default: /create_account
*/
	Links *Links `yaml:"links,omitempty"`

	/*Tiles - Descr: A list of links to other services to show on the landing page after log in. Default: <nil>
*/
	Tiles interface{} `yaml:"tiles,omitempty"`

	/*SelfServiceLinksEnabled - Descr: Enable self-service account creation and password resets links. Default: <nil>
*/
	SelfServiceLinksEnabled interface{} `yaml:"self_service_links_enabled,omitempty"`

	/*Branding - Descr: This name is used on the UAA Pages and in account management related communication in UAA Default: <nil>
*/
	Branding *Branding `yaml:"branding,omitempty"`

	/*Logout - Descr: A list of URLs. When this list is non null, including empty, and disable=false, logout redirects are allowed, but limited to the whitelist URLs. If a redirect parameter value is not white listed, redirect will be to the default URL. Default: <nil>
*/
	Logout *Logout `yaml:"logout,omitempty"`

	/*Prompt - Descr: The text used to prompt for a username during login Default: Email
*/
	Prompt *Prompt `yaml:"prompt,omitempty"`

	/*HomeRedirect - Descr: URL for configuring a custom home page Default: <nil>
*/
	HomeRedirect interface{} `yaml:"home_redirect,omitempty"`

	/*InvitationsEnabled - Descr: Allows users to send invitations to email addresses outside the system and invite them to create an account. Disabled by default. Default: <nil>
*/
	InvitationsEnabled interface{} `yaml:"invitations_enabled,omitempty"`

	/*Url - Descr: Set if you have an external login server.
The UAA uses this link on by its email service to create links
The UAA uses this as a base domain for internal hostnames so that subdomain can be detected
This defaults to the uaa.url property, and if not set, to login.<domain>
 Default: <nil>
*/
	Url interface{} `yaml:"url,omitempty"`

	/*EntityId - Descr: Deprecated. Use login.saml.entityid Default: <nil>
*/
	EntityId interface{} `yaml:"entity_id,omitempty"`

	/*SignupsEnabled - Descr: Deprecated. Use login.self_service_links_enabled. Instructs UAA to use 'enable account creation flow'. Enabled by default. Default: true
*/
	SignupsEnabled interface{} `yaml:"signups_enabled,omitempty"`

	/*Protocol - Descr: Scheme to use for HTTP communication (http/https) Default: https
*/
	Protocol interface{} `yaml:"protocol,omitempty"`

}