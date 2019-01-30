package main

import (
	"log"
	"strings"

	"github.com/aliensteam51/functional"
)

func add3(a int, b int, c int) int {
	return a + b + c
}

func main() {
	// Specific definitions
	var intMap func(func(int) int, []int) []int
	var floatFilter func(func(float64) bool, []float64) []float64
	var stringReduce func(func(string, string) string, string, []string) string
	var intReduce func(func(int, int) int, int, []int) int
	var intPartial func(func(int, int, int) int, int, int) func(int) int
	var stringPipe func(...interface{}) func(string) string
	var intCurry func(func(int, int, int) int) func(int) func(int) func(int) int
	var stringCompose func(func(string) string, func(string) string) func(string) string

	functional.MakeSpecificFunc(&floatFilter, functional.GenericFilter)
	functional.MakeSpecificFunc(&intMap, functional.GenericMap)
	functional.MakeSpecificFunc(&stringReduce, functional.GenericReduce)
	functional.MakeSpecificFunc(&intReduce, functional.GenericReduce)
	functional.MakeSpecificFunc(&intPartial, functional.GenericPartial)
	functional.MakeSpecificFunc(&stringPipe, functional.GenericPipe)
	functional.MakeSpecificFunc(&intCurry, functional.GenericCurry)
	functional.MakeSpecificFunc(&stringCompose, functional.GenericCompose)

	result := intMap(func(i int) int {
		return int(i * 2.0)
	}, []int{1, 2, 3, 4})
	log.Printf("intMap %v\n", result)

	fresult := floatFilter(func(x float64) bool {
		return x > 1.0
	}, []float64{0.1, 1.1, 1.02, -1.0})
	log.Printf("floatFilter %v\n", fresult)

	sresult := stringReduce(func(acc string, next string) string {
		return acc + next
	}, "begin", []string{" and middle", " and end"})
	log.Printf("stringReduce %v\n", sresult)

	iresult := intReduce(func(acc int, next int) int {
		return acc + next
	}, 0, []int{1, 2})
	log.Printf("intReduce %v\n", iresult)

	partialResult := intPartial(add3, 1, 2)
	log.Printf("presult %v\n", partialResult(3))

	pipe := stringPipe(strings.TrimSpace, strings.ToLower)
	pipeResult := pipe("  do it now DDDKDKDKKD   ")
	log.Printf("pipe %v result\n", pipeResult)

	curry := intCurry(add3)
	curryNext1 := curry(1)
	curryNext2 := curryNext1(2)
	cresult := curryNext2(3)
	log.Printf("curry %v result\n", cresult)

	compose := stringCompose(strings.ToLower, strings.TrimSpace)
	log.Printf("compose %v result\n", compose("WHAT IS DAT ccccc    "))
}
