package main

import (
	"encoding/binary"
	"os"
)

const (
	STATE_FILE_PATH = "/tmp/gspc"
)

/*
 SERIALIZATION
 -------------

 This is very basic because we only need to store two ints. The first
 8 bytes are an int64 representing the current space, and the next 8
 bytes are an int64 representing the previous space.

*/

func setCurrentSpace(space uint) error {
	prev, err := currentSpace()
	if err != nil {
		prev = 1
	}

	f, err := os.Create(STATE_FILE_PATH)
	if err != nil {
		return err
	}
	defer f.Close()

	binary.Write(f, binary.BigEndian, uint64(space))
	binary.Write(f, binary.BigEndian, uint64(prev))

	return nil
}

func prevSpace() (uint, error) {
	b, err := os.ReadFile(STATE_FILE_PATH)
	if err != nil {
		return 0, err
	}

	return uint(binary.BigEndian.Uint64(b[8:])), nil
}

func currentSpace() (uint, error) {
	b, err := os.ReadFile(STATE_FILE_PATH)
	if err != nil {
		return 0, err
	}

	return uint(binary.BigEndian.Uint64(b[:8])), nil
}
