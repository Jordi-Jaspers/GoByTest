package iteration

const REPITITIONS = 5

func Repeat(char string) string {
	var repeated string
	for i := 0; i < REPITITIONS; i++ {
		repeated += char
	}
	return repeated
}
