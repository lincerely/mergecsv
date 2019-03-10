package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleMergeCSV() {
	dataA := generateData([]string{"A", "B", "C", "D"}, 3, false)
	dataB := generateData([]string{"A", "D", "E"}, 3, true)

	out := mergeCSV(dataA, dataB)
	fmt.Println(out)
	// Output:
	// [[A B C D E] [1A 1B 1C 1D 1E] [2A 2B 2C 2D 2E] [3A 3B 3C 3D 3E]]
}

func BenchmarkMergeCSV1(b *testing.B)     { benchmarkMergeCSV(1, b) }
func BenchmarkMergeCSV10(b *testing.B)    { benchmarkMergeCSV(10, b) }
func BenchmarkMergeCSV100(b *testing.B)   { benchmarkMergeCSV(100, b) }
func BenchmarkMergeCSV1000(b *testing.B)  { benchmarkMergeCSV(1000, b) }
func BenchmarkMergeCSV10000(b *testing.B) { benchmarkMergeCSV(10000, b) }

func benchmarkMergeCSV(rowCount int, b *testing.B) {
	dataA := generateData([]string{"A", "B", "C", "D"}, rowCount, false)
	dataB := generateData([]string{"A", "D", "E"}, rowCount, true)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		mergeCSV(dataA, dataB)
	}
}

func generateData(headers []string, rowCount int, shuffledRow bool) [][]string {
	rowCount++
	out := make([][]string, rowCount)
	randIdxs := rand.Perm(rowCount - 1)

	for y, _ := range out {
		if y == 0 {
			out[y] = headers
			continue
		}

		var rowIdx int
		if shuffledRow {
			rowIdx = randIdxs[y-1] + 1
		} else {
			rowIdx = y
		}
		var row = make([]string, len(headers))
		for x, v := range headers {
			row[x] = fmt.Sprintf("%d%s", rowIdx, v)
		}
		out[y] = row
	}

	return out
}
