package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alecthomas/kingpin/v2"
	"github.com/kubecub/github-label-syncer/pkg/exporter"
)


var (
    // Required arguments.
    owner = kingpin.Arg("owner", "Owner of the repository.").Required().String()
    repo  = kingpin.Arg("repo", "Repository whose wanted labels.").Required().String()
    
    // Optional flags.
    yaml  = kingpin.Flag("yaml", "Use the YAML format.").Short('y').Bool()
    json  = kingpin.Flag("json", "Use the JSON format.").Short('j').Bool()
    table = kingpin.Flag("table", "Use the table format.").Short('t').Bool()
    
    // TODO: Add support for these formats.
    xml   = kingpin.Flag("xml", "Use the XML format.").Bool()
    toml  = kingpin.Flag("toml", "Use the TOML format.").Bool()
    ini   = kingpin.Flag("ini", "Use the INI format.").Bool()
    csv   = kingpin.Flag("csv", "Use the CSV format.").Bool()
)

var (
	goVersion = "unknown"	// Populated by goreleaser during build

	// Populated by goreleaser during build
	version = "master"
	commit  = "?"
	date    = ""
)

func main() {
	kingpin.Parse()
	client, err := exporter.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	labels, err := client.ListLabels(context.Background(), *owner, *repo)
	if err != nil {
		log.Fatal(err)
	}

	if *yaml {
		b, err := exporter.LabelsToYAML(labels)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		return
	}

	if *json {
		b, err := exporter.LabelsToJSON(labels)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		return
	}

	if *table {
		b, err := exporter.LabelsToTable(labels)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		return
	}
}
