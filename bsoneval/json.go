package bsoneval

import (
    "fmt"
    "strconv"
)

func AppendKey(dst []byte, key string) []byte {
    if len(dst) > 1 {
        dst = append(dst, ',')
    }
    dst = AppendString(dst, key)
    return append(dst, ':')
}

func AppendBool(dst []byte, val bool) []byte {
    return strconv.AppendBool(dst, val)
}

func AppendInt(dst []byte, val int) []byte {
    return strconv.AppendInt(dst, int64(val), 10)
}

func AppendString(dst []byte, s string) []byte {
    // Start with a double quote.
    dst = append(dst, '"')
    // Loop through each character in the string.
    for i := 0; i < len(s); i++ {
        //Special string processing - avoided the copy over to here..
        _ = i
    }
    // The string has no need for encoding an therefore is directly
    // appended to the byte slice.
    dst = append(dst, s...)
    // End with a double quote
    return append(dst, '"')
}

func Json_main () {
    ba := make([]byte, 0, 500)
    ba = append(ba, 0)

    ba = AppendInt(AppendKey(ba, "Int"), 10)
    fmt.Println(string(ba[1:]))
    ba = AppendBool(AppendKey(ba, "Bool"), true)
    fmt.Println(string(ba[1:]))
    ba = AppendString(AppendKey(ba, "Str"), "ValueStr")
    fmt.Println(string(ba[1:]))
}
