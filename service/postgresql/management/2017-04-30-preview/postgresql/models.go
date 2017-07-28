package postgresql

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// OperationOrigin enumerates the values for operation origin.
type OperationOrigin string

const (
	// NotSpecified specifies the not specified state for operation origin.
	NotSpecified OperationOrigin = "NotSpecified"
	// System specifies the system state for operation origin.
	System OperationOrigin = "system"
	// User specifies the user state for operation origin.
	User OperationOrigin = "user"
)

// ServerState enumerates the values for server state.
type ServerState string

const (
	// Disabled specifies the disabled state for server state.
	Disabled ServerState = "Disabled"
	// Dropping specifies the dropping state for server state.
	Dropping ServerState = "Dropping"
	// Ready specifies the ready state for server state.
	Ready ServerState = "Ready"
)

// ServerVersion enumerates the values for server version.
type ServerVersion string

const (
	// NineFullStopFive specifies the nine full stop five state for server
	// version.
	NineFullStopFive ServerVersion = "9.5"
	// NineFullStopSix specifies the nine full stop six state for server
	// version.
	NineFullStopSix ServerVersion = "9.6"
)

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Basic specifies the basic state for sku tier.
	Basic SkuTier = "Basic"
	// Standard specifies the standard state for sku tier.
	Standard SkuTier = "Standard"
)

// SslEnforcementEnum enumerates the values for ssl enforcement enum.
type SslEnforcementEnum string

const (
	// SslEnforcementEnumDisabled specifies the ssl enforcement enum disabled
	// state for ssl enforcement enum.
	SslEnforcementEnumDisabled SslEnforcementEnum = "Disabled"
	// SslEnforcementEnumEnabled specifies the ssl enforcement enum enabled
	// state for ssl enforcement enum.
	SslEnforcementEnumEnabled SslEnforcementEnum = "Enabled"
)

// Configuration is represents a Configuration.
type Configuration struct {
	autorest.Response        `json:"-"`
	ID                       *string `json:"id,omitempty"`
	Name                     *string `json:"name,omitempty"`
	Type                     *string `json:"type,omitempty"`
	*ConfigurationProperties `json:"properties,omitempty"`
}

// ConfigurationListResult is a list of server configurations.
type ConfigurationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Configuration `json:"value,omitempty"`
}

// ConfigurationProperties is the properties of a configuration.
type ConfigurationProperties struct {
	Value         *string `json:"value,omitempty"`
	Description   *string `json:"description,omitempty"`
	DefaultValue  *string `json:"defaultValue,omitempty"`
	DataType      *string `json:"dataType,omitempty"`
	AllowedValues *string `json:"allowedValues,omitempty"`
	Source        *string `json:"source,omitempty"`
}

// Database is represents a Database.
type Database struct {
	autorest.Response   `json:"-"`
	ID                  *string `json:"id,omitempty"`
	Name                *string `json:"name,omitempty"`
	Type                *string `json:"type,omitempty"`
	*DatabaseProperties `json:"properties,omitempty"`
}

// DatabaseListResult is a List of databases.
type DatabaseListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Database `json:"value,omitempty"`
}

// DatabaseProperties is the properties of a database.
type DatabaseProperties struct {
	Charset   *string `json:"charset,omitempty"`
	Collation *string `json:"collation,omitempty"`
}

// FirewallRule is represents a server firewall rule.
type FirewallRule struct {
	autorest.Response       `json:"-"`
	ID                      *string `json:"id,omitempty"`
	Name                    *string `json:"name,omitempty"`
	Type                    *string `json:"type,omitempty"`
	*FirewallRuleProperties `json:"properties,omitempty"`
}

// FirewallRuleListResult is a list of firewall rules.
type FirewallRuleListResult struct {
	autorest.Response `json:"-"`
	Value             *[]FirewallRule `json:"value,omitempty"`
}

// FirewallRuleProperties is the properties of a server firewall rule.
type FirewallRuleProperties struct {
	StartIPAddress *string `json:"startIpAddress,omitempty"`
	EndIPAddress   *string `json:"endIpAddress,omitempty"`
}

// LogFile is represents a log file.
type LogFile struct {
	ID                 *string `json:"id,omitempty"`
	Name               *string `json:"name,omitempty"`
	Type               *string `json:"type,omitempty"`
	*LogFileProperties `json:"properties,omitempty"`
}

// LogFileListResult is a list of log files.
type LogFileListResult struct {
	autorest.Response `json:"-"`
	Value             *[]LogFile `json:"value,omitempty"`
}

// LogFileProperties is the properties of a log file.
type LogFileProperties struct {
	Name             *string    `json:"name,omitempty"`
	SizeInKB         *int64     `json:"sizeInKB,omitempty"`
	CreatedTime      *date.Time `json:"createdTime,omitempty"`
	LastModifiedTime *date.Time `json:"lastModifiedTime,omitempty"`
	Type             *string    `json:"type,omitempty"`
	URL              *string    `json:"url,omitempty"`
}

// Operation is rEST API operation definition.
type Operation struct {
	Name       *string                             `json:"name,omitempty"`
	Display    *OperationDisplay                   `json:"display,omitempty"`
	Origin     OperationOrigin                     `json:"origin,omitempty"`
	Properties *map[string]*map[string]interface{} `json:"properties,omitempty"`
}

// OperationDisplay is display metadata associated with the operation.
type OperationDisplay struct {
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Description *string `json:"description,omitempty"`
}

// OperationListResult is a list of resource provider operations.
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
}

// ProxyResource is resource properties.
type ProxyResource struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// Server is represents a server.
type Server struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Sku               *Sku                `json:"sku,omitempty"`
	*ServerProperties `json:"properties,omitempty"`
}

// ServerForCreate is represents a server to be created.
type ServerForCreate struct {
	Sku        *Sku                       `json:"sku,omitempty"`
	Properties *ServerPropertiesForCreate `json:"properties,omitempty"`
	Location   *string                    `json:"location,omitempty"`
	Tags       *map[string]*string        `json:"tags,omitempty"`
}

// ServerListResult is a list of servers.
type ServerListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Server `json:"value,omitempty"`
}

// ServerProperties is the properties of a server.
type ServerProperties struct {
	AdministratorLogin       *string            `json:"administratorLogin,omitempty"`
	StorageMB                *int64             `json:"storageMB,omitempty"`
	Version                  ServerVersion      `json:"version,omitempty"`
	SslEnforcement           SslEnforcementEnum `json:"sslEnforcement,omitempty"`
	UserVisibleState         ServerState        `json:"userVisibleState,omitempty"`
	FullyQualifiedDomainName *string            `json:"fullyQualifiedDomainName,omitempty"`
}

// ServerPropertiesForCreate is the properties used to create a new server.
type ServerPropertiesForCreate struct {
	StorageMB      *int64             `json:"storageMB,omitempty"`
	Version        ServerVersion      `json:"version,omitempty"`
	SslEnforcement SslEnforcementEnum `json:"sslEnforcement,omitempty"`
}

// ServerPropertiesForDefaultCreate is the properties used to create a new
// server.
type ServerPropertiesForDefaultCreate struct {
	StorageMB                  *int64             `json:"storageMB,omitempty"`
	Version                    ServerVersion      `json:"version,omitempty"`
	SslEnforcement             SslEnforcementEnum `json:"sslEnforcement,omitempty"`
	AdministratorLogin         *string            `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword *string            `json:"administratorLoginPassword,omitempty"`
}

// ServerPropertiesForRestore is the properties to a new server by restoring
// from a backup.
type ServerPropertiesForRestore struct {
	StorageMB          *int64             `json:"storageMB,omitempty"`
	Version            ServerVersion      `json:"version,omitempty"`
	SslEnforcement     SslEnforcementEnum `json:"sslEnforcement,omitempty"`
	SourceServerID     *string            `json:"sourceServerId,omitempty"`
	RestorePointInTime *date.Time         `json:"restorePointInTime,omitempty"`
}

// ServerUpdateParameters is parameters allowd to update for a server.
type ServerUpdateParameters struct {
	Sku                               *Sku `json:"sku,omitempty"`
	*ServerUpdateParametersProperties `json:"properties,omitempty"`
	Tags                              *map[string]*string `json:"tags,omitempty"`
}

// ServerUpdateParametersProperties is the properties that can be updated for a
// server.
type ServerUpdateParametersProperties struct {
	StorageMB                  *int64             `json:"storageMB,omitempty"`
	AdministratorLoginPassword *string            `json:"administratorLoginPassword,omitempty"`
	Version                    ServerVersion      `json:"version,omitempty"`
	SslEnforcement             SslEnforcementEnum `json:"sslEnforcement,omitempty"`
}

// Sku is billing information related properties of a server.
type Sku struct {
	Name     *string `json:"name,omitempty"`
	Tier     SkuTier `json:"tier,omitempty"`
	Capacity *int32  `json:"capacity,omitempty"`
	Size     *string `json:"size,omitempty"`
	Family   *string `json:"family,omitempty"`
}

// TrackedResource is resource properties including location and tags for track
// resources.
type TrackedResource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}
