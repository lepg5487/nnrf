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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A NFInstanceIDDocumentApiController binds http requests to an api service and writes the service results to the http response
type NFInstanceIDDocumentApiController struct {
	service NFInstanceIDDocumentApiServicer
}

// NewNFInstanceIDDocumentApiController creates a default api controller
func NewNFInstanceIDDocumentApiController(s NFInstanceIDDocumentApiServicer) Router {
	return &NFInstanceIDDocumentApiController{ service: s }
}

// Routes returns all of the api route for the NFInstanceIDDocumentApiController
func (c *NFInstanceIDDocumentApiController) Routes() Routes {
	return Routes{ 
		{
			"DeregisterNFInstance",
			strings.ToUpper("Delete"),
			"/nnrf-nfm/v1/nf-instances/{nfInstanceID}",
			c.DeregisterNFInstance,
		},
		{
			"GetNFInstance",
			strings.ToUpper("Get"),
			"/nnrf-nfm/v1/nf-instances/{nfInstanceID}",
			c.GetNFInstance,
		},
		{
			"RegisterNFInstance",
			strings.ToUpper("Put"),
			"/nnrf-nfm/v1/nf-instances/{nfInstanceID}",
			c.RegisterNFInstance,
		},
		{
			"UpdateNFInstance",
			strings.ToUpper("Patch"),
			"/nnrf-nfm/v1/nf-instances/{nfInstanceID}",
			c.UpdateNFInstance,
		},
	}
}

// DeregisterNFInstance - Deregisters a given NF Instance
func (c *NFInstanceIDDocumentApiController) DeregisterNFInstance(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	nfInstanceID := params["nfInstanceID"]
	result, err := c.service.DeregisterNFInstance(nfInstanceID)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// GetNFInstance - Read the profile of a given NF Instance
func (c *NFInstanceIDDocumentApiController) GetNFInstance(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	nfInstanceID := params["nfInstanceID"]
	result, err := c.service.GetNFInstance(nfInstanceID)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// RegisterNFInstance - Register a new NF Instance
func (c *NFInstanceIDDocumentApiController) RegisterNFInstance(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	nfInstanceID := params["nfInstanceID"]
	nfProfile := &NfProfile{}
	if err := json.NewDecoder(r.Body).Decode(&nfProfile); err != nil {
		w.WriteHeader(500)
		return
	}
	
	contentEncoding := r.Header.Get("contentEncoding")
	result, err := c.service.RegisterNFInstance(nfInstanceID, *nfProfile, contentEncoding)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// UpdateNFInstance - Update NF Instance profile
func (c *NFInstanceIDDocumentApiController) UpdateNFInstance(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	nfInstanceID := params["nfInstanceID"]
	patchItem := &[]PatchItem{}
	if err := json.NewDecoder(r.Body).Decode(&patchItem); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.UpdateNFInstance(nfInstanceID, *patchItem)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}