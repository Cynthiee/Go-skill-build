package main

import (
	"fmt"

	"piscine"
)

func main() {
	fmt.Println(piscine.CountAlpha("Hello world"))
	fmt.Println(piscine.CountAlpha("H e l l o"))
	fmt.Println(piscine.CountAlpha("H1e2l3l4o"))
	fmt.Printf("*******\n")
	fmt.Println(piscine.CountChar("Hello World", 'l'))
	fmt.Println(piscine.CountChar("5  balloons", 5))
	fmt.Println(piscine.CountChar("   ", ' '))
	fmt.Println(piscine.CountChar("The 7 deadly sins", '7'))
	fmt.Printf("*******\n")
	fmt.Print(piscine.PrintIf("abcdefz"))
	fmt.Print(piscine.PrintIf("abc"))
	fmt.Print(piscine.PrintIf(""))
	fmt.Print(piscine.PrintIf("14"))
	fmt.Printf("*******\n")
	fmt.Print(piscine.PrintIfNot("abcdefz"))
	fmt.Print(piscine.PrintIfNot("abc"))
	fmt.Print(piscine.PrintIfNot(""))
	fmt.Print(piscine.PrintIfNot("14"))
}
