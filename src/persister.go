package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var lock sync.Mutex

// Marshal is a function that marshals the object into an io.Reader.
// By default, it uses the JSON marshaller which will turn the v object
// into nicely formatted JSON for us
var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// Unmarshal is a function that unmarshals the data from the reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// Save saves a representation of v to the file at path.
func Save(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func Load(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return Unmarshal(f, v)
}

type obj struct {
	Name   string
	Number int
	When   time.Time
}

func main() {
	o := &obj{
		Name:   "Mat",
		Number: 47,
		When:   time.Now(),
	}
	if err := Save("./file.tmp", o); err != nil {
		log.Fatalln(err)
	}
	// load it back
	var o2 obj
	if err := Load("./file.tmp", &o2); err != nil {
		log.Fatalln(err)
	}
	// o and o2 are now the same
	// and check out file.tmp - you'll see the JSON file
	fmt.Println(o2)
}
