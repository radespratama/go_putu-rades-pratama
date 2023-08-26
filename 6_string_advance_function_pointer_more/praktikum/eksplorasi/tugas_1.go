package main

import (
	"fmt"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode string

	for _, char := range s.name {
		encodedChar := char + 1
		nameEncode += string(encodedChar)
	}

	s.nameEncode = nameEncode
	return s.nameEncode
}

func (s *student) Decode() string {
	var nameDecode string

	for _, char := range s.nameEncode {
		decodedChar := char - 1
		nameDecode += string(decodedChar)
	}

	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student`s name " + a.name + "is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student`s name " + a.name + "is : " + c.Decode())
	}
}
