package main 

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main(){

	/*
	Step :-
		Create a file 
		open a file if error then print it out 
		Write text inside that file
		Read the data from the file
	*/


	file, err := os.Create("sample.txt")

	if err != nil{
		log.Fatal(err)
	}

	file.WriteString("This is sample text.")

	defer file.Close()

	stream, err := ioutil.ReadFile("sample.txt")

	if err != nil{
		log.Fatal(err)
	}

	readString := string(stream)
	fmt.Println(readString)

}