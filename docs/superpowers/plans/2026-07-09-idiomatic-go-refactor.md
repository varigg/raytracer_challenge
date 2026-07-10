# Idiomatic Go Refactor Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Refactor the ray tracer to idiomatic Go (value semantics, single object interface, deduplicated CLI commands, proper error handling), then apply performance optimizations last, with benchmark measurements recorded at every checkpoint.

**Architecture:** Three phases of semantic-preserving refactors (hygiene/correctness → value semantics for `Tuple`/`Color`/`Matrix` → interface consolidation and CLI dedup), followed by a performance phase (invariant caching, determinant hoisting, parallel render). A benchmark harness is added first so every phase's effect is measured against the recorded baseline in `docs/benchmarks.md`.

**Tech Stack:** Go 1.22, cobra, testify. No new dependencies.

## Global Constraints

- Work happens on branch `idiomatic-go-refactor` in worktree `.worktrees/idiomatic-go-refactor` (create via superpowers:using-git-worktrees at execution start).
- Conventional Commits subjects; concise prose bodies explaining why. **NEVER add AI attribution** (no `Co-Authored-By`, no "Generated with" footers) — user's global CLAUDE.md overrides any default.
- `go test ./...` must be green after every task. Run it **unpiped** (never `go test ... | tail` in an `&&` chain — tail masks failures).
- Performance-affecting changes are confined to Phase 4 (Tasks 10–13). Phases 1–3 are behavior-preserving refactors (value semantics will incidentally reduce allocations; that's measured, not tuned).
- Benchmark command, used verbatim at every checkpoint: `go test ./... -run '^$' -bench . -benchmem`. Append full output to `docs/benchmarks.md` under a dated heading naming the checkpoint.
- The rendered images are the end-to-end oracle: after tasks that touch the render path, run `go run . scene` and visually compare `scene.png` against a copy saved before the refactor.
- Existing 100+ tests in `pkg/` are the safety net; mechanical test updates (dropping `*`/`&`, renames) are part of each task, never deletions of assertions.

---

## Phase 0 — Baseline

### Task 1: Benchmark harness and baseline measurements

**Files:**
- Create: `pkg/core/bench_test.go`
- Create: `pkg/objects/bench_test.go`
- Create: `pkg/shader/bench_test.go`
- Create: `pkg/scene/bench_test.go`
- Create: `docs/benchmarks.md`

**Interfaces:**
- Consumes: current (pointer-based) public API.
- Produces: benchmarks named `BenchmarkMatrixInvert`, `BenchmarkMatrixTimes`, `BenchmarkTupleOps`, `BenchmarkSphereIntersect`, `BenchmarkSphereNormalAt`, `BenchmarkLighting`, `BenchmarkCameraRender`. Later tasks update their *signatures* mechanically but must keep these names so results stay comparable across checkpoints.

- [ ] **Step 1: Save a reference render**

```bash
cd .worktrees/idiomatic-go-refactor
go run . scene && cp scene.png /tmp/scene-baseline.png
```

- [ ] **Step 2: Write the benchmarks**

`pkg/core/bench_test.go`:

```go
package core_test

import (
	"testing"

	"github.com/varigg/raytracer-challenge/pkg/core"
)

var benchMatrix = core.NewMatrix([][]float64{
	{-5, 2, 6, -8},
	{1, -5, 1, 8},
	{7, 7, -6, -7},
	{1, -3, 7, 4},
})

var (
	sinkMatrix *core.Matrix
	sinkTuple  *core.Tuple
)

func BenchmarkMatrixInvert(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkMatrix = benchMatrix.Invert()
	}
}

func BenchmarkMatrixTimes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkMatrix = benchMatrix.Times(benchMatrix)
	}
}

func BenchmarkTupleOps(b *testing.B) {
	v := core.NewVector(1, -1, 2)
	n := core.NewVector(0, 1, 0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkTuple = v.Add(v).Multiply(0.5).Reflect(n).Normalize()
	}
}
```

`pkg/objects/bench_test.go`:

```go
package objects_test

import (
	"testing"

	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
)

var (
	sinkIntersections []objects.Intersection
	sinkNormal        *core.Tuple
)

func BenchmarkSphereIntersect(b *testing.B) {
	s := objects.NewSphere()
	s.SetTransform(core.ScalingMatrix(2, 2, 2))
	r := objects.NewRay(core.NewPoint(0, 0, -5), core.NewVector(0, 0, 1))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkIntersections = s.Intersect(r)
	}
}

func BenchmarkSphereNormalAt(b *testing.B) {
	s := objects.NewSphere()
	s.SetTransform(core.TranslationMatrix(0, 1, 0))
	p := core.NewPoint(0, 1.70711, -0.70711)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkNormal = s.NormalAt(p)
	}
}
```

`pkg/shader/bench_test.go`:

```go
package shader_test

import (
	"testing"

	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

var sinkColor *core.Color

func BenchmarkLighting(b *testing.B) {
	m := shader.NewMaterial()
	light := shader.NewLight(core.NewPoint(0, 0, -10), core.NewColor(1, 1, 1))
	position := core.NewPoint(0, 0, 0)
	eye := core.NewVector(0, 0, -1)
	normal := core.NewVector(0, 0, -1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkColor = light.Lighting(m, position, eye, normal)
	}
}
```

`pkg/scene/bench_test.go`:

```go
package scene_test

import (
	"math"
	"testing"

	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/scene"
)

var sinkCanvas *core.Canvas

func BenchmarkCameraRender(b *testing.B) {
	w := scene.NewDefaultWorld()
	c := scene.NewCamera(100, 50, math.Pi/3)
	c.Transform = scene.ViewTransform(
		core.NewPoint(0, 0, -5), core.NewPoint(0, 0, 0), core.NewVector(0, 1, 0))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkCanvas = c.Render(w)
	}
}
```

- [ ] **Step 3: Run benchmarks, record baseline**

```bash
go test ./... -run '^$' -bench . -benchmem
```

Expected: all seven benchmarks report ns/op, B/op, allocs/op; no failures. Create `docs/benchmarks.md`:

```markdown
# Benchmark Log

Environment: <output of `go version`>, <CPU from /proc/cpuinfo model name>.
All runs: `go test ./... -run '^$' -bench . -benchmem`.

## 2026-07-09 — Baseline (pre-refactor, main)

<paste full benchmark output>
```

- [ ] **Step 4: Verify tests still green**

Run: `go test ./...`
Expected: `ok` for all four `pkg/` packages.

- [ ] **Step 5: Commit**

```bash
git add pkg/core/bench_test.go pkg/objects/bench_test.go pkg/shader/bench_test.go pkg/scene/bench_test.go docs/benchmarks.md
git commit -m "test: add benchmark harness and record baseline"
```

---

## Phase 1 — Hygiene and correctness

### Task 2: Dead code, naming, and repo hygiene

**Files:**
- Modify: `.gitignore`, `go.mod`/`go.sum`, `pkg/core/tuple.go:9-28`, `pkg/core/canvas.go`, `pkg/core/matrix.go`, `pkg/core/utils.go`, `pkg/objects/sphere.go:21-23`, `pkg/scene/camera.go:40`, `cmd/raytracer/clock.go:23`, `cmd/raytracer/sphere.go:17-18,56-57`

**Interfaces:**
- Consumes: nothing new.
- Produces: `core.Identity(size int) *Matrix` is the only identity constructor (callers of `Identity4()` switch to `Identity(4)`). Unexported `maxColorValue = 255` in canvas.go. `equals`/`epsilon` stay unexported in `pkg/core`.

- [ ] **Step 1: Apply the deletions and renames**

1. `.gitignore`: add a line `/raytrace` (anchored — an unanchored name would also match directories).
2. Run `go mod tidy` (fixes cobra/testify being marked `// indirect`).
3. `pkg/core/tuple.go`: delete the commented-out interface sketch (lines 9–28).
4. `pkg/core/canvas.go`: delete `Pixels()` (unused, off-by-one). Replace `const MAX_COLORS = 256` with `const maxColorValue = 255`; in `ToPPM` use `maxColorValue` directly (drop the `-1`s); in `SavePNG` replace the literal `255` with `maxColorValue`.
5. `pkg/core/matrix.go`: delete `Identity4()` and `Matrix.Print()` (unused). Keep `Identity(size)`.
6. `pkg/scene/camera.go:40`: `Transform: core.Identity4(),` → `Transform: core.Identity(4),`.
7. `pkg/objects/sphere.go:21-23`: replace both `core.Identity4()` calls with `core.Identity(4)`.
8. `pkg/core/utils.go`: rename `EPSILON` → `epsilon` (no external references exist); body of `equals` becomes `return math.Abs(f1-f2) < epsilon`.
9. `cmd/raytracer/clock.go:23`: `for _ = range 12` → `for range 12`.
10. `cmd/raytracer/sphere.go`: `Use: "shadow "` → `"shadow"` with `Aliases: []string{"chapter5"}` (the silhouette *is* chapter 5); `Use: "sphere "` → `"sphere"` with `Aliases: []string{"chapter6"}` (lit sphere is chapter 6 — currently both claim `chapter5`).

- [ ] **Step 2: Verify**

Run: `go build ./... && go vet ./...` then `go test ./...`
Expected: clean build, all tests pass. `git status` no longer lists `raytrace`.

- [ ] **Step 3: Commit**

```bash
git add -A
git commit -m "chore: remove dead code, fix naming and CLI metadata"
```

### Task 3: Canvas I/O error handling

**Files:**
- Modify: `pkg/core/canvas.go`
- Modify: `cmd/raytracer/clock.go`, `cmd/raytracer/projectile.go`, `cmd/raytracer/scene.go`, `cmd/raytracer/simple_scene.go`, `cmd/raytracer/sphere.go` (switch `Run` → `RunE`, return save errors)
- Test: `pkg/core/canvas_test.go`

**Interfaces:**
- Produces: `func (c *Canvas) SavePNG(filename string) error`, `func (c *Canvas) SavePPM(filename string) error` (no longer prints on success), `func (c *Canvas) ToPPM(w io.Writer) error` (now checks write errors). All cobra commands use `RunE`.

- [ ] **Step 1: Write the failing tests** (append to `pkg/core/canvas_test.go`)

```go
func TestSavePNGReturnsErrorForBadPath(t *testing.T) {
	c := core.NewCanvas(2, 2)
	err := c.SavePNG(filepath.Join(t.TempDir(), "no-such-dir", "out.png"))
	assert.Error(t, err)
}

func TestSavePNGWritesFile(t *testing.T) {
	c := core.NewCanvas(2, 2)
	path := filepath.Join(t.TempDir(), "out.png")
	assert.NoError(t, c.SavePNG(path))
	info, err := os.Stat(path)
	assert.NoError(t, err)
	assert.Greater(t, info.Size(), int64(0))
}
```

Add `os` and `path/filepath` to the test file's imports.

- [ ] **Step 2: Run to verify failure**

Run: `go test ./pkg/core/ -run TestSavePNG -v`
Expected: FAIL — `TestSavePNGReturnsErrorForBadPath` panics or fails to compile against the current `func SavePNG(string)` (no return). Compile error is the expected failure mode here.

- [ ] **Step 3: Implement**

Replace `ToPPM`, `SavePPM`, `SavePNG` in `pkg/core/canvas.go` (imports become `bufio`, `fmt`, `image`, `image/png`, `io`, `os`, `strconv`):

```go
func (c *Canvas) ToPPM(w io.Writer) error {
	bw := bufio.NewWriter(w)
	fmt.Fprintf(bw, "P3\n%d %d\n%d\n", c.Width, c.Height, maxColorValue)
	var lineLen int
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			rgb := c.Get(x, y).ToRGBA(maxColorValue)
			for _, v := range []uint8{rgb.R, rgb.G, rgb.B} {
				s := strconv.Itoa(int(v))
				if lineLen > 0 && lineLen+1+len(s) > 70 {
					bw.WriteByte('\n')
					lineLen = 0
				}
				if lineLen > 0 {
					bw.WriteByte(' ')
					lineLen++
				}
				bw.WriteString(s)
				lineLen += len(s)
			}
		}
		bw.WriteByte('\n')
		lineLen = 0
	}
	return bw.Flush()
}

func (c *Canvas) SavePPM(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("create %s: %w", fileName, err)
	}
	defer file.Close()
	return c.ToPPM(file)
}

func (c *Canvas) SavePNG(filename string) error {
	img := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			img.Set(x, y, c.Get(x, y).ToRGBA(maxColorValue))
		}
	}
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create %s: %w", filename, err)
	}
	defer f.Close()
	return png.Encode(f, img)
}
```

- [ ] **Step 4: Propagate to cobra commands**

In every `cmd/raytracer/*.go` command, change `Run: func(cmd *cobra.Command, args []string) {` to `RunE: func(cmd *cobra.Command, args []string) error {` and end the body with `return canvas.SavePNG("<name>.png")` (or `SavePPM`). Specifics:
- `clock.go`: `if err := c.SavePPM("clock.ppm"); err != nil { return err }` then `return c.SavePNG("clock.png")`.
- `projectile.go`: `projectileCmd` returns `nil` at the end (no file output); `projectileGraphCmd` ends `return canvas.SavePPM("canvas.ppm")`.
- `scene.go`: `return canvas.SavePNG("scene.png")`.
- `simple_scene.go`: `return canvas.SavePNG("simple_scene.png")`.
- `sphere.go`: both commands `return canvas.SavePNG(...)`.

- [ ] **Step 5: Verify**

Run: `go test ./...` then `go run . clock`
Expected: tests pass; `clock.ppm` and `clock.png` created; no spurious "successfully written" print.

- [ ] **Step 6: Commit**

```bash
git add -A
git commit -m "fix: return errors from canvas save paths"
```

Body: mention the `defer f.Close()`-before-error-check nil-pointer bug this removes.

---

## Phase 2 — Value semantics

### Task 4: Color as a value type

**Files:**
- Modify: `pkg/core/color.go` (full rewrite), `pkg/core/canvas.go`, `pkg/shader/material.go`, `pkg/shader/light.go`, `pkg/scene/world.go`, `pkg/scene/camera.go`, all of `cmd/raytracer/`, test files that compile against these
- Test: existing `pkg/core/color_test.go` et al. (mechanical updates)

**Interfaces:**
- Produces: `type Color struct { R, G, B float64 }` with **value** receivers/returns: `NewColor(r, g, b float64) Color`, `(c Color) Add/Subtract/HadamardProduct(o Color) Color`, `(c Color) Multiply(s float64) Color`, `(c Color) Equals(o Color) bool`, `(c Color) ToRGBA(maxValue int) color.RGBA`. `ConvertFloatToColorValue` keeps its exported name/signature (tests use it). Canvas: `Get(x, y int) Color`, `Set(x, y int, color Color)`. Shader: `Material.Color Color` (unchanged field type, but no more `*core.NewColor(...)` deref needed), `Light.Intensity Color`, `NewLight(position *core.Tuple, color Color) *Light`, `Lighting(...) Color`. Scene: `ShadeHit/ColorAt` return `Color`.

- [ ] **Step 1: Rewrite `pkg/core/color.go`**

```go
package core

import (
	"image/color"
	"math"
)

type Color struct {
	R, G, B float64
}

func NewColor(r, g, b float64) Color {
	return Color{R: r, G: g, B: b}
}

func (c Color) Add(o Color) Color {
	return Color{c.R + o.R, c.G + o.G, c.B + o.B}
}

func (c Color) Subtract(o Color) Color {
	return Color{c.R - o.R, c.G - o.G, c.B - o.B}
}

func (c Color) Multiply(s float64) Color {
	return Color{c.R * s, c.G * s, c.B * s}
}

func (c Color) HadamardProduct(o Color) Color {
	return Color{c.R * o.R, c.G * o.G, c.B * o.B}
}

func (c Color) Equals(o Color) bool {
	return equals(c.R, o.R) && equals(c.G, o.G) && equals(c.B, o.B)
}

func (c Color) ToRGBA(maxValue int) color.RGBA {
	return color.RGBA{
		R: ConvertFloatToColorValue(c.R, maxValue),
		G: ConvertFloatToColorValue(c.G, maxValue),
		B: ConvertFloatToColorValue(c.B, maxValue),
		A: 0xFF,
	}
}

func ConvertFloatToColorValue(f float64, maxValue int) uint8 {
	f = math.Max(0, math.Min(1, f))
	return uint8(math.Round(f * float64(maxValue)))
}
```

- [ ] **Step 2: Fix compile errors outward, package by package**

Run `go build ./...` and fix mechanically. The complete pattern list:
- `pkg/core/canvas.go`: `Get` returns `Color` (`return c.pixels[x][y]`), `Set` takes `Color` (`c.pixels[x][y] = color`).
- `pkg/shader/material.go`: `Color: NewColor(1, 1, 1)` loses the `*`... i.e. `Color: core.NewColor(1, 1, 1)` now assigns directly.
- `pkg/shader/light.go`: `Intensity Color` value field; `black := core.Color{}`; `Lighting` returns `Color`; drop all `*Color` in signatures.
- `pkg/scene/world.go`: `ShadeHit`/`ColorAt` return `core.Color`; `core.NewColor(0, 0, 0)` return stays as-is (now a value). `m.Color = *core.NewColor(.8, 1, .6)` → `m.Color = core.NewColor(.8, 1, .6)`.
- `cmd/raytracer/*.go`: every `*core.NewColor(...)` → `core.NewColor(...)`; `DrawSquare`/`canvas.Set` callers pass values.
- Tests: drop `*`/`&` where the compiler complains; assertions comparing colors keep working because values compare cleanly with `assert.Equal`.
- `pkg/shader/bench_test.go`: `var sinkColor core.Color`.

- [ ] **Step 3: Verify**

Run: `go test ./...` and `go run . scene && go run . simple-scene`
Expected: green; `scene.png` visually identical to `/tmp/scene-baseline.png`.

- [ ] **Step 4: Commit**

```bash
git add -A
git commit -m "refactor: make Color a value type with exported fields"
```

### Task 5: Tuple as a value type

**Files:**
- Modify: `pkg/core/tuple.go` (full rewrite), `pkg/core/matrix.go` (`MultiplyWithTuple`), `pkg/objects/ray.go`, `pkg/objects/sphere.go`, `pkg/objects/object.go`, `pkg/objects/intersection.go`, `pkg/shader/light.go`, `pkg/scene/*.go`, `cmd/raytracer/*.go`, all affected tests

**Interfaces:**
- Produces: `type Tuple struct{ X, Y, Z, W float64 }` with value receivers/returns throughout: `NewTuple/NewPoint/NewVector(...) Tuple`, `(t Tuple) Add/Subtract(o Tuple) Tuple`, `Negate() Tuple`, `Multiply/Divide(s float64) Tuple`, `Magnitude() float64`, `Normalize() Tuple`, `Dot(o Tuple) float64`, `Cross(o Tuple) Tuple`, `Equals(o Tuple) bool`, `Reflect(normal Tuple) Tuple`, `Transform(m *Matrix) Tuple`. `NewPointFromString`/`NewVectorFromString` return `Tuple` (error handling comes in Task 8). Matrix: `MultiplyWithTuple(t Tuple) Tuple`. `Ray` stays a pointer type but its fields become values: `Ray{ Origin, Direction Tuple }`, `NewRay(origin, direction Tuple) *Ray`, `Position(t float64) Tuple`. Interfaces: `NormalAt(Tuple) Tuple` in both `Object` and `Intersecter`. Shader: `Light.Position Tuple`, `NewLight(position Tuple, color Color) *Light`, `Lighting(m *Material, point, eyeV, normalV Tuple) Color`. Scene: `Computations.Point/EyeV/NormalV Tuple`, `ViewTransform(from, to, up Tuple) *Matrix`.

- [ ] **Step 1: Rewrite `pkg/core/tuple.go`**

```go
package core

import (
	"math"
	"strconv"
	"strings"
)

type Tuple struct {
	X, Y, Z, W float64
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: w}
}

func NewPoint(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 1)
}

func NewVector(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 0)
}

func NewVectorFromString(v string) Tuple {
	x, y, z := stringToCoordinates(v)
	return NewVector(x, y, z)
}

func NewPointFromString(v string) Tuple {
	x, y, z := stringToCoordinates(v)
	return NewPoint(x, y, z)
}

func stringToCoordinates(v string) (float64, float64, float64) {
	coords := strings.Split(v, ",")
	x, _ := strconv.ParseFloat(strings.TrimSpace(coords[0]), 64)
	y, _ := strconv.ParseFloat(strings.TrimSpace(coords[1]), 64)
	z, _ := strconv.ParseFloat(strings.TrimSpace(coords[2]), 64)
	return x, y, z
}

func (t Tuple) IsVector() bool { return t.W == 0 }
func (t Tuple) IsPoint() bool  { return t.W == 1 }

func (t Tuple) Add(o Tuple) Tuple {
	return Tuple{t.X + o.X, t.Y + o.Y, t.Z + o.Z, t.W + o.W}
}

func (t Tuple) Subtract(o Tuple) Tuple {
	return Tuple{t.X - o.X, t.Y - o.Y, t.Z - o.Z, t.W - o.W}
}

func (t Tuple) Negate() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

func (t Tuple) Multiply(s float64) Tuple {
	return Tuple{t.X * s, t.Y * s, t.Z * s, t.W * s}
}

func (t Tuple) Divide(s float64) Tuple {
	return Tuple{t.X / s, t.Y / s, t.Z / s, t.W / s}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

func (t Tuple) Normalize() Tuple {
	return t.Divide(t.Magnitude())
}

func (t Tuple) Dot(o Tuple) float64 {
	return t.X*o.X + t.Y*o.Y + t.Z*o.Z + t.W*o.W
}

// Cross only makes sense for vectors (W = 0).
func (t Tuple) Cross(o Tuple) Tuple {
	return NewVector(
		t.Y*o.Z-t.Z*o.Y,
		t.Z*o.X-t.X*o.Z,
		t.X*o.Y-t.Y*o.X,
	)
}

func (t Tuple) Equals(o Tuple) bool {
	return equals(t.X, o.X) && equals(t.Y, o.Y) &&
		equals(t.Z, o.Z) && equals(t.W, o.W)
}

func (t Tuple) Transform(m *Matrix) Tuple {
	return m.MultiplyWithTuple(t)
}

func (t Tuple) Reflect(normal Tuple) Tuple {
	return t.Subtract(normal.Multiply(2).Multiply(t.Dot(normal)))
}
```

(Note `Normalize` now computes `Magnitude` once instead of four times — a correctness-neutral simplification that falls out of the rewrite.)

- [ ] **Step 2: Update `pkg/core/matrix.go` `MultiplyWithTuple`**

```go
func (m *Matrix) MultiplyWithTuple(t Tuple) Tuple {
	row := func(i int) Tuple {
		return NewTuple((*m)[i][0], (*m)[i][1], (*m)[i][2], (*m)[i][3])
	}
	return NewTuple(row(0).Dot(t), row(1).Dot(t), row(2).Dot(t), row(3).Dot(t))
}
```

- [ ] **Step 3: Fix compile errors outward**

Run `go build ./...` repeatedly. Complete pattern list:
- `pkg/objects/ray.go`: fields `Origin, Direction Tuple`; `NewRay(origin, direction Tuple) *Ray`; `Position`/`Transform` return values.
- `pkg/objects/sphere.go`: `var center = NewPoint(0, 0, 0)` becomes a value; `NormalAt(worldPoint core.Tuple) core.Tuple`; `worldNormal.W = 0` still works (local value).
- `pkg/objects/object.go` and `intersection.go`: `NormalAt(core.Tuple) core.Tuple`.
- `pkg/shader/light.go`: signatures per Interfaces block above.
- `pkg/scene/computations.go`: value fields; `camera.go`: `RayForPixel` internals lose pointers; `ViewTransform` takes values.
- `cmd/raytracer/*.go`: `projectile`/`tick` structs hold values; everything else is deref-dropping.
- Benchmarks: `sinkTuple core.Tuple`, `sinkNormal core.Tuple`.
- Tests: mechanical `*`/`&` removal.

- [ ] **Step 4: Verify**

Run: `go test ./...` then `go run . scene`
Expected: green; `scene.png` matches baseline.

- [ ] **Step 5: Commit**

```bash
git add -A
git commit -m "refactor: make Tuple a value type"
```

### Task 6: Matrix value receivers and constructor cleanup

**Files:**
- Modify: `pkg/core/matrix.go`, `pkg/core/tuple.go` (`Transform` param), `pkg/objects/ray.go`, `pkg/objects/sphere.go`, `pkg/objects/object.go`, `pkg/scene/camera.go`, `pkg/scene/world.go`, `cmd/raytracer/*.go`, tests
- Test: `pkg/core/matrix_test.go` (mechanical)

**Interfaces:**
- Produces: `type Matrix [][]float64` kept, but all constructors return `Matrix` (not `*Matrix`) and all methods use value receivers: `NewMatrix(data [][]float64) Matrix` (panics on non-square input instead of returning nil), `NewEmptyMatrix(size int) Matrix`, `Identity(size int) Matrix`, `(m Matrix) Size/Equals/Times/MultiplyWithTuple/Transpose/Determinant/SubMatrix/Minor/Cofactor/IsInvertible/Invert/Translate/Scale/RotateX/RotateY/RotateZ`, `TranslationMatrix/ScalingMatrix/RotationMatrixX/Y/Z/Shearing(...) Matrix`. Downstream fields/params change `*core.Matrix` → `core.Matrix`: `Sphere.transform/invert`, `SetTransform(m core.Matrix) error` (error dropped in Task 7), `GetTransform()/GetInverseTransform() core.Matrix`, `Camera.Transform core.Matrix`, `ViewTransform(...) core.Matrix`, `Ray.Transform(m core.Matrix) *Ray`, `Tuple.Transform(m Matrix) Tuple`.

- [ ] **Step 1: Rewrite matrix.go mechanically**

Three rules applied to every declaration:
1. Receiver `(m *Matrix)` → `(m Matrix)`; body `(*m)[i][j]` → `m[i][j]`, `len(*m)` → `len(m)`.
2. Return type `*Matrix` → `Matrix`; `m := Matrix(data); return &m` → `return Matrix(data)`.
3. `Equals` drops the nil-pointer comparison; size guard becomes `if len(m) != len(o) { return false }`.

The two non-mechanical functions:

```go
// NewMatrix returns a square Matrix wrapping data. Panics if data is not square.
func NewMatrix(data [][]float64) Matrix {
	if len(data) == 0 || len(data) != len(data[0]) {
		panic("matrix data must be square")
	}
	return Matrix(data)
}
```

`Identity` writes `m[j][j] = 1.0` directly.

- [ ] **Step 2: Fix compile errors outward**

`go build ./...`; apply the type changes listed in the Interfaces block. In tests, `m := *core.NewMatrix(...)` → `m := core.NewMatrix(...)` (indexing `m[0][0]` keeps working since `Matrix` is a slice). `var sinkMatrix core.Matrix` in the bench file.

- [ ] **Step 3: Verify and record checkpoint**

Run: `go test ./...` then `go run . scene` (compare image), then:

```bash
go test ./... -run '^$' -bench . -benchmem
```

Append output to `docs/benchmarks.md` under `## 2026-07-09 — Checkpoint: after value semantics (Tasks 4–6)`. Expect allocs/op to drop sharply for `BenchmarkTupleOps` and `BenchmarkLighting`; note the deltas in one sentence.

- [ ] **Step 4: Commit**

```bash
git add -A
git commit -m "refactor: value receivers and value returns for Matrix"
```

---

## Phase 3 — API consolidation and CLI dedup

### Task 7: Single Object interface, accessor renames, no fake errors

**Files:**
- Modify: `pkg/objects/object.go`, `pkg/objects/intersection.go`, `pkg/objects/sphere.go`, `pkg/scene/computations.go`, `pkg/scene/world.go`, `cmd/raytracer/scene.go`, `cmd/raytracer/simple_scene.go`, `cmd/raytracer/sphere.go`, `CLAUDE.md`
- Test: `pkg/objects/sphere_test.go`, `pkg/scene/world_test.go`, `pkg/scene/world_colorat_shadehit_test.go` (mechanical renames + two type assertions)

**Interfaces:**
- Produces: exactly one interface in `pkg/objects/object.go`:

```go
type Object interface {
	Intersect(*Ray) []Intersection
	NormalAt(core.Tuple) core.Tuple
	Material() *shader.Material
}
```

`Intersecter` is deleted; `Intersection.Object` and `scene.Computations.Object` are typed `Object`. Sphere API: unexported fields `transform`, `invert`, `material`; methods `SetTransform(m core.Matrix)` (**no error return**), `Transform() core.Matrix`, `InverseTransform() core.Matrix`, `Material() *shader.Material`, `SetMaterial(m *shader.Material)`. Transform accessors are *not* on the interface (nothing consumes them polymorphically).

- [ ] **Step 1: Apply the interface and sphere changes**

`pkg/objects/sphere.go` becomes:

```go
type Sphere struct {
	transform core.Matrix
	// invert caches transform.Invert(); every intersection needs it.
	invert   core.Matrix
	material *shader.Material
}

func NewSphere() *Sphere {
	return &Sphere{
		transform: core.Identity(4),
		invert:    core.Identity(4).Invert(),
		material:  shader.NewMaterial(),
	}
}

func (s *Sphere) SetTransform(m core.Matrix) {
	s.transform = m
	s.invert = m.Invert()
}

func (s *Sphere) Transform() core.Matrix        { return s.transform }
func (s *Sphere) InverseTransform() core.Matrix { return s.invert }
func (s *Sphere) Material() *shader.Material    { return s.material }
func (s *Sphere) SetMaterial(m *shader.Material) { s.material = m }
```

(`Intersect` and `NormalAt` unchanged from Task 6 state.) Delete `Intersecter` from `intersection.go`; `Intersection.Object Object`.

- [ ] **Step 2: Fix call sites**

- `pkg/scene/computations.go`: `Object objects.Object`.
- `pkg/scene/world.go`: `comps.Object.Material()`; `NewDefaultWorld` loses the error dance — `innerSphere.SetTransform(core.ScalingMatrix(.5, .5, .5))` plain, `outerSphere.SetMaterial(m)`.
- `cmd/raytracer/scene.go` & `simple_scene.go` & `sphere.go`: `x.Material = mat` → `x.SetMaterial(mat)`; `x.Material.Specular = 0` → mutate the `*shader.Material` before `SetMaterial`, e.g. build `m := shader.NewMaterial(); m.Color = ...; m.Specular = 0; floor.SetMaterial(m)`. Keep `leftWall.SetMaterial(floor.Material())` semantics *as-is for now* (deliberate share; Task 9's builder gives each object its own material and fixes the latent aliasing).
- Tests: `GetMaterial()` → `Material()`, `GetTransform()` → `Transform()`; drop `err :=` at `sphere_test.go:77`. Two interface-typed accesses need assertions: `world_test.go:65` → `w.Objects[0].(*objects.Sphere).Transform()`, `world_colorat_shadehit_test.go:43-44` → same pattern.

- [ ] **Step 3: Update CLAUDE.md**

Replace the `pkg/objects` bullet's two-interface warning with: "`pkg/objects` — `Ray`, `Sphere`, intersection logic, and the single `Object` interface (`Intersect`/`NormalAt`/`Material`) consumed by `scene.World` and `scene.Computations`."

- [ ] **Step 4: Verify**

Run: `go test ./...` then `go run . scene` (image matches baseline).

- [ ] **Step 5: Commit**

```bash
git add -A
git commit -m "refactor: collapse Object and Intersecter into one interface"
```

### Task 8: Parse errors for tuple-from-string

**Files:**
- Modify: `pkg/core/tuple.go`, `cmd/raytracer/projectile.go`
- Test: `pkg/core/tuple_test.go`

**Interfaces:**
- Produces: `func NewPointFromString(v string) (Tuple, error)`, `func NewVectorFromString(v string) (Tuple, error)`. Sole callers are the two projectile commands (already `RunE` from Task 3).

- [ ] **Step 1: Write failing tests** (append to `pkg/core/tuple_test.go`)

```go
func TestNewPointFromStringParsesCoordinates(t *testing.T) {
	p, err := core.NewPointFromString(" 1, -2.5, 3 ")
	assert.NoError(t, err)
	assert.True(t, p.Equals(core.NewPoint(1, -2.5, 3)))
}

func TestNewPointFromStringRejectsWrongArity(t *testing.T) {
	_, err := core.NewPointFromString("1,2")
	assert.Error(t, err)
}

func TestNewVectorFromStringRejectsGarbage(t *testing.T) {
	_, err := core.NewVectorFromString("1,two,3")
	assert.Error(t, err)
}
```

- [ ] **Step 2: Run to verify failure**

Run: `go test ./pkg/core/ -run FromString -v`
Expected: compile failure (single return value today) — that is the failing state.

- [ ] **Step 3: Implement**

```go
func NewVectorFromString(v string) (Tuple, error) {
	x, y, z, err := stringToCoordinates(v)
	if err != nil {
		return Tuple{}, err
	}
	return NewVector(x, y, z), nil
}

func NewPointFromString(v string) (Tuple, error) {
	x, y, z, err := stringToCoordinates(v)
	if err != nil {
		return Tuple{}, err
	}
	return NewPoint(x, y, z), nil
}

func stringToCoordinates(v string) (x, y, z float64, err error) {
	coords := strings.Split(v, ",")
	if len(coords) != 3 {
		return 0, 0, 0, fmt.Errorf("expected 3 comma-separated coordinates, got %q", v)
	}
	dests := []*float64{&x, &y, &z}
	for i, c := range coords {
		*dests[i], err = strconv.ParseFloat(strings.TrimSpace(c), 64)
		if err != nil {
			return 0, 0, 0, fmt.Errorf("invalid coordinate %q in %q: %w", c, v, err)
		}
	}
	return x, y, z, nil
}
```

Add `fmt` to tuple.go imports. In `cmd/raytracer/projectile.go`, both `RunE` bodies parse all four inputs up front:

```go
position, err := core.NewPointFromString(origin)
if err != nil {
	return fmt.Errorf("--origin: %w", err)
}
vel, err := core.NewVectorFromString(velocity)
if err != nil {
	return fmt.Errorf("--velocity: %w", err)
}
grav, err := core.NewVectorFromString(gravity)
if err != nil {
	return fmt.Errorf("--gravity: %w", err)
}
wnd, err := core.NewVectorFromString(wind)
if err != nil {
	return fmt.Errorf("--wind: %w", err)
}
p := projectile{position: position, velocity: vel}
env := environment{gravity: grav, wind: wnd}
```

- [ ] **Step 4: Verify**

Run: `go test ./...` then `go run . projectile --origin "1,2"` (expect a clean error message, exit 1) and `go run . projectile` (expect distance output).

- [ ] **Step 5: Commit**

```bash
git add -A
git commit -m "feat: validate coordinate flags instead of silently zeroing"
```

### Task 9: CLI dedup — scene builders, wall renderer, Canvas.DrawSquare, --output flag

**Files:**
- Create: `cmd/raytracer/helpers.go`
- Modify: `pkg/core/canvas.go`, `cmd/raytracer/root.go`, `cmd/raytracer/scene.go`, `cmd/raytracer/simple_scene.go`, `cmd/raytracer/sphere.go`, `cmd/raytracer/clock.go`, `cmd/raytracer/projectile.go`
- Test: `pkg/core/canvas_test.go`

**Interfaces:**
- Produces: `func (c *Canvas) DrawSquare(x, y int, color Color)` in `pkg/core` (bounds-clipped). In `cmd/raytracer`: `newMaterial(c core.Color, diffuse, specular float64) *shader.Material`, `newSphere(transform core.Matrix, mat *shader.Material) *objects.Sphere`, `renderOnWall(shape *objects.Sphere, shade func(hit *objects.Intersection, ray *objects.Ray) core.Color) *core.Canvas`, `saveCanvas(c *core.Canvas, defaultName string) error`, persistent `--output` flag on root.

- [ ] **Step 1: Failing test for DrawSquare clipping** (append to `pkg/core/canvas_test.go`)

```go
func TestDrawSquareClipsAtCanvasEdge(t *testing.T) {
	c := core.NewCanvas(4, 4)
	red := core.NewColor(1, 0, 0)
	c.DrawSquare(0, 0, red) // must not panic on negative neighbors
	assert.True(t, c.Get(0, 0).Equals(red))
	assert.True(t, c.Get(1, 1).Equals(red))
	assert.True(t, c.Get(2, 2).Equals(core.NewColor(0, 0, 0)))
}
```

Run: `go test ./pkg/core/ -run DrawSquare -v` — expect compile failure (method undefined).

- [ ] **Step 2: Implement DrawSquare on Canvas** (`pkg/core/canvas.go`)

```go
// DrawSquare paints a 3x3 square centered on (x, y), clipped to the canvas bounds.
func (c *Canvas) DrawSquare(x, y int, color Color) {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			px, py := x+dx, y+dy
			if px >= 0 && px < c.Width && py >= 0 && py < c.Height {
				c.Set(px, py, color)
			}
		}
	}
}
```

Delete `DrawSquare` from `projectile.go`; call sites in `clock.go`/`projectile.go` become `c.DrawSquare(x, y, color)`. Run the test — PASS.

- [ ] **Step 3: Add `--output` and `saveCanvas`** (`cmd/raytracer/root.go`)

```go
var outputFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&outputFile, "output", "",
		"output image file; .ppm extension selects PPM, anything else PNG (default per command)")
}

// saveCanvas writes c to --output if set, else to defaultName, picking the
// format from the file extension.
func saveCanvas(c *core.Canvas, defaultName string) error {
	name := outputFile
	if name == "" {
		name = defaultName
	}
	if strings.HasSuffix(name, ".ppm") {
		return c.SavePPM(name)
	}
	return c.SavePNG(name)
}
```

(No `-o` shorthand: `projectile` already uses `-o` for `--origin`.) Every command's final save becomes `return saveCanvas(canvas, "<default>")`: `clock.png`, `canvas.ppm` (projectile-graph), `scene.png`, `simple_scene.png`, `shadow.png`, `sphere.png`. `clock.go` drops the dual PPM+PNG save in favor of one `saveCanvas(c, "clock.png")`.

- [ ] **Step 4: Create `cmd/raytracer/helpers.go`**

```go
package raytracer

import (
	"github.com/varigg/raytracer-challenge/pkg/core"
	"github.com/varigg/raytracer-challenge/pkg/objects"
	"github.com/varigg/raytracer-challenge/pkg/shader"
)

func newMaterial(c core.Color, diffuse, specular float64) *shader.Material {
	m := shader.NewMaterial()
	m.Color = c
	m.Diffuse = diffuse
	m.Specular = specular
	return m
}

func newSphere(transform core.Matrix, mat *shader.Material) *objects.Sphere {
	s := objects.NewSphere()
	s.SetTransform(transform)
	s.SetMaterial(mat)
	return s
}

// renderOnWall shoots a ray from a fixed origin at every pixel of a 7x7 wall
// at z=10 and paints pixels where shape is hit, using shade for the color.
func renderOnWall(shape *objects.Sphere, shade func(hit *objects.Intersection, ray *objects.Ray) core.Color) *core.Canvas {
	const pixels = 500
	const wallZ, wallSize = 10.0, 7.0
	canvas := core.NewCanvas(pixels, pixels)
	rayOrigin := core.NewPoint(0, 0, -5)
	pixelSize := wallSize / float64(pixels)
	half := wallSize / 2
	for y := 0; y < pixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < pixels; x++ {
			worldX := -half + pixelSize*float64(x)
			target := core.NewPoint(worldX, worldY, wallZ)
			ray := objects.NewRay(rayOrigin, target.Subtract(rayOrigin).Normalize())
			if hit := objects.Hit(shape.Intersect(ray)); hit != nil {
				canvas.Set(x, y, shade(hit, ray))
			}
		}
	}
	return canvas
}
```

- [ ] **Step 5: Rewrite the scene commands**

`cmd/raytracer/sphere.go` — both `RunE` bodies collapse to:

```go
// shadow command
shape := objects.NewSphere()
shape.SetTransform(core.ScalingMatrix(1, 0.5, 1))
canvas := renderOnWall(shape, func(_ *objects.Intersection, _ *objects.Ray) core.Color {
	return core.NewColor(0.9, 0, 0)
})
return saveCanvas(canvas, "shadow.png")

// sphere command
shape := objects.NewSphere()
shape.SetMaterial(newMaterial(core.NewColor(1, 0.2, 1), 0.9, 0.9))
light := shader.NewLight(core.NewPoint(-10, 10, -10), core.NewColor(1, 1, 1))
canvas := renderOnWall(shape, func(hit *objects.Intersection, ray *objects.Ray) core.Color {
	point := ray.Position(hit.T)
	normal := hit.Object.NormalAt(point)
	eye := ray.Direction.Negate()
	return light.Lighting(hit.Object.Material(), point, eye, normal)
})
return saveCanvas(canvas, "sphere.png")
```

`cmd/raytracer/scene.go` `RunE` body, using the fluent transform chain (apply-on-top order: scale, then rotate, then translate — equivalent to the old nested `Times` pyramid):

```go
wallMaterial := func() *shader.Material {
	return newMaterial(core.NewColor(1, 0.9, 0.9), 0.9, 0)
}
wallTransform := func(rotY float64) core.Matrix {
	return core.ScalingMatrix(10, 0.01, 10).RotateX(math.Pi / 2).RotateY(rotY).Translate(0, 0, 5)
}

floor := newSphere(core.ScalingMatrix(10, 0.01, 10), wallMaterial())
leftWall := newSphere(wallTransform(-math.Pi/4), wallMaterial())
rightWall := newSphere(wallTransform(math.Pi/4), wallMaterial())
middle := newSphere(core.TranslationMatrix(-0.5, 1, 0.5),
	newMaterial(core.NewColor(0.1, 1, 0.5), 0.7, 0.3))
right := newSphere(core.ScalingMatrix(0.5, 0.5, 0.5).Translate(1.5, 0.5, -0.5),
	newMaterial(core.NewColor(0.5, 1, 0.1), 0.7, 0.3))
left := newSphere(core.ScalingMatrix(0.33, 0.33, 0.33).Translate(-1.5, 0.333, -0.75),
	newMaterial(core.NewColor(1, 0.8, 0.1), 0.7, 0.3))

world := scene.NewWorld()
world.Add(floor, leftWall, rightWall, middle, right, left)
world.Light = shader.NewLight(core.NewPoint(-10, 10, -10), core.NewColor(1, 1, 1))

camera := scene.NewCamera(400, 200, math.Pi/3)
camera.Transform = scene.ViewTransform(core.NewPoint(0, 1.5, -5),
	core.NewPoint(0, 1, 0), core.NewVector(0, 1, 0))
return saveCanvas(camera.Render(world), "scene.png")
```

Each object now owns its material — this fixes the latent `leftWall.Material = floor.Material` aliasing bug. `simple_scene.go` uses the same helpers (`newSphere(core.ScalingMatrix(1.5, 1.5, 1.5).Translate(-0.5, 0, 0), newMaterial(core.NewColor(1, 0.2, 1), 0.9, 0.9))`).

- [ ] **Step 6: Verify end-to-end**

Run: `go test ./...` then:

```bash
go run . scene && go run . simple-scene && go run . sphere && go run . shadow && go run . clock && go run . projectile-graph
go run . scene --output /tmp/claude-scene.ppm
```

Expected: all outputs created; `scene.png` visually identical to `/tmp/scene-baseline.png`; the `--output` run writes a PPM.

- [ ] **Step 7: Record Phase 3 checkpoint benchmarks**

```bash
go test ./... -run '^$' -bench . -benchmem
```

Append to `docs/benchmarks.md` under `## 2026-07-09 — Checkpoint: after refactor phases 1–3`. Numbers should be ~flat vs. the Task 6 checkpoint (this phase was structural, not performance).

- [ ] **Step 8: Commit**

```bash
git add -A
git commit -m "refactor: dedupe CLI scene setup and add --output flag"
```

Body: note the material-aliasing fix and DrawSquare bounds clipping.

---

## Phase 4 — Performance (measured, one change per task)

### Task 10: Hoist determinant out of Matrix.Invert

**Files:**
- Modify: `pkg/core/matrix.go:164-174`
- Test: existing `pkg/core/matrix_test.go` inversion tests

- [ ] **Step 1: Implement**

```go
func (m Matrix) Invert() Matrix {
	size := len(m)
	det := m.Determinant()
	if det == 0 {
		panic("matrix is not invertible")
	}
	result := NewEmptyMatrix(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			result[j][i] = m.Cofactor(i, j) / det
		}
	}
	return result
}
```

(Previously `m.Determinant()` — itself O(n!) cofactor expansion — ran once per matrix element: 16 recomputations for a 4×4. Singular matrices now fail loudly instead of yielding `±Inf` cells.)

- [ ] **Step 2: Verify and measure**

Run: `go test ./pkg/core/` (inversion tests still pass), then:

```bash
go test ./pkg/core/ -run '^$' -bench BenchmarkMatrixInvert -benchmem -count 3
```

Append to `docs/benchmarks.md` under `## 2026-07-09 — Perf: Invert determinant hoist`. Expected: roughly an order of magnitude faster than the previous checkpoint's `BenchmarkMatrixInvert`.

- [ ] **Step 3: Commit**

```bash
git add -A
git commit -m "perf: compute determinant once in Matrix.Invert"
```

### Task 11: Cache the inverse-transpose on Sphere

**Files:**
- Modify: `pkg/objects/sphere.go`
- Test: existing `pkg/objects/sphere_test.go` normal tests

- [ ] **Step 1: Implement**

Add a field and maintain it in both places the transform is set:

```go
type Sphere struct {
	transform core.Matrix
	// invert and invTranspose cache transform.Invert() and its transpose;
	// they are read on every intersection / normal computation.
	invert       core.Matrix
	invTranspose core.Matrix
	material     *shader.Material
}

func NewSphere() *Sphere {
	s := &Sphere{material: shader.NewMaterial()}
	s.SetTransform(core.Identity(4))
	return s
}

func (s *Sphere) SetTransform(m core.Matrix) {
	s.transform = m
	s.invert = m.Invert()
	s.invTranspose = s.invert.Transpose()
}
```

`NormalAt` line `s.invert.Transpose().MultiplyWithTuple(objectNormal)` → `s.invTranspose.MultiplyWithTuple(objectNormal)`.

- [ ] **Step 2: Verify and measure**

Run: `go test ./...`, then:

```bash
go test ./pkg/objects/ -run '^$' -bench BenchmarkSphereNormalAt -benchmem -count 3
```

Append results under `## 2026-07-09 — Perf: cached inverse-transpose`. Expected: `BenchmarkSphereNormalAt` drops by the cost of one 4×4 transpose + allocation per call.

- [ ] **Step 3: Commit**

```bash
git add -A
git commit -m "perf: cache inverse-transpose for sphere normals"
```

### Task 12: Cache camera inverse transform and ray origin

**Files:**
- Modify: `pkg/scene/camera.go`
- Modify: `cmd/raytracer/scene.go`, `cmd/raytracer/simple_scene.go`, `pkg/scene/camera_test.go`, `pkg/scene/bench_test.go` (call-site rename)

**Interfaces:**
- Produces: `Camera.Transform` field is replaced by `SetTransform(m core.Matrix)` / `Transform() core.Matrix`; internally `inverse core.Matrix` and `origin core.Tuple` are cached. `RayForPixel` no longer inverts anything.

- [ ] **Step 1: Implement**

```go
type Camera struct {
	HSize      int
	VSize      int
	FOV        float64
	PixelSize  float64
	HalfWidth  float64
	HalfHeight float64

	transform core.Matrix
	// inverse and origin cache transform.Invert() and the camera position;
	// RayForPixel reads them for every pixel.
	inverse core.Matrix
	origin  core.Tuple
}

func (c *Camera) SetTransform(m core.Matrix) {
	c.transform = m
	c.inverse = m.Invert()
	c.origin = c.inverse.MultiplyWithTuple(core.NewPoint(0, 0, 0))
}

func (c *Camera) Transform() core.Matrix { return c.transform }
```

`NewCamera` calls `c.SetTransform(core.Identity(4))` after computing pixel size. `RayForPixel` body:

```go
func (c *Camera) RayForPixel(x, y int) *objects.Ray {
	xOffset := (float64(x) + .5) * c.PixelSize
	yOffset := (float64(y) + .5) * c.PixelSize
	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset
	pixel := c.inverse.MultiplyWithTuple(core.NewPoint(worldX, worldY, -1))
	direction := pixel.Subtract(c.origin).Normalize()
	return objects.NewRay(c.origin, direction)
}
```

Call sites: `camera.Transform = scene.ViewTransform(...)` → `camera.SetTransform(scene.ViewTransform(...))` in `scene.go`, `simple_scene.go`, `camera_test.go`, `bench_test.go`.

- [ ] **Step 2: Verify and measure**

Run: `go test ./...` and `go run . scene` (image must match baseline), then:

```bash
go test ./pkg/scene/ -run '^$' -bench BenchmarkCameraRender -benchmem -count 3
```

Append under `## 2026-07-09 — Perf: cached camera inverse`. Expected: the largest single improvement in the plan — the per-pixel `Invert()` (two of them, via `RayForPixel`) dominated `Render`.

- [ ] **Step 3: Commit**

```bash
git add -A
git commit -m "perf: precompute camera inverse transform and origin"
```

### Task 13: Parallel Render and final measurements

**Files:**
- Modify: `pkg/scene/camera.go` (Render), `docs/benchmarks.md`

- [ ] **Step 1: Implement scanline-parallel Render**

```go
func (c *Camera) Render(w *World) *core.Canvas {
	image := core.NewCanvas(c.HSize, c.VSize)
	var wg sync.WaitGroup
	for y := 0; y < c.VSize; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			for x := 0; x < c.HSize; x++ {
				image.Set(x, y, w.ColorAt(c.RayForPixel(x, y)))
			}
		}(y)
	}
	wg.Wait()
	return image
}
```

Safe without locks: each goroutine owns one scanline and every `(x, y)` element of the pixel grid is written exactly once.

- [ ] **Step 2: Verify with the race detector**

Run: `go test -race ./pkg/scene/` and `go test ./...`
Expected: PASS, no race reports. Then `go run . scene` — image must match baseline exactly (parallelism must not change output).

- [ ] **Step 3: Final measurements and summary**

```bash
go test ./... -run '^$' -bench . -benchmem -count 3
```

Append under `## 2026-07-09 — Final (after Phase 4)`, then add a closing summary section to `docs/benchmarks.md`: a small table with one row per benchmark, columns `Baseline`, `After value semantics`, `Final`, and a sentence naming the two dominant wins (camera inverse caching, value-type allocation elimination).

- [ ] **Step 4: Commit**

```bash
git add -A
git commit -m "perf: render scanlines concurrently"
```

Body: include the headline speedup numbers from the summary table.

---

## Completion

After Task 13: run the full suite one last time (`go test -race ./...`), regenerate all six images and eyeball them, then use superpowers:finishing-a-development-branch to merge/PR. Per user conventions: squash merge, single-paragraph body, worktree removed and branch deleted after merge.

Explicitly out of scope (deferred, not forgotten): fixed-size `[4][4]float64` matrix representation (breaks the 2×2/3×3 book tests for submatrix/determinant), shadows/chapter 8 features, and any rendering behavior changes.
