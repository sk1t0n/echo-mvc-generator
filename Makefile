.PHONY: run_tests, run_cover_tests

run_tests:
	go test ./... -coverprofile=c.out

run_cover_tests:
	go tool cover -func=c.out
