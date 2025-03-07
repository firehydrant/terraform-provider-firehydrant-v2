// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	tfTypes "github.com/firehydrant/terraform-provider-firehydrant/internal/provider/types"
	"github.com/firehydrant/terraform-provider-firehydrant/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *ChecklistTemplatesDataSourceModel) RefreshFromSharedChecklistTemplateEntityPaginated(resp *shared.ChecklistTemplateEntityPaginated) {
	if resp != nil {
		r.Data = []tfTypes.Checklists{}
		if len(r.Data) > len(resp.Data) {
			r.Data = r.Data[:len(resp.Data)]
		}
		for dataCount, dataItem := range resp.Data {
			var data1 tfTypes.Checklists
			data1.Checks = []tfTypes.ChecklistCheckEntity{}
			for checksCount, checksItem := range dataItem.Checks {
				var checks1 tfTypes.ChecklistCheckEntity
				checks1.Description = types.StringPointerValue(checksItem.Description)
				checks1.ID = types.StringPointerValue(checksItem.ID)
				checks1.Name = types.StringPointerValue(checksItem.Name)
				checks1.Status = types.BoolPointerValue(checksItem.Status)
				if checksCount+1 > len(data1.Checks) {
					data1.Checks = append(data1.Checks, checks1)
				} else {
					data1.Checks[checksCount].Description = checks1.Description
					data1.Checks[checksCount].ID = checks1.ID
					data1.Checks[checksCount].Name = checks1.Name
					data1.Checks[checksCount].Status = checks1.Status
				}
			}
			data1.CreatedAt = types.StringPointerValue(dataItem.CreatedAt)
			data1.Description = types.StringPointerValue(dataItem.Description)
			data1.ID = types.StringPointerValue(dataItem.ID)
			data1.Name = types.StringPointerValue(dataItem.Name)
			if dataItem.Owner == nil {
				data1.Owner = nil
			} else {
				data1.Owner = &tfTypes.TeamEntity{}
				if dataItem.Owner.CreatedAt != nil {
					data1.Owner.CreatedAt = types.StringValue(dataItem.Owner.CreatedAt.Format(time.RFC3339Nano))
				} else {
					data1.Owner.CreatedAt = types.StringNull()
				}
				if dataItem.Owner.CreatedBy == nil {
					data1.Owner.CreatedBy = nil
				} else {
					data1.Owner.CreatedBy = &tfTypes.AuthorEntity{}
					data1.Owner.CreatedBy.Email = types.StringPointerValue(dataItem.Owner.CreatedBy.Email)
					data1.Owner.CreatedBy.ID = types.StringPointerValue(dataItem.Owner.CreatedBy.ID)
					data1.Owner.CreatedBy.Name = types.StringPointerValue(dataItem.Owner.CreatedBy.Name)
					data1.Owner.CreatedBy.Source = types.StringPointerValue(dataItem.Owner.CreatedBy.Source)
				}
				data1.Owner.Description = types.StringPointerValue(dataItem.Owner.Description)
				data1.Owner.Functionalities = []tfTypes.FunctionalityEntity{}
				for functionalitiesCount, functionalitiesItem := range dataItem.Owner.Functionalities {
					var functionalities1 tfTypes.FunctionalityEntity
					functionalities1.ActiveIncidents = []types.String{}
					for _, v := range functionalitiesItem.ActiveIncidents {
						functionalities1.ActiveIncidents = append(functionalities1.ActiveIncidents, types.StringValue(v))
					}
					functionalities1.AlertOnAdd = types.BoolPointerValue(functionalitiesItem.AlertOnAdd)
					functionalities1.AutoAddRespondingTeam = types.BoolPointerValue(functionalitiesItem.AutoAddRespondingTeam)
					if functionalitiesItem.CreatedAt != nil {
						functionalities1.CreatedAt = types.StringValue(functionalitiesItem.CreatedAt.Format(time.RFC3339Nano))
					} else {
						functionalities1.CreatedAt = types.StringNull()
					}
					functionalities1.Description = types.StringPointerValue(functionalitiesItem.Description)
					functionalities1.ExternalResources = []tfTypes.ExternalResourceEntity{}
					for externalResourcesCount, externalResourcesItem := range functionalitiesItem.ExternalResources {
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
						if externalResourcesCount+1 > len(functionalities1.ExternalResources) {
							functionalities1.ExternalResources = append(functionalities1.ExternalResources, externalResources1)
						} else {
							functionalities1.ExternalResources[externalResourcesCount].ConnectionID = externalResources1.ConnectionID
							functionalities1.ExternalResources[externalResourcesCount].ConnectionName = externalResources1.ConnectionName
							functionalities1.ExternalResources[externalResourcesCount].ConnectionType = externalResources1.ConnectionType
							functionalities1.ExternalResources[externalResourcesCount].CreatedAt = externalResources1.CreatedAt
							functionalities1.ExternalResources[externalResourcesCount].Name = externalResources1.Name
							functionalities1.ExternalResources[externalResourcesCount].RemoteID = externalResources1.RemoteID
							functionalities1.ExternalResources[externalResourcesCount].RemoteURL = externalResources1.RemoteURL
							functionalities1.ExternalResources[externalResourcesCount].UpdatedAt = externalResources1.UpdatedAt
						}
					}
					functionalities1.ID = types.StringPointerValue(functionalitiesItem.ID)
					if len(functionalitiesItem.Labels) > 0 {
						functionalities1.Labels = make(map[string]types.String)
						for key, value := range functionalitiesItem.Labels {
							functionalities1.Labels[key] = types.StringValue(value)
						}
					}
					functionalities1.Links = []tfTypes.LinksEntity{}
					for linksCount, linksItem := range functionalitiesItem.Links {
						var links1 tfTypes.LinksEntity
						links1.HrefURL = types.StringPointerValue(linksItem.HrefURL)
						links1.IconURL = types.StringPointerValue(linksItem.IconURL)
						links1.ID = types.StringPointerValue(linksItem.ID)
						links1.Name = types.StringPointerValue(linksItem.Name)
						if linksCount+1 > len(functionalities1.Links) {
							functionalities1.Links = append(functionalities1.Links, links1)
						} else {
							functionalities1.Links[linksCount].HrefURL = links1.HrefURL
							functionalities1.Links[linksCount].IconURL = links1.IconURL
							functionalities1.Links[linksCount].ID = links1.ID
							functionalities1.Links[linksCount].Name = links1.Name
						}
					}
					functionalities1.Name = types.StringPointerValue(functionalitiesItem.Name)
					if functionalitiesItem.Owner == nil {
						functionalities1.Owner = nil
					} else {
						functionalities1.Owner = &tfTypes.TeamEntity1{}
					}
					functionalities1.Slug = types.StringPointerValue(functionalitiesItem.Slug)
					if functionalitiesItem.UpdatedAt != nil {
						functionalities1.UpdatedAt = types.StringValue(functionalitiesItem.UpdatedAt.Format(time.RFC3339Nano))
					} else {
						functionalities1.UpdatedAt = types.StringNull()
					}
					if functionalitiesItem.UpdatedBy == nil {
						functionalities1.UpdatedBy = nil
					} else {
						functionalities1.UpdatedBy = &tfTypes.AuthorEntity{}
						functionalities1.UpdatedBy.Email = types.StringPointerValue(functionalitiesItem.UpdatedBy.Email)
						functionalities1.UpdatedBy.ID = types.StringPointerValue(functionalitiesItem.UpdatedBy.ID)
						functionalities1.UpdatedBy.Name = types.StringPointerValue(functionalitiesItem.UpdatedBy.Name)
						functionalities1.UpdatedBy.Source = types.StringPointerValue(functionalitiesItem.UpdatedBy.Source)
					}
					if functionalitiesCount+1 > len(data1.Owner.Functionalities) {
						data1.Owner.Functionalities = append(data1.Owner.Functionalities, functionalities1)
					} else {
						data1.Owner.Functionalities[functionalitiesCount].ActiveIncidents = functionalities1.ActiveIncidents
						data1.Owner.Functionalities[functionalitiesCount].AlertOnAdd = functionalities1.AlertOnAdd
						data1.Owner.Functionalities[functionalitiesCount].AutoAddRespondingTeam = functionalities1.AutoAddRespondingTeam
						data1.Owner.Functionalities[functionalitiesCount].CreatedAt = functionalities1.CreatedAt
						data1.Owner.Functionalities[functionalitiesCount].Description = functionalities1.Description
						data1.Owner.Functionalities[functionalitiesCount].ExternalResources = functionalities1.ExternalResources
						data1.Owner.Functionalities[functionalitiesCount].ID = functionalities1.ID
						data1.Owner.Functionalities[functionalitiesCount].Labels = functionalities1.Labels
						data1.Owner.Functionalities[functionalitiesCount].Links = functionalities1.Links
						data1.Owner.Functionalities[functionalitiesCount].Name = functionalities1.Name
						data1.Owner.Functionalities[functionalitiesCount].Owner = functionalities1.Owner
						data1.Owner.Functionalities[functionalitiesCount].Slug = functionalities1.Slug
						data1.Owner.Functionalities[functionalitiesCount].UpdatedAt = functionalities1.UpdatedAt
						data1.Owner.Functionalities[functionalitiesCount].UpdatedBy = functionalities1.UpdatedBy
					}
				}
				data1.Owner.ID = types.StringPointerValue(dataItem.Owner.ID)
				data1.Owner.Memberships = []tfTypes.MembershipEntity{}
				for membershipsCount, membershipsItem := range dataItem.Owner.Memberships {
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
					if membershipsCount+1 > len(data1.Owner.Memberships) {
						data1.Owner.Memberships = append(data1.Owner.Memberships, memberships1)
					} else {
						data1.Owner.Memberships[membershipsCount].DefaultIncidentRole = memberships1.DefaultIncidentRole
						data1.Owner.Memberships[membershipsCount].Schedule = memberships1.Schedule
						data1.Owner.Memberships[membershipsCount].User = memberships1.User
					}
				}
				if dataItem.Owner.MsTeamsChannel == nil {
					data1.Owner.MsTeamsChannel = nil
				} else {
					data1.Owner.MsTeamsChannel = &tfTypes.IntegrationsMicrosoftTeamsV2ChannelEntity{}
					data1.Owner.MsTeamsChannel.ChannelID = types.StringPointerValue(dataItem.Owner.MsTeamsChannel.ChannelID)
					data1.Owner.MsTeamsChannel.ChannelName = types.StringPointerValue(dataItem.Owner.MsTeamsChannel.ChannelName)
					data1.Owner.MsTeamsChannel.ChannelURL = types.StringPointerValue(dataItem.Owner.MsTeamsChannel.ChannelURL)
					data1.Owner.MsTeamsChannel.ID = types.StringPointerValue(dataItem.Owner.MsTeamsChannel.ID)
					data1.Owner.MsTeamsChannel.MsTeamID = types.StringPointerValue(dataItem.Owner.MsTeamsChannel.MsTeamID)
					data1.Owner.MsTeamsChannel.Status = types.StringPointerValue(dataItem.Owner.MsTeamsChannel.Status)
					data1.Owner.MsTeamsChannel.TeamName = types.StringPointerValue(dataItem.Owner.MsTeamsChannel.TeamName)
				}
				data1.Owner.Name = types.StringPointerValue(dataItem.Owner.Name)
				data1.Owner.OwnedFunctionalities = []tfTypes.FunctionalityEntity{}
				for ownedFunctionalitiesCount, ownedFunctionalitiesItem := range dataItem.Owner.OwnedFunctionalities {
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
						for key1, value1 := range ownedFunctionalitiesItem.Labels {
							ownedFunctionalities1.Labels[key1] = types.StringValue(value1)
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
					if ownedFunctionalitiesCount+1 > len(data1.Owner.OwnedFunctionalities) {
						data1.Owner.OwnedFunctionalities = append(data1.Owner.OwnedFunctionalities, ownedFunctionalities1)
					} else {
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].ActiveIncidents = ownedFunctionalities1.ActiveIncidents
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].AlertOnAdd = ownedFunctionalities1.AlertOnAdd
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].AutoAddRespondingTeam = ownedFunctionalities1.AutoAddRespondingTeam
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].CreatedAt = ownedFunctionalities1.CreatedAt
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].Description = ownedFunctionalities1.Description
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].ExternalResources = ownedFunctionalities1.ExternalResources
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].ID = ownedFunctionalities1.ID
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].Labels = ownedFunctionalities1.Labels
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].Links = ownedFunctionalities1.Links
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].Name = ownedFunctionalities1.Name
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].Owner = ownedFunctionalities1.Owner
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].Slug = ownedFunctionalities1.Slug
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].UpdatedAt = ownedFunctionalities1.UpdatedAt
						data1.Owner.OwnedFunctionalities[ownedFunctionalitiesCount].UpdatedBy = ownedFunctionalities1.UpdatedBy
					}
				}
				data1.Owner.OwnedRunbooks = []tfTypes.SlimRunbookEntity{}
				for ownedRunbooksCount, ownedRunbooksItem := range dataItem.Owner.OwnedRunbooks {
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
					if ownedRunbooksCount+1 > len(data1.Owner.OwnedRunbooks) {
						data1.Owner.OwnedRunbooks = append(data1.Owner.OwnedRunbooks, ownedRunbooks1)
					} else {
						data1.Owner.OwnedRunbooks[ownedRunbooksCount].AttachmentRule = ownedRunbooks1.AttachmentRule
						data1.Owner.OwnedRunbooks[ownedRunbooksCount].Categories = ownedRunbooks1.Categories
						data1.Owner.OwnedRunbooks[ownedRunbooksCount].CreatedAt = ownedRunbooks1.CreatedAt
						data1.Owner.OwnedRunbooks[ownedRunbooksCount].Description = ownedRunbooks1.Description
						data1.Owner.OwnedRunbooks[ownedRunbooksCount].ID = ownedRunbooks1.ID
						data1.Owner.OwnedRunbooks[ownedRunbooksCount].Name = ownedRunbooks1.Name
						data1.Owner.OwnedRunbooks[ownedRunbooksCount].Owner = ownedRunbooks1.Owner
						data1.Owner.OwnedRunbooks[ownedRunbooksCount].Summary = ownedRunbooks1.Summary
						data1.Owner.OwnedRunbooks[ownedRunbooksCount].Type = ownedRunbooks1.Type
						data1.Owner.OwnedRunbooks[ownedRunbooksCount].UpdatedAt = ownedRunbooks1.UpdatedAt
					}
				}
				data1.Owner.SignalsIcalURL = types.StringPointerValue(dataItem.Owner.SignalsIcalURL)
				if dataItem.Owner.SlackChannel == nil {
					data1.Owner.SlackChannel = nil
				} else {
					data1.Owner.SlackChannel = &tfTypes.IntegrationsSlackSlackChannelEntity{}
					data1.Owner.SlackChannel.ID = types.StringPointerValue(dataItem.Owner.SlackChannel.ID)
					data1.Owner.SlackChannel.Name = types.StringPointerValue(dataItem.Owner.SlackChannel.Name)
					data1.Owner.SlackChannel.SlackChannelID = types.StringPointerValue(dataItem.Owner.SlackChannel.SlackChannelID)
				}
				data1.Owner.Slug = types.StringPointerValue(dataItem.Owner.Slug)
				if dataItem.Owner.UpdatedAt != nil {
					data1.Owner.UpdatedAt = types.StringValue(dataItem.Owner.UpdatedAt.Format(time.RFC3339Nano))
				} else {
					data1.Owner.UpdatedAt = types.StringNull()
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
				r.Data[dataCount].Checks = data1.Checks
				r.Data[dataCount].CreatedAt = data1.CreatedAt
				r.Data[dataCount].Description = data1.Description
				r.Data[dataCount].ID = data1.ID
				r.Data[dataCount].Name = data1.Name
				r.Data[dataCount].Owner = data1.Owner
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
