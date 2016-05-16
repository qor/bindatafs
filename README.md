# BindataFS

Compile qor templates into binary with [go-bindata](https://github.com/jteeuwen/go-bindata)

[![GoDoc](https://godoc.org/github.com/qor/bindatafs?status.svg)](https://godoc.org/github.com/qor/bindatafs)

## Usage

Install BindataFS

```sh
$ go install github.com/qor/bindatafs
```

Initialize BindataFS for your project, `config/bindatafs` is the path you want to store bindatafs related files

```sh
$ bindatafs config/bindatafs
```

Use Bindatafs for QOR Admin

```go
import "<your_project>/config/bindatafs"

func main() {
  Admin = admin.New(&qor.Config{DB: db.Publish.DraftDB()})
  Admin.SetAssetFS(bindatafs.AssetFS)
}
```

Compiling QOR templates

```sh
go run main.go --compile-qor-templates
```

Run with compiled templates

```sh
go run -tags 'bindatafs' main.go
```

Run normally

```sh
go run main.go
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
