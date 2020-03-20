package text

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/mickep76/color"
)

func Split(s string, del string) []string {
	out := []string{}
	for _, v := range strings.Split(s, del) {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func InList(a string, l []string) bool {
	for _, b := range l {
		if a == b {
			return true
		}
	}
	return false
}

type Table struct {
	Headers []string
	Rows    [][]string
}

type Tables []*Table

func (ts Tables) PrintTable(output io.Writer, fields []string, noColor bool, fmtColors []string) {
	if len(ts) < 1 {
		return
	}

	w := tabwriter.NewWriter(output, 0, 0, 2, ' ', 0)
	for i, c := range ts[0].Headers {
		if len(fields) > 0 && !InList(ts[0].Headers[i], fields) {
			continue
		}
		if noColor {
			fmt.Fprintf(w, "%s\t", c)
		} else if i == 0 {
			fmt.Fprintf(w, "%s%s%s\t", fmtColors[0], c, color.Reset)
		} else {
			fmt.Fprintf(w, "%s%s%s\t", fmtColors[2], c, color.Reset)
		}
	}
	fmt.Fprintln(w)

	for _, t := range ts {
		for _, r := range t.Rows {
			for i, c := range r {
				if len(fields) > 0 && !InList(t.Headers[i], fields) {
					continue
				}
				if noColor {
					fmt.Fprintf(w, "%s\t", c)
				} else if i == 0 {
					fmt.Fprintf(w, "%s%s%s\t", fmtColors[1], c, color.Reset)
				} else {
					fmt.Fprintf(w, "%s%s%s\t", fmtColors[3], c, color.Reset)
				}
			}
			fmt.Fprintln(w)
		}
	}

	w.Flush()
}

func (ts Tables) PrintVertTable(output io.Writer, fields []string, noColor bool, fmtColors []string) {
	if len(ts) < 1 {
		return
	}

	w := tabwriter.NewWriter(output, 0, 0, 2, ' ', 0)
	for _, t := range ts {
		for _, r := range t.Rows {
			for i, c := range r {
				if len(fields) > 0 && !InList(t.Headers[i], fields) {
					continue
				}
				if noColor {
					fmt.Fprintf(w, "%s\t: %s\n", t.Headers[i], c)
				} else if i == 0 {
					fmt.Fprintf(w, "%s%s\t: %s%s%s\n", fmtColors[0], t.Headers[i], fmtColors[1], c, color.Reset)
				} else {
					fmt.Fprintf(w, "%s%s\t: %s%s%s\n", fmtColors[2], t.Headers[i], fmtColors[3], c, color.Reset)
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
