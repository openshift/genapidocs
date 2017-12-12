package main

import (
    "flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
    "strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/go-openapi/loads"
	"github.com/openshift/genapidocs/pkg/apidocs"
)

func writeAPIDocs(root, base string) error {
    err := os.RemoveAll(root)
	if err != nil {
		return err
	}

	doc, err := loads.JSONSpec(strings.Join([]string{base, "api/swagger-spec/openshift-openapi-spec.json"}, "/"))
	if err != nil {
		return err
	}

	pages, err := apidocs.BuildPages(doc.Spec())
	if err != nil {
		return err
	}

	err = pages.Write(root)
	if err != nil {
		return err
	}

	topics := apidocs.BuildTopics(pages)

	b, err := yaml.Marshal(topics)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(root, "_topic_map.yml"), b, 0666)
}

func main() {
    root := flag.String("root", "", "location where generated api docs will be saved (--root=./api-docs)")
    base := flag.String("base", ".", fmt.Sprintf("base path to directory containing %q", "api/swagger-spec/openshift-openapi-spec.json. Defaults to the current working directory."))	

    flag.Parse()

    if len(*root) == 0 {
        fmt.Fprintf(os.Stderr, "You must specify a --root flag denoting the location where the generated topic map will be saved: --root=./output\n")    
        os.Exit(1)
    }

	err := writeAPIDocs(*root, *base)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
		os.Exit(1)
	}
}
