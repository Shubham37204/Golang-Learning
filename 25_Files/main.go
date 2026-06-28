package main

import (
	"fmt"
	"os"
)

func main() {

	//read the file
	// f, err := os.Open("a.txt") //Opens the file in read-only mode.
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close() //defer:Run this just before the function exits.

	// buff := make([]byte, 10) //This creates a slice of bytes.
	// //Why byte? Because computers read files as bytes.

	// d, err := f.Read(buff) //This reads data from the file into the buffer.
	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < len(buff); i++ {
	// 	println(d, string(buff[i]))
	// }

	// f,err := os.ReadFile("a.txt")
	// 	if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf(string(f))

	//read folders
	// dir, err := os.Open(".")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileinfo,err :=dir.ReadDir(1)
	// for _,f1 := range fileinfo{
	// 	fmt.Println(f1.Name())
	// }

	//create a file
	// f, err := os.Create("b.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	//f.WriteString("hello world")
	// bytes := []byte("hello golang")
	// f.Write(bytes)

	//read and write to another file
	// src, err := os.Open("a.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer src.Close()

	// dest, err := os.Create("b.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dest.Close()

	// reader := bufio.NewReader(src)
	// writer := bufio.NewWriter(dest)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(err)
	// 	}
	// }
	// writer.Flush()
	// fmt.Println("all done")

	//delete a file
	err := os.Remove("b.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("file deleted")

}
