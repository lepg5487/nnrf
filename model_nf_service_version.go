/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service. © 2019, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type NfServiceVersion struct {

	ApiVersionInUri string `json:"apiVersionInUri"`

	ApiFullVersion string `json:"apiFullVersion"`

	Expiry time.Time `json:"expiry,omitempty"`
}
