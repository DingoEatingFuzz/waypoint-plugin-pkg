package registry

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/dingoeatingfuzz/waypoint-plugin-pkg/builder"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

type RegistryConfig struct {
	Path string `hcl:"path,attr"`
	Host string `hcl:"host,attr"`
}

type Registry struct {
	config RegistryConfig
}

// Implement Configurable
func (r *Registry) Config() (interface{}, error) {
	return &r.config, nil
}

// Implement Registry
func (r *Registry) PushFunc() interface{} {
	// return a function which will be called by Waypoint
	return r.push
}

// A PushFunc does not have a strict signature, you can define the parameters
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
// In addition to default input parameters the builder.Binary from the Build step
// can also be injected.
//
// The output parameters for PushFunc must be a Struct which can
// be serialzied to Protocol Buffers binary format and an error.
// This Output Value will be made available for other functions
// as an input parameter.
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (r *Registry) push(ctx context.Context, ui terminal.UI, binary *builder.Binary) (*Artifact, error) {
	u := ui.Status()
	defer u.Close()

	filePath := r.config.Path

	locationParts := strings.Split(binary.Location, "/")
	name := locationParts[len(locationParts)-1]
	dest := path.Join(filePath, name)

	u.Update(fmt.Sprintf("Copying binary from %s to %s", binary.Location, dest))

	err := CopyFile(binary.Location, dest)
	if err != nil {
		return nil, err
	}

	return &Artifact{
		Source: name,
		Host:   r.config.Host,
	}, nil
}

// From waypoint internal
//
// CopyFile copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file. The file mode will be copied from the source and
// the copied data is synced/flushed to stable storage.
func CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}
