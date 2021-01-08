package main

import (
	"github.com/dingoeatingfuzz/waypoint-plugin-pkg/builder"
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
)

func main() {
	sdk.Main(sdk.WithComponents(
		&builder.Builder{},
	))
}
