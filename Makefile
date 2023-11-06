
test:
	go test -v ./...

publish:
	GOPROXY=proxy.golang.org go list -m github.com/scalingparrots/go-eth-tools@v0.1.0