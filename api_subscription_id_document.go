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

// A SubscriptionIDDocumentApiController binds http requests to an api service and writes the service results to the http response
type SubscriptionIDDocumentApiController struct {
	service SubscriptionIDDocumentApiServicer
}

// NewSubscriptionIDDocumentApiController creates a default api controller
func NewSubscriptionIDDocumentApiController(s SubscriptionIDDocumentApiServicer) Router {
	return &SubscriptionIDDocumentApiController{ service: s }
}

// Routes returns all of the api route for the SubscriptionIDDocumentApiController
func (c *SubscriptionIDDocumentApiController) Routes() Routes {
	return Routes{ 
		{
			"RemoveSubscription",
			strings.ToUpper("Delete"),
			"/nnrf-nfm/v1/subscriptions/{subscriptionID}",
			c.RemoveSubscription,
		},
		{
			"UpdateSubscription",
			strings.ToUpper("Patch"),
			"/nnrf-nfm/v1/subscriptions/{subscriptionID}",
			c.UpdateSubscription,
		},
	}
}

// RemoveSubscription - Deletes a subscription
func (c *SubscriptionIDDocumentApiController) RemoveSubscription(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	subscriptionID := params["subscriptionID"]
	result, err := c.service.RemoveSubscription(subscriptionID)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}

// UpdateSubscription - Updates a subscription
func (c *SubscriptionIDDocumentApiController) UpdateSubscription(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	subscriptionID := params["subscriptionID"]
	patchItem := &[]PatchItem{}
	if err := json.NewDecoder(r.Body).Decode(&patchItem); err != nil {
		w.WriteHeader(500)
		return
	}
	
	result, err := c.service.UpdateSubscription(subscriptionID, *patchItem)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	EncodeJSONResponse(result, nil, w)
}
