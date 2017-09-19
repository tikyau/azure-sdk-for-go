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

package connectionpolicies

import original "github.com/Azure/azure-sdk-for-go/service/sql/management/2014-04-01/connectionPolicies"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type ServerConnectionPoliciesClient = original.ServerConnectionPoliciesClient
type ServerConnectionType = original.ServerConnectionType

const (
	Default		ServerConnectionType	= original.Default
	Proxy		ServerConnectionType	= original.Proxy
	Redirect	ServerConnectionType	= original.Redirect
)

type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ServerConnectionPolicy = original.ServerConnectionPolicy
type ServerConnectionPolicyProperties = original.ServerConnectionPolicyProperties

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
func NewServerConnectionPoliciesClient(subscriptionID string) ServerConnectionPoliciesClient {
	return original.NewServerConnectionPoliciesClient(subscriptionID)
}
func NewServerConnectionPoliciesClientWithBaseURI(baseURI string, subscriptionID string) ServerConnectionPoliciesClient {
	return original.NewServerConnectionPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
