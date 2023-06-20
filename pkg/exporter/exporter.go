// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package exporter

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/olekukonko/tablewriter"
	"gopkg.in/ini.v1"
	"sigs.k8s.io/yaml"
)

type LabelsToObject interface {
	// Convert labels to JSON format
	LabelsToJSON(labels []*Label) ([]byte, error)
	// Convert labels to YAML format
	LabelsToYAML(labels []*Label) ([]byte, error)
	// Convert labels to table format
	LabelsToTable(labels []*Label) ([]byte, error)

	// Convert labels to XML format
	LabelsToXML(labels []*Label) ([]byte, error)
	// Convert labels to TOML format
	LabelsToTOML(labels []*Label) ([]byte, error)
	// Convert labels to INI format
	LabelsToINI(labels []*Label) ([]byte, error)
	// Convert labels to CSV format
	LabelsToCSV(labels []*Label) ([]byte, error)
}

type Labels struct {
	Name        string `json:"name" yaml:"name" toml:"name" ini:"name" xml:"name"`
	Description string `json:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty" ini:"description,omitempty" xml:"description,omitempty"`
	Color       string `json:"color,omitempty" yaml:"color,omitempty" toml:"color,omitempty" ini:"color,omitempty" xml:"color,omitempty"`
}

func (l Labels) LabelsToJSON(labels []*Label) ([]byte, error) {
	fmt.Println("hits labels to json")
	return json.Marshal(labels)
}

func (l Labels) LabelsToYAML(labels []*Label) ([]byte, error) {
	fmt.Println("hits labels to yaml")
	return yaml.Marshal(labels)
}

func (l Labels) LabelsToTable(labels []*Label) ([]byte, error) {
	fmt.Println("hits labels to table")
	labelRows := make([][]string, 0, len(labels))
	for _, l := range labels {
		labelRows = append(labelRows, []string{l.Name, l.Description, l.Color})
	}

	b := &bytes.Buffer{}
	t := tablewriter.NewWriter(b)
	t.SetHeader([]string{"Name", "Description", "Color"})
	t.AppendBulk(labelRows)
	t.Render()

	return b.Bytes(), nil
}

func (l Labels) LabelsToXML(labels []*Label) ([]byte, error) {
	fmt.Println("hits labels to xml")
	type XMLLabels struct {
		XMLName xml.Name `xml:"labels"`
		Labels  []*Label `xml:"label"`
	}
	xmlLabels := &XMLLabels{Labels: labels}
	return xml.MarshalIndent(xmlLabels, "", "  ")
}

func (l Labels) LabelsToTOML(labels []*Label) ([]byte, error) {
	fmt.Println("hits labels to toml")
	type TOMLLabel struct {
		Name        string `toml:"name"`
		Description string `toml:"description"`
		Color       string `toml:"color"`
	}
	type TOMLLabels struct {
		Labels []*TOMLLabel `toml:"label"`
	}
	tomlLabels := &TOMLLabels{}
	for _, label := range labels {
		tomlLabels.Labels = append(tomlLabels.Labels, &TOMLLabel{
			Name:        label.Name,
			Description: label.Description,
			Color:       label.Color,
		})
	}
	return nil, nil
}

func (l Labels) LabelsToINI(labels []*Label) ([]byte, error) {
	fmt.Println("hits labels to ini")
	cfg := ini.Empty()
	for _, label := range labels {
		section, err := cfg.NewSection(label.Name)
		if err != nil {
			return nil, err
		}
		section.NewKey("description", label.Description)
		section.NewKey("color", label.Color)
	}
	return nil, nil
}

func (l Labels) LabelsToCSV(labels []*Label) ([]byte, error) {
	fmt.Println("hits labels to csv")
	b := &bytes.Buffer{}
	w := csv.NewWriter(b)
	if err := w.Write([]string{"Name", "Description", "Color"}); err != nil {
		return nil, err
	}
	for _, label := range labels {
		if err := w.Write([]string{label.Name, label.Description, label.Color}); err != nil {
			return nil, err
		}
	}
	w.Flush()
	return b.Bytes(), nil
}
