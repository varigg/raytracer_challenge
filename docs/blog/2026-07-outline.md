# Writing Plan: "Debugging a Naive Raytracer Implementation"

**Audience & framing.** Go developers who know basic profiling but haven't
internalized where allocation and recomputation costs hide. The through-line is
**performance debugging as hypothesis testing**: establish a baseline, form a
cost model, change one thing, and check that the measurement matches the
theory. The ray tracer is the vehicle, not the subject. (Title thought:
"Debugging" slightly undersells it — consider *"A 530× Ray Tracer Speedup, One
Benchmark at a Time"* as an alternative or subtitle.)

**Invariant to establish early and repeat.** Every change was gated on
`scene.png` being **byte-identical** to the pre-refactor render — an oracle
stronger than the test suite. This is the article's best transferable
technique.

**Suggested length:** 2,500–3,500 words, one code pair + one number per
section, full data in an appendix.

---

## 1. Setup: the naive implementation (~300 words)

- One paragraph on the Ray Tracer Challenge book and the render pipeline
  (camera → ray per pixel → intersect → Phong shading).
- The confession that frames everything: a 400×200 scene took ~370 ms per
  frame at benchmark scale, and the code contained its own diagnosis — quote
  the original `camera.go` comments (`// Todo: precompute inverse`,
  `// this can be precomputed`). Hook: *the bottleneck was documented and
  ignored; this post is about proving it mattered.*

## 2. Measure before touching anything (~350 words)

- The benchmark harness: seven benchmarks spanning the layers (`TupleOps` →
  `Lighting` → `SphereNormalAt`/`Intersect` → `MatrixInvert` →
  `CameraRender`), names held stable across the entire effort so every later
  number is comparable.
- Show the baseline table row for `CameraRender` (369,954,233 ns/op) and the
  sink-variable pattern (`sinkCanvas = c.Render(w)`) with a sentence on why
  benchmarks need it (dead-code elimination).
- The byte-identical render oracle: `cmp scene.png baseline.png` after every
  change.
- Code pointer: `pkg/*/bench_test.go` at commit `463f35a`.

## 3. Improvement 1 — Value semantics: stop heap-allocating arithmetic (~500 words)

- **Before/after pair:** `func (t *Tuple) Add(t2 *Tuple) *Tuple { return
  &Tuple{...} }` vs `func (t Tuple) Add(o Tuple) Tuple`
  (before: `git show 24039c5^:pkg/core/tuple.go`; after: `24039c5`, plus
  `e2c9f7a` for Color).
- **Why:** every `Add`/`Multiply`/`Normalize` in the per-pixel hot path
  allocated 32 bytes; escape analysis can't save you when the API forces
  pointers. Explain the *semantic* argument too — small immutable types are
  values — so the perf win falls out of a correctness-of-design fix.
- **Numbers:** `TupleOps` 48.4→8.8 ns/op, `Lighting` 71.3→16.2 ns/op, both to
  **0 B/op, 0 allocs/op**.
- **The twist that sets up section 5:** the full render only improved ~20%
  (370→299 ms). Allocation wasn't the bottleneck — foreshadow that the reader
  (like the author) is optimizing the wrong layer. This is the article's key
  tension.

## 4. Improvement 2 — Do the math: hoisting a determinant (~400 words)

- **Before/after:** `Invert` calling `m.Determinant()` inside the
  16-iteration loop vs computing `det` once (`3358fbe`); mention the bonus
  contract fix (panic on singular instead of silent `±Inf`).
- **The cost model as the centerpiece:** cofactor-expansion inversion costs
  ~16×(4+1)=80 3×3-determinant units before, 4+16=20 after → predicted 4×;
  measured 57,884→14,640 ns/op = 3.95×. *The measurement confirming the
  back-of-envelope is the lesson* — and note honestly that the original plan
  guessed "an order of magnitude" and the math corrected it.

## 5. Improvement 3 — The documented bottleneck: cache the camera inverse (~500 words)

- **Before/after:** `RayForPixel` calling `c.Transform.Invert()` (twice, per
  pixel!) vs `SetTransform` caching `inverse` and `origin` once (`023e29c`).
  Include the original TODO comments in the "before" excerpt — they're the
  emotional core.
- **Why it dominated:** 80,000 pixels × 2 cofactor-expansion inversions each,
  each inversion ~15–58 µs — everything else was noise by comparison.
  **CameraRender: 299 ms → 0.97 ms (~310×).**
- **Design point:** the fix required replacing an exported `Transform` field
  with `SetTransform`/`Transform()` accessors — caching needs an invariant,
  and invariants need encapsulation. Same pattern applied to the sphere's
  inverse-transpose (`6d7f10c`, `NormalAt` 139→14.9 ns/op, 0 allocs) — cover
  that here as a compact "same disease, smaller organ" subsection rather than
  its own section.

## 6. Improvement 4 — Parallelism last, and why it "only" gave 1.4× (~400 words)

- **Before/after:** sequential pixel loop vs goroutine-per-scanline with
  `sync.WaitGroup` (`0950ab3`); the three-line safety argument (each
  goroutine owns a row; every pixel written exactly once; camera state is
  read-only after the caching fix — which is *why* this became trivially
  race-free).
- **The humility number:** 966→701 µs on a 6-core/12-thread machine. Explain
  via workload granularity: a 100×50 benchmark canvas is 50 goroutines ×
  ~14 µs — spawn/schedule overhead eats the theoretical gain; larger frames
  fare better.
- **The ordering lesson:** parallelizing the *original* code would have
  burned 12 cores inverting the same matrix 160,000 times. Fix the work, then
  distribute it.

## 7. Closing: the progression table and what's next (~250 words)

- The full Baseline → After-value-semantics → Final table from
  `docs/benchmarks.md`, headline: **~530× end to end**.
- Rank the wins: camera caching (310×) ≫ value semantics (allocation
  elimination) > determinant hoist (4×) > parallelism (1.4×) — i.e., the
  least glamorous change won, and the most glamorous came last and mattered
  least.
- Tease the next bottleneck (already identified in the log): per-ray slice
  allocations in `World.Intersect` / `Sphere.Intersect`, ~19.7k allocs per
  frame — sequel bait for the shadows chapter.

## Appendix / reproducibility box

- Repo link, the exact benchmark command
  (`go test ./... -run '^$' -bench . -benchmem -count 3`), a pointer to
  `docs/benchmarks.md`, and the commit SHAs per section so readers can
  `git show` every before/after themselves.

---

**Drafting order suggestion:** write sections 5 → 3 → 4 → 6 first (they're
self-contained and evidence-rich), then 2, then the intro and closing last
once you know what the piece actually argues. All before-code is reachable
via `git show <sha>^:<path>` since every phase merged with a merge commit —
SHAs remain valid.

**Commit SHA reference:**

| Section | Change | Commit |
|---|---|---|
| 2 | Benchmark harness + baseline | `463f35a` |
| 3 | Color value type | `e2c9f7a` |
| 3 | Tuple value type | `24039c5` |
| 3 | Matrix value receivers | `cfe5cfa` |
| 4 | Determinant hoist | `3358fbe` |
| 5 | Sphere inverse-transpose cache | `6d7f10c` |
| 5 | Camera inverse + origin cache | `023e29c` |
| 6 | Scanline-parallel Render | `0950ab3` |
