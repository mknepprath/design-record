package main

import (
	"encoding/json"
	"slices"
	"sort"
)

const dataPath = "_data/record.json"

type siteDoc struct {
	Path      string   `json:"path"`
	Kind      string   `json:"kind"`
	Layer     string   `json:"layer"`
	Ref       string   `json:"ref,omitempty"`
	Title     string   `json:"title"`
	Status    string   `json:"status,omitempty"`
	Owner     string   `json:"owner,omitempty"`
	Date      string   `json:"date"`
	Source    string   `json:"source,omitempty"`
	Projects  []string `json:"projects"`
	Backlinks []string `json:"backlinks"`
}

// renderData builds _data/record.json, which the Pages layouts read for
// page headers, the home page tables and backlinks.
func renderData(docs []doc) (string, error) {
	backlinks := map[string][]string{}
	for _, d := range docs {
		targets := linkTargets(d.Path, d.Body)
		targets = append(targets, d.Evidence()...)
		targets = append(targets, d.Lists["supersedes"]...)
		for _, p := range d.Projects() {
			targets = append(targets, "projects/"+p+".md")
		}
		for _, t := range targets {
			if t != d.Path && !slices.Contains(backlinks[t], d.Path) {
				backlinks[t] = append(backlinks[t], d.Path)
			}
		}
	}

	out := []siteDoc{}
	for _, d := range docs {
		bl := backlinks[d.Path]
		sort.Strings(bl)
		if bl == nil {
			bl = []string{}
		}
		projects := d.Projects()
		if projects == nil {
			projects = []string{}
		}
		out = append(out, siteDoc{
			Path: d.Path, Kind: d.Kind.Name, Layer: string(d.Kind.Layer), Ref: d.Ref(),
			Title: d.Title(), Status: d.Status(), Owner: d.Fields["owner"], Date: d.Date(),
			Source: d.Fields["source"], Projects: projects, Backlinks: bl,
		})
	}
	b, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b) + "\n", nil
}
