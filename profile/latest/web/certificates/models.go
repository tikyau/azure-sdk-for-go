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

package certificates

import original "github.com/Azure/azure-sdk-for-go/service/web/management/2016-03-01/certificates"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type KeyVaultSecretStatus = original.KeyVaultSecretStatus

const (
	AzureServiceUnauthorizedToAccessKeyVault	KeyVaultSecretStatus	= original.AzureServiceUnauthorizedToAccessKeyVault
	CertificateOrderFailed				KeyVaultSecretStatus	= original.CertificateOrderFailed
	ExternalPrivateKey				KeyVaultSecretStatus	= original.ExternalPrivateKey
	Initialized					KeyVaultSecretStatus	= original.Initialized
	KeyVaultDoesNotExist				KeyVaultSecretStatus	= original.KeyVaultDoesNotExist
	KeyVaultSecretDoesNotExist			KeyVaultSecretStatus	= original.KeyVaultSecretDoesNotExist
	OperationNotPermittedOnKeyVault			KeyVaultSecretStatus	= original.OperationNotPermittedOnKeyVault
	Succeeded					KeyVaultSecretStatus	= original.Succeeded
	Unknown						KeyVaultSecretStatus	= original.Unknown
	UnknownError					KeyVaultSecretStatus	= original.UnknownError
	WaitingOnCertificateOrder			KeyVaultSecretStatus	= original.WaitingOnCertificateOrder
)

type Certificate = original.Certificate
type CertificateProperties = original.CertificateProperties
type Collection = original.Collection
type HostingEnvironmentProfile = original.HostingEnvironmentProfile
type Resource = original.Resource
type GroupClient = original.GroupClient

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
func NewGroupClient(subscriptionID string) GroupClient {
	return original.NewGroupClient(subscriptionID)
}
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return original.NewGroupClientWithBaseURI(baseURI, subscriptionID)
}
