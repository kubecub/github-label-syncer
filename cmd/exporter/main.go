package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/kubecub/github-label-syncer/pkg/exporter"

	"github.com/kubecub/github-label-syncer/internal/exporter/cmd"
)

func main() {
	command := cmd.NewDefaultIAMCtlCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

var (
	owner = kingpin.Arg("owner", "Owner of the repository.").Required().String()
	repo  = kingpin.Arg("repo", "Repository whose wanted labels.").Required().String()
	yaml  = kingpin.Flag("yaml", "Use the YAML format.").Short('y').Bool()
	json  = kingpin.Flag("json", "Use the JSON format.").Short('j').Bool()
	table = kingpin.Flag("table", "Use the table format.").Short('t').Bool()

	// TODO: xml, toml, ini, csv
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
