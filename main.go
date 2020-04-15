package main

import (
	"encoding/hex"
	"io"
	"log"
	"os"

	"github.com/minio/highwayhash"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) != 2 {
		log.Fatalln("usage: highwayhashsum <file>")
	}

	filename := os.Args[1]

	fileinfo, err := os.Stat(filename)
	if err != nil {
		log.Fatalln(err)
	}
	if !fileinfo.Mode().IsRegular() {
		log.Fatalln("Not a regular file: ", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Failed to open file: ", err.Error())
	}
	defer file.Close()

	key := []byte("12345678901234567890123456789012")
	hash, err := highwayhash.New(key)
	if err != nil {
		log.Fatalln("Failed to create HighwayHash instance: ", err.Error())
	}

	if _, err = io.Copy(hash, file); err != nil {
		log.Fatalln("Failed to read from file: ", err.Error())
	}

	checksum := hash.Sum(nil)
	log.Printf("%v  %v\n", hex.EncodeToString(checksum), filename)
}
