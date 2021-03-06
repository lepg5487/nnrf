/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service. © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ProblemDetails struct {

	Type string `json:"type,omitempty"`

	Title string `json:"title,omitempty"`

	Status int32 `json:"status,omitempty"`

	Detail string `json:"detail,omitempty"`

	Instance string `json:"instance,omitempty"`

	Cause string `json:"cause,omitempty"`

	InvalidParams []InvalidParam `json:"invalidParams,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}
