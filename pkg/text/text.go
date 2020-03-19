package text

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/mickep76/color"
)

type Table struct {
	Headers []string
	Rows    [][]string
}

type Tables []*Table

func inList(a string, l []string) bool {
	for _, b := range l {
		if a == b {
			return true
		}
	}
	return false
}

func (ts Tables) PrintTable(output io.Writer, fields []string) {
	if len(ts) < 1 {
		return
	}

	w := tabwriter.NewWriter(output, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, fmt.Sprintf("%s%s\t%s%s%s", color.LightCyan, ts[0].Headers[0], color.Cyan, strings.Join(ts[0].Headers[1:], "\t"), color.Reset))
	for _, t := range ts {
		for _, r := range t.Rows {
			fmt.Fprintln(w, fmt.Sprintf("%s%s\t%s%s%s", color.LightYellow, r[0], color.Yellow, strings.Join(r[1:], "\t"), color.Reset))
		}
	}

	w.Flush()
}

func (ts Tables) PrintVertTable(output io.Writer, fields []string) {
	if len(ts) < 1 {
		return
	}

	w := tabwriter.NewWriter(output, 0, 0, 2, ' ', 0)
	for _, t := range ts {
		for _, r := range t.Rows {
			for i, c := range r {
				if len(fields) > 0 && !inList(t.Headers[i], fields) {
					continue
				}
				if i == 0 {
					fmt.Fprintf(w, "%s%s\t: %s%s%s\n", color.LightCyan, t.Headers[i], color.LightYellow, c, color.Reset)
				} else {
					fmt.Fprintf(w, "%s%s\t: %s%s%s\n", color.Cyan, t.Headers[i], color.Yellow, c, color.Reset)
				}
			}
			fmt.Fprintln(w)
		}
	}

	w.Flush()
}

func (ts Tables) PrintCSV(output io.Writer, fields []string) {
	if len(ts) < 1 {
		return
	}

	w := csv.NewWriter(output)
	w.Write(ts[0].Headers)
	for _, t := range ts {
		w.WriteAll(t.Rows)
	}
}
