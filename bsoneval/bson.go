package bsoneval

import (
	"fmt"
	"strconv"
)

const (
	BsonInt32Type = 0x10
	BsonBoolType  = 0x08
	BsonStrType   = 0x02
)

func BsonAppendBool(dst []byte, key string, v bool) []byte {
	dst = append(dst, BsonBoolType)
	dst = append(dst, key...)
	dst = append(dst, 0)
	if v {
		dst = append(dst, 0x01)
	} else {
		dst = append(dst, 0x00)
	}
	return dst
}

func BsonAppendInt(dst []byte, key string, v int) []byte {
	dst = append(dst, BsonInt32Type)
	dst = append(dst, key...)
	dst = append(dst, 0)
	dst = append(dst, byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
	return dst
}

func BsonAppendString(dst []byte, k string, v string) []byte {
	dst = append(dst, BsonStrType)
	dst = append(dst, k...)
	dst = append(dst, 0)
	l := len(v)
	dst = append(dst, byte(l), byte(l>>8), byte(l>>16), byte(l>>24))
	dst = append(dst, v...)
	dst = append(dst, 0)
	return dst
}

func decodeBsonInt(src []byte, dst *[]byte) int {
	bytesDecoded := 0
	sz := len(src)
	opCode := src[0]
	if opCode != BsonInt32Type {
		return bytesDecoded
	}
	if len(*dst) > 0 {
		*dst = append(*dst, ',')
	}
	*dst = append(*dst, '"')
	i := 1 //consume the opcode here
	bytesDecoded++
	for ; byte(src[i]) != 0x0 && i < sz; i++ {
		*dst = append(*dst, src[i])
		bytesDecoded++
	}
	if i == sz || i+5 > sz {
		//we have reached the end or we do NOT have enough bytes
		//in the trailer..
		//Error - shd really take out whatever we have added to dst
		return 0
	}
	*dst = append(*dst, '"', ':')
	i++ //Consume the string termination '\0'
	bytesDecoded++
	val := int32(src[i]) + int32(src[i+1])<<8 + int32(src[i+2])<<16 + int32(src[i+3])<<24
	*dst = strconv.AppendInt(*dst, int64(val), 10)
	bytesDecoded += 4
	return bytesDecoded
}

func decodeBsonBool(src []byte, dst *[]byte) int {
	bytesDecoded := 0
	sz := len(src)
	opCode := src[0]
	if opCode != BsonBoolType {
		return bytesDecoded
	}
	if len(*dst) > 0 {
		*dst = append(*dst, ',')
	}
	*dst = append(*dst, '"')
	i := 1 //consume the opcode here
	bytesDecoded++
	for ; byte(src[i]) != 0x0 && i < sz; i++ {
		*dst = append(*dst, src[i])
		bytesDecoded++
	}
	if i == sz || i+1 > sz {
		//we have reached the end or we do NOT have enough bytes
		//in the trailer..
		//Error - shd really take out whatever we have added to dst
		return 0
	}
	*dst = append(*dst, '"', ':')
	i++ //Consume the string termination '\0'
	bytesDecoded++
	if src[i] == 0x0 {
		*dst = append(*dst, []byte("false")...)
	} else {
		//Slight deviation from standard - we are assuming
		//any non-zero as true
		*dst = append(*dst, []byte("true")...)
	}
	bytesDecoded++
	return bytesDecoded
}

func decodeBsonStr(src []byte, dst *[]byte) int {
	bytesDecoded := 0
	sz := len(src)
	opCode := src[0]
	if opCode != BsonStrType {
		return bytesDecoded
	}
	if len(*dst) > 0 {
		*dst = append(*dst, ',')
	}
	*dst = append(*dst, '"')
	i := int(1) //consume the opcode here
	bytesDecoded++
	for ; byte(src[i]) != 0x0 && i < sz; i++ {
		*dst = append(*dst, src[i])
		bytesDecoded++
	}
	if i == sz || i+4 > sz {
		//we have reached the end or we do NOT have enough bytes
		//in the trailer..
		//Error - shd really take out whatever we have added to dst
		return 0
	}
	*dst = append(*dst, '"', ':')
	i++ //Consume the string termination '\0'
	bytesDecoded++
	valLen := int32(src[i]) + int32(src[i+1])<<8 + int32(src[i+2])<<16 + int32(src[i+3])<<24
	bytesDecoded += 4
	i += 4
	if i+int(valLen) > sz {
		//Length of the Value is larger than the src !!
		return 0
	}
	*dst = append(*dst, '"')
	for ; byte(src[i]) != 0x0 && i < sz; i++ {
		*dst = append(*dst, src[i])
		bytesDecoded++
		valLen--
		if valLen < 0 {
			//More characters than expected ??
			return 0
		}
	}
	*dst = append(*dst, '"')
	bytesDecoded++ // terminating \0
	return bytesDecoded
}

func decodeBson(src []byte) (string, error) {
	result := make([]byte, 0, 500)
	var err error
	curPos := 0
	for {
		if curPos >= len(src) || err != nil {
			//No more Records
			break
		}
		opCode := src[curPos]
		switch opCode {
		case BsonInt32Type:
			curPos += decodeBsonInt(src[curPos:], &result)
		case BsonBoolType:
			curPos += decodeBsonBool(src[curPos:], &result)
		case BsonStrType:
			curPos += decodeBsonStr(src[curPos:], &result)
		default:
			err = fmt.Errorf("invalid Opcode found %X", opCode)
		}
	}
	return string(result), err
}

func Bson_main() {
	ba := make([]byte, 0, 500)
	ba = append(ba, 0)

	ba = BsonAppendInt(ba, "Keeyyy", 642121)
	//fmt.Printf("%q\n", ba[1:])

	ba = BsonAppendInt(ba, "Int", 64)
	//fmt.Printf("%q\n", ba[1:])

	ba = BsonAppendBool(ba, "Bool", true)
	//fmt.Printf("%q\n", ba[1:])

	ba = BsonAppendBool(ba, "BadCode", false)
	//fmt.Printf("%q\n", ba[1:])

	ba = BsonAppendString(ba, "Str", "ValueStr")
	//fmt.Printf("%q\n", ba[1:])

	ba = BsonAppendString(ba, "MoreStr", "MoreValueStr")
	//fmt.Printf("%q\n", ba[1:])
	ba = BsonAppendInt(ba, "Keeyyy", 642121)
	//fmt.Printf("%q\n", ba[1:])

	ba = BsonAppendInt(ba, "Int", 64)
	//fmt.Printf("%q\n", ba[1:])

	ba = BsonAppendBool(ba, "Bool", true)
	//fmt.Printf("%q\n", ba[1:])

	ba = BsonAppendBool(ba, "BadCode", false)
	//fmt.Printf("%q\n", ba[1:])
	fmt.Println(decodeBson(ba[1:]))

}
