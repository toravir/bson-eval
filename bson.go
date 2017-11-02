package main

import (
    "fmt"
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
    return dst
}

func bson_main () {
    ba := make([]byte, 0, 500)
    ba = append(ba, 0)

    ba = BsonAppendInt(ba, "Int", 64)
    fmt.Printf("%q\n",ba[1:])

    ba = BsonAppendBool(ba, "Bool", true)
    fmt.Printf("%q\n",ba[1:])

    ba = BsonAppendString(ba, "Str", "ValueStr")
    fmt.Printf("%q\n",ba[1:])
}
