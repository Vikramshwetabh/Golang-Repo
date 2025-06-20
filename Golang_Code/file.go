package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	// file, err := os.Open("example.txt") // Open the file in the files directory

	// if err != nil {
	// 	panic(err) // Handle the error if the file cannot be opened
	// }

	// fileInfo, err := file.Stat() // Get the file information

	// if err != nil {
	// 	panic(err) // Handle the error if file information cannot be retrieved
	// }
	// fmt.Println("file name", fileInfo.Name())                 // Print the file name
	// fmt.Println("file size", fileInfo.Size())                 // Print the file size
	// fmt.Println("file mode", fileInfo.Mode())                 // Print the file mode (permissions)
	// fmt.Println("file is directory", fileInfo.IsDir())        // Check if the file is a directory
	// fmt.Println("file modification time", fileInfo.ModTime()) // Print the last modification time of the file

	//reading a file
	// f, err := os.Open("example.txt") // Open the file in the files directory
	// if err != nil {
	// 	panic(err) // Handle the error if the file cannot be opened
	// }

	// defer f.Close()
	// buf := make([]byte, 12) //created a array

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i])) // Print the data read from the file . string(buf[i]) converts the byte to a string
	// }

	//2nd method to read a file. Only works for small files
	// file, err := os.ReadFile("example.txt") // Read the entire file into memory
	// if err != nil {
	// 	panic(err) // Handle the error if the file cannot be read
	// }
	// fmt.Println("file content:", string(file)) // Print the content of the file as a string

	//read a folder
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close() // Close the directory when done

	// fileInfo, err := dir.Readdir(-1) // Read all files in the directory
	// for _, f1 := range fileInfo {
	// 	fmt.Println("file name:", f1.Name()) // Print the name of each file in the directory
	// }

	// //creating a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// f.WriteString("hellllo2")

	//reading and writing in another file

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	desinationFile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer desinationFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(desinationFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}
	}
	writer.Flush()                      // Flush the buffered writer to ensure all data is written to the file
	println("file copied successfully") // Print a success message

	// delete a file

	// sourceFile, err := os.Open("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// err := os.Remove("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(" file deleted successfully")

}
