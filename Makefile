.PHONY: run
run: 
	go run .


.PHONY: test
test:
	go test -coverprofile cov.out ./...
