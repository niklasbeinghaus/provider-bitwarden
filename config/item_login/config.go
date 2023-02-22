package item_login

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("bitwarden_item_login", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "item_login"

		r.References["folder"] = config.Reference{
            Type: "github.com/niklasbeinghaus/provider-bitwarde/apis/folder/v1alpha1.Folder",
        }
    })
}
