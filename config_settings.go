package baseSetup

import (
	"github.com/hexya-addons/base/basetypes"
	"github.com/hexya-erp/hexya/src/actions"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/views"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
)

var fields_ConfigSettings = map[string]models.FieldDefinition{
	"Company": fields.Many2One{
		RelationModel: h.Company(),
		String:        "Company",
		Required:      true,
		Default: func(env models.Environment) interface{} {
			return h.User().NewSet(env).CurrentUser().Company()
		}},
	"UserDefaultRights": fields.Boolean{
		String: "Default Access Rights",
	},
	"ExternalEmailServerDefault": fields.Boolean{
		String: "External Email Servers",
	},
}

func configSettings_ConfigFields(rs m.ConfigSettingsSet) basetypes.ConfigFieldsMap {
	res := rs.Super().ConfigFields()
	res[h.ConfigSettings().Fields().UserDefaultRights()] = "base_setup.default_user_rights"
	res[h.ConfigSettings().Fields().ExternalEmailServerDefault()] = "base_setup.default_external_email_server"
	return res
}

// OpenCompany opens the current user's company form view
func configSettings_OpenCompany(rs m.ConfigSettingsSet) *actions.Action {
	return &actions.Action{
		Type:     actions.ActionActWindow,
		Name:     "My Company",
		ViewMode: "form",
		Model:    "Company",
		ResID:    h.User().NewSet(rs.Env()).CurrentUser().Company().ID(),
		Target:   "current",
	}
}

// OpenDefaultUser opens the default user form view
func configSettings_OpenDefaultUser(rs m.ConfigSettingsSet) *actions.Action {
	action := actions.Registry.MustGetByXMLID("base_action_res_users")
	action.ResID = h.User().NewSet(rs.Env()).GetRecord("base_default_user").ID()
	action.Views = []views.ViewTuple{{ID: "base_view_users_form", Type: "form"}}
	return action
}

func init() {
	h.ConfigSettings().AddFields(fields_ConfigSettings)
	h.ConfigSettings().Methods().ConfigFields().Extend(configSettings_ConfigFields)
	h.ConfigSettings().NewMethod("OpenCompany", configSettings_OpenCompany)
	h.ConfigSettings().NewMethod("OpenDefaultUser", configSettings_OpenDefaultUser)
}
