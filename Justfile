# List available recipes
# Run the tests of the project
test:
    go test -count=1 -p 1 -shuffle=on -coverprofile=coverage.txt -covermode=atomic ./...

list:
    @just --list
