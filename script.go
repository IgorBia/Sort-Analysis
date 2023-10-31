package main

import ("fmt")

func insertSort(list []int) []int {
	var j=1

	for j<(len(list)){
		key := list[j]
		i:=j-1
		for ;(i > -1) && (list[i]> key);{
				list[i+1]=list[i]
				i-=1
		}
		list[i+1] = key
		j+=1
	}
	return list
}

func bubbleSort(list []int) []int {
	var change=true
	for ;change==true;{
		change = false
		for i:=0;i<len(list)-1;i++{
			key:=list[i]
			if list[i]>list[i+1]{
				list[i]=list[i+1]
				list[i+1]=key
				change=true
			}
		}
	}
	return list
}

func main() {
	var tablica []int = []int{100,23,45,12,6785,3}
	fmt.Println("Starting with: ", tablica, "\n")
	fmt.Println("Ending with: ", bubbleSort(tablica))
}
