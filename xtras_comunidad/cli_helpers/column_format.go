package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 8, 8, 0, '\t', 0)

	fmt.Fprintf(w, "\n %s\t%s\t%s\t",
		"Col-1", "Col-2", "Col-3")

	fmt.Fprintf(w, "\n %s\t%s\t%s\t",
		"----", "----", "----")

	for i := 0; i < 5; i++ {
		fmt.Fprintf(
			w, "\n %d\t%d\t%d\t", i, i+1, i+2)
	}

	w.Flush()

	fmt.Println()
}
