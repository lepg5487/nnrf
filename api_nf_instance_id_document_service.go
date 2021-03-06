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
	"errors"
)

// NFInstanceIDDocumentApiService is a service that implents the logic for the NFInstanceIDDocumentApiServicer
// This service should implement the business logic for every endpoint for the NFInstanceIDDocumentApi API. 
// Include any external packages or services that will be required by this service.
type NFInstanceIDDocumentApiService struct {
}

// NewNFInstanceIDDocumentApiService creates a default api service
func NewNFInstanceIDDocumentApiService() NFInstanceIDDocumentApiServicer {
	return &NFInstanceIDDocumentApiService{}
}

// DeregisterNFInstance - Deregisters a given NF Instance
func (s *NFInstanceIDDocumentApiService) DeregisterNFInstance(nfInstanceID string) (interface{}, error) {
	// TODO - update DeregisterNFInstance with the required logic for this service method.
	// Add api_nf_instance_id_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'DeregisterNFInstance' not implemented")
}

// GetNFInstance - Read the profile of a given NF Instance
func (s *NFInstanceIDDocumentApiService) GetNFInstance(nfInstanceID string) (interface{}, error) {
	// TODO - update GetNFInstance with the required logic for this service method.
	// Add api_nf_instance_id_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetNFInstance' not implemented")
}

// RegisterNFInstance - Register a new NF Instance
func (s *NFInstanceIDDocumentApiService) RegisterNFInstance(nfInstanceID string, nfProfile NfProfile, contentEncoding string) (interface{}, error) {
	// TODO - update RegisterNFInstance with the required logic for this service method.
	// Add api_nf_instance_id_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'RegisterNFInstance' not implemented")
}

// UpdateNFInstance - Update NF Instance profile
func (s *NFInstanceIDDocumentApiService) UpdateNFInstance(nfInstanceID string, patchItem []PatchItem) (interface{}, error) {
	// TODO - update UpdateNFInstance with the required logic for this service method.
	// Add api_nf_instance_id_document_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'UpdateNFInstance' not implemented")
}
