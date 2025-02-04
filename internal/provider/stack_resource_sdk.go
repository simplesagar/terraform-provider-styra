// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	tfTypes "github.com/Styra/terraform-provider-styra/internal/provider/types"
	"github.com/Styra/terraform-provider-styra/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *StackResourceModel) ToSharedStacksV1StacksPostRequest() *shared.StacksV1StacksPostRequest {
	description := r.Description.ValueString()
	name := r.Name.ValueString()
	readOnly := r.ReadOnly.ValueBool()
	var sourceControl *shared.StacksV1SourceControlConfig
	if r.SourceControl != nil {
		var origin *shared.GitV1GitRepoConfig
		if r.SourceControl.Origin != nil {
			commit := r.SourceControl.Origin.Commit.ValueString()
			credentials := r.SourceControl.Origin.Credentials.ValueString()
			path := r.SourceControl.Origin.Path.ValueString()
			reference := r.SourceControl.Origin.Reference.ValueString()
			var sshCredentials *shared.GitV1SSHCredentials
			if r.SourceControl.Origin.SSHCredentials != nil {
				passphrase := r.SourceControl.Origin.SSHCredentials.Passphrase.ValueString()
				privateKey := r.SourceControl.Origin.SSHCredentials.PrivateKey.ValueString()
				sshCredentials = &shared.GitV1SSHCredentials{
					Passphrase: passphrase,
					PrivateKey: privateKey,
				}
			}
			url := r.SourceControl.Origin.URL.ValueString()
			origin = &shared.GitV1GitRepoConfig{
				Commit:         commit,
				Credentials:    credentials,
				Path:           path,
				Reference:      reference,
				SSHCredentials: sshCredentials,
				URL:            url,
			}
		}
		var stackOrigin *shared.GitV1GitRepoConfig
		if r.SourceControl.StackOrigin != nil {
			commit1 := r.SourceControl.StackOrigin.Commit.ValueString()
			credentials1 := r.SourceControl.StackOrigin.Credentials.ValueString()
			path1 := r.SourceControl.StackOrigin.Path.ValueString()
			reference1 := r.SourceControl.StackOrigin.Reference.ValueString()
			var sshCredentials1 *shared.GitV1SSHCredentials
			if r.SourceControl.StackOrigin.SSHCredentials != nil {
				passphrase1 := r.SourceControl.StackOrigin.SSHCredentials.Passphrase.ValueString()
				privateKey1 := r.SourceControl.StackOrigin.SSHCredentials.PrivateKey.ValueString()
				sshCredentials1 = &shared.GitV1SSHCredentials{
					Passphrase: passphrase1,
					PrivateKey: privateKey1,
				}
			}
			url1 := r.SourceControl.StackOrigin.URL.ValueString()
			stackOrigin = &shared.GitV1GitRepoConfig{
				Commit:         commit1,
				Credentials:    credentials1,
				Path:           path1,
				Reference:      reference1,
				SSHCredentials: sshCredentials1,
				URL:            url1,
			}
		}
		useWorkspaceSettings := r.SourceControl.UseWorkspaceSettings.ValueBool()
		sourceControl = &shared.StacksV1SourceControlConfig{
			Origin:               origin,
			StackOrigin:          stackOrigin,
			UseWorkspaceSettings: useWorkspaceSettings,
		}
	}
	typeVar := r.Type.ValueString()
	var typeParameters *shared.TypeParameters
	if r.TypeParameters != nil {
		typeParameters = &shared.TypeParameters{}
	}
	out := shared.StacksV1StacksPostRequest{
		Description:    description,
		Name:           name,
		ReadOnly:       readOnly,
		SourceControl:  sourceControl,
		Type:           typeVar,
		TypeParameters: typeParameters,
	}
	return &out
}

func (r *StackResourceModel) RefreshFromSharedStacksV1StackConfig(resp shared.StacksV1StackConfig) {
	if resp.Authz == nil {
		r.Authz = nil
	} else {
		r.Authz = &tfTypes.SystemsV1AuthzConfig{}
		r.Authz.RoleBindings = []tfTypes.SystemsV1V1RoleBindingConfig{}
		if len(r.Authz.RoleBindings) > len(resp.Authz.RoleBindings) {
			r.Authz.RoleBindings = r.Authz.RoleBindings[:len(resp.Authz.RoleBindings)]
		}
		for roleBindingsCount, roleBindingsItem := range resp.Authz.RoleBindings {
			var roleBindings1 tfTypes.SystemsV1V1RoleBindingConfig
			roleBindings1.ID = types.StringValue(roleBindingsItem.ID)
			roleBindings1.RoleName = types.StringValue(roleBindingsItem.RoleName)
			if roleBindingsCount+1 > len(r.Authz.RoleBindings) {
				r.Authz.RoleBindings = append(r.Authz.RoleBindings, roleBindings1)
			} else {
				r.Authz.RoleBindings[roleBindingsCount].ID = roleBindings1.ID
				r.Authz.RoleBindings[roleBindingsCount].RoleName = roleBindings1.RoleName
			}
		}
	}
	r.Datasources = []tfTypes.SystemsV1DatasourceConfig{}
	if len(r.Datasources) > len(resp.Datasources) {
		r.Datasources = r.Datasources[:len(resp.Datasources)]
	}
	for datasourcesCount, datasourcesItem := range resp.Datasources {
		var datasources1 tfTypes.SystemsV1DatasourceConfig
		datasources1.Category = types.StringValue(datasourcesItem.Category)
		datasources1.ID = types.StringValue(datasourcesItem.ID)
		datasources1.Optional = types.BoolPointerValue(datasourcesItem.Optional)
		if datasourcesItem.Status == nil {
			datasources1.Status = nil
		} else {
			datasources1.Status = &tfTypes.MetaV1Status{}
			datasources1.Status.Code = types.StringValue(datasourcesItem.Status.Code)
			datasources1.Status.Message = types.StringValue(datasourcesItem.Status.Message)
			datasources1.Status.Timestamp = types.StringValue(datasourcesItem.Status.Timestamp.Format(time.RFC3339Nano))
		}
		if datasourcesCount+1 > len(r.Datasources) {
			r.Datasources = append(r.Datasources, datasources1)
		} else {
			r.Datasources[datasourcesCount].Category = datasources1.Category
			r.Datasources[datasourcesCount].ID = datasources1.ID
			r.Datasources[datasourcesCount].Optional = datasources1.Optional
			r.Datasources[datasourcesCount].Status = datasources1.Status
		}
	}
	r.Description = types.StringValue(resp.Description)
	if len(resp.Errors) > 0 {
		r.Errors = make(map[string]tfTypes.SystemsV1AgentErrors)
		for systemsV1AgentErrorsKey, systemsV1AgentErrorsValue := range resp.Errors {
			var systemsV1AgentErrorsResult tfTypes.SystemsV1AgentErrors
			systemsV1AgentErrorsResult.Errors = []tfTypes.MetaV1Status{}
			for errorsCount, errorsItem := range systemsV1AgentErrorsValue.Errors {
				var errors2 tfTypes.MetaV1Status
				errors2.Code = types.StringValue(errorsItem.Code)
				errors2.Message = types.StringValue(errorsItem.Message)
				errors2.Timestamp = types.StringValue(errorsItem.Timestamp.Format(time.RFC3339Nano))
				if errorsCount+1 > len(systemsV1AgentErrorsResult.Errors) {
					systemsV1AgentErrorsResult.Errors = append(systemsV1AgentErrorsResult.Errors, errors2)
				} else {
					systemsV1AgentErrorsResult.Errors[errorsCount].Code = errors2.Code
					systemsV1AgentErrorsResult.Errors[errorsCount].Message = errors2.Message
					systemsV1AgentErrorsResult.Errors[errorsCount].Timestamp = errors2.Timestamp
				}
			}
			systemsV1AgentErrorsResult.Waiting = types.BoolValue(systemsV1AgentErrorsValue.Waiting)
			r.Errors[systemsV1AgentErrorsKey] = systemsV1AgentErrorsResult
		}
	}
	r.ID = types.StringValue(resp.ID)
	r.MatchingSystems = []types.String{}
	for _, v := range resp.MatchingSystems {
		r.MatchingSystems = append(r.MatchingSystems, types.StringValue(v))
	}
	if resp.Metadata.CreatedAt != nil {
		r.Metadata.CreatedAt = types.StringValue(resp.Metadata.CreatedAt.Format(time.RFC3339Nano))
	} else {
		r.Metadata.CreatedAt = types.StringNull()
	}
	r.Metadata.CreatedBy = types.StringPointerValue(resp.Metadata.CreatedBy)
	r.Metadata.CreatedThrough = types.StringPointerValue(resp.Metadata.CreatedThrough)
	if resp.Metadata.LastModifiedAt != nil {
		r.Metadata.LastModifiedAt = types.StringValue(resp.Metadata.LastModifiedAt.Format(time.RFC3339Nano))
	} else {
		r.Metadata.LastModifiedAt = types.StringNull()
	}
	r.Metadata.LastModifiedBy = types.StringPointerValue(resp.Metadata.LastModifiedBy)
	r.Metadata.LastModifiedThrough = types.StringPointerValue(resp.Metadata.LastModifiedThrough)
	r.MigrationHistory = []tfTypes.SystemsV1MigrationRecord{}
	if len(r.MigrationHistory) > len(resp.MigrationHistory) {
		r.MigrationHistory = r.MigrationHistory[:len(resp.MigrationHistory)]
	}
	for migrationHistoryCount, migrationHistoryItem := range resp.MigrationHistory {
		var migrationHistory1 tfTypes.SystemsV1MigrationRecord
		migrationHistory1.From = types.StringValue(migrationHistoryItem.From)
		migrationHistory1.InitiatedBy = types.StringValue(migrationHistoryItem.InitiatedBy)
		migrationHistory1.InitiatingUser = types.StringValue(migrationHistoryItem.InitiatingUser)
		migrationHistory1.MigratedAt = types.StringValue(migrationHistoryItem.MigratedAt.Format(time.RFC3339Nano))
		migrationHistory1.Recovered = types.BoolPointerValue(migrationHistoryItem.Recovered)
		migrationHistory1.To = types.StringValue(migrationHistoryItem.To)
		if migrationHistoryCount+1 > len(r.MigrationHistory) {
			r.MigrationHistory = append(r.MigrationHistory, migrationHistory1)
		} else {
			r.MigrationHistory[migrationHistoryCount].From = migrationHistory1.From
			r.MigrationHistory[migrationHistoryCount].InitiatedBy = migrationHistory1.InitiatedBy
			r.MigrationHistory[migrationHistoryCount].InitiatingUser = migrationHistory1.InitiatingUser
			r.MigrationHistory[migrationHistoryCount].MigratedAt = migrationHistory1.MigratedAt
			r.MigrationHistory[migrationHistoryCount].Recovered = migrationHistory1.Recovered
			r.MigrationHistory[migrationHistoryCount].To = migrationHistory1.To
		}
	}
	r.MinimumOpaVersion = types.StringPointerValue(resp.MinimumOpaVersion)
	r.Name = types.StringValue(resp.Name)
	r.Policies = []tfTypes.SystemsV1PolicyConfig{}
	if len(r.Policies) > len(resp.Policies) {
		r.Policies = r.Policies[:len(resp.Policies)]
	}
	for policiesCount, policiesItem := range resp.Policies {
		var policies1 tfTypes.SystemsV1PolicyConfig
		policies1.Created = types.StringPointerValue(policiesItem.Created)
		policies1.Enforcement.Enforced = types.BoolValue(policiesItem.Enforcement.Enforced)
		policies1.Enforcement.Type = types.StringValue(policiesItem.Enforcement.Type)
		policies1.ID = types.StringValue(policiesItem.ID)
		policies1.Modules = []tfTypes.SystemsV1Module{}
		for modulesCount, modulesItem := range policiesItem.Modules {
			var modules1 tfTypes.SystemsV1Module
			modules1.Name = types.StringValue(modulesItem.Name)
			modules1.Placeholder = types.BoolPointerValue(modulesItem.Placeholder)
			modules1.ReadOnly = types.BoolValue(modulesItem.ReadOnly)
			if modulesItem.Rules == nil {
				modules1.Rules = nil
			} else {
				modules1.Rules = &tfTypes.PoliciesV1RuleCounts{}
				modules1.Rules.Allow = types.Int64Value(int64(modulesItem.Rules.Allow))
				modules1.Rules.Deny = types.Int64Value(int64(modulesItem.Rules.Deny))
				modules1.Rules.Enforce = types.Int64Value(int64(modulesItem.Rules.Enforce))
				modules1.Rules.Ignore = types.Int64Value(int64(modulesItem.Rules.Ignore))
				modules1.Rules.Monitor = types.Int64Value(int64(modulesItem.Rules.Monitor))
				modules1.Rules.Notify = types.Int64Value(int64(modulesItem.Rules.Notify))
				modules1.Rules.Other = types.Int64Value(int64(modulesItem.Rules.Other))
				modules1.Rules.Test = types.Int64Value(int64(modulesItem.Rules.Test))
				modules1.Rules.Total = types.Int64Value(int64(modulesItem.Rules.Total))
			}
			if modulesCount+1 > len(policies1.Modules) {
				policies1.Modules = append(policies1.Modules, modules1)
			} else {
				policies1.Modules[modulesCount].Name = modules1.Name
				policies1.Modules[modulesCount].Placeholder = modules1.Placeholder
				policies1.Modules[modulesCount].ReadOnly = modules1.ReadOnly
				policies1.Modules[modulesCount].Rules = modules1.Rules
			}
		}
		if policiesItem.Rules == nil {
			policies1.Rules = nil
		} else {
			policies1.Rules = &tfTypes.PoliciesV1RuleCounts{}
			policies1.Rules.Allow = types.Int64Value(int64(policiesItem.Rules.Allow))
			policies1.Rules.Deny = types.Int64Value(int64(policiesItem.Rules.Deny))
			policies1.Rules.Enforce = types.Int64Value(int64(policiesItem.Rules.Enforce))
			policies1.Rules.Ignore = types.Int64Value(int64(policiesItem.Rules.Ignore))
			policies1.Rules.Monitor = types.Int64Value(int64(policiesItem.Rules.Monitor))
			policies1.Rules.Notify = types.Int64Value(int64(policiesItem.Rules.Notify))
			policies1.Rules.Other = types.Int64Value(int64(policiesItem.Rules.Other))
			policies1.Rules.Test = types.Int64Value(int64(policiesItem.Rules.Test))
			policies1.Rules.Total = types.Int64Value(int64(policiesItem.Rules.Total))
		}
		policies1.Type = types.StringValue(policiesItem.Type)
		if policiesCount+1 > len(r.Policies) {
			r.Policies = append(r.Policies, policies1)
		} else {
			r.Policies[policiesCount].Created = policies1.Created
			r.Policies[policiesCount].Enforcement = policies1.Enforcement
			r.Policies[policiesCount].ID = policies1.ID
			r.Policies[policiesCount].Modules = policies1.Modules
			r.Policies[policiesCount].Rules = policies1.Rules
			r.Policies[policiesCount].Type = policies1.Type
		}
	}
	r.ReadOnly = types.BoolValue(resp.ReadOnly)
	if resp.SourceControl == nil {
		r.SourceControl = nil
	} else {
		r.SourceControl = &tfTypes.StacksV1SourceControlConfig{}
		if resp.SourceControl.Origin == nil {
			r.SourceControl.Origin = nil
		} else {
			r.SourceControl.Origin = &tfTypes.GitV1GitRepoConfig{}
			r.SourceControl.Origin.Commit = types.StringValue(resp.SourceControl.Origin.Commit)
			r.SourceControl.Origin.Credentials = types.StringValue(resp.SourceControl.Origin.Credentials)
			r.SourceControl.Origin.Path = types.StringValue(resp.SourceControl.Origin.Path)
			r.SourceControl.Origin.Reference = types.StringValue(resp.SourceControl.Origin.Reference)
			if resp.SourceControl.Origin.SSHCredentials == nil {
				r.SourceControl.Origin.SSHCredentials = nil
			} else {
				r.SourceControl.Origin.SSHCredentials = &tfTypes.GitV1SSHCredentials{}
				r.SourceControl.Origin.SSHCredentials.Passphrase = types.StringValue(resp.SourceControl.Origin.SSHCredentials.Passphrase)
				r.SourceControl.Origin.SSHCredentials.PrivateKey = types.StringValue(resp.SourceControl.Origin.SSHCredentials.PrivateKey)
			}
			r.SourceControl.Origin.URL = types.StringValue(resp.SourceControl.Origin.URL)
		}
		if resp.SourceControl.StackOrigin == nil {
			r.SourceControl.StackOrigin = nil
		} else {
			r.SourceControl.StackOrigin = &tfTypes.GitV1GitRepoConfig{}
			r.SourceControl.StackOrigin.Commit = types.StringValue(resp.SourceControl.StackOrigin.Commit)
			r.SourceControl.StackOrigin.Credentials = types.StringValue(resp.SourceControl.StackOrigin.Credentials)
			r.SourceControl.StackOrigin.Path = types.StringValue(resp.SourceControl.StackOrigin.Path)
			r.SourceControl.StackOrigin.Reference = types.StringValue(resp.SourceControl.StackOrigin.Reference)
			if resp.SourceControl.StackOrigin.SSHCredentials == nil {
				r.SourceControl.StackOrigin.SSHCredentials = nil
			} else {
				r.SourceControl.StackOrigin.SSHCredentials = &tfTypes.GitV1SSHCredentials{}
				r.SourceControl.StackOrigin.SSHCredentials.Passphrase = types.StringValue(resp.SourceControl.StackOrigin.SSHCredentials.Passphrase)
				r.SourceControl.StackOrigin.SSHCredentials.PrivateKey = types.StringValue(resp.SourceControl.StackOrigin.SSHCredentials.PrivateKey)
			}
			r.SourceControl.StackOrigin.URL = types.StringValue(resp.SourceControl.StackOrigin.URL)
		}
		r.SourceControl.UseWorkspaceSettings = types.BoolValue(resp.SourceControl.UseWorkspaceSettings)
	}
	r.Status = types.StringValue(resp.Status)
	r.Type = types.StringValue(resp.Type)
	if resp.TypeParameters == nil {
		r.TypeParameters = nil
	} else {
		r.TypeParameters = &tfTypes.TypeParameters{}
	}
}

func (r *StackResourceModel) ToSharedStacksV1StacksPutRequest() *shared.StacksV1StacksPutRequest {
	description := r.Description.ValueString()
	name := r.Name.ValueString()
	readOnly := r.ReadOnly.ValueBool()
	var sourceControl *shared.StacksV1SourceControlConfig
	if r.SourceControl != nil {
		var origin *shared.GitV1GitRepoConfig
		if r.SourceControl.Origin != nil {
			commit := r.SourceControl.Origin.Commit.ValueString()
			credentials := r.SourceControl.Origin.Credentials.ValueString()
			path := r.SourceControl.Origin.Path.ValueString()
			reference := r.SourceControl.Origin.Reference.ValueString()
			var sshCredentials *shared.GitV1SSHCredentials
			if r.SourceControl.Origin.SSHCredentials != nil {
				passphrase := r.SourceControl.Origin.SSHCredentials.Passphrase.ValueString()
				privateKey := r.SourceControl.Origin.SSHCredentials.PrivateKey.ValueString()
				sshCredentials = &shared.GitV1SSHCredentials{
					Passphrase: passphrase,
					PrivateKey: privateKey,
				}
			}
			url := r.SourceControl.Origin.URL.ValueString()
			origin = &shared.GitV1GitRepoConfig{
				Commit:         commit,
				Credentials:    credentials,
				Path:           path,
				Reference:      reference,
				SSHCredentials: sshCredentials,
				URL:            url,
			}
		}
		var stackOrigin *shared.GitV1GitRepoConfig
		if r.SourceControl.StackOrigin != nil {
			commit1 := r.SourceControl.StackOrigin.Commit.ValueString()
			credentials1 := r.SourceControl.StackOrigin.Credentials.ValueString()
			path1 := r.SourceControl.StackOrigin.Path.ValueString()
			reference1 := r.SourceControl.StackOrigin.Reference.ValueString()
			var sshCredentials1 *shared.GitV1SSHCredentials
			if r.SourceControl.StackOrigin.SSHCredentials != nil {
				passphrase1 := r.SourceControl.StackOrigin.SSHCredentials.Passphrase.ValueString()
				privateKey1 := r.SourceControl.StackOrigin.SSHCredentials.PrivateKey.ValueString()
				sshCredentials1 = &shared.GitV1SSHCredentials{
					Passphrase: passphrase1,
					PrivateKey: privateKey1,
				}
			}
			url1 := r.SourceControl.StackOrigin.URL.ValueString()
			stackOrigin = &shared.GitV1GitRepoConfig{
				Commit:         commit1,
				Credentials:    credentials1,
				Path:           path1,
				Reference:      reference1,
				SSHCredentials: sshCredentials1,
				URL:            url1,
			}
		}
		useWorkspaceSettings := r.SourceControl.UseWorkspaceSettings.ValueBool()
		sourceControl = &shared.StacksV1SourceControlConfig{
			Origin:               origin,
			StackOrigin:          stackOrigin,
			UseWorkspaceSettings: useWorkspaceSettings,
		}
	}
	typeVar := r.Type.ValueString()
	var typeParameters *shared.StacksV1StacksPutRequestTypeParameters
	if r.TypeParameters != nil {
		typeParameters = &shared.StacksV1StacksPutRequestTypeParameters{}
	}
	out := shared.StacksV1StacksPutRequest{
		Description:    description,
		Name:           name,
		ReadOnly:       readOnly,
		SourceControl:  sourceControl,
		Type:           typeVar,
		TypeParameters: typeParameters,
	}
	return &out
}

func (r *StackResourceModel) RefreshFromSharedStacksV1StacksPutResponse(resp *shared.StacksV1StacksPutResponse) {
	if resp != nil {
	}
}
