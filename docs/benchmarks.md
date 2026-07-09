# Benchmark Log

Environment: go version go1.26.1 linux/amd64, AMD Ryzen 5 5600X 6-Core Processor.
All runs: `go test ./... -run '^$' -bench . -benchmem`.

## 2026-07-09 — Baseline (pre-refactor, main)

```
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/core
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkMatrixInvert-12    	   18457	     65621 ns/op	   63608 B/op	    2726 allocs/op
BenchmarkMatrixTimes-12     	 5786070	       206.5 ns/op	     248 B/op	       6 allocs/op
BenchmarkTupleOps-12        	24660932	        48.36 ns/op	      64 B/op	       2 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/core	4.528s
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/objects
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkSphereIntersect-12    	 8853871	       145.8 ns/op	     152 B/op	       5 allocs/op
BenchmarkSphereNormalAt-12     	 5168281	       237.2 ns/op	     344 B/op	       9 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/objects	2.900s
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/scene
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkCameraRender-12    	       3	 369954233 ns/op	320388621 B/op	13704861 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/scene	2.253s
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/shader
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkLighting-12    	16650871	        71.28 ns/op	      88 B/op	       3 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/shader	1.268s
```
