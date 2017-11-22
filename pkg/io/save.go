package io

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

const (
	saveFileName = "larn.sav" // To be prepended with a unique id
)

// randGen for generating save file unique random numbers
var randGen *rand.Rand

func init() {
	// Setup the random number generator
	randGen = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// NewGame creeates a new game file with an initial time and seed
func NewGame() (string, error) {
	ts := time.Now()
	r := randGen.Uint64()
	data := make([]byte, 10)

	// Save the timestamp into data
	binary.LittleEndian.PutUint64(data, uint64(ts.UnixNano()))

	filename := fmt.Sprintf("%v-%s", r, saveFileName)
	return filename, ioutil.WriteFile(filename, data, 0644)
}
