// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package apimanagement

import original "github.com/Azure/azure-sdk-for-go/service/apimanagement/management/2016-10-10/apimanagement"

type PolicySnippetsClient = original.PolicySnippetsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type PolicyScopeContract = original.PolicyScopeContract

const (
	All		PolicyScopeContract	= original.All
	API		PolicyScopeContract	= original.API
	Operation	PolicyScopeContract	= original.Operation
	Product		PolicyScopeContract	= original.Product
	Tenant		PolicyScopeContract	= original.Tenant
)

type ErrorBodyContract = original.ErrorBodyContract
type ErrorFieldContract = original.ErrorFieldContract
type PolicySnippetContract = original.PolicySnippetContract
type PolicySnippetsCollection = original.PolicySnippetsCollection
type RegionContract = original.RegionContract
type RegionListResult = original.RegionListResult
type RegionsClient = original.RegionsClient

func NewPolicySnippetsClient(subscriptionID string) PolicySnippetsClient {
	return original.NewPolicySnippetsClient(subscriptionID)
}
func NewPolicySnippetsClientWithBaseURI(baseURI string, subscriptionID string) PolicySnippetsClient {
	return original.NewPolicySnippetsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent()
}
func Version() string {
	return original.Version()
}
func NewRegionsClient(subscriptionID string) RegionsClient {
	return original.NewRegionsClient(subscriptionID)
}
func NewRegionsClientWithBaseURI(baseURI string, subscriptionID string) RegionsClient {
	return original.NewRegionsClientWithBaseURI(baseURI, subscriptionID)
}
