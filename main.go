package main

import (
	"fmt"
	"strings"
	"time"
)

var base = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var baseLength = len(base)

func decode(encoding string) int {
	numID := 0

	for i := 0; i < len(encoding); i++ {
		numID = numID*baseLength + strings.Index(base, string(encoding[i]))
	}

	return numID
}

func encode(numID int) string {
	var encodedValueBuffer strings.Builder

	for numID > 0 {
		encodedValueBuffer.WriteString(string(base[numID%baseLength]))
		numID /= baseLength
	}

	return encodedValueBuffer.String()
}

func main() {
	start := time.Now()

	encoded := encode(91381)
	fmt.Printf("%s encoded\n", encoded)
	fmt.Printf("%d decoded\n", decode(encoded))

	elapsed := time.Since(start)
	fmt.Printf("Encoding took %s\n", elapsed)
}
