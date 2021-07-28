package iteration

const repeatCount = 5

func Repeat(s string) string {
	var repeat string

	for i := 0; i < repeatCount; i++ {
		repeat += s
	}

	return repeat
}
