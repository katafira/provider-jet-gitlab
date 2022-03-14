/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	branch "github.com/crossplane-contrib/provider-jet-gitlab/internal/controller/branch/branch"
	group "github.com/crossplane-contrib/provider-jet-gitlab/internal/controller/group/group"
	sharegroup "github.com/crossplane-contrib/provider-jet-gitlab/internal/controller/groupingroupmembership/sharegroup"
	sharegroupgroupinprojectmembership "github.com/crossplane-contrib/provider-jet-gitlab/internal/controller/groupinprojectmembership/sharegroup"
	project "github.com/crossplane-contrib/provider-jet-gitlab/internal/controller/project/project"
	providerconfig "github.com/crossplane-contrib/provider-jet-gitlab/internal/controller/providerconfig"
	user "github.com/crossplane-contrib/provider-jet-gitlab/internal/controller/user/user"
	membership "github.com/crossplane-contrib/provider-jet-gitlab/internal/controller/useringroupmembership/membership"
	membershipuserinprojectmembership "github.com/crossplane-contrib/provider-jet-gitlab/internal/controller/userinprojectmembership/membership"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		branch.Setup,
		group.Setup,
		sharegroup.Setup,
		sharegroupgroupinprojectmembership.Setup,
		project.Setup,
		providerconfig.Setup,
		user.Setup,
		membership.Setup,
		membershipuserinprojectmembership.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
