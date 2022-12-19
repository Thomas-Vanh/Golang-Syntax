package main

import (
	"strings"
)

func main() {
	dna := "ATGATCTCGTAA"
	println(dna_trans(dna))

}

func dna_trans(dna_str string) (rna string) {
	rna1 := strings.ReplaceAll(dna_str, "A", "u") //string
	rna2 := strings.ReplaceAll(rna1, "T", "a")    //string
	rna3 := strings.ReplaceAll(rna2, "C", "g")    //string
	rna4 := strings.ReplaceAll(rna3, "G", "c")    //string
	rna = rna4
	return rna
}
