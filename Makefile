projectile: cmd/projectile foundation
	go build -o projectile cmd/projectile/main.go
clock: cmd/analog_clock/main.go foundation
	go build -o clock cmd/analog_clock/main.go
sphere_shadow: cmd/sphere_shadow/main.go foundation
	go build -o sphere_shadow cmd/sphere_shadow/main.go
test: foundation
	go test ./...
