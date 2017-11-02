package main

import (
    "testing"
)

func BenchmarkBsonInt (b *testing.B) {
    ba := make([]byte, 0, 500)
    for i:=0;i<b.N;i++ {
        BsonAppendInt(ba, "Int", 10)
    }
}

func BenchmarkBsonBool (b *testing.B) {
    ba := make([]byte, 0, 500)
    for i:=0;i<b.N;i++ {
        BsonAppendBool(ba, "Bool", true)
    }
}

func BenchmarkBsonStr (b *testing.B) {
    ba := make([]byte, 0, 500)
    for i:=0;i<b.N;i++ {
        BsonAppendString(ba, "Str", "ValueStr")
    }
}
