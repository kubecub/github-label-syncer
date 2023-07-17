// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"

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
	xml   = kingpin.Flag("xml", "Use the XML format.").Short('x').Bool()

	// TODO: Add support for these formats.
	toml = kingpin.Flag("toml", "Use the TOML format.").Bool()
	ini  = kingpin.Flag("ini", "Use the INI format.").Bool()
	csv  = kingpin.Flag("csv", "Use the CSV format.").Bool()
)

var (
	// New flags.
	file  = kingpin.Flag("file", "Export labels to file.").Short('f').String()
	token = kingpin.Flag("token", "GitHub token.").Envar("GITHUB_TOKEN").String()
)

var (
	goVersion = "unknown" // Populated by goreleaser during build

	// Populated by goreleaser during build.
	version = "master"
	commit  = "?"
	date    = ""
)

func main() {
	exporters := exporter.Labels{}
	// var exporters exporter.LabelsToObject
	if buildInfo, available := debug.ReadBuildInfo(); available {
		goVersion = buildInfo.GoVersion

		if date == "" {
			version = buildInfo.Main.Version
			commit = fmt.Sprintf("(unknown, mod sum: %q)", buildInfo.Main.Sum)
			date = "(unknown)"
		}
	}

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
		b, err := exporters.LabelsToYAML(labels)
		if err != nil {
			log.Fatal(err)
		}
		if *file != "" {
			err = os.WriteFile(*file, b, 0o644)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Labels exported to %s\n", *file)
		} else {
			fmt.Println(string(b))
		}
		return
	}

	if *json {
		b, err := exporters.LabelsToJSON(labels)
		if err != nil {
			log.Fatal(err)
		}
		if *file != "" {
			err = os.WriteFile(*file, b, 0o644)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Labels exported to %s\n", *file)
		} else {
			fmt.Println(string(b))
		}
		return
	}

	if *table {
		b, err := exporters.LabelsToTable(labels)
		if err != nil {
			log.Fatal(err)
		}
		if *file != "" {
			err = os.WriteFile(*file, b, 0o644)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Labels exported to %s\n", *file)
		} else {
			fmt.Println(string(b))
		}
		return
	}

	if *xml {
		b, err := exporters.LabelsToXML(labels)
		if err != nil {
			log.Fatal(err)
		}
		if *file != "" {
			err = os.WriteFile(*file, b, 0o644)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Labels exported to %s\n", *file)
		} else {
			fmt.Println(string(b))
		}
		return
	}

	if *toml {
		b, err := exporters.LabelsToTOML(labels)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		return
	}

	if *ini {
		b, err := exporters.LabelsToINI(labels)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		return
	}

	if *csv {
		b, err := exporters.LabelsToCSV(labels)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		return
	}
}
