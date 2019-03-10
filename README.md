# Mergecsv
> Merge two csv table by matching common fields

Usage: `mergecsv file1.csv file2.csv > merged.csv`

# Install

`go get -u github.com/lincerely/mergecsv`

# Benchmark

```
goos: darwin
goarch: amd64
pkg: github.com/lincerely/mergecsv
BenchmarkMergeCSV1-8             2000000               964 ns/op
BenchmarkMergeCSV10-8             300000              4734 ns/op
BenchmarkMergeCSV100-8             50000             38401 ns/op
BenchmarkMergeCSV1000-8             3000            403214 ns/op
BenchmarkMergeCSV10000-8             200           6181605 ns/op
PASS
ok      github.com/lincerely/mergecsv   9.866s
```
