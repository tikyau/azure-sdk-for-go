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

package disk

import original "github.com/Azure/azure-sdk-for-go/service/compute/management/2017-03-30/disk"

type DisksClient = original.DisksClient
type SnapshotsClient = original.SnapshotsClient
type AccessLevel = original.AccessLevel

const (
	None	AccessLevel	= original.None
	Read	AccessLevel	= original.Read
)

type CreateOption = original.CreateOption

const (
	Attach		CreateOption	= original.Attach
	Copy		CreateOption	= original.Copy
	Empty		CreateOption	= original.Empty
	FromImage	CreateOption	= original.FromImage
	Import		CreateOption	= original.Import
)

type OperatingSystemTypes = original.OperatingSystemTypes

const (
	Linux	OperatingSystemTypes	= original.Linux
	Windows	OperatingSystemTypes	= original.Windows
)

type StorageAccountTypes = original.StorageAccountTypes

const (
	PremiumLRS	StorageAccountTypes	= original.PremiumLRS
	StandardLRS	StorageAccountTypes	= original.StandardLRS
)

type AccessURI = original.AccessURI
type AccessURIOutput = original.AccessURIOutput
type AccessURIRaw = original.AccessURIRaw
type APIError = original.APIError
type APIErrorBase = original.APIErrorBase
type CreationData = original.CreationData
type EncryptionSettings = original.EncryptionSettings
type GrantAccessData = original.GrantAccessData
type ImageDiskReference = original.ImageDiskReference
type InnerError = original.InnerError
type KeyVaultAndKeyReference = original.KeyVaultAndKeyReference
type KeyVaultAndSecretReference = original.KeyVaultAndSecretReference
type ListType = original.ListType
type Model = original.Model
type OperationStatusResponse = original.OperationStatusResponse
type Properties = original.Properties
type Resource = original.Resource
type ResourceUpdate = original.ResourceUpdate
type Sku = original.Sku
type Snapshot = original.Snapshot
type SnapshotList = original.SnapshotList
type SnapshotUpdate = original.SnapshotUpdate
type SourceVault = original.SourceVault
type UpdateProperties = original.UpdateProperties
type UpdateType = original.UpdateType

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient

func NewSnapshotsClient(subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClient(subscriptionID)
}
func NewSnapshotsClientWithBaseURI(baseURI string, subscriptionID string) SnapshotsClient {
	return original.NewSnapshotsClientWithBaseURI(baseURI, subscriptionID)
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
func NewDisksClient(subscriptionID string) DisksClient {
	return original.NewDisksClient(subscriptionID)
}
func NewDisksClientWithBaseURI(baseURI string, subscriptionID string) DisksClient {
	return original.NewDisksClientWithBaseURI(baseURI, subscriptionID)
}
