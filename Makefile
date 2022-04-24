projectile: cmd/projectile foundation
	go build -o dist/projectile cmd/projectile/main.go
clock: cmd/analog_clock/main.go foundation
	go build -o dist/clock cmd/analog_clock/main.go
sphere_shadow: cmd/sphere_shadow/main.go foundation
	go build -o dist/sphere_shadow cmd/sphere_shadow/main.go
test: foundation
	go test ./...
