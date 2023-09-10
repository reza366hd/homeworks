package twentysix

import (
	"log"
	"os"
)

func twentysix() {

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 6)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}
