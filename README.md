# BindataFS

BindataFS could be used to compile templates into binary utilizing [go-bindata](https://github.com/containous/go-bindata)

[![GoDoc](https://godoc.org/github.com/qor/bindatafs?status.svg)](https://godoc.org/github.com/qor/bindatafs)

## Install [BindataFS](https://github.com/qor/bindatafs)

```sh
$ go get -u -f github.com/qor/bindatafs/...
```

Initialize BindataFS for your project, set the path you want to store BindataFS related files, e.g. `config/bindatafs`:

```sh
$ bindatafs config/bindatafs
```

## Usage

```go
import "<your_project>/config/bindatafs"

func main() {
  assetFS := bindatafs.AssetFS

  // Register view paths into AssetFS
  assetFS.RegisterPath("<your_project>/app/views")
  assetFS.RegisterPath("<your_project>/vender/plugin/views")

  // Compile templates under registered view paths into binary
  assetFS.Compile()

  // Get file content with name
  fileContent, ok := assetFS.Asset("home/index.tmpl")
}
```

You need to compile templates into a go file with method `Compile` before run `go build`, and if any templates changed, you need to recompile it.

If you started your application with tag `bindatafs`, AssetFS will access file from generated go file or current binary

```sh
go run -tags 'bindatafs' main.go
```

Or it will access its content from registered view paths of your filesystem, which is easier for your local development

```sh
go run main.go
```

### Using NameSpace

Although you could initalize several assetfs packages to hold templates from different view paths (templates name might has confliction) for different use cases, Bindatafs provide an easier solution for you.

```go
func main() {
  // Generate a sub AssetFS with namespace
  adminAssetFS := assetFS.NameSpace("admin_related_files")

  // Register view paths into this name space
  adminAssetFS.RegisterPath("<your_project>/app/admin_views")

  // Access file that under registered views paths of current name space
  adminAssetFS.Asset("admin_view.tmpl")
}
```

### Use BindataFS with [QOR Admin](https://github.com/qor/admin)

```go
import "<your_project>/config/bindatafs"

func main() {
  Admin = admin.New(&qor.Config{DB: db.Publish.DraftDB()})
  Admin.SetAssetFS(bindatafs.AssetFS.NameSpace("admin"))
}
```

### Use BindataFS with [QOR Render](https://github.com/qor/render)

```go
import  "github.com/qor/render"

func main() {
  View := render.New()
  View.SetAssetFS(assetFS.NameSpace("views"))
}
```

### Use BindataFS with [QOR Widget](https://github.com/qor/widget)

```go
import  "github.com/qor/widget"

func main() {
  Widgets := widget.New(&widget.Config{DB: db.DB})
	Widgets.SetAssetFS(assetFS.NameSpace("widgets"))
}
```

### Use BindataFS with static files

```go
func main() {
  mux := http.NewServeMux()

  // this will add all files under public into a generated go file, which will be included into the binary
  assetFS := assetFS.FileServer(http.Dir("public"))

  // If you only want to include specified paths, you could use it like this
  assetFS := assetFS.FileServer(http.Dir("public"), "javascripts", "stylesheets", "images")

  for _, path := range []string{"javascripts", "stylesheets", "images"} {
    mux.Handle(fmt.Sprintf("/%s/", path), assetFS)
  }
}
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
