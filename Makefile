.PHONY: test

test:
	go test -count=1 -mod=readonly ./...