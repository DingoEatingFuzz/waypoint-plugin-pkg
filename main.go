package main

import (
	"fmt"

	"github.com/dingoeatingfuzz/waypoint-plugin-pkg/builder"
	"github.com/dingoeatingfuzz/waypoint-plugin-pkg/registry"
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
	"github.com/hashicorp/waypoint/builtin/files"
)

func main() {
	sdk.Main(
		sdk.WithComponents(
			&builder.Builder{},
			&registry.Registry{},
		),
		sdk.WithMappers(PkgFilesMapper),
	)
}

func PkgFilesMapper(src *registry.Artifact) *files.Files {
	return &files.Files{
		Path: fmt.Sprintf("%s/%s", src.Host, src.Source),
	}
}
