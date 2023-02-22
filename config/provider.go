/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/niklasbeinghaus/provider-bitwarden/config/folder"
	"github.com/niklasbeinghaus/provider-bitwarden/config/item_login"
	"github.com/niklasbeinghaus/provider-bitwarden/config/item_secure_note"
)

const (
	resourcePrefix = "bitwarden"
	modulePath     = "github.com/niklasbeinghaus/provider-bitwarden"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		folder.Configure,
		item_login.Configure,
		item_secure_note.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
