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
	// this is not what i thought it was
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
	buffer := new(bytes.Buffer) // 0 len but ready to write
	binary.Write(buffer, binary.LittleEndian, &i)
	fmt.Printf("% x\n", buffer.Bytes())
}

// modified utf 8
// see here: https://docs.oracle.com/javase/8/docs/api/java/io/DataInput.html#modified-utf-8
// http://grepcode.com/file/repository.grepcode.com/java/root/jdk/openjdk/6-b14/java/io/DataOutputStream.java
func pr(r0 rune) []byte {
	res := make([]byte, 0)
	if r0 >= 0x0001 && r0 <= 0x007f {
		fmt.Println("case1")
		a := r0
		fmt.Println(a)
		res = append(res, byte(a))
	} else if (r0 >= 0x0080 && r0 <= 0x07ff) || r0 == 0x0000 {
		fmt.Println("case2")
		a := (0xC0 | ((r0 >> 6) & 0x1F))
		b := (0x80 | ((r0 >> 0) & 0x3F))
		fmt.Println(a, b)
		res = append(res, byte(a))
		res = append(res, byte(b))
	} else {
		fmt.Println("case3")
		a := (0xE0 | ((r0 >> 12) & 0x0F))
		b := (0x80 | ((r0 >> 6) & 0x3F))
		c := (0x80 | ((r0 >> 0) & 0x3F))
		fmt.Println(a, b, c)
		res = append(res, byte(a))
		res = append(res, byte(b))
		res = append(res, byte(c))
	}
	return res
}

func readModUTF8(b []byte) {
	var bunk [4]byte
	counter := 0
	for i := 0; i < len(b); i++ {
		c := b[i] & 0xff
		tmp := c >> 4
		if tmp >= 0 && tmp < 7 {
			bunk[counter] = c
			counter++
		} else if tmp == 12 || tmp == 13 {

		}
	}
}

func testpr() {
	const placeOfInterest = `⌘`
	runes := []rune(placeOfInterest)
	fmt.Println(runes)
	myBytes := pr(runes[0])
	readModUTF8(myBytes)
}

func main() {
	//testWrite()
	//b := stringTocstring("tony")
	//fmt.Printf("res % x\n", b)
	//b2 := stringToBsonString("tony")
	//fmt.Printf("res2 % x\n", b2)
	b := newBsonString("tony", "iscool")
	fmt.Println(b)

	myBytes := testpr()
}
