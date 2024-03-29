package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"gopkg.in/yaml.v2"
)

type StructuredOutput struct {
	Namespaces ItemList `yaml:"namespaces,omitempty" json:"namespaces,omitempty"`
	Nodes      ItemList `yaml:"nodes,omitempty" json:"nodes,omitempty"`
	Labels     ItemList `yaml:"labels,omitempty" json:"labels,omitempty"`
	Pods       ItemList `yaml:"pods,omitempty" json:"pods,omitempty"`
}

func showResults() {
	var so StructuredOutput
	w := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 1, '\t', 0)

	if output == "standard" {
		fmt.Printf("kurt: KUbernetes Restart Tracker")
	}

	if printNS || printAll {
		so.Namespaces = returnSortedLimit(namespaceTracker, limitFlag, false, nil)
		if output == "standard" {
			fmt.Println("\n\n==========")
			fmt.Fprintf(w, "\n Namespace\tRestarts\t")
			fmt.Fprintf(w, "\n \t\t")
			for _, v := range so.Namespaces {
				fmt.Fprintf(w, "\n %v\t%v\t", v.Name, v.Count)
			}
			w.Flush()
		}
	}

	if printNode || printAll {
		so.Nodes = returnSortedLimit(nodeTracker, limitFlag, false, nil)
		if output == "standard" {
			fmt.Println("\n\n==========")
			fmt.Fprintf(w, "\n Node\tRestarts\t")
			fmt.Fprintf(w, "\n \t\t")
			for _, v := range so.Nodes {
				fmt.Fprintf(w, "\n %v\t%v\t", v.Name, v.Count)
			}
			w.Flush()
		}
	}

	if printLabel || printAll {
		if len(labelTracker) > 0 {
			so.Labels = returnSortedLimit(labelTracker, limitFlag, false, nil)
			if output == "standard" {
				fmt.Println("\n\n==========")
				fmt.Fprintf(w, "\n Label\tRestarts\t")
				fmt.Fprintf(w, "\n \t\t")
				for _, v := range so.Labels {
					fmt.Fprintf(w, "\n %v\t%v\t", v.Name, v.Count)
				}
				w.Flush()
			}
		}
	}

	if printPods || printAll {
		so.Pods = returnSortedLimit(podTracker, limitFlag, true, containerTracker)
		if output == "standard" {
			fmt.Println("\n\n==========")
			fmt.Fprintf(w, "\n Pod\tNamespace\tRestarts\t")
			fmt.Fprintf(w, "\n \t\t\t")
			for _, v := range so.Pods {
				fmt.Fprintf(w, "\n %v\t%v\t%v\t", v.Name, v.Namespace, v.Count)
				if v.Containers != nil && ishowContainers {
					for _, vv := range v.Containers {
						fmt.Fprintf(w, "\n    └%v: %v\t\t\t", vv.Name, vv.Count)
					}
				}
			}
			w.Flush()
		}
	}
	switch output {
	case "json":
		j, _ := json.MarshalIndent(so, "", "  ")
		fmt.Println(string(j))
	case "yaml":
		y, _ := yaml.Marshal(so)
		fmt.Println(string(y))
	default:
		fmt.Printf("\n")
	}

}

// sorting results
// https://stackoverflow.com/a/18695740
type Container struct {
	Name  string `yaml:"name" json:"name"`
	Count int32  `yaml:"count" json:"count"`
}

type Containers []Container

func (p Containers) Len() int           { return len(p) }
func (p Containers) Less(i, j int) bool { return p[i].Count < p[j].Count }
func (p Containers) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Item struct {
	Name       string     `yaml:"name" json:"name"`
	Count      int32      `yaml:"count" json:"count"`
	Namespace  string     `yaml:"namespace,omitempty" json:"namespace,omitempty"`
	Containers Containers `yaml:"containers,omitempty" json:"containers,omitempty"`
}

type ItemList []Item

func (p ItemList) Len() int           { return len(p) }
func (p ItemList) Less(i, j int) bool { return p[i].Count < p[j].Count }
func (p ItemList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func returnSortedLimit(data map[string]int32, limit int, parseNS bool, containers map[string]map[string]int32) ItemList {
	il := make(ItemList, len(data))
	i := 0
	for k, v := range data {
		if parseNS {
			// split the Name so we can display the pod an namespace separately
			// only used for pod items
			s := strings.Split(k, ":")
			il[i] = Item{s[1], v, s[0], createContainerSlice(containers[k])}
		} else {
			il[i] = Item{k, v, "", nil}
		}
		i++
	}
	sort.Sort(sort.Reverse(il))
	if len(il) <= limit || limit == 0 {
		return il
	} else {
		return il[0:limit]
	}
}

func createContainerSlice(containers map[string]int32) []Container {
	if containers != nil {
		c := make(Containers, len(containers))
		i := 0
		for k, v := range containers {
			c[i] = Container{k, v}
			i++
		}
		sort.Sort(sort.Reverse(c))
		return c
	}
	return nil
}
