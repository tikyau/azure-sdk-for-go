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

package backup

import original "github.com/Azure/azure-sdk-for-go/service/recoveryservices/management/2016-12-01/backup"

type VaultConfigsClient = original.VaultConfigsClient
type EnhancedSecurityState = original.EnhancedSecurityState

const (
	Disabled	EnhancedSecurityState	= original.Disabled
	Enabled		EnhancedSecurityState	= original.Enabled
	Invalid		EnhancedSecurityState	= original.Invalid
)

type SkuName = original.SkuName

const (
	RS0		SkuName	= original.RS0
	Standard	SkuName	= original.Standard
)

type StorageModelType = original.StorageModelType

const (
	StorageModelTypeGeoRedundant		StorageModelType	= original.StorageModelTypeGeoRedundant
	StorageModelTypeInvalid			StorageModelType	= original.StorageModelTypeInvalid
	StorageModelTypeLocallyRedundant	StorageModelType	= original.StorageModelTypeLocallyRedundant
)

type StorageType = original.StorageType

const (
	StorageTypeGeoRedundant		StorageType	= original.StorageTypeGeoRedundant
	StorageTypeInvalid		StorageType	= original.StorageTypeInvalid
	StorageTypeLocallyRedundant	StorageType	= original.StorageTypeLocallyRedundant
)

type StorageTypeState = original.StorageTypeState

const (
	StorageTypeStateInvalid		StorageTypeState	= original.StorageTypeStateInvalid
	StorageTypeStateLocked		StorageTypeState	= original.StorageTypeStateLocked
	StorageTypeStateUnlocked	StorageTypeState	= original.StorageTypeStateUnlocked
)

type TriggerType = original.TriggerType

const (
	ForcedUpgrade	TriggerType	= original.ForcedUpgrade
	UserTriggered	TriggerType	= original.UserTriggered
)

type VaultUpgradeState = original.VaultUpgradeState

const (
	Failed		VaultUpgradeState	= original.Failed
	InProgress	VaultUpgradeState	= original.InProgress
	Unknown		VaultUpgradeState	= original.Unknown
	Upgraded	VaultUpgradeState	= original.Upgraded
)

type Resource = original.Resource
type Sku = original.Sku
type StorageConfig = original.StorageConfig
type StorageConfigProperties = original.StorageConfigProperties
type TrackedResource = original.TrackedResource
type UpgradeDetails = original.UpgradeDetails
type Vault = original.Vault
type VaultConfig = original.VaultConfig
type VaultConfigProperties = original.VaultConfigProperties
type VaultExtendedInfo = original.VaultExtendedInfo
type VaultExtendedInfoResource = original.VaultExtendedInfoResource
type VaultProperties = original.VaultProperties
type StorageConfigsClient = original.StorageConfigsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient

func NewStorageConfigsClient(subscriptionID string) StorageConfigsClient {
	return original.NewStorageConfigsClient(subscriptionID)
}
func NewStorageConfigsClientWithBaseURI(baseURI string, subscriptionID string) StorageConfigsClient {
	return original.NewStorageConfigsClientWithBaseURI(baseURI, subscriptionID)
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
func NewVaultConfigsClient(subscriptionID string) VaultConfigsClient {
	return original.NewVaultConfigsClient(subscriptionID)
}
func NewVaultConfigsClientWithBaseURI(baseURI string, subscriptionID string) VaultConfigsClient {
	return original.NewVaultConfigsClientWithBaseURI(baseURI, subscriptionID)
}
