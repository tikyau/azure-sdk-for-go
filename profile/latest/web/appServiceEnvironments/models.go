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

package appserviceenvironments

import original "github.com/Azure/azure-sdk-for-go/service/web/management/2016-09-01/appServiceEnvironments"

type GroupClient = original.GroupClient
type AccessControlEntryAction = original.AccessControlEntryAction

const (
	Deny	AccessControlEntryAction	= original.Deny
	Permit	AccessControlEntryAction	= original.Permit
)

type AutoHealActionType = original.AutoHealActionType

const (
	CustomAction	AutoHealActionType	= original.CustomAction
	LogEvent	AutoHealActionType	= original.LogEvent
	Recycle		AutoHealActionType	= original.Recycle
)

type ComputeModeOptions = original.ComputeModeOptions

const (
	Dedicated	ComputeModeOptions	= original.Dedicated
	Dynamic		ComputeModeOptions	= original.Dynamic
	Shared		ComputeModeOptions	= original.Shared
)

type ConnectionStringType = original.ConnectionStringType

const (
	APIHub		ConnectionStringType	= original.APIHub
	Custom		ConnectionStringType	= original.Custom
	DocDb		ConnectionStringType	= original.DocDb
	EventHub	ConnectionStringType	= original.EventHub
	MySQL		ConnectionStringType	= original.MySQL
	NotificationHub	ConnectionStringType	= original.NotificationHub
	PostgreSQL	ConnectionStringType	= original.PostgreSQL
	RedisCache	ConnectionStringType	= original.RedisCache
	ServiceBus	ConnectionStringType	= original.ServiceBus
	SQLAzure	ConnectionStringType	= original.SQLAzure
	SQLServer	ConnectionStringType	= original.SQLServer
)

type HostingEnvironmentStatus = original.HostingEnvironmentStatus

const (
	Deleting	HostingEnvironmentStatus	= original.Deleting
	Preparing	HostingEnvironmentStatus	= original.Preparing
	Ready		HostingEnvironmentStatus	= original.Ready
	Scaling		HostingEnvironmentStatus	= original.Scaling
)

type HostType = original.HostType

const (
	Repository	HostType	= original.Repository
	Standard	HostType	= original.Standard
)

type InternalLoadBalancingMode = original.InternalLoadBalancingMode

const (
	None		InternalLoadBalancingMode	= original.None
	Publishing	InternalLoadBalancingMode	= original.Publishing
	Web		InternalLoadBalancingMode	= original.Web
)

type ManagedPipelineMode = original.ManagedPipelineMode

const (
	Classic		ManagedPipelineMode	= original.Classic
	Integrated	ManagedPipelineMode	= original.Integrated
)

type OperationStatus = original.OperationStatus

const (
	Created		OperationStatus	= original.Created
	Failed		OperationStatus	= original.Failed
	InProgress	OperationStatus	= original.InProgress
	Succeeded	OperationStatus	= original.Succeeded
	TimedOut	OperationStatus	= original.TimedOut
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCanceled	ProvisioningState	= original.ProvisioningStateCanceled
	ProvisioningStateDeleting	ProvisioningState	= original.ProvisioningStateDeleting
	ProvisioningStateFailed		ProvisioningState	= original.ProvisioningStateFailed
	ProvisioningStateInProgress	ProvisioningState	= original.ProvisioningStateInProgress
	ProvisioningStateSucceeded	ProvisioningState	= original.ProvisioningStateSucceeded
)

type ScmType = original.ScmType

const (
	ScmTypeBitbucketGit	ScmType	= original.ScmTypeBitbucketGit
	ScmTypeBitbucketHg	ScmType	= original.ScmTypeBitbucketHg
	ScmTypeCodePlexGit	ScmType	= original.ScmTypeCodePlexGit
	ScmTypeCodePlexHg	ScmType	= original.ScmTypeCodePlexHg
	ScmTypeDropbox		ScmType	= original.ScmTypeDropbox
	ScmTypeExternalGit	ScmType	= original.ScmTypeExternalGit
	ScmTypeExternalHg	ScmType	= original.ScmTypeExternalHg
	ScmTypeGitHub		ScmType	= original.ScmTypeGitHub
	ScmTypeLocalGit		ScmType	= original.ScmTypeLocalGit
	ScmTypeNone		ScmType	= original.ScmTypeNone
	ScmTypeOneDrive		ScmType	= original.ScmTypeOneDrive
	ScmTypeTfs		ScmType	= original.ScmTypeTfs
	ScmTypeVSO		ScmType	= original.ScmTypeVSO
)

type SiteAvailabilityState = original.SiteAvailabilityState

const (
	DisasterRecoveryMode	SiteAvailabilityState	= original.DisasterRecoveryMode
	Limited			SiteAvailabilityState	= original.Limited
	Normal			SiteAvailabilityState	= original.Normal
)

type SiteLoadBalancing = original.SiteLoadBalancing

const (
	LeastRequests		SiteLoadBalancing	= original.LeastRequests
	LeastResponseTime	SiteLoadBalancing	= original.LeastResponseTime
	RequestHash		SiteLoadBalancing	= original.RequestHash
	WeightedRoundRobin	SiteLoadBalancing	= original.WeightedRoundRobin
	WeightedTotalTraffic	SiteLoadBalancing	= original.WeightedTotalTraffic
)

type SslState = original.SslState

const (
	Disabled	SslState	= original.Disabled
	IPBasedEnabled	SslState	= original.IPBasedEnabled
	SniEnabled	SslState	= original.SniEnabled
)

type StatusOptions = original.StatusOptions

const (
	StatusOptionsPending	StatusOptions	= original.StatusOptionsPending
	StatusOptionsReady	StatusOptions	= original.StatusOptionsReady
)

type UsageState = original.UsageState

const (
	UsageStateExceeded	UsageState	= original.UsageStateExceeded
	UsageStateNormal	UsageState	= original.UsageStateNormal
)

type WorkerSizeOptions = original.WorkerSizeOptions

const (
	Default	WorkerSizeOptions	= original.Default
	Large	WorkerSizeOptions	= original.Large
	Medium	WorkerSizeOptions	= original.Medium
	Small	WorkerSizeOptions	= original.Small
)

type AddressResponse = original.AddressResponse
type APIDefinitionInfo = original.APIDefinitionInfo
type AppServiceEnvironment = original.AppServiceEnvironment
type AppServicePlan = original.AppServicePlan
type AppServicePlanProperties = original.AppServicePlanProperties
type AppServicePlanCollection = original.AppServicePlanCollection
type AutoHealActions = original.AutoHealActions
type AutoHealCustomAction = original.AutoHealCustomAction
type AutoHealRules = original.AutoHealRules
type AutoHealTriggers = original.AutoHealTriggers
type Capability = original.Capability
type CloningInfo = original.CloningInfo
type Collection = original.Collection
type ConnStringInfo = original.ConnStringInfo
type CorsSettings = original.CorsSettings
type CsmUsageQuota = original.CsmUsageQuota
type CsmUsageQuotaCollection = original.CsmUsageQuotaCollection
type ErrorEntity = original.ErrorEntity
type Experiments = original.Experiments
type HandlerMapping = original.HandlerMapping
type HostingEnvironmentDiagnostics = original.HostingEnvironmentDiagnostics
type HostingEnvironmentProfile = original.HostingEnvironmentProfile
type HostNameSslState = original.HostNameSslState
type IPSecurityRestriction = original.IPSecurityRestriction
type ListHostingEnvironmentDiagnostics = original.ListHostingEnvironmentDiagnostics
type ListOperation = original.ListOperation
type LocalizableString = original.LocalizableString
type MetricAvailabilily = original.MetricAvailabilily
type MetricDefinition = original.MetricDefinition
type MetricDefinitionProperties = original.MetricDefinitionProperties
type NameValuePair = original.NameValuePair
type NetworkAccessControlEntry = original.NetworkAccessControlEntry
type Operation = original.Operation
type PushSettings = original.PushSettings
type RampUpRule = original.RampUpRule
type RequestsBasedTrigger = original.RequestsBasedTrigger
type Resource = original.Resource
type ResourceMetric = original.ResourceMetric
type ResourceMetricAvailability = original.ResourceMetricAvailability
type ResourceMetricCollection = original.ResourceMetricCollection
type ResourceMetricDefinition = original.ResourceMetricDefinition
type ResourceMetricDefinitionProperties = original.ResourceMetricDefinitionProperties
type ResourceMetricDefinitionCollection = original.ResourceMetricDefinitionCollection
type ResourceMetricName = original.ResourceMetricName
type ResourceMetricProperty = original.ResourceMetricProperty
type ResourceMetricValue = original.ResourceMetricValue
type ResourceType = original.ResourceType
type Site = original.Site
type SiteProperties = original.SiteProperties
type SiteConfig = original.SiteConfig
type SiteLimits = original.SiteLimits
type SiteMachineKey = original.SiteMachineKey
type SkuCapacity = original.SkuCapacity
type SkuDescription = original.SkuDescription
type SkuInfo = original.SkuInfo
type SkuInfoCollection = original.SkuInfoCollection
type SlotSwapStatus = original.SlotSwapStatus
type SlowRequestsBasedTrigger = original.SlowRequestsBasedTrigger
type StampCapacity = original.StampCapacity
type StampCapacityCollection = original.StampCapacityCollection
type StatusCodesBasedTrigger = original.StatusCodesBasedTrigger
type Usage = original.Usage
type UsageProperties = original.UsageProperties
type UsageCollection = original.UsageCollection
type VirtualApplication = original.VirtualApplication
type VirtualDirectory = original.VirtualDirectory
type VirtualIPMapping = original.VirtualIPMapping
type VirtualNetworkProfile = original.VirtualNetworkProfile
type WebAppCollection = original.WebAppCollection
type WorkerPool = original.WorkerPool
type WorkerPoolCollection = original.WorkerPoolCollection
type WorkerPoolResource = original.WorkerPoolResource

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient

func NewGroupClient(subscriptionID string) GroupClient {
	return original.NewGroupClient(subscriptionID)
}
func NewGroupClientWithBaseURI(baseURI string, subscriptionID string) GroupClient {
	return original.NewGroupClientWithBaseURI(baseURI, subscriptionID)
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
