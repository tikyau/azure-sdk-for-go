package scheduler

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// JobCollectionsClient is the client for the JobCollections methods of the
// Scheduler service.
type JobCollectionsClient struct {
	ManagementClient
}

// NewJobCollectionsClient creates an instance of the JobCollectionsClient
// client.
func NewJobCollectionsClient(subscriptionID string) JobCollectionsClient {
	return NewJobCollectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewJobCollectionsClientWithBaseURI creates an instance of the
// JobCollectionsClient client.
func NewJobCollectionsClientWithBaseURI(baseURI string, subscriptionID string) JobCollectionsClient {
	return JobCollectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate provisions a new job collection or updates an existing job
// collection.
//
// resourceGroupName is the resource group name. jobCollectionName is the job
// collection name. jobCollection is the job collection definition.
func (client JobCollectionsClient) CreateOrUpdate(resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition) (result JobCollectionDefinition, err error) {
	if err := validation.Validate([]validation.Validation{
		{jobCollection,
			[]validation.Constraint{{"ID", validation.ReadOnly, true, nil},
				{"Type", validation.ReadOnly, true, nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "scheduler.JobCollectionsClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, jobCollectionName, jobCollection)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client JobCollectionsClient) CreateOrUpdatePreparer(resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", pathParameters),
		autorest.WithJSON(jobCollection),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) CreateOrUpdateResponder(resp *http.Response) (result JobCollectionDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a job collection. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The
// channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the resource group name. jobCollectionName is the job
// collection name.
func (client JobCollectionsClient) Delete(resourceGroupName string, jobCollectionName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, jobCollectionName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client JobCollectionsClient) DeletePreparer(resourceGroupName string, jobCollectionName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Disable disables all of the jobs in the job collection. This method may
// poll for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the resource group name. jobCollectionName is the job
// collection name.
func (client JobCollectionsClient) Disable(resourceGroupName string, jobCollectionName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.DisablePreparer(resourceGroupName, jobCollectionName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Disable", nil, "Failure preparing request")
	}

	resp, err := client.DisableSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Disable", resp, "Failure sending request")
	}

	result, err = client.DisableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Disable", resp, "Failure responding to request")
	}

	return
}

// DisablePreparer prepares the Disable request.
func (client JobCollectionsClient) DisablePreparer(resourceGroupName string, jobCollectionName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/disable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DisableSender sends the Disable request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) DisableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DisableResponder handles the response to the Disable request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) DisableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Enable enables all of the jobs in the job collection. This method may poll
// for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the resource group name. jobCollectionName is the job
// collection name.
func (client JobCollectionsClient) Enable(resourceGroupName string, jobCollectionName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.EnablePreparer(resourceGroupName, jobCollectionName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Enable", nil, "Failure preparing request")
	}

	resp, err := client.EnableSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Enable", resp, "Failure sending request")
	}

	result, err = client.EnableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Enable", resp, "Failure responding to request")
	}

	return
}

// EnablePreparer prepares the Enable request.
func (client JobCollectionsClient) EnablePreparer(resourceGroupName string, jobCollectionName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}/enable", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// EnableSender sends the Enable request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) EnableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// EnableResponder handles the response to the Enable request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) EnableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a job collection.
//
// resourceGroupName is the resource group name. jobCollectionName is the job
// collection name.
func (client JobCollectionsClient) Get(resourceGroupName string, jobCollectionName string) (result JobCollectionDefinition, err error) {
	req, err := client.GetPreparer(resourceGroupName, jobCollectionName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client JobCollectionsClient) GetPreparer(resourceGroupName string, jobCollectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) GetResponder(resp *http.Response) (result JobCollectionDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets all job collections under specified resource group.
//
// resourceGroupName is the resource group name.
func (client JobCollectionsClient) ListByResourceGroup(resourceGroupName string) (result JobCollectionListResult, err error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListByResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListByResourceGroup", resp, "Failure sending request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client JobCollectionsClient) ListByResourceGroupPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) ListByResourceGroupResponder(resp *http.Response) (result JobCollectionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client JobCollectionsClient) ListByResourceGroupNextResults(lastResults JobCollectionListResult) (result JobCollectionListResult, err error) {
	req, err := lastResults.JobCollectionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListByResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListByResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListByResourceGroup", resp, "Failure responding to next results request")
	}

	return
}

// ListBySubscription gets all job collections under specified subscription.
func (client JobCollectionsClient) ListBySubscription() (result JobCollectionListResult, err error) {
	req, err := client.ListBySubscriptionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListBySubscription", nil, "Failure preparing request")
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListBySubscription", resp, "Failure sending request")
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client JobCollectionsClient) ListBySubscriptionPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Scheduler/jobCollections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) ListBySubscriptionResponder(resp *http.Response) (result JobCollectionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscriptionNextResults retrieves the next set of results, if any.
func (client JobCollectionsClient) ListBySubscriptionNextResults(lastResults JobCollectionListResult) (result JobCollectionListResult, err error) {
	req, err := lastResults.JobCollectionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListBySubscription", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListBySubscription", resp, "Failure sending next results request")
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "ListBySubscription", resp, "Failure responding to next results request")
	}

	return
}

// Patch patches an existing job collection.
//
// resourceGroupName is the resource group name. jobCollectionName is the job
// collection name. jobCollection is the job collection definition.
func (client JobCollectionsClient) Patch(resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition) (result JobCollectionDefinition, err error) {
	if err := validation.Validate([]validation.Validation{
		{jobCollection,
			[]validation.Constraint{{"ID", validation.ReadOnly, true, nil},
				{"Type", validation.ReadOnly, true, nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "scheduler.JobCollectionsClient", "Patch")
	}

	req, err := client.PatchPreparer(resourceGroupName, jobCollectionName, jobCollection)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Patch", nil, "Failure preparing request")
	}

	resp, err := client.PatchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Patch", resp, "Failure sending request")
	}

	result, err = client.PatchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scheduler.JobCollectionsClient", "Patch", resp, "Failure responding to request")
	}

	return
}

// PatchPreparer prepares the Patch request.
func (client JobCollectionsClient) PatchPreparer(resourceGroupName string, jobCollectionName string, jobCollection JobCollectionDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"jobCollectionName": autorest.Encode("path", jobCollectionName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Scheduler/jobCollections/{jobCollectionName}", pathParameters),
		autorest.WithJSON(jobCollection),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client JobCollectionsClient) PatchSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client JobCollectionsClient) PatchResponder(resp *http.Response) (result JobCollectionDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
