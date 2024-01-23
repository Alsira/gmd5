package main

import (
	"fmt"
	"crypto/md5"
	"flag"
	"os"
	"io"
	"encoding/hex"
)

func main() {

	// Setup the command line parsing
	var filepath string

	flag.StringVar(&filepath, "file", "", "This is the file to hash")
	flag.Parse()

	if !flag.Parsed() {
		fmt.Printf("Error: cannot parse :(")
		os.Exit(-1)
	}

	if filepath == "" {
		fmt.Printf("Error: no file entered :(")
		os.Exit(-1)
	}

	// Use the file to start calculating the hash
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer file.Close()

	// Otherwise, we have an open file, stream it into the hasher
	hash := md5.New()
	io.Copy(hash, file)

	// Get the hashed data
	data := hash.Sum(nil)

	// Output the hash and exit
	fmt.Printf("MD5 hash is : %s\n", hex.EncodeToString(data))
	os.Exit(0)
}
