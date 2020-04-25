/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service. © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineResponse200 struct {

	// List of the URI of NF instances. It has two members whose names are item and self. The item one contains an array of URIs.
	Links map[string]LinksValueSchema `json:"_links,omitempty"`
}
