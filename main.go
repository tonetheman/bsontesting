package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type bsonDocument struct {
	doclen int32
	// followed by lots of stuff

	eoo byte // needs to be null
}

type bsonString struct {
	bson_type byte
	// followed by a variable len cstring
	// which is a modified utf8 string
	// no null bytes
	// followed by a single null byte
	key []byte
	// followed by string
	val []byte
}

func (b bsonString) cvt2json() string {
	return ""
}

func newBsonString(k string, v string) bsonString {
	return bsonString{bson_type: 2, key: stringTocstring(k),
		val: stringToBsonString(v)}
}

// creates a cstring
// string with a null byte at the end
func stringTocstring(s string) []byte {
	tmpB := new(bytes.Buffer)
	tmpB.WriteString(s)
	tmpB.WriteByte(0)
	return tmpB.Bytes()
}

// creates a len prefixed bson string
func stringToBsonString(s string) []byte {
	ll := int32(len(s))
	// add 1 for the null at the end
	ll++
	tmpB := new(bytes.Buffer)
	binary.Write(tmpB, binary.LittleEndian, &ll)
	tmpB.WriteString(s)
	tmpB.WriteByte(0)
	return tmpB.Bytes()
}

func testWrite() {
	var i int32 = 888
	buffer := new(bytes.Buffer) // 0 len but read to write
	binary.Write(buffer, binary.LittleEndian, &i)
	fmt.Printf("% x\n", buffer.Bytes())
}

func main() {
	//testWrite()
	//b := stringTocstring("tony")
	//fmt.Printf("res % x\n", b)
	//b2 := stringToBsonString("tony")
	//fmt.Printf("res2 % x\n", b2)
	b := newBsonString("tony", "iscool")
	fmt.Println(b)
}
