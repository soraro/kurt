package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"
)

func showResults() {
	if limitFlag < 0 {
		log.Fatal("FATAL CONFIGURATION: --limit flag value must not be negative.")
	}

	w := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 1, '\t', 0)

	fmt.Printf("kurt: KUbernetes Restart Tracker")

	if printNS || printAll {
		fmt.Println("\n\n==========")
		fmt.Fprintf(w, "\n Namespace\tRestarts\t")
		fmt.Fprintf(w, "\n \t\t")
		for _, v := range returnSortedLimit(namespaceTracker, limitFlag) {
			fmt.Fprintf(w, "\n %v\t%v\t", v.Key, v.Value)
		}
		w.Flush()
	}

	if printNode || printAll {
		fmt.Println("\n\n==========")
		fmt.Fprintf(w, "\n Node\tRestarts\t")
		fmt.Fprintf(w, "\n \t\t")
		for _, v := range returnSortedLimit(nodeTracker, limitFlag) {
			fmt.Fprintf(w, "\n %v\t%v\t", v.Key, v.Value)
		}
		w.Flush()
	}

	if printLabel || printAll {
		if len(labelTracker) > 0 {
			fmt.Println("\n\n==========")
			fmt.Fprintf(w, "\n Labels\tRestarts\t")
			fmt.Fprintf(w, "\n \t\t")
			for _, v := range returnSortedLimit(labelTracker, limitFlag) {
				fmt.Fprintf(w, "\n %v\t%v\t", v.Key, v.Value)
			}
			w.Flush()
		}
	}

	if printPods || printAll {
		fmt.Println("\n\n==========")
		fmt.Fprintf(w, "\n Pods\tRestarts\t")
		fmt.Fprintf(w, "\n \t\t")
		for _, v := range returnSortedLimit(podTracker, limitFlag) {
			fmt.Fprintf(w, "\n %v\t%v\t", v.Key, v.Value)
		}
		w.Flush()
	}

	fmt.Printf("\n")
}

// sorting results
// https://stackoverflow.com/a/18695740
type Pair struct {
	Key   string
	Value int32
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func returnSorted(data map[string]int32) PairList {
	pl := make(PairList, len(data))
	i := 0
	for k, v := range data {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func returnSortedLimit(data map[string]int32, limit int) PairList {
	pl := make(PairList, len(data))
	i := 0
	for k, v := range data {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	if len(pl) <= limit || limit == 0 {
		return pl
	} else {
		return pl[0:limit]
	}
}
