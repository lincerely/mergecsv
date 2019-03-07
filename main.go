package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Printf("Usage:\t %s file1 file2\n", os.Args[0])
		os.Exit(1)
	}

	filepathA := os.Args[1]
	filepathB := os.Args[2]

	fileA, err := os.Open(filepathA)
	checkErr(err)

	fileB, err := os.Open(filepathB)
	checkErr(err)

	readerA := csv.NewReader(bufio.NewReader(fileA))
	readerB := csv.NewReader(bufio.NewReader(fileB))

	dataA, err := readerA.ReadAll()
	checkErr(err)

	dataB, err := readerB.ReadAll()
	checkErr(err)

	// matching headers
	var headerA2B []int
	for idxA, hA := range dataA[0] {
		headerA2B = append(headerA2B, -1)
		for idxB, hB := range dataB[0] {
			if hA == hB {
				headerA2B[idxA] = idxB
				continue
			}
		}
	}

	// get unique identifier for each row in two tables
	var keyA []string
	var keyB = make(map[string]int, 0)

	for _, valAs := range dataA {
		var uid = ""
		for colIdxA, valA := range valAs {
			if headerA2B[colIdxA] >= 0 {
				uid += valA
			}
		}
		keyA = append(keyA, uid)
	}

	for rowIdxB, valBs := range dataB {
		var uid = ""
		for _, v := range headerA2B {
			if v >= 0 {
				uid += valBs[v]
			}
		}
		keyB[uid] = rowIdxB
	}

	// map by unique identifier
	var out [][]string
	for rowIdxA, valAs := range dataA {
		var outRow = valAs
		for colIdxB, valB := range dataB[keyB[keyA[rowIdxA]]] {
			if colIdxB > 0 {
				outRow = append(outRow, valB)
			}
		}
		out = append(out, outRow)
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(out)

	err = w.Error()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
