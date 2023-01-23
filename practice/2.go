package main

import "fmt"

func biggestLenOfSubString(str string) int {
	len := 0
	
	for i:=0; i<len(str); i++ {
		if str[i] != str[i+1] {
			len++
		}
		else{
			
		}
	}

	return max

}

func main() {
	str := "abeccdd"
	str2 := "abeccddcexya"

	fmt.Println(biggestLenOfSubString(str))
	fmt.Println(biggestLenOfSubString(str2))

}