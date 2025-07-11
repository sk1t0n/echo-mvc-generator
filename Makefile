.PHONY: run_tests

run_tests:
	go test github.com/sk1t0n/echo-mvc-generator/cmd
	go test github.com/sk1t0n/echo-mvc-generator/lib
