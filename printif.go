package piscine

// Write a function that takes a string as an argument and returns the letter G followed by a newline \n if the argument length is more or equal than 3, otherwise returns Invalid Input followed by a newline \n.

// If it's an empty string return G followed by a newline \n.

func PrintIf(str string) string {
	if len(str) != 0 || len(str) >= 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}
