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

## 2026-07-09 — Checkpoint: after value semantics (Tasks 4–6)

```
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/core
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkMatrixInvert-12    	   20614	     57640 ns/op	   55904 B/op	    2405 allocs/op
BenchmarkMatrixTimes-12     	 7043257	       170.1 ns/op	     224 B/op	       5 allocs/op
BenchmarkTupleOps-12        	137089489	         8.783 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/core	5.249s
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/objects
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkSphereIntersect-12    	11606737	       104.6 ns/op	     136 B/op	       3 allocs/op
BenchmarkSphereNormalAt-12     	 8429895	       138.9 ns/op	     224 B/op	       5 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/objects	2.648s
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/scene
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkCameraRender-12    	       4	 298647798 ns/op	281152124 B/op	12051358 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/scene	2.398s
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/shader
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkLighting-12    	74946117	        16.17 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/shader	1.236s
```

Since baseline, `BenchmarkTupleOps` dropped from 48.36 ns/op (64 B/op, 2 allocs/op) to 8.783 ns/op (0 B/op, 0 allocs/op) and `BenchmarkLighting` dropped from 71.28 ns/op (88 B/op, 3 allocs/op) to 16.17 ns/op (0 B/op, 0 allocs/op), confirming the value-type conversions for `Tuple`/`Color` eliminated all heap allocation on those hot paths; `Matrix`-heavy benchmarks (`Invert`, `Times`, `SphereIntersect`, `SphereNormalAt`, `CameraRender`) also improved modestly since `Matrix` stayed a slice type and value receivers only avoid one extra pointer indirection/allocation for the header itself.

## 2026-07-09 — Checkpoint: after refactor phases 1–3

```
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/core
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkMatrixInvert-12    	   20330	     57884 ns/op	   55904 B/op	    2405 allocs/op
BenchmarkMatrixTimes-12     	 7554115	       159.4 ns/op	     224 B/op	       5 allocs/op
BenchmarkTupleOps-12        	135679488	         8.850 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/core	5.249s
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/objects
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkSphereIntersect-12    	11402106	       104.6 ns/op	     136 B/op	       3 allocs/op
BenchmarkSphereNormalAt-12     	 8703980	       136.8 ns/op	     224 B/op	       5 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/objects	2.641s
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/scene
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkCameraRender-12    	       4	 298219020 ns/op	281153410 B/op	12051358 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/scene	2.394s
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/shader
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkLighting-12    	74433103	        16.19 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/shader	1.230s
```

Numbers are ~flat vs. the Task 6 checkpoint, as expected — Phase 3 (Task 7–9) was a structural dedup of the CLI layer and the `Object` interface, with no changes to the hot math/render paths.

## 2026-07-09 — Perf: Invert determinant hoist

```
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/core
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkMatrixInvert-12    	   81429	     14552 ns/op	   14144 B/op	     605 allocs/op
BenchmarkMatrixInvert-12    	   81880	     14623 ns/op	   14144 B/op	     605 allocs/op
BenchmarkMatrixInvert-12    	   81908	     14745 ns/op	   14144 B/op	     605 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/core	4.046s
```

`BenchmarkMatrixInvert` dropped from 57884 ns/op to 14640 ns/op, a 4× improvement from hoisting the determinant computation out of the double loop and eliminating 15 redundant O(n!) cofactor expansions.

## 2026-07-09 — Perf: cached inverse-transpose

```
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/objects
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkSphereNormalAt-12    	80721944	        15.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkSphereNormalAt-12    	81468397	        14.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkSphereNormalAt-12    	82179948	        15.04 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/objects	3.718s
```

`BenchmarkSphereNormalAt` dropped from 136.8 ns/op to ~15 ns/op, a ~9× improvement from caching the inverse-transpose matrix to avoid redundant computation on every normal calculation.

## 2026-07-09 — Perf: cached camera inverse

```
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/scene
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkCameraRender-12    	    1176	    959564 ns/op	 1371187 B/op	   19696 allocs/op
BenchmarkCameraRender-12    	    1242	    955439 ns/op	 1371184 B/op	   19696 allocs/op
BenchmarkCameraRender-12    	    1242	    984178 ns/op	 1371188 B/op	   19696 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/scene	3.847s
```

`BenchmarkCameraRender` dropped from 298219020 ns/op (281153410 B/op, 12051358 allocs/op) at the phases 1–3 checkpoint to ~966000 ns/op (1371186 B/op, 19696 allocs/op), a ~310× speedup, ~205× fewer bytes allocated, and ~612× fewer allocations, from hoisting the per-pixel `Camera.Transform.Invert()` (and its redundant duplicate for the origin) out of `RayForPixel` into a single cache populated by `SetTransform`.

## 2026-07-09 — Final (after Phase 4)

```
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/core
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkMatrixInvert-12    	   79803	     14552 ns/op	   14144 B/op	     605 allocs/op
BenchmarkMatrixInvert-12    	   79508	     14620 ns/op	   14144 B/op	     605 allocs/op
BenchmarkMatrixInvert-12    	   79656	     14730 ns/op	   14144 B/op	     605 allocs/op
BenchmarkMatrixTimes-12     	 7464505	       160.4 ns/op	     224 B/op	       5 allocs/op
BenchmarkMatrixTimes-12     	 7413759	       168.1 ns/op	     224 B/op	       5 allocs/op
BenchmarkMatrixTimes-12     	 7377282	       164.9 ns/op	     224 B/op	       5 allocs/op
BenchmarkTupleOps-12        	133143684	         8.996 ns/op	       0 B/op	       0 allocs/op
BenchmarkTupleOps-12        	132990475	         8.884 ns/op	       0 B/op	       0 allocs/op
BenchmarkTupleOps-12        	135889078	         8.836 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/core	14.426s
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/objects
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkSphereIntersect-12    	11612594	       104.3 ns/op	     136 B/op	       3 allocs/op
BenchmarkSphereIntersect-12    	11473344	       104.1 ns/op	     136 B/op	       3 allocs/op
BenchmarkSphereIntersect-12    	11500155	       106.5 ns/op	     136 B/op	       3 allocs/op
BenchmarkSphereNormalAt-12     	80172000	        14.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkSphereNormalAt-12     	79947649	        14.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkSphereNormalAt-12     	81475776	        14.90 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/objects	7.607s
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/scene
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkCameraRender-12    	    1899	    700588 ns/op	 1374877 B/op	   19797 allocs/op
BenchmarkCameraRender-12    	    1783	    697973 ns/op	 1374882 B/op	   19797 allocs/op
BenchmarkCameraRender-12    	    1706	    704208 ns/op	 1374878 B/op	   19797 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/scene	3.989s
goos: linux
goarch: amd64
pkg: github.com/varigg/raytracer-challenge/pkg/shader
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkLighting-12    	73796630	        16.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkLighting-12    	74143182	        16.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkLighting-12    	74821405	        16.20 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/varigg/raytracer-challenge/pkg/shader	3.702s
```

`BenchmarkCameraRender` dropped further from ~966000 ns/op (1371186 B/op, 19696 allocs/op) to ~701000 ns/op (1374878 B/op, 19797 allocs/op), a further ~1.4× speedup from scanline-level parallelism (goroutine per row, `sync.WaitGroup`) on a 12-vCPU host; the small residual allocation increase (~3693 B/op, ~101 allocs/op) is per-goroutine closure/scheduling overhead. The modest speedup relative to core count reflects workload granularity: the benchmark canvas is 100×50, so `Render` spawns only 50 goroutines with ~14µs of work each, and goroutine spawn/schedule/`WaitGroup` overhead is significant at that scale. All other benchmarks (`MatrixInvert`, `MatrixTimes`, `TupleOps`, `SphereIntersect`, `SphereNormalAt`, `Lighting`) are unchanged within noise, as expected — Task 13 touched only `Camera.Render`.

## Summary

| Benchmark | Baseline | After value semantics | Final |
|---|---|---|---|
| BenchmarkMatrixInvert | 65621 ns/op | 57640 ns/op | 14552–14730 ns/op |
| BenchmarkMatrixTimes | 206.5 ns/op | 170.1 ns/op | 160.4–168.1 ns/op |
| BenchmarkTupleOps | 48.36 ns/op | 8.783 ns/op | 8.836–8.996 ns/op |
| BenchmarkSphereIntersect | 145.8 ns/op | 104.6 ns/op | 104.1–106.5 ns/op |
| BenchmarkSphereNormalAt | 237.2 ns/op | 138.9 ns/op | 14.87–14.91 ns/op |
| BenchmarkCameraRender | 369954233 ns/op | 298647798 ns/op | 697973–704208 ns/op |
| BenchmarkLighting | 71.28 ns/op | 16.17 ns/op | 16.20–16.55 ns/op |

Two wins dominate the overall trajectory: caching the camera's inverse transform and ray origin in `SetTransform` (instead of recomputing `Matrix.Invert()` on every pixel in `RayForPixel`) collapsed `BenchmarkCameraRender` by ~310× on its own, and the `Tuple`/`Color` value-type conversions eliminated heap allocation entirely on the innermost per-pixel math (`BenchmarkTupleOps`, `BenchmarkLighting` both go to 0 B/op, 0 allocs/op), which is what let the scanline-parallel `Render` in this task scale cleanly without lock contention or GC pressure swamping the goroutines. Scanline parallelism itself contributes a further, comparatively modest ~1.4× on this 6-core/12-thread host — a consequence of workload granularity: the 100×50 benchmark canvas yields only 50 goroutines with ~14µs of work each, so goroutine spawn/schedule/`WaitGroup` overhead eats into the theoretical parallel gain.
