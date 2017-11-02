package bsoneval

import (
    "testing"
)

func BenchmarkJsonInt (b *testing.B) {
    ba := make([]byte, 0, 500)
    for i:=0;i<b.N;i++ {
        AppendInt(AppendKey(ba, "Int"), 10)
    }
}

func BenchmarkJsonBool (b *testing.B) {
    ba := make([]byte, 0, 500)
    for i:=0;i<b.N;i++ {
        AppendBool(AppendKey(ba, "Bool"), true)
    }
}

func BenchmarkJsonStr (b *testing.B) {
    ba := make([]byte, 0, 500)
    for i:=0;i<b.N;i++ {
        AppendString(AppendKey(ba, "StringKey"), "StringValue")
    }
}
