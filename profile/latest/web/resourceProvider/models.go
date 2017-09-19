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

package resourceprovider

import original "github.com/Azure/azure-sdk-for-go/service/web/management/2016-03-01/resourceProvider"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type AppServicePlanRestrictions = original.AppServicePlanRestrictions

const (
	Basic		AppServicePlanRestrictions	= original.Basic
	Free		AppServicePlanRestrictions	= original.Free
	None		AppServicePlanRestrictions	= original.None
	Premium		AppServicePlanRestrictions	= original.Premium
	Shared		AppServicePlanRestrictions	= original.Shared
	Standard	AppServicePlanRestrictions	= original.Standard
)

type CheckNameResourceTypes = original.CheckNameResourceTypes

const (
	HostingEnvironment	CheckNameResourceTypes	= original.HostingEnvironment
	Site			CheckNameResourceTypes	= original.Site
	Slot			CheckNameResourceTypes	= original.Slot
)

type InAvailabilityReasonType = original.InAvailabilityReasonType

const (
	AlreadyExists	InAvailabilityReasonType	= original.AlreadyExists
	Invalid		InAvailabilityReasonType	= original.Invalid
)

type SkuName = original.SkuName

const (
	SkuNameBasic	SkuName	= original.SkuNameBasic
	SkuNameDynamic	SkuName	= original.SkuNameDynamic
	SkuNameFree	SkuName	= original.SkuNameFree
	SkuNameIsolated	SkuName	= original.SkuNameIsolated
	SkuNamePremium	SkuName	= original.SkuNamePremium
	SkuNameShared	SkuName	= original.SkuNameShared
	SkuNameStandard	SkuName	= original.SkuNameStandard
)

type ValidateResourceTypes = original.ValidateResourceTypes

const (
	ValidateResourceTypesServerFarm	ValidateResourceTypes	= original.ValidateResourceTypesServerFarm
	ValidateResourceTypesSite	ValidateResourceTypes	= original.ValidateResourceTypesSite
)

type Capability = original.Capability
type CsmMoveResourceEnvelope = original.CsmMoveResourceEnvelope
type GeoRegion = original.GeoRegion
type GeoRegionProperties = original.GeoRegionProperties
type GeoRegionCollection = original.GeoRegionCollection
type GlobalCsmSkuDescription = original.GlobalCsmSkuDescription
type PremierAddOnOffer = original.PremierAddOnOffer
type PremierAddOnOfferProperties = original.PremierAddOnOfferProperties
type PremierAddOnOfferCollection = original.PremierAddOnOfferCollection
type Resource = original.Resource
type ResourceNameAvailability = original.ResourceNameAvailability
type ResourceNameAvailabilityRequest = original.ResourceNameAvailabilityRequest
type SkuCapacity = original.SkuCapacity
type SkuInfos = original.SkuInfos
type SourceControl = original.SourceControl
type SourceControlProperties = original.SourceControlProperties
type SourceControlCollection = original.SourceControlCollection
type User = original.User
type UserProperties = original.UserProperties
type ValidateProperties = original.ValidateProperties
type ValidateRequest = original.ValidateRequest
type ValidateResponse = original.ValidateResponse
type ValidateResponseError = original.ValidateResponseError

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