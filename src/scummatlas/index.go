package scummatlas

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type RoomIndex struct {
	number int
	offset []byte
}

type ScriptIndex struct {
	roomNumber int
	roomOffset int
}

func ReadXoredFile(fileName string, code byte) (out []byte, err error) {
	out, err = ioutil.ReadFile(fileName)
	for i, _ := range out {
		out[i] = out[i] ^ 0x69
	}
	return out, err
}

func complementaryByte(in []byte) (out []byte) {
	for i, _ := range in {
		in[i] = in[i] ^ 0xFF
	}
	return in
}

func ParseRoomNames(data []byte) []string {
	var names []string
	if string(data[0:4]) == "RNAM" {
		currentIndex := 8
		for currentIndex < len(data) {
			roomNumber := int(data[currentIndex])
			if roomNumber == 0 {
				break
			}
			name := data[currentIndex+1 : currentIndex+10]
			names = append(names, string(complementaryByte(name)))
			currentIndex += 10
		}
	}
	return names
}

func ParseRoomIndex(data []byte) (index []ScriptIndex) {
	var out []ScriptIndex
	numEntries := int(binary.LittleEndian.Uint16(data[8:10]))
	fmt.Println("Num entries: ", numEntries)

	currentIndex := 10
	for currentIndex < len(data) {
		roomNumber := int(data[currentIndex])
		roomOffset := int(binary.LittleEndian.Uint16(data[currentIndex+1 : currentIndex+4]))
		out = append(out, ScriptIndex{roomNumber, roomOffset})
		currentIndex += 5
	}
	return out
}

func ParseScriptsIndex(data []byte) (index []ScriptIndex) {
	var out []ScriptIndex
	numEntries := int(binary.LittleEndian.Uint16(data[8:10]))
	fmt.Println("Num entries: ", numEntries)

	currentIndex := 10
	for currentIndex < len(data) {
		roomNumber := int(data[currentIndex])
		roomOffset := int(binary.LittleEndian.Uint16(data[currentIndex+1 : currentIndex+4]))
		out = append(out, ScriptIndex{roomNumber, roomOffset})
		currentIndex += 5
	}
	return out
}