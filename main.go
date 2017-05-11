package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	exitAfterCompile := flag.Bool("exit-after-compile", true, "Exit after compile templates")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("invalid argument")
		os.Exit(1)
	}

	fmt.Println("Initalizing BindataFS...")

	destPath := args[0]
	funcMap := map[string]interface{}{
		"package_path": func() string {
			return destPath
		},
		"package_name": func() string {
			return path.Base(destPath)
		},
		"exit_after_compile": func() bool {
			return *exitAfterCompile
		},
	}

	hasExists := false
	for _, gopath := range strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator)) {
		sourcePath := filepath.Join(gopath, "src/github.com/qor/bindatafs/templates")
		_, err := os.Stat(sourcePath)
		if err == nil {
			hasExists = true
		}
		err = filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
			if err == nil {
				var relativePath = strings.TrimPrefix(path, sourcePath)

				if info.IsDir() {
					err = os.MkdirAll(filepath.Join(destPath, relativePath), os.ModePerm)
				} else if info.Mode().IsRegular() {
					if source, err := ioutil.ReadFile(path); err == nil {
						var tmpl *template.Template
						if tmpl, err = template.New("").Funcs(funcMap).Parse(string(source)); err == nil {
							var result = bytes.NewBufferString("")
							if err = tmpl.Execute(result, ""); err != nil {
								return err
							}
							source = result.Bytes()
						} else {
							return err
						}
						if err = ioutil.WriteFile(filepath.Join(destPath, strings.TrimSuffix(relativePath, ".template")), source, os.ModePerm); err != nil {
							fmt.Println(err)
						}
					}
				}
			}
			return err
		})

		if hasExists && err == nil {
			fmt.Printf("copy from %s to %s\n", sourcePath, destPath)
			break
		}

		if err != nil {
			fmt.Println("failed to copy files:", err)
		}
	}
}
