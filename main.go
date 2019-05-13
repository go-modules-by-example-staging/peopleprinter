package main

import (
	"fmt"
	"github.com/go-modules-by-example-staging/goinfo/designers"
	"strings"
)

func main() {
	fmt.Printf("The designers of Go: %v\n", strings.Join(designers.Names(), ", "))
}
