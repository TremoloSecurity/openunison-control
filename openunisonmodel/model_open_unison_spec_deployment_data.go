/*
 * OpenUnison CRD
 *
 * OpenUnison ScaleJS Register API
 *
 * API version: v6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package openunisonmodel

type OpenUnisonSpecDeploymentData struct {
	TokenrequestApi *OpenUnisonSpecDeploymentDataTokenrequestApi `json:"tokenrequest_api,omitempty"`
	ReadinessProbeCommand []string `json:"readiness_probe_command,omitempty"`
	LivenessProbeCommand []string `json:"liveness_probe_command,omitempty"`
	NodeSelectors []OpenUnisonSpecAnnotations `json:"node_selectors,omitempty"`
	PullSecret string `json:"pull_secret,omitempty"`
	Resources *OpenUnisonSpecDeploymentDataResources `json:"resources,omitempty"`
}
