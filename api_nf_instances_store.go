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

// A NFInstancesStoreApiController binds http requests to an api service and writes the service results to the http response
type NFInstancesStoreApiController struct {
	service NFInstancesStoreApiServicer
}

// NewNFInstancesStoreApiController creates a default api controller
func NewNFInstancesStoreApiController(s NFInstancesStoreApiServicer) Router {
	return &NFInstancesStoreApiController{ service: s }
}

// Routes returns all of the api route for the NFInstancesStoreApiController
func (c *NFInstancesStoreApiController) Routes() Routes {
	return Routes{ 
		{
			"GetNFInstances",
			strings.ToUpper("Get"),
			"/nnrf-nfm/v1/nf-instances",
			c.GetNFInstances,
		},
		{
			"OptionsNFInstances",
			strings.ToUpper("Options"),
			"/nnrf-nfm/v1/nf-instances",
			c.OptionsNFInstances,
		},
	}
}

// GetNFInstances - Retrieves a collection of NF Instances
func (c *NFInstancesStoreApiController) GetNFInstances(w http.ResponseWriter, r *http.Request) { 
	query := r.URL.Query()
	targetnftype := c.query.Get("target-nf-type")
	//targetnftype := c.Query("target-nf-type")
	requesternftype := query.Get("requester-nf-type")
	targetnfinstanceid := query.Get("target-nf-instance-id")
	requesterplmnlist := query.Get("requester-plmn-list")
	limit := query.Get("limit")
	//result, err := c.service.GetNFInstances(targetnftype, limit)

	if query != "" {
		var searchResult SearchResult
		searchResult.NfInstances = []NFProfile
		searchResult.ValidityPeriod(1)
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// OptionsNFInstances - Discover communication options supported by NRF for NF Instances
func (c *NFInstancesStoreApiController) OptionsNFInstances(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.OptionsNFInstances()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}
