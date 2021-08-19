package iteration

import (
	"fmt"
	"strings"
)

func Repeat(char string, repititions int) string {
	var repeated string
	for i := 0; i < repititions; i++ {
		repeated += char
	}
	return repeated
}

func main() {
	fmt.Println(Repeat("test", 5))
}
