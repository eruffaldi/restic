package dns

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// RecordSetsClient is the client for managing DNS zones and record.
type RecordSetsClient struct {
	ManagementClient
}

// NewRecordSetsClient creates an instance of the RecordSetsClient client.
func NewRecordSetsClient(subscriptionID string) RecordSetsClient {
	return NewRecordSetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecordSetsClientWithBaseURI creates an instance of the RecordSetsClient client.
func NewRecordSetsClientWithBaseURI(baseURI string, subscriptionID string) RecordSetsClient {
	return RecordSetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a RecordSet within a DNS zone.
//
// resourceGroupName is the name of the resource group. zoneName is the name of the zone without a terminating dot.
// recordType is the type of DNS record. relativeRecordSetName is the name of the RecordSet, relative to the name of
// the zone. parameters is parameters supplied to the CreateOrUpdate operation. ifMatch is the etag of RecordSet.
// ifNoneMatch is defines the If-None-Match condition. Set to '*' to force Create-If-Not-Exist. Other values will be
// ignored.
func (client RecordSetsClient) CreateOrUpdate(resourceGroupName string, zoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, ifMatch string, ifNoneMatch string) (result RecordSet, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, zoneName, recordType, relativeRecordSetName, parameters, ifMatch, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RecordSetsClient) CreateOrUpdatePreparer(resourceGroupName string, zoneName string, recordType RecordType, relativeRecordSetName string, parameters RecordSet, ifMatch string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"zoneName":              autorest.Encode("path", zoneName),
	}

	const APIVersion = "2015-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) CreateOrUpdateResponder(resp *http.Response) (result RecordSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete removes a RecordSet from a DNS zone.
//
// resourceGroupName is the name of the resource group. zoneName is the name of the zone without a terminating dot.
// recordType is the type of DNS record. relativeRecordSetName is the name of the RecordSet, relative to the name of
// the zone. ifMatch is defines the If-Match condition. The delete operation will be performed only if the ETag of the
// zone on the server matches this value.
func (client RecordSetsClient) Delete(resourceGroupName string, zoneName string, recordType RecordType, relativeRecordSetName string, ifMatch string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, zoneName, recordType, relativeRecordSetName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RecordSetsClient) DeletePreparer(resourceGroupName string, zoneName string, recordType RecordType, relativeRecordSetName string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"zoneName":              autorest.Encode("path", zoneName),
	}

	const APIVersion = "2015-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a RecordSet.
//
// resourceGroupName is the name of the resource group. zoneName is the name of the zone without a terminating dot.
// recordType is the type of DNS record. relativeRecordSetName is the name of the RecordSet, relative to the name of
// the zone.
func (client RecordSetsClient) Get(resourceGroupName string, zoneName string, recordType RecordType, relativeRecordSetName string) (result RecordSet, err error) {
	req, err := client.GetPreparer(resourceGroupName, zoneName, recordType, relativeRecordSetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecordSetsClient) GetPreparer(resourceGroupName string, zoneName string, recordType RecordType, relativeRecordSetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":            autorest.Encode("path", recordType),
		"relativeRecordSetName": relativeRecordSetName,
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"zoneName":              autorest.Encode("path", zoneName),
	}

	const APIVersion = "2015-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}/{recordType}/{relativeRecordSetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) GetResponder(resp *http.Response) (result RecordSet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the RecordSets of a specified type in a DNS zone.
//
// resourceGroupName is the name of the resource group that contains the zone. zoneName is the name of the zone from
// which to enumerate RecordsSets. recordType is the type of record sets to enumerate. top is query parameters. If null
// is passed returns the default number of zones. filter is the filter to apply on the operation.
func (client RecordSetsClient) List(resourceGroupName string, zoneName string, recordType RecordType, top string, filter string) (result RecordSetListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName, zoneName, recordType, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RecordSetsClient) ListPreparer(resourceGroupName string, zoneName string, recordType RecordType, top string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"recordType":        autorest.Encode("path", recordType),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"zoneName":          autorest.Encode("path", zoneName),
	}

	const APIVersion = "2015-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(top) > 0 {
		queryParameters["$top"] = autorest.Encode("query", top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}/{recordType}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) ListResponder(resp *http.Response) (result RecordSetListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client RecordSetsClient) ListNextResults(lastResults RecordSetListResult) (result RecordSetListResult, err error) {
	req, err := lastResults.RecordSetListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client RecordSetsClient) ListComplete(resourceGroupName string, zoneName string, recordType RecordType, top string, filter string, cancel <-chan struct{}) (<-chan RecordSet, <-chan error) {
	resultChan := make(chan RecordSet)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceGroupName, zoneName, recordType, top, filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ListAll lists all RecordSets in a DNS zone.
//
// resourceGroupName is the name of the resource group that contains the zone. zoneName is the name of the zone from
// which to enumerate RecordSets. top is query parameters. If null is passed returns the default number of zones.
// filter is the filter to apply on the operation.
func (client RecordSetsClient) ListAll(resourceGroupName string, zoneName string, top string, filter string) (result RecordSetListResult, err error) {
	req, err := client.ListAllPreparer(resourceGroupName, zoneName, top, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListAll", resp, "Failure sending request")
		return
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client RecordSetsClient) ListAllPreparer(resourceGroupName string, zoneName string, top string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"zoneName":          autorest.Encode("path", zoneName),
	}

	const APIVersion = "2015-05-04-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(top) > 0 {
		queryParameters["$top"] = autorest.Encode("query", top)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnszones/{zoneName}/recordsets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client RecordSetsClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client RecordSetsClient) ListAllResponder(resp *http.Response) (result RecordSetListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllNextResults retrieves the next set of results, if any.
func (client RecordSetsClient) ListAllNextResults(lastResults RecordSetListResult) (result RecordSetListResult, err error) {
	req, err := lastResults.RecordSetListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListAll", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListAll", resp, "Failure sending next results request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dns.RecordSetsClient", "ListAll", resp, "Failure responding to next results request")
	}

	return
}

// ListAllComplete gets all elements from the list without paging.
func (client RecordSetsClient) ListAllComplete(resourceGroupName string, zoneName string, top string, filter string, cancel <-chan struct{}) (<-chan RecordSet, <-chan error) {
	resultChan := make(chan RecordSet)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListAll(resourceGroupName, zoneName, top, filter)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListAllNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
