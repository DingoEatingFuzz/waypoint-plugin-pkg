# Pkg Waypoint Plugin

A WIP waypoint plugin for Vercel's nodejs binary build tool, [Pkg](https://github.com/vercel/pkg).

### Builder Plugin

The Pkg builder plugin will call Pkg on your behalf in a somewhat opinionated way. Since the goal of waypoint is to build and deploy a service, a target (`macos`, `linux`, or `windows`) is required. You are also not allowed to specify an entrypoint, the Pkg builder plugin expects a `package.json` file which will be used to determine the entrypoint as well as associated scripts and assets to bundle.

**Example Waypoint HCL**
```hcl
build {
  use "pkg" {
    target = "linux"
  }
}
```

### Registry Plugin

Idk, do I need this? My vision here is that the binary should be either copied or uploaded using an existing Files or S3 registry plugin. Maybe that's not how Waypoint is supposed to work?

### Deploy Plugin

Pkg just builds a binary and has no opinions how that binary should be deployed. Try using Nomad with the `exec` driver.

## Steps

1. To scaffold a new plugin use the `./clone.sh` script passing the destination folder and the Go package
for your new plugin as parameters

```shell
./clone.sh myplugin ../destination_folder github.com/myorg/mypackage
```

2. You can then run the Makefile to compile the new plugin, the `Makefile` will build the plugin for all architectures.

```shell
cd ../destination_folder

make
```

```shell
Build Protos
protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./builder/output.proto
protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./registry/output.proto
protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./platform/output.proto
protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./release/output.proto

Compile Plugin
# Clear the output
rm -rf ./bin
GOOS=linux GOARCH=amd64 go build -o ./bin/linux_amd64/waypoint-plugin-mytest ./main.go 
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin_amd64/waypoint-plugin-mytest ./main.go 
GOOS=windows GOARCH=amd64 go build -o ./bin/windows_amd64/waypoint-plugin-mytest.exe ./main.go 
GOOS=windows GOARCH=386 go build -o ./bin/windows_386/waypoint-plugin-mytest.exe ./main.go 
```

## Building with Docker

To build plugins for release you can use the `build-docker` Makefile target, this will 
build your plugin for all architectures and create zipped artifacts which can be uploaded
to an artifact manager such as GitHub releases.

The built artifacts will be output in the `./releases` folder.

```shell
make build-docker

rm -rf ./releases
DOCKER_BUILDKIT=1 docker build --output releases --progress=plain .
#1 [internal] load .dockerignore
#1 transferring context: 2B done
#1 DONE 0.0s

#...

#14 [export_stage 1/1] COPY --from=build /go/plugin/bin/*.zip .
#14 DONE 0.1s

#15 exporting to client
#15 copying files 36.45MB 0.1s done
#15 DONE 0.1s
```

## Building and releasing with GitHub Actions

When cloning the template a default GitHub Action is created at the path `.github/workflows/build-plugin.yaml`. You can use this action to automatically build and release your plugin.

The action has two main phases:
1. **Build** - This phase builds the plugin binaries for all the supported architectures. It is triggered when pushing
   to a branch or on pull requests.
1. **Release** - This phase creates a new GitHub release containing the built plugin. It is triggered when pushing tags
   which starting with `v`, for example `v0.1.0`.

You can enable this action by clicking on the `Actions` tab in your GitHub repository and enabling GitHub Actions.
