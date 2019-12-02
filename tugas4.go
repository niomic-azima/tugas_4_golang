package main

import "fmt"

func main() {
	var tinggiMahasiswa = map[string]int{
		"aldo":  182,
		"yosep": 178,
	}
	fmt.Println("Aldo :", tinggiMahasiswa["aldo"], "cm")
	fmt.Println("Yosep :", tinggiMahasiswa["yosep"], "cm")
}
