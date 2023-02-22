package item_secure_note

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("bitwarden_secure_note", func(r *config.Resource) {
        r.ShortGroup = "item_secure_note"

		r.References["folder"] = config.Reference{
            Type: "github.com/niklasbeinghaus/provider-bitwarden/apis/folder/v1alpha1.Folder",
        }
    })
}
