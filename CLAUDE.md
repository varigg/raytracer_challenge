# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

A Go implementation of the ray tracer from "The Ray Tracer Challenge" book. Each `cmd/raytracer` CLI subcommand corresponds to a chapter's exercise/demo scene (some carry chapter aliases, e.g. `clock`/`chapter4`, `scene`/`chapter7`).

## Commands

- Build: `make build` (equivalent to `go build -o raytrace -ldflags="-s -w"`)
- Test all: `make test` (equivalent to `go test ./...`)
- Test a single package: `go test ./pkg/core/...`
- Test a single test: `go test ./pkg/scene/... -run TestWorld_ShadeHit`
- Run a scene: `go run . <command>`, e.g. `go run . scene` or `go run . clock` (outputs a `.ppm`/`.png` to the repo root)

## Architecture

Code is layered bottom-up across four packages, each depending only on the ones below it:

- `pkg/core` — math primitives with no dependency on the rest of the tree: `Tuple` (points/vectors, W=1 vs W=0), `Matrix` (transforms: translation/scaling/rotation, inversion, multiplication), `Color`, and `Canvas` (pixel grid, saves to PPM or PNG).
- `pkg/objects` — `Ray`, `Sphere`, and intersection logic. Two related interfaces live here: `Object` (in `object.go`, used by `scene.World.Objects`) and `Intersecter` (in `intersection.go`, used by `scene.Computations.Object`) — they describe overlapping but not identical method sets, so don't assume one can be swapped for the other without checking both call sites.
- `pkg/shader` — `Material` (Phong parameters: ambient/diffuse/specular/shininess) and `Light` (point light + `Lighting()`, the Phong reflection model implementation).
- `pkg/scene` — ties the above together: `World` (light + objects, `Intersect`/`ColorAt`/`ShadeHit`), `Camera` (`RayForPixel`, `Render` walks every pixel and calls `World.ColorAt`), `ViewTransform`, and `Computations` (per-hit precomputed point/eye vector/normal, used to feed `ShadeHit`).
- `cmd/raytracer` — Cobra commands, one per book chapter/scene, each hand-assembling a `World`/`Camera` and writing an image file. `root.go` wires commands together; `main.go` just calls `raytracer.Execute()`.

The render pipeline for any scene command: build spheres with materials/transforms → add them to a `World` with a `Light` → build a `Camera` with a `ViewTransform` → `camera.Render(world)` produces a `Canvas` → save to PNG/PPM. Ray/object intersection always transforms the incoming ray into object space via the object's precomputed inverse transform (see `Sphere.invert`) rather than transforming the object's geometry.
