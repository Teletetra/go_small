package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	// fmt.Println("yoo yo uo")
	f, err := os.Open("message_read/messages.txt")

	if err != nil {
		log.Fatal("error", "error", err)
	}
	str := ""
	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			break
		}
		data = data[:n]
		// explained?
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			str += string(data[:i])
			data = data[i+1:]
			fmt.Printf("read:%s\n", str)
			str = ""
		}
		fmt.Printf("reead line : %s\n", string(data[:n]))
	}
}

// what is make in file handling
// panic: runtime error: slice bounds out of range [:8] with capacity 4

// goroutine 1 [running]:
// main.main()
//         D:/new_lang/GoLang/real_proj/network/message_read/main.go:32 +0x2d6
// exit status 2
