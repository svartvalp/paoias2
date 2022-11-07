package main

import (
	"github.com/svartvalp/paoias2/internal/filter"
	"github.com/svartvalp/paoias2/internal/writer"
)

func main() {
	f := filter.New(
		[]int{6, 0, -4, -3, 5, 6, -6, -13, 7, 44, 64, 44, 7, -13, -6, 6, 5, -3, -4, 0, 6},
		100,
		21,
		2,
		0.1,
	)
	res := f.Run()
	wr := writer.New()
	err := wr.Write("res.txt", res)
	if err != nil {
		panic(err)
	}
}
