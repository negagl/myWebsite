# Go source files
SRC = cmd/main.go
# Output binary
OUT = bin/myWebsite

# Clean the builds
clean:
	rm -f $(OUT)

# Build the Go application
build: clean
	go build -o $(OUT) $(SRC)

# Test the application
test:
	go test -v ./...

# Run the application
run: build
	./$(OUT)