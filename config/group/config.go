package group

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("gitlab_group", func(r *config.Resource) {

        // we need to override the default group that terrajet generated for
        // this resource, which would be "github"
        r.ShortGroup = "group"
        r.ExternalName = config.IdentifierFromProvider
        // r.References["parent_id"] = config.Reference{
        //     Type: "github.com/crossplane-contrib/provider-jet-gitlab/apis/group/v1alpha1.Group",
        // }
    })
}
