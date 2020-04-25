/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service. © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type N2InterfaceAmfInfo struct {

	Ipv4EndpointAddress []string `json:"ipv4EndpointAddress,omitempty"`

	Ipv6EndpointAddress []Ipv6Addr `json:"ipv6EndpointAddress,omitempty"`

	AmfName string `json:"amfName,omitempty"`
}
