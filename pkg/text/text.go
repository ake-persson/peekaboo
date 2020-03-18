package text

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

type Table struct {
	Headers []string
	Rows    [][]string
}

type Tables []*Table

func (ts Tables) PrintTable(output io.Writer) {
	if len(ts) < 1 {
		return
	}

	w := tabwriter.NewWriter(output, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, strings.Join(ts[0].Headers, "\t"))
	for _, t := range ts {
		for _, r := range t.Rows {
			fmt.Fprintln(w, strings.Join(r, "\t"))
		}
	}

	w.Flush()
}

func (ts Tables) PrintVertTable(output io.Writer) {
	if len(ts) < 1 {
		return
	}

	w := tabwriter.NewWriter(output, 0, 0, 2, ' ', 0)
	for _, t := range ts {
		for _, r := range t.Rows {
			for i, c := range r {
				fmt.Fprintf(w, "%s\t%s\n", t.Headers[i], c)
			}
			fmt.Fprintln(w)
		}
	}

	w.Flush()
}

func (ts Tables) PrintCSV(output io.Writer) {
	if len(ts) < 1 {
		return
	}

	w := csv.NewWriter(output)
	w.Write(ts[0].Headers)
	for _, t := range ts {
		w.WriteAll(t.Rows)
	}
}
