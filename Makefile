test:
	go test -v ./...

bench:
	go test -bench=. ./benchmarking/benchmarking_test.go