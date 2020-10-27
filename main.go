package main

import (
	"encoding/hex"
	"io"
	"log"
	"os"
	"sync"

	"github.com/minio/highwayhash"
)

var (
	rc  int
	key = []byte("12345678901234567890123456789012")
	wg  = new(sync.WaitGroup)
)

func main() {

	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatalln("usage: highwayhashsum <file> [<file> ...]")
	}

	filenames := os.Args[1:]

	for _, filename := range filenames {
		fileinfo, err := os.Stat(filename)
		if err != nil {
			log.Println(err)
			rc = 1
			continue
		}

		if !fileinfo.Mode().IsRegular() {
			log.Println("Not a regular file: ", filename)
			rc = 1
			continue
		}

		wg.Add(1)
		go hash(filename)
	}

	wg.Wait()
	os.Exit(rc)
}

func hash(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Println("Failed to open file: ", err.Error())
		rc = 1
		wg.Done()
		return
	}
	defer file.Close()

	hash, err := highwayhash.New(key)
	if err != nil {
		log.Println("Failed to create HighwayHash instance: ", err.Error())
		rc = 1
		wg.Done()
		return
	}

	if _, err = io.Copy(hash, file); err != nil {
		log.Println("Failed to read from file: ", err.Error())
		rc = 1
		wg.Done()
		return
	}

	checksum := hash.Sum(nil)
	log.Printf("%v  %v\n", hex.EncodeToString(checksum), filename)
	wg.Done()
}
