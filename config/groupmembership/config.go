package groupmembership

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("gitlab_group_membership", func(r *config.Resource) {

        // we need to override the default group that terrajet generated for
        // this resource, which would be "github"
// ShortGroup is the short name of the API group of this CRD. The full
// CRD API group is calculated by adding the group suffix of the provider.
// For example, ShortGroup could be `ec2` where group suffix of the
// provider is `aws.crossplane.io` and in that case, the full group would
// be `ec2.aws.crossplane.io`
        r.ShortGroup = "groupmembership"
        r.ExternalName = config.IdentifierFromProvider
        // r.References["parent_id"] = config.Reference{
        //     Type: "github.com/crossplane-contrib/provider-jet-gitlab/apis/groupmembership/v1alpha1.GroupMembership",
        // }
    })
}
