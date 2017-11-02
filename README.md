# bson-eval
Code to compare writing in JSON vs BSON

Results:

```
$ go test -bench .
goos: darwin
goarch: amd64
pkg: json-bson
BenchmarkBsonInt-8    	500000000	         3.63 ns/op
BenchmarkBsonBool-8   	500000000	         3.04 ns/op
BenchmarkBsonStr-8    	200000000	         5.80 ns/op
BenchmarkJsonInt-8    	100000000	        22.6 ns/op
BenchmarkJsonBool-8   	100000000	        14.4 ns/op
BenchmarkJsonStr-8    	50000000	        25.1 ns/op
PASS
ok  	json-bson	10.866s
```
