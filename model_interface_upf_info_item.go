/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service. © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InterfaceUpfInfoItem struct {

	InterfaceType UpInterfaceType `json:"interfaceType"`

	Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses,omitempty"`

	Ipv6EndpointAddresses []Ipv6Addr `json:"ipv6EndpointAddresses,omitempty"`

	EndpointFqdn string `json:"endpointFqdn,omitempty"`

	NetworkInstance string `json:"networkInstance,omitempty"`
}