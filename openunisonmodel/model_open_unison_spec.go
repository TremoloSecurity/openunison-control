/*
 * OpenUnison CRD
 *
 * OpenUnison ScaleJS Register API
 *
 * API version: v6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package openunisonmodel

type OpenUnisonSpec struct {
	Image string `json:"image,omitempty"`
	Replicas int32 `json:"replicas,omitempty"`
	EnableActivemq bool `json:"enable_activemq,omitempty"`
	ActivemqImage string `json:"activemq_image,omitempty"`
	DestSecret string `json:"dest_secret,omitempty"`
	SourceSecret string `json:"source_secret,omitempty"`
	SecretData []string `json:"secret_data,omitempty"`
	MyvdConfigmap string `json:"myvd_configmap,omitempty"`
	Openshift *OpenUnisonSpecOpenshift `json:"openshift,omitempty"`
	Hosts []OpenUnisonSpecHosts `json:"hosts,omitempty"`
	DeploymentData *OpenUnisonSpecDeploymentData `json:"deployment_data,omitempty"`
	NonSecretData []OpenUnisonSpecAnnotations `json:"non_secret_data,omitempty"`
	OpenunisonNetworkConfiguration *OpenUnisonSpecOpenunisonNetworkConfiguration `json:"openunison_network_configuration,omitempty"`
	SamlRemoteIdp []OpenUnisonSpecSamlRemoteIdp `json:"saml_remote_idp,omitempty"`
	RunSql string `json:"run_sql,omitempty"`
	SqlCheckQuery string `json:"sql_check_query,omitempty"`
	KeyStore *OpenUnisonSpecKeyStore `json:"key_store,omitempty"`
}
