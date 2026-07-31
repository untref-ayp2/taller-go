.PHONY: build test test-v fmt lint clean

build:
	go build ./...

test:
	go test ./...

test-v:
	go test -v ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

clean:
	rm -f *_test.txt _*.txt _origen.txt _destino.txt _bitacora.txt _entrada.txt _salida.txt _nums.txt _test_vacio.txt
