package main

import "fmt"


func bubbleSort(arr []int) []int{
	len := len(arr)
	for i:=0; i<len-1; i++ {
		for j:=0; j<len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func main2() {
	arr1:=[]int{1,2,3,4,3,8}
	arr1 = bubbleSort(arr1[:])



	for i:=0; i< len(arr1)-1; i++ {
			if arr1[i] == arr1[i+1] {
				fmt.Println(arr1[i])
			}

	}
}