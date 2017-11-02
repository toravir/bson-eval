# bson-eval
Code to compare writing in JSON vs BSON

Results:

```
$ go test -bench . -benchmem -v
goos: darwin
goarch: amd64
pkg: json-bson/bson-eval
BenchmarkBsonInt-8    	500000000	         3.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkBsonBool-8   	500000000	         3.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkBsonStr-8    	200000000	         6.04 ns/op	       0 B/op	       0 allocs/op
BenchmarkJsonInt-8    	50000000	        24.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkJsonBool-8   	100000000	        14.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkJsonStr-8    	50000000	        25.6 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	json-bson/bson-eval	9.981s
```
