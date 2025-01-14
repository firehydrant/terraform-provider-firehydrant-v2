// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/provider/types"
	"github.com/speakeasy/terraform-provider-firehydrant-terraform-sdk/internal/sdk/models/shared"
	"time"
)

func (r *IncidentTypesDataSourceModel) RefreshFromSharedIncidentTypeEntityPaginated(resp *shared.IncidentTypeEntityPaginated) {
	if resp != nil {
		r.Data = []tfTypes.IncidentTypeEntity{}
		if len(r.Data) > len(resp.Data) {
			r.Data = r.Data[:len(resp.Data)]
		}
		for dataCount, dataItem := range resp.Data {
			var data1 tfTypes.IncidentTypeEntity
			if dataItem.CreatedAt != nil {
				data1.CreatedAt = types.StringValue(dataItem.CreatedAt.Format(time.RFC3339Nano))
			} else {
				data1.CreatedAt = types.StringNull()
			}
			data1.ID = types.StringPointerValue(dataItem.ID)
			data1.Name = types.StringPointerValue(dataItem.Name)
			if dataItem.Template == nil {
				data1.Template = nil
			} else {
				data1.Template = &tfTypes.IncidentTypeEntityTemplateEntity{}
				data1.Template.CustomFields = types.StringPointerValue(dataItem.Template.CustomFields)
				data1.Template.CustomerImpactSummary = types.StringPointerValue(dataItem.Template.CustomerImpactSummary)
				data1.Template.Description = types.StringPointerValue(dataItem.Template.Description)
				data1.Template.Impacts = []tfTypes.IncidentTypeEntityTemplateImpactEntity{}
				for impactsCount, impactsItem := range dataItem.Template.Impacts {
					var impacts1 tfTypes.IncidentTypeEntityTemplateImpactEntity
					impacts1.ConditionID = types.StringPointerValue(impactsItem.ConditionID)
					impacts1.ConditionName = types.StringPointerValue(impactsItem.ConditionName)
					impacts1.ID = types.StringPointerValue(impactsItem.ID)
					impacts1.Name = types.StringPointerValue(impactsItem.Name)
					if impactsItem.Type != nil {
						impacts1.Type = types.StringValue(string(*impactsItem.Type))
					} else {
						impacts1.Type = types.StringNull()
					}
					if impactsCount+1 > len(data1.Template.Impacts) {
						data1.Template.Impacts = append(data1.Template.Impacts, impacts1)
					} else {
						data1.Template.Impacts[impactsCount].ConditionID = impacts1.ConditionID
						data1.Template.Impacts[impactsCount].ConditionName = impacts1.ConditionName
						data1.Template.Impacts[impactsCount].ID = impacts1.ID
						data1.Template.Impacts[impactsCount].Name = impacts1.Name
						data1.Template.Impacts[impactsCount].Type = impacts1.Type
					}
				}
				data1.Template.IncidentName = types.StringPointerValue(dataItem.Template.IncidentName)
				if len(dataItem.Template.Labels) > 0 {
					data1.Template.Labels = make(map[string]types.String)
					for key, value := range dataItem.Template.Labels {
						data1.Template.Labels[key] = types.StringValue(value)
					}
				}
				data1.Template.Priority = types.StringPointerValue(dataItem.Template.Priority)
				data1.Template.PrivateIncident = types.BoolPointerValue(dataItem.Template.PrivateIncident)
				data1.Template.RunbookIds = []types.String{}
				for _, v := range dataItem.Template.RunbookIds {
					data1.Template.RunbookIds = append(data1.Template.RunbookIds, types.StringValue(v))
				}
				data1.Template.Severity = types.StringPointerValue(dataItem.Template.Severity)
				data1.Template.Summary = types.StringPointerValue(dataItem.Template.Summary)
				data1.Template.TagList = []types.String{}
				for _, v := range dataItem.Template.TagList {
					data1.Template.TagList = append(data1.Template.TagList, types.StringValue(v))
				}
				data1.Template.TeamIds = []types.String{}
				for _, v := range dataItem.Template.TeamIds {
					data1.Template.TeamIds = append(data1.Template.TeamIds, types.StringValue(v))
				}
			}
			if dataItem.TemplateValues == nil {
				data1.TemplateValues = nil
			} else {
				data1.TemplateValues = &tfTypes.IncidentTypeEntityTemplateValuesEntity{}
				data1.TemplateValues.Environments = []tfTypes.IncidentTypeEntityTemplateImpactEntity{}
				for environmentsCount, environmentsItem := range dataItem.TemplateValues.Environments {
					var environments1 tfTypes.IncidentTypeEntityTemplateImpactEntity
					environments1.ConditionID = types.StringPointerValue(environmentsItem.ConditionID)
					environments1.ConditionName = types.StringPointerValue(environmentsItem.ConditionName)
					environments1.ID = types.StringPointerValue(environmentsItem.ID)
					environments1.Name = types.StringPointerValue(environmentsItem.Name)
					if environmentsItem.Type != nil {
						environments1.Type = types.StringValue(string(*environmentsItem.Type))
					} else {
						environments1.Type = types.StringNull()
					}
					if environmentsCount+1 > len(data1.TemplateValues.Environments) {
						data1.TemplateValues.Environments = append(data1.TemplateValues.Environments, environments1)
					} else {
						data1.TemplateValues.Environments[environmentsCount].ConditionID = environments1.ConditionID
						data1.TemplateValues.Environments[environmentsCount].ConditionName = environments1.ConditionName
						data1.TemplateValues.Environments[environmentsCount].ID = environments1.ID
						data1.TemplateValues.Environments[environmentsCount].Name = environments1.Name
						data1.TemplateValues.Environments[environmentsCount].Type = environments1.Type
					}
				}
				data1.TemplateValues.Functionalities = []tfTypes.IncidentTypeEntityTemplateImpactEntity{}
				for functionalitiesCount, functionalitiesItem := range dataItem.TemplateValues.Functionalities {
					var functionalities1 tfTypes.IncidentTypeEntityTemplateImpactEntity
					functionalities1.ConditionID = types.StringPointerValue(functionalitiesItem.ConditionID)
					functionalities1.ConditionName = types.StringPointerValue(functionalitiesItem.ConditionName)
					functionalities1.ID = types.StringPointerValue(functionalitiesItem.ID)
					functionalities1.Name = types.StringPointerValue(functionalitiesItem.Name)
					if functionalitiesItem.Type != nil {
						functionalities1.Type = types.StringValue(string(*functionalitiesItem.Type))
					} else {
						functionalities1.Type = types.StringNull()
					}
					if functionalitiesCount+1 > len(data1.TemplateValues.Functionalities) {
						data1.TemplateValues.Functionalities = append(data1.TemplateValues.Functionalities, functionalities1)
					} else {
						data1.TemplateValues.Functionalities[functionalitiesCount].ConditionID = functionalities1.ConditionID
						data1.TemplateValues.Functionalities[functionalitiesCount].ConditionName = functionalities1.ConditionName
						data1.TemplateValues.Functionalities[functionalitiesCount].ID = functionalities1.ID
						data1.TemplateValues.Functionalities[functionalitiesCount].Name = functionalities1.Name
						data1.TemplateValues.Functionalities[functionalitiesCount].Type = functionalities1.Type
					}
				}
				if dataItem.TemplateValues.Runbooks == nil {
					data1.TemplateValues.Runbooks = nil
				} else {
					data1.TemplateValues.Runbooks = &tfTypes.TeamEntity1{}
				}
				data1.TemplateValues.Services = []tfTypes.IncidentTypeEntityTemplateImpactEntity{}
				for servicesCount, servicesItem := range dataItem.TemplateValues.Services {
					var services1 tfTypes.IncidentTypeEntityTemplateImpactEntity
					services1.ConditionID = types.StringPointerValue(servicesItem.ConditionID)
					services1.ConditionName = types.StringPointerValue(servicesItem.ConditionName)
					services1.ID = types.StringPointerValue(servicesItem.ID)
					services1.Name = types.StringPointerValue(servicesItem.Name)
					if servicesItem.Type != nil {
						services1.Type = types.StringValue(string(*servicesItem.Type))
					} else {
						services1.Type = types.StringNull()
					}
					if servicesCount+1 > len(data1.TemplateValues.Services) {
						data1.TemplateValues.Services = append(data1.TemplateValues.Services, services1)
					} else {
						data1.TemplateValues.Services[servicesCount].ConditionID = services1.ConditionID
						data1.TemplateValues.Services[servicesCount].ConditionName = services1.ConditionName
						data1.TemplateValues.Services[servicesCount].ID = services1.ID
						data1.TemplateValues.Services[servicesCount].Name = services1.Name
						data1.TemplateValues.Services[servicesCount].Type = services1.Type
					}
				}
				data1.TemplateValues.Teams = []tfTypes.TeamEntity{}
				for teamsCount, teamsItem := range dataItem.TemplateValues.Teams {
					var teams1 tfTypes.TeamEntity
					if teamsItem.CreatedAt != nil {
						teams1.CreatedAt = types.StringValue(teamsItem.CreatedAt.Format(time.RFC3339Nano))
					} else {
						teams1.CreatedAt = types.StringNull()
					}
					if teamsItem.CreatedBy == nil {
						teams1.CreatedBy = nil
					} else {
						teams1.CreatedBy = &tfTypes.AuthorEntity{}
						teams1.CreatedBy.Email = types.StringPointerValue(teamsItem.CreatedBy.Email)
						teams1.CreatedBy.ID = types.StringPointerValue(teamsItem.CreatedBy.ID)
						teams1.CreatedBy.Name = types.StringPointerValue(teamsItem.CreatedBy.Name)
						teams1.CreatedBy.Source = types.StringPointerValue(teamsItem.CreatedBy.Source)
					}
					teams1.Description = types.StringPointerValue(teamsItem.Description)
					teams1.Functionalities = []tfTypes.FunctionalityEntity{}
					for functionalitiesCount1, functionalitiesItem1 := range teamsItem.Functionalities {
						var functionalities3 tfTypes.FunctionalityEntity
						functionalities3.ActiveIncidents = []types.String{}
						for _, v := range functionalitiesItem1.ActiveIncidents {
							functionalities3.ActiveIncidents = append(functionalities3.ActiveIncidents, types.StringValue(v))
						}
						functionalities3.AlertOnAdd = types.BoolPointerValue(functionalitiesItem1.AlertOnAdd)
						functionalities3.AutoAddRespondingTeam = types.BoolPointerValue(functionalitiesItem1.AutoAddRespondingTeam)
						if functionalitiesItem1.CreatedAt != nil {
							functionalities3.CreatedAt = types.StringValue(functionalitiesItem1.CreatedAt.Format(time.RFC3339Nano))
						} else {
							functionalities3.CreatedAt = types.StringNull()
						}
						functionalities3.Description = types.StringPointerValue(functionalitiesItem1.Description)
						functionalities3.ExternalResources = []tfTypes.ExternalResourceEntity{}
						for externalResourcesCount, externalResourcesItem := range functionalitiesItem1.ExternalResources {
							var externalResources1 tfTypes.ExternalResourceEntity
							externalResources1.ConnectionID = types.StringPointerValue(externalResourcesItem.ConnectionID)
							externalResources1.ConnectionName = types.StringPointerValue(externalResourcesItem.ConnectionName)
							externalResources1.ConnectionType = types.StringPointerValue(externalResourcesItem.ConnectionType)
							if externalResourcesItem.CreatedAt != nil {
								externalResources1.CreatedAt = types.StringValue(externalResourcesItem.CreatedAt.Format(time.RFC3339Nano))
							} else {
								externalResources1.CreatedAt = types.StringNull()
							}
							externalResources1.Name = types.StringPointerValue(externalResourcesItem.Name)
							externalResources1.RemoteID = types.StringPointerValue(externalResourcesItem.RemoteID)
							externalResources1.RemoteURL = types.StringPointerValue(externalResourcesItem.RemoteURL)
							if externalResourcesItem.UpdatedAt != nil {
								externalResources1.UpdatedAt = types.StringValue(externalResourcesItem.UpdatedAt.Format(time.RFC3339Nano))
							} else {
								externalResources1.UpdatedAt = types.StringNull()
							}
							if externalResourcesCount+1 > len(functionalities3.ExternalResources) {
								functionalities3.ExternalResources = append(functionalities3.ExternalResources, externalResources1)
							} else {
								functionalities3.ExternalResources[externalResourcesCount].ConnectionID = externalResources1.ConnectionID
								functionalities3.ExternalResources[externalResourcesCount].ConnectionName = externalResources1.ConnectionName
								functionalities3.ExternalResources[externalResourcesCount].ConnectionType = externalResources1.ConnectionType
								functionalities3.ExternalResources[externalResourcesCount].CreatedAt = externalResources1.CreatedAt
								functionalities3.ExternalResources[externalResourcesCount].Name = externalResources1.Name
								functionalities3.ExternalResources[externalResourcesCount].RemoteID = externalResources1.RemoteID
								functionalities3.ExternalResources[externalResourcesCount].RemoteURL = externalResources1.RemoteURL
								functionalities3.ExternalResources[externalResourcesCount].UpdatedAt = externalResources1.UpdatedAt
							}
						}
						functionalities3.ID = types.StringPointerValue(functionalitiesItem1.ID)
						if len(functionalitiesItem1.Labels) > 0 {
							functionalities3.Labels = make(map[string]types.String)
							for key1, value1 := range functionalitiesItem1.Labels {
								functionalities3.Labels[key1] = types.StringValue(value1)
							}
						}
						functionalities3.Links = []tfTypes.LinksEntity{}
						for linksCount, linksItem := range functionalitiesItem1.Links {
							var links1 tfTypes.LinksEntity
							links1.HrefURL = types.StringPointerValue(linksItem.HrefURL)
							links1.IconURL = types.StringPointerValue(linksItem.IconURL)
							links1.ID = types.StringPointerValue(linksItem.ID)
							links1.Name = types.StringPointerValue(linksItem.Name)
							if linksCount+1 > len(functionalities3.Links) {
								functionalities3.Links = append(functionalities3.Links, links1)
							} else {
								functionalities3.Links[linksCount].HrefURL = links1.HrefURL
								functionalities3.Links[linksCount].IconURL = links1.IconURL
								functionalities3.Links[linksCount].ID = links1.ID
								functionalities3.Links[linksCount].Name = links1.Name
							}
						}
						functionalities3.Name = types.StringPointerValue(functionalitiesItem1.Name)
						if functionalitiesItem1.Owner == nil {
							functionalities3.Owner = nil
						} else {
							functionalities3.Owner = &tfTypes.TeamEntity1{}
						}
						functionalities3.Slug = types.StringPointerValue(functionalitiesItem1.Slug)
						if functionalitiesItem1.UpdatedAt != nil {
							functionalities3.UpdatedAt = types.StringValue(functionalitiesItem1.UpdatedAt.Format(time.RFC3339Nano))
						} else {
							functionalities3.UpdatedAt = types.StringNull()
						}
						if functionalitiesItem1.UpdatedBy == nil {
							functionalities3.UpdatedBy = nil
						} else {
							functionalities3.UpdatedBy = &tfTypes.AuthorEntity{}
							functionalities3.UpdatedBy.Email = types.StringPointerValue(functionalitiesItem1.UpdatedBy.Email)
							functionalities3.UpdatedBy.ID = types.StringPointerValue(functionalitiesItem1.UpdatedBy.ID)
							functionalities3.UpdatedBy.Name = types.StringPointerValue(functionalitiesItem1.UpdatedBy.Name)
							functionalities3.UpdatedBy.Source = types.StringPointerValue(functionalitiesItem1.UpdatedBy.Source)
						}
						if functionalitiesCount1+1 > len(teams1.Functionalities) {
							teams1.Functionalities = append(teams1.Functionalities, functionalities3)
						} else {
							teams1.Functionalities[functionalitiesCount1].ActiveIncidents = functionalities3.ActiveIncidents
							teams1.Functionalities[functionalitiesCount1].AlertOnAdd = functionalities3.AlertOnAdd
							teams1.Functionalities[functionalitiesCount1].AutoAddRespondingTeam = functionalities3.AutoAddRespondingTeam
							teams1.Functionalities[functionalitiesCount1].CreatedAt = functionalities3.CreatedAt
							teams1.Functionalities[functionalitiesCount1].Description = functionalities3.Description
							teams1.Functionalities[functionalitiesCount1].ExternalResources = functionalities3.ExternalResources
							teams1.Functionalities[functionalitiesCount1].ID = functionalities3.ID
							teams1.Functionalities[functionalitiesCount1].Labels = functionalities3.Labels
							teams1.Functionalities[functionalitiesCount1].Links = functionalities3.Links
							teams1.Functionalities[functionalitiesCount1].Name = functionalities3.Name
							teams1.Functionalities[functionalitiesCount1].Owner = functionalities3.Owner
							teams1.Functionalities[functionalitiesCount1].Slug = functionalities3.Slug
							teams1.Functionalities[functionalitiesCount1].UpdatedAt = functionalities3.UpdatedAt
							teams1.Functionalities[functionalitiesCount1].UpdatedBy = functionalities3.UpdatedBy
						}
					}
					teams1.ID = types.StringPointerValue(teamsItem.ID)
					teams1.Memberships = []tfTypes.MembershipEntity{}
					for membershipsCount, membershipsItem := range teamsItem.Memberships {
						var memberships1 tfTypes.MembershipEntity
						if membershipsItem.DefaultIncidentRole == nil {
							memberships1.DefaultIncidentRole = nil
						} else {
							memberships1.DefaultIncidentRole = &tfTypes.IncidentRoleEntity{}
							if membershipsItem.DefaultIncidentRole.CreatedAt != nil {
								memberships1.DefaultIncidentRole.CreatedAt = types.StringValue(membershipsItem.DefaultIncidentRole.CreatedAt.Format(time.RFC3339Nano))
							} else {
								memberships1.DefaultIncidentRole.CreatedAt = types.StringNull()
							}
							memberships1.DefaultIncidentRole.Description = types.StringPointerValue(membershipsItem.DefaultIncidentRole.Description)
							if membershipsItem.DefaultIncidentRole.DiscardedAt != nil {
								memberships1.DefaultIncidentRole.DiscardedAt = types.StringValue(membershipsItem.DefaultIncidentRole.DiscardedAt.Format(time.RFC3339Nano))
							} else {
								memberships1.DefaultIncidentRole.DiscardedAt = types.StringNull()
							}
							memberships1.DefaultIncidentRole.ID = types.StringPointerValue(membershipsItem.DefaultIncidentRole.ID)
							memberships1.DefaultIncidentRole.Name = types.StringPointerValue(membershipsItem.DefaultIncidentRole.Name)
							memberships1.DefaultIncidentRole.Summary = types.StringPointerValue(membershipsItem.DefaultIncidentRole.Summary)
							if membershipsItem.DefaultIncidentRole.UpdatedAt != nil {
								memberships1.DefaultIncidentRole.UpdatedAt = types.StringValue(membershipsItem.DefaultIncidentRole.UpdatedAt.Format(time.RFC3339Nano))
							} else {
								memberships1.DefaultIncidentRole.UpdatedAt = types.StringNull()
							}
						}
						if membershipsItem.Schedule == nil {
							memberships1.Schedule = nil
						} else {
							memberships1.Schedule = &tfTypes.ScheduleEntity{}
							memberships1.Schedule.Discarded = types.BoolPointerValue(membershipsItem.Schedule.Discarded)
							memberships1.Schedule.ID = types.StringPointerValue(membershipsItem.Schedule.ID)
							memberships1.Schedule.Integration = types.StringPointerValue(membershipsItem.Schedule.Integration)
							memberships1.Schedule.Name = types.StringPointerValue(membershipsItem.Schedule.Name)
						}
						if membershipsItem.User == nil {
							memberships1.User = nil
						} else {
							memberships1.User = &tfTypes.UserEntity{}
							if membershipsItem.User.CreatedAt != nil {
								memberships1.User.CreatedAt = types.StringValue(membershipsItem.User.CreatedAt.Format(time.RFC3339Nano))
							} else {
								memberships1.User.CreatedAt = types.StringNull()
							}
							memberships1.User.Email = types.StringPointerValue(membershipsItem.User.Email)
							memberships1.User.ID = types.StringPointerValue(membershipsItem.User.ID)
							memberships1.User.Name = types.StringPointerValue(membershipsItem.User.Name)
							memberships1.User.SignalsEnabledNotificationTypes = []types.String{}
							for _, v := range membershipsItem.User.SignalsEnabledNotificationTypes {
								memberships1.User.SignalsEnabledNotificationTypes = append(memberships1.User.SignalsEnabledNotificationTypes, types.StringValue(v))
							}
							memberships1.User.SlackLinked = types.BoolPointerValue(membershipsItem.User.SlackLinked)
							memberships1.User.SlackUserID = types.StringPointerValue(membershipsItem.User.SlackUserID)
							if membershipsItem.User.UpdatedAt != nil {
								memberships1.User.UpdatedAt = types.StringValue(membershipsItem.User.UpdatedAt.Format(time.RFC3339Nano))
							} else {
								memberships1.User.UpdatedAt = types.StringNull()
							}
						}
						if membershipsCount+1 > len(teams1.Memberships) {
							teams1.Memberships = append(teams1.Memberships, memberships1)
						} else {
							teams1.Memberships[membershipsCount].DefaultIncidentRole = memberships1.DefaultIncidentRole
							teams1.Memberships[membershipsCount].Schedule = memberships1.Schedule
							teams1.Memberships[membershipsCount].User = memberships1.User
						}
					}
					if teamsItem.MsTeamsChannel == nil {
						teams1.MsTeamsChannel = nil
					} else {
						teams1.MsTeamsChannel = &tfTypes.IntegrationsMicrosoftTeamsV2ChannelEntity{}
						teams1.MsTeamsChannel.ChannelID = types.StringPointerValue(teamsItem.MsTeamsChannel.ChannelID)
						teams1.MsTeamsChannel.ChannelName = types.StringPointerValue(teamsItem.MsTeamsChannel.ChannelName)
						teams1.MsTeamsChannel.ChannelURL = types.StringPointerValue(teamsItem.MsTeamsChannel.ChannelURL)
						teams1.MsTeamsChannel.ID = types.StringPointerValue(teamsItem.MsTeamsChannel.ID)
						teams1.MsTeamsChannel.MsTeamID = types.StringPointerValue(teamsItem.MsTeamsChannel.MsTeamID)
						teams1.MsTeamsChannel.Status = types.StringPointerValue(teamsItem.MsTeamsChannel.Status)
						teams1.MsTeamsChannel.TeamName = types.StringPointerValue(teamsItem.MsTeamsChannel.TeamName)
					}
					teams1.Name = types.StringPointerValue(teamsItem.Name)
					teams1.OwnedFunctionalities = []tfTypes.FunctionalityEntity{}
					for ownedFunctionalitiesCount, ownedFunctionalitiesItem := range teamsItem.OwnedFunctionalities {
						var ownedFunctionalities1 tfTypes.FunctionalityEntity
						ownedFunctionalities1.ActiveIncidents = []types.String{}
						for _, v := range ownedFunctionalitiesItem.ActiveIncidents {
							ownedFunctionalities1.ActiveIncidents = append(ownedFunctionalities1.ActiveIncidents, types.StringValue(v))
						}
						ownedFunctionalities1.AlertOnAdd = types.BoolPointerValue(ownedFunctionalitiesItem.AlertOnAdd)
						ownedFunctionalities1.AutoAddRespondingTeam = types.BoolPointerValue(ownedFunctionalitiesItem.AutoAddRespondingTeam)
						if ownedFunctionalitiesItem.CreatedAt != nil {
							ownedFunctionalities1.CreatedAt = types.StringValue(ownedFunctionalitiesItem.CreatedAt.Format(time.RFC3339Nano))
						} else {
							ownedFunctionalities1.CreatedAt = types.StringNull()
						}
						ownedFunctionalities1.Description = types.StringPointerValue(ownedFunctionalitiesItem.Description)
						ownedFunctionalities1.ExternalResources = []tfTypes.ExternalResourceEntity{}
						for externalResourcesCount1, externalResourcesItem1 := range ownedFunctionalitiesItem.ExternalResources {
							var externalResources3 tfTypes.ExternalResourceEntity
							externalResources3.ConnectionID = types.StringPointerValue(externalResourcesItem1.ConnectionID)
							externalResources3.ConnectionName = types.StringPointerValue(externalResourcesItem1.ConnectionName)
							externalResources3.ConnectionType = types.StringPointerValue(externalResourcesItem1.ConnectionType)
							if externalResourcesItem1.CreatedAt != nil {
								externalResources3.CreatedAt = types.StringValue(externalResourcesItem1.CreatedAt.Format(time.RFC3339Nano))
							} else {
								externalResources3.CreatedAt = types.StringNull()
							}
							externalResources3.Name = types.StringPointerValue(externalResourcesItem1.Name)
							externalResources3.RemoteID = types.StringPointerValue(externalResourcesItem1.RemoteID)
							externalResources3.RemoteURL = types.StringPointerValue(externalResourcesItem1.RemoteURL)
							if externalResourcesItem1.UpdatedAt != nil {
								externalResources3.UpdatedAt = types.StringValue(externalResourcesItem1.UpdatedAt.Format(time.RFC3339Nano))
							} else {
								externalResources3.UpdatedAt = types.StringNull()
							}
							if externalResourcesCount1+1 > len(ownedFunctionalities1.ExternalResources) {
								ownedFunctionalities1.ExternalResources = append(ownedFunctionalities1.ExternalResources, externalResources3)
							} else {
								ownedFunctionalities1.ExternalResources[externalResourcesCount1].ConnectionID = externalResources3.ConnectionID
								ownedFunctionalities1.ExternalResources[externalResourcesCount1].ConnectionName = externalResources3.ConnectionName
								ownedFunctionalities1.ExternalResources[externalResourcesCount1].ConnectionType = externalResources3.ConnectionType
								ownedFunctionalities1.ExternalResources[externalResourcesCount1].CreatedAt = externalResources3.CreatedAt
								ownedFunctionalities1.ExternalResources[externalResourcesCount1].Name = externalResources3.Name
								ownedFunctionalities1.ExternalResources[externalResourcesCount1].RemoteID = externalResources3.RemoteID
								ownedFunctionalities1.ExternalResources[externalResourcesCount1].RemoteURL = externalResources3.RemoteURL
								ownedFunctionalities1.ExternalResources[externalResourcesCount1].UpdatedAt = externalResources3.UpdatedAt
							}
						}
						ownedFunctionalities1.ID = types.StringPointerValue(ownedFunctionalitiesItem.ID)
						if len(ownedFunctionalitiesItem.Labels) > 0 {
							ownedFunctionalities1.Labels = make(map[string]types.String)
							for key2, value2 := range ownedFunctionalitiesItem.Labels {
								ownedFunctionalities1.Labels[key2] = types.StringValue(value2)
							}
						}
						ownedFunctionalities1.Links = []tfTypes.LinksEntity{}
						for linksCount1, linksItem1 := range ownedFunctionalitiesItem.Links {
							var links3 tfTypes.LinksEntity
							links3.HrefURL = types.StringPointerValue(linksItem1.HrefURL)
							links3.IconURL = types.StringPointerValue(linksItem1.IconURL)
							links3.ID = types.StringPointerValue(linksItem1.ID)
							links3.Name = types.StringPointerValue(linksItem1.Name)
							if linksCount1+1 > len(ownedFunctionalities1.Links) {
								ownedFunctionalities1.Links = append(ownedFunctionalities1.Links, links3)
							} else {
								ownedFunctionalities1.Links[linksCount1].HrefURL = links3.HrefURL
								ownedFunctionalities1.Links[linksCount1].IconURL = links3.IconURL
								ownedFunctionalities1.Links[linksCount1].ID = links3.ID
								ownedFunctionalities1.Links[linksCount1].Name = links3.Name
							}
						}
						ownedFunctionalities1.Name = types.StringPointerValue(ownedFunctionalitiesItem.Name)
						if ownedFunctionalitiesItem.Owner == nil {
							ownedFunctionalities1.Owner = nil
						} else {
							ownedFunctionalities1.Owner = &tfTypes.TeamEntity1{}
						}
						ownedFunctionalities1.Slug = types.StringPointerValue(ownedFunctionalitiesItem.Slug)
						if ownedFunctionalitiesItem.UpdatedAt != nil {
							ownedFunctionalities1.UpdatedAt = types.StringValue(ownedFunctionalitiesItem.UpdatedAt.Format(time.RFC3339Nano))
						} else {
							ownedFunctionalities1.UpdatedAt = types.StringNull()
						}
						if ownedFunctionalitiesItem.UpdatedBy == nil {
							ownedFunctionalities1.UpdatedBy = nil
						} else {
							ownedFunctionalities1.UpdatedBy = &tfTypes.AuthorEntity{}
							ownedFunctionalities1.UpdatedBy.Email = types.StringPointerValue(ownedFunctionalitiesItem.UpdatedBy.Email)
							ownedFunctionalities1.UpdatedBy.ID = types.StringPointerValue(ownedFunctionalitiesItem.UpdatedBy.ID)
							ownedFunctionalities1.UpdatedBy.Name = types.StringPointerValue(ownedFunctionalitiesItem.UpdatedBy.Name)
							ownedFunctionalities1.UpdatedBy.Source = types.StringPointerValue(ownedFunctionalitiesItem.UpdatedBy.Source)
						}
						if ownedFunctionalitiesCount+1 > len(teams1.OwnedFunctionalities) {
							teams1.OwnedFunctionalities = append(teams1.OwnedFunctionalities, ownedFunctionalities1)
						} else {
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].ActiveIncidents = ownedFunctionalities1.ActiveIncidents
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].AlertOnAdd = ownedFunctionalities1.AlertOnAdd
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].AutoAddRespondingTeam = ownedFunctionalities1.AutoAddRespondingTeam
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].CreatedAt = ownedFunctionalities1.CreatedAt
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].Description = ownedFunctionalities1.Description
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].ExternalResources = ownedFunctionalities1.ExternalResources
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].ID = ownedFunctionalities1.ID
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].Labels = ownedFunctionalities1.Labels
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].Links = ownedFunctionalities1.Links
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].Name = ownedFunctionalities1.Name
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].Owner = ownedFunctionalities1.Owner
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].Slug = ownedFunctionalities1.Slug
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].UpdatedAt = ownedFunctionalities1.UpdatedAt
							teams1.OwnedFunctionalities[ownedFunctionalitiesCount].UpdatedBy = ownedFunctionalities1.UpdatedBy
						}
					}
					teams1.OwnedRunbooks = []tfTypes.SlimRunbookEntity{}
					for ownedRunbooksCount, ownedRunbooksItem := range teamsItem.OwnedRunbooks {
						var ownedRunbooks1 tfTypes.SlimRunbookEntity
						if ownedRunbooksItem.AttachmentRule == nil {
							ownedRunbooks1.AttachmentRule = nil
						} else {
							ownedRunbooks1.AttachmentRule = &tfTypes.RulesRuleEntity{}
							if ownedRunbooksItem.AttachmentRule.Logic == nil {
								ownedRunbooks1.AttachmentRule.Logic = nil
							} else {
								ownedRunbooks1.AttachmentRule.Logic = &tfTypes.TeamEntity1{}
							}
							if ownedRunbooksItem.AttachmentRule.UserData == nil {
								ownedRunbooks1.AttachmentRule.UserData = nil
							} else {
								ownedRunbooks1.AttachmentRule.UserData = &tfTypes.FHTypesGenericEntity{}
								ownedRunbooks1.AttachmentRule.UserData.Label = types.StringPointerValue(ownedRunbooksItem.AttachmentRule.UserData.Label)
								ownedRunbooks1.AttachmentRule.UserData.Type = types.StringPointerValue(ownedRunbooksItem.AttachmentRule.UserData.Type)
								ownedRunbooks1.AttachmentRule.UserData.Value = types.StringPointerValue(ownedRunbooksItem.AttachmentRule.UserData.Value)
							}
						}
						ownedRunbooks1.Categories = types.StringPointerValue(ownedRunbooksItem.Categories)
						if ownedRunbooksItem.CreatedAt != nil {
							ownedRunbooks1.CreatedAt = types.StringValue(ownedRunbooksItem.CreatedAt.Format(time.RFC3339Nano))
						} else {
							ownedRunbooks1.CreatedAt = types.StringNull()
						}
						ownedRunbooks1.Description = types.StringPointerValue(ownedRunbooksItem.Description)
						ownedRunbooks1.ID = types.StringPointerValue(ownedRunbooksItem.ID)
						ownedRunbooks1.Name = types.StringPointerValue(ownedRunbooksItem.Name)
						if ownedRunbooksItem.Owner == nil {
							ownedRunbooks1.Owner = nil
						} else {
							ownedRunbooks1.Owner = &tfTypes.TeamEntity1{}
						}
						ownedRunbooks1.Summary = types.StringPointerValue(ownedRunbooksItem.Summary)
						ownedRunbooks1.Type = types.StringPointerValue(ownedRunbooksItem.Type)
						if ownedRunbooksItem.UpdatedAt != nil {
							ownedRunbooks1.UpdatedAt = types.StringValue(ownedRunbooksItem.UpdatedAt.Format(time.RFC3339Nano))
						} else {
							ownedRunbooks1.UpdatedAt = types.StringNull()
						}
						if ownedRunbooksCount+1 > len(teams1.OwnedRunbooks) {
							teams1.OwnedRunbooks = append(teams1.OwnedRunbooks, ownedRunbooks1)
						} else {
							teams1.OwnedRunbooks[ownedRunbooksCount].AttachmentRule = ownedRunbooks1.AttachmentRule
							teams1.OwnedRunbooks[ownedRunbooksCount].Categories = ownedRunbooks1.Categories
							teams1.OwnedRunbooks[ownedRunbooksCount].CreatedAt = ownedRunbooks1.CreatedAt
							teams1.OwnedRunbooks[ownedRunbooksCount].Description = ownedRunbooks1.Description
							teams1.OwnedRunbooks[ownedRunbooksCount].ID = ownedRunbooks1.ID
							teams1.OwnedRunbooks[ownedRunbooksCount].Name = ownedRunbooks1.Name
							teams1.OwnedRunbooks[ownedRunbooksCount].Owner = ownedRunbooks1.Owner
							teams1.OwnedRunbooks[ownedRunbooksCount].Summary = ownedRunbooks1.Summary
							teams1.OwnedRunbooks[ownedRunbooksCount].Type = ownedRunbooks1.Type
							teams1.OwnedRunbooks[ownedRunbooksCount].UpdatedAt = ownedRunbooks1.UpdatedAt
						}
					}
					teams1.SignalsIcalURL = types.StringPointerValue(teamsItem.SignalsIcalURL)
					if teamsItem.SlackChannel == nil {
						teams1.SlackChannel = nil
					} else {
						teams1.SlackChannel = &tfTypes.IntegrationsSlackSlackChannelEntity{}
						teams1.SlackChannel.ID = types.StringPointerValue(teamsItem.SlackChannel.ID)
						teams1.SlackChannel.Name = types.StringPointerValue(teamsItem.SlackChannel.Name)
						teams1.SlackChannel.SlackChannelID = types.StringPointerValue(teamsItem.SlackChannel.SlackChannelID)
					}
					teams1.Slug = types.StringPointerValue(teamsItem.Slug)
					if teamsItem.UpdatedAt != nil {
						teams1.UpdatedAt = types.StringValue(teamsItem.UpdatedAt.Format(time.RFC3339Nano))
					} else {
						teams1.UpdatedAt = types.StringNull()
					}
					if teamsCount+1 > len(data1.TemplateValues.Teams) {
						data1.TemplateValues.Teams = append(data1.TemplateValues.Teams, teams1)
					} else {
						data1.TemplateValues.Teams[teamsCount].CreatedAt = teams1.CreatedAt
						data1.TemplateValues.Teams[teamsCount].CreatedBy = teams1.CreatedBy
						data1.TemplateValues.Teams[teamsCount].Description = teams1.Description
						data1.TemplateValues.Teams[teamsCount].Functionalities = teams1.Functionalities
						data1.TemplateValues.Teams[teamsCount].ID = teams1.ID
						data1.TemplateValues.Teams[teamsCount].Memberships = teams1.Memberships
						data1.TemplateValues.Teams[teamsCount].MsTeamsChannel = teams1.MsTeamsChannel
						data1.TemplateValues.Teams[teamsCount].Name = teams1.Name
						data1.TemplateValues.Teams[teamsCount].OwnedFunctionalities = teams1.OwnedFunctionalities
						data1.TemplateValues.Teams[teamsCount].OwnedRunbooks = teams1.OwnedRunbooks
						data1.TemplateValues.Teams[teamsCount].SignalsIcalURL = teams1.SignalsIcalURL
						data1.TemplateValues.Teams[teamsCount].SlackChannel = teams1.SlackChannel
						data1.TemplateValues.Teams[teamsCount].Slug = teams1.Slug
						data1.TemplateValues.Teams[teamsCount].UpdatedAt = teams1.UpdatedAt
					}
				}
			}
			if dataItem.UpdatedAt != nil {
				data1.UpdatedAt = types.StringValue(dataItem.UpdatedAt.Format(time.RFC3339Nano))
			} else {
				data1.UpdatedAt = types.StringNull()
			}
			if dataCount+1 > len(r.Data) {
				r.Data = append(r.Data, data1)
			} else {
				r.Data[dataCount].CreatedAt = data1.CreatedAt
				r.Data[dataCount].ID = data1.ID
				r.Data[dataCount].Name = data1.Name
				r.Data[dataCount].Template = data1.Template
				r.Data[dataCount].TemplateValues = data1.TemplateValues
				r.Data[dataCount].UpdatedAt = data1.UpdatedAt
			}
		}
		if resp.Pagination == nil {
			r.Pagination = nil
		} else {
			r.Pagination = &tfTypes.PaginationEntity{}
			if resp.Pagination.Count != nil {
				r.Pagination.Count = types.Int64Value(int64(*resp.Pagination.Count))
			} else {
				r.Pagination.Count = types.Int64Null()
			}
			if resp.Pagination.Items != nil {
				r.Pagination.Items = types.Int64Value(int64(*resp.Pagination.Items))
			} else {
				r.Pagination.Items = types.Int64Null()
			}
			if resp.Pagination.Last != nil {
				r.Pagination.Last = types.Int64Value(int64(*resp.Pagination.Last))
			} else {
				r.Pagination.Last = types.Int64Null()
			}
			if resp.Pagination.Next != nil {
				r.Pagination.Next = types.Int64Value(int64(*resp.Pagination.Next))
			} else {
				r.Pagination.Next = types.Int64Null()
			}
			if resp.Pagination.Page != nil {
				r.Pagination.Page = types.Int64Value(int64(*resp.Pagination.Page))
			} else {
				r.Pagination.Page = types.Int64Null()
			}
			if resp.Pagination.Pages != nil {
				r.Pagination.Pages = types.Int64Value(int64(*resp.Pagination.Pages))
			} else {
				r.Pagination.Pages = types.Int64Null()
			}
			if resp.Pagination.Prev != nil {
				r.Pagination.Prev = types.Int64Value(int64(*resp.Pagination.Prev))
			} else {
				r.Pagination.Prev = types.Int64Null()
			}
		}
	}
}
