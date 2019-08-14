package storagesync

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SyncGroupsClient is the microsoft Storage Sync Service API
type SyncGroupsClient struct {
	BaseClient
}

// NewSyncGroupsClient creates an instance of the SyncGroupsClient client.
func NewSyncGroupsClient(subscriptionID string) SyncGroupsClient {
	return NewSyncGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSyncGroupsClientWithBaseURI creates an instance of the SyncGroupsClient client.
func NewSyncGroupsClientWithBaseURI(baseURI string, subscriptionID string) SyncGroupsClient {
	return SyncGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create create a new SyncGroup.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
// parameters - sync Group Body
func (client SyncGroupsClient) Create(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, parameters SyncGroupCreateParameters) (result SyncGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SyncGroupsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.SyncGroupsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.SyncGroupsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagesync.SyncGroupsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.SyncGroupsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SyncGroupsClient) CreatePreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string, parameters SyncGroupCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SyncGroupsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SyncGroupsClient) CreateResponder(resp *http.Response) (result SyncGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a given SyncGroup.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
func (client SyncGroupsClient) Delete(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SyncGroupsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.SyncGroupsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.SyncGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "storagesync.SyncGroupsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.SyncGroupsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SyncGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SyncGroupsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SyncGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a given SyncGroup.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
// syncGroupName - name of Sync Group resource.
func (client SyncGroupsClient) Get(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string) (result SyncGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SyncGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.SyncGroupsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, storageSyncServiceName, syncGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.SyncGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagesync.SyncGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.SyncGroupsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SyncGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string, syncGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":          autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups/{syncGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SyncGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SyncGroupsClient) GetResponder(resp *http.Response) (result SyncGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByStorageSyncService get a SyncGroup List.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// storageSyncServiceName - name of Storage Sync Service resource.
func (client SyncGroupsClient) ListByStorageSyncService(ctx context.Context, resourceGroupName string, storageSyncServiceName string) (result SyncGroupArray, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SyncGroupsClient.ListByStorageSyncService")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagesync.SyncGroupsClient", "ListByStorageSyncService", err.Error())
	}

	req, err := client.ListByStorageSyncServicePreparer(ctx, resourceGroupName, storageSyncServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.SyncGroupsClient", "ListByStorageSyncService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByStorageSyncServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagesync.SyncGroupsClient", "ListByStorageSyncService", resp, "Failure sending request")
		return
	}

	result, err = client.ListByStorageSyncServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagesync.SyncGroupsClient", "ListByStorageSyncService", resp, "Failure responding to request")
	}

	return
}

// ListByStorageSyncServicePreparer prepares the ListByStorageSyncService request.
func (client SyncGroupsClient) ListByStorageSyncServicePreparer(ctx context.Context, resourceGroupName string, storageSyncServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"storageSyncServiceName": autorest.Encode("path", storageSyncServiceName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StorageSync/storageSyncServices/{storageSyncServiceName}/syncGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByStorageSyncServiceSender sends the ListByStorageSyncService request. The method will close the
// http.Response Body if it receives an error.
func (client SyncGroupsClient) ListByStorageSyncServiceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByStorageSyncServiceResponder handles the response to the ListByStorageSyncService request. The method always
// closes the http.Response Body.
func (client SyncGroupsClient) ListByStorageSyncServiceResponder(resp *http.Response) (result SyncGroupArray, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
