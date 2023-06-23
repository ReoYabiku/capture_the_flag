package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("ksnctf/onion.txt")
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 16; i++ {
		decoded, err := base64.StdEncoding.DecodeString(string(data))
		if err != nil {
			fmt.Println(err)
		}
		data = decoded
	}

	fmt.Println(string(data))
}
