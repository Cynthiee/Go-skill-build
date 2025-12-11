package piscine

// Write a function that takes a string and returns:

// "G\n" if the string's length is less than 3 (including empty string).

// "Invalid Input\n" otherwise.

func PrintIfNot(str string) string {
	if len(str) == 0 || len(str) < 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}
