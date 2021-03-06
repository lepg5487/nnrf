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
	"net/http"
)


// NFInstanceIDDocumentApiRouter defines the required methods for binding the api requests to a responses for the NFInstanceIDDocumentApi
// The NFInstanceIDDocumentApiRouter implementation should parse necessary information from the http request, 
// pass the data to a NFInstanceIDDocumentApiServicer to perform the required actions, then write the service results to the http response.
type NFInstanceIDDocumentApiRouter interface { 
	DeregisterNFInstance(http.ResponseWriter, *http.Request)
	GetNFInstance(http.ResponseWriter, *http.Request)
	RegisterNFInstance(http.ResponseWriter, *http.Request)
	UpdateNFInstance(http.ResponseWriter, *http.Request)
}
// NFInstancesStoreApiRouter defines the required methods for binding the api requests to a responses for the NFInstancesStoreApi
// The NFInstancesStoreApiRouter implementation should parse necessary information from the http request, 
// pass the data to a NFInstancesStoreApiServicer to perform the required actions, then write the service results to the http response.
type NFInstancesStoreApiRouter interface { 
	GetNFInstances(http.ResponseWriter, *http.Request)
	OptionsNFInstances(http.ResponseWriter, *http.Request)
}
// SubscriptionIDDocumentApiRouter defines the required methods for binding the api requests to a responses for the SubscriptionIDDocumentApi
// The SubscriptionIDDocumentApiRouter implementation should parse necessary information from the http request, 
// pass the data to a SubscriptionIDDocumentApiServicer to perform the required actions, then write the service results to the http response.
type SubscriptionIDDocumentApiRouter interface { 
	RemoveSubscription(http.ResponseWriter, *http.Request)
	UpdateSubscription(http.ResponseWriter, *http.Request)
}
// SubscriptionsCollectionApiRouter defines the required methods for binding the api requests to a responses for the SubscriptionsCollectionApi
// The SubscriptionsCollectionApiRouter implementation should parse necessary information from the http request, 
// pass the data to a SubscriptionsCollectionApiServicer to perform the required actions, then write the service results to the http response.
type SubscriptionsCollectionApiRouter interface { 
	CreateSubscription(http.ResponseWriter, *http.Request)
}


// NFInstanceIDDocumentApiServicer defines the api actions for the NFInstanceIDDocumentApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type NFInstanceIDDocumentApiServicer interface { 
	DeregisterNFInstance(string) (interface{}, error)
	GetNFInstance(string) (interface{}, error)
	RegisterNFInstance(string, NfProfile, string) (interface{}, error)
	UpdateNFInstance(string, []PatchItem) (interface{}, error)
}


// NFInstancesStoreApiServicer defines the api actions for the NFInstancesStoreApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type NFInstancesStoreApiServicer interface { 
	GetNFInstances(NfType, int32) (interface{}, error)
	OptionsNFInstances() (interface{}, error)
}


// SubscriptionIDDocumentApiServicer defines the api actions for the SubscriptionIDDocumentApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type SubscriptionIDDocumentApiServicer interface { 
	RemoveSubscription(string) (interface{}, error)
	UpdateSubscription(string, []PatchItem) (interface{}, error)
}


// SubscriptionsCollectionApiServicer defines the api actions for the SubscriptionsCollectionApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type SubscriptionsCollectionApiServicer interface { 
	CreateSubscription(SubscriptionData) (interface{}, error)
}
