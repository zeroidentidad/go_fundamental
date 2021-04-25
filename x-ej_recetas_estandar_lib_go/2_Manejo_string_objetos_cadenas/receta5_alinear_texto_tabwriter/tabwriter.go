package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {

	w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ', tabwriter.AlignRight)
	// ver metodos incorporados tabwriter.
	fmt.Fprintln(w, "usuario\tnombre\tapellido\t")
	fmt.Fprintln(w, "rand\tRadomir\tSohlich\t")
	fmt.Fprintln(w, "rand2\tFulanio\tDeTal\t")
	fmt.Fprintln(w, "zeroX\tJesus\tFerrer\t")
	w.Flush()

}
