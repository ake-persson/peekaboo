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

func PrintTable(out io.Writer, fields []string, noColor bool, fmtColors []string, rows [][]string) {
	if len(rows) < 1 {
		return
	}

	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	for i, c := range rows[0] {
		if len(fields) > 0 && !InList(rows[0][i], fields) {
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

	for _, r := range rows {
		for i, c := range r {
			if len(fields) > 0 && !InList(rows[0][i], fields) {
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

	w.Flush()
}

func PrintVertTable(out io.Writer, fields []string, noColor bool, fmtColors []string, rows [][]string) {
	if len(rows) < 1 {
		return
	}

	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	for _, r := range rows {
		for i, c := range r {
			if len(fields) > 0 && !InList(rows[0][i], fields) {
				continue
			}
			if noColor {
				fmt.Fprintf(w, "%s\t: %s\n", rows[0][i], c)
			} else if i == 0 {
				fmt.Fprintf(w, "%s%s\t: %s%s%s\n", fmtColors[0], rows[0][i], fmtColors[1], c, color.Reset)
			} else {
				fmt.Fprintf(w, "%s%s\t: %s%s%s\n", fmtColors[2], rows[0][i], fmtColors[3], c, color.Reset)
			}
		}
		fmt.Fprintln(w)
	}

	w.Flush()
}

func PrintCSV(out io.Writer, fields []string, rows [][]string) {
	w := csv.NewWriter(out)
	// Support fields
	// w.Write(rows[0])
	//	for _, r := range rows {
	//		w.WriteAll(r)
	//	}
	w.Flush()
}
