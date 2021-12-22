.PHOHY: build, mclient,clean

build:
	go build -v ./cmd/main.go

mclient:
	go build -v ./cmd/client.go


clean :
	rm -f main client
