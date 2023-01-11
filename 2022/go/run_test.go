package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkDay(b *testing.B) {
	for d := 0; d < len(INPUTS); d++ { // for each day
		s, err := os.ReadFile(INPUTS[d])
		if err != nil {
			panic(err)
		}
		b.Run(fmt.Sprintf("DAY_%d", d+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ { // running it a 1000 times
				runDay(d, FUNCTIONS[d], string(s))
			}
		})
	}
}
