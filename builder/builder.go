package builder

import (
	"context"
  "strings"
  "io/ioutil"
  "os"
  "os/exec"
  "path"

	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
  PackageJSON "github.com/cloudrecipes/packagejson"
)

type BuildConfig struct {
	Target  string `hcl:"target"`
	Options  []string `hcl:"options,optional"`
  Output   string `hcl:"output,optional"`
  OutPath  string `hcl:"out_path,optional"`
  Build    bool `hcl:"build,optional"`
  Public   bool `hcl:"public,optional"`
}

type Builder struct {
	config BuildConfig
}

func (b *Builder) Config() (interface{}, error) {
	return &b.config, nil
}

func (b *Builder) ConfigSet(config interface{}) error {
  // All config options are optional.
  // Any bad values are handled by pkg itself.
  // Maybe in the future some static checking can be done here to speed things up, idk.

	return nil
}

func (b *Builder) BuildFunc() interface{} {
	// return a function which will be called by Waypoint
	return b.build
}

// A BuildFunc does not have a strict signature, you can define the parameters
// you need based on the Available parameters that the Waypoint SDK provides.
// Waypoint will automatically inject parameters as specified
// in the signature at run time.
//
// Available input parameters:
// - context.Context
// - *component.Source
// - *component.JobInfo
// - *component.DeploymentConfig
// - *datadir.Project
// - *datadir.App
// - *datadir.Component
// - hclog.Logger
// - terminal.UI
// - *component.LabelSet
//
// The output parameters for BuildFunc must be a Struct which can
// be serialzied to Protocol Buffers binary format and an error.
// This Output Value will be made available for other functions
// as an input parameter.
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (b *Builder) build(ctx context.Context, ui terminal.UI) (*Binary, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Building application")

  // Always run using the local package.json
  // Enforce a single target (and therefore a single binary)
  args := []string{".", "--targets", b.config.Target}

  // Flags
  if b.config.Build {
    args = append(args, "--build")
  }
  if b.config.Public {
    args = append(args, "--public")
  }

  // Outputs
  if b.config.Output != "" {
    args = append(args, "--output")
    args = append(args, b.config.Output)
  }
  if b.config.OutPath != "" {
    args = append(args, "--out-path")
    args = append(args, b.config.OutPath)
  }

  // Options
  if b.config.Options != nil {
    args = append(args, "--options")
    args = append(args, strings.Join(args, ","))
  }

  c := exec.Command("pkg", args...)

  pwd, _ := os.Getwd()
  u.Step(terminal.StatusWarn, pwd)
  u.Step(terminal.StatusWarn, strings.Join(args, " "))

  err := c.Run()
  if err != nil {
    u.Step(terminal.StatusError, "Pkg failed to build the binary.")
    return nil, err
  }

  u.Step(terminal.StatusOK, "Pkg successfully built the binary.")

  name := b.config.Output
  if (b.config.Output == "") {
    // The assumed name comes from the package.json file. Read it.
    // This also assumes that waypoint is being run where the package.json lives
    // so uh....that should change at some point...
    packageJsonRaw, err := ioutil.ReadFile("./package.json")
    if err != nil {
      u.Step(terminal.StatusError, "An Output was not specified and there was no package.json file")
      return nil, err
    }

    packageJson, parseErr := PackageJSON.Parse(packageJsonRaw)
    if err != nil {
      return nil, parseErr
    }

    name = packageJson.Name
  }

  // return &Files {
  //   Path: path.Join("./", name),
  // }, nil

	return &Binary{
    Location: path.Join("./", name),
  }, nil
}
