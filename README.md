A tiny webapp which generates UUID version 7.



## Installation

To install Gin package, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed (**version 1.13+ is required**), then you can use the below Go command to install Gin and others packages

```sh
$ go mod vendor
```
2. Then start the webapp

```sh
$ go run ./main.go
```
The webapp should be accessed at http://localhost:8080/generate?version=7. The response should be an UUID version 7. Other versions are on the roadmap.
## ROADMAP
**UUID Generator**
- [ ] Version 1
- [ ] Version 2
- [ ] Version 3
- [ ] Version 4
- [ ] Version 5
- [ ] Version 6
- [x] Version 7
- [ ] Version 8

**UUID Validator (Decoder)**

- [ ] Version 1
- [ ] Version 2
- [ ] Version 3
- [ ] Version 4
- [ ] Version 5
- [ ] Version 6
- [ ] Version 7
- [ ] Version 8

**Benchmarking**
- [ ] Version 1
- [ ] Version 2
- [ ] Version 3
- [x] Version 4
- [ ] Version 5
- [ ] Version 6
- [x] Version 7
- [ ] Version 8

**Benchmark results (on my PC)**

*UUID version 4, golang v1.17*

- goos: linux
- goarch: amd64
- pkg: uuid-draft/benchmark
- cpu: Intel(R) Core(TM) i7-2640M CPU @ 2.80GHz
- BenchmarkUUIDv4-4
- 1276732
- 938.8 ns/op 
- 64 B/op	       
- 2 allocs/op
- PASS ok  
- uuid-draft/benchmark	2.158s

*UUID version 7, golang v1.17*
- goos: linux
- goarch: amd64
- pkg: uuid-draft/benchmark
- cpu: Intel(R) Core(TM) i7-2640M CPU @ 2.80GHz
- BenchmarkV7-4
- 343354	      
- 3483 ns/op	    
- 208 B/op	       
- 9 allocs/op
- PASS ok  
- uuid-draft/benchmark	1.242s