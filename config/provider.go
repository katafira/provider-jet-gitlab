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

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/crossplane-contrib/provider-jet-gitlab/config/project"
	"github.com/crossplane-contrib/provider-jet-gitlab/config/branch"
	"github.com/crossplane-contrib/provider-jet-gitlab/config/group"
	"github.com/crossplane-contrib/provider-jet-gitlab/config/useringroupmembership"
	"github.com/crossplane-contrib/provider-jet-gitlab/config/groupinprojectmembership"
	"github.com/crossplane-contrib/provider-jet-gitlab/config/groupingroupmembership"
	"github.com/crossplane-contrib/provider-jet-gitlab/config/userinprojectmembership"
	"github.com/crossplane-contrib/provider-jet-gitlab/config/user"
)

const (
	resourcePrefix = "gitlab"
	modulePath     = "github.com/crossplane-contrib/provider-jet-gitlab"
)

//go:embed schema.json
var providerSchema string

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
	    tjconfig.WithDefaultResourceFn(defaultResourceFn),
	    tjconfig.WithIncludeList([]string{
	        "gitlab_project$",
  	      "gitlab_branch$",
					"gitlab_group$",
					"gitlab_group_membership$",
					"gitlab_project_share_group$",
					"gitlab_group_share_group$",
					"gitlab_project_membership$",
					"gitlab_user$",
	    }))

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		project.Configure,
    branch.Configure,
		group.Configure,
		useringroupmembership.Configure,
		groupinprojectmembership.Configure,
		groupingroupmembership.Configure,
		userinprojectmembership.Configure,
		user.Configure,
	}	{
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
