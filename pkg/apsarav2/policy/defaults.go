// Copyright 2019 Yunion
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

package policy

import (
	common_policy "yunion.io/x/onecloud/pkg/cloudcommon/policy"
	"yunion.io/x/onecloud/pkg/util/rbacutils"
	"yunion.io/x/pkg/util/rbacscope"
)

const (
	PolicyActionPerform = common_policy.PolicyActionPerform
	PolicyActionList    = common_policy.PolicyActionList
	PolicyActionGet     = common_policy.PolicyActionGet
	PolicyActionCreate  = common_policy.PolicyActionCreate
	PolicyActionUpdate  = common_policy.PolicyActionUpdate
	PolicyActionDelete  = common_policy.PolicyActionDelete
)

var (
	predefinedDefaultPolicies = []rbacutils.SRbacPolicy{
		{
			Auth:  false,
			Scope: rbacscope.ScopeSystem,
			Rules: []rbacutils.SRbacRule{
				{
					Service:  "wz-api",
					Resource: "irs_orders",
					Action:   PolicyActionList,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_orders",
					Action:   PolicyActionGet,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_orders",
					Action:   PolicyActionUpdate,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_orders",
					Action:   PolicyActionPerform,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_orders",
					Action:   PolicyActionDelete,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_billing_reports",
					Action:   PolicyActionList,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_billing_reports",
					Action:   PolicyActionGet,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_billing_reports",
					Action:   PolicyActionUpdate,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_billing_reports",
					Action:   PolicyActionPerform,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_billing_reports",
					Action:   PolicyActionDelete,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_unit_reports",
					Action:   PolicyActionList,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_unit_reports",
					Action:   PolicyActionGet,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_unit_reports",
					Action:   PolicyActionUpdate,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_unit_reports",
					Action:   PolicyActionPerform,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_unit_reports",
					Action:   PolicyActionDelete,
					Result:   rbacutils.Allow,
				},

				{
					Service:  "wz-api",
					Resource: "irs_unified_ids",
					Action:   PolicyActionList,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_unified_ids",
					Action:   PolicyActionCreate,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_unified_ids",
					Action:   PolicyActionGet,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_unified_ids",
					Action:   PolicyActionUpdate,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_unified_ids",
					Action:   PolicyActionPerform,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_unified_ids",
					Action:   PolicyActionDelete,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_upload_resources",
					Action:   PolicyActionList,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_upload_resources",
					Action:   PolicyActionGet,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_upload_resources",
					Action:   PolicyActionUpdate,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_upload_resources",
					Action:   PolicyActionPerform,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_upload_resources",
					Action:   PolicyActionDelete,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "billing_accounts",
					Action:   PolicyActionCreate,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "billing_accounts",
					Action:   PolicyActionList,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "billing_accounts",
					Action:   PolicyActionGet,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "billing_accounts",
					Action:   PolicyActionUpdate,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "billing_accounts",
					Action:   PolicyActionPerform,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "billing_accounts",
					Action:   PolicyActionDelete,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_upload_resource_logs",
					Action:   PolicyActionList,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_upload_resource_logs",
					Action:   PolicyActionGet,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_depts",
					Action:   PolicyActionCreate,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_depts",
					Action:   PolicyActionList,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_depts",
					Action:   PolicyActionGet,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_depts",
					Action:   PolicyActionUpdate,
					Result:   rbacutils.Allow,
				},
				{
					Service:  "wz-api",
					Resource: "irs_depts",
					Action:   PolicyActionPerform,
					Result:   rbacutils.Allow,
				},
			},
		},
	}
)

func init() {
	common_policy.AppendDefaultPolicies(predefinedDefaultPolicies)
}
