package main

import "fmt"

func main() {
	var tinggi int8 = 10
	var sisiA int8 = 8
	var sisiB int8 = 14

	var luasTrapesium int16 = int16(tinggi) * (int16(sisiA) + int16(sisiB)) / 2

	fmt.Println("====================================")
	fmt.Printf("Jadi, luas dari trapesium adalah : %d\n", luasTrapesium)
	fmt.Println("====================================")
}
