main_package := ./cmd/lichess-tui/
binary_name := lichess-tui
build_dir := build

.PHONY: build
build: format
	go build -o $(build_dir)/$(binary_name) $(main_package)

.PHONY: run
run: format
	go run $(main_package)

.PHONY: clean
clean:
	rm -rf $(build_dir)

.PHONY: install
install:
	go install $(main_package)

.PHONY: format
format:
	gofumpt -w .
	golines -w .
	goimports -w .

.PHONY: tidy
tidy: format
	go mod tidy
