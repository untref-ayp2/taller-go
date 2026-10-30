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
	@golangci-lint version 2>/dev/null | grep -q 'version 2' || { echo 'Error: se requiere golangci-lint v2.x. Instalalo: https://golangci-lint.run/welcome/install/'; exit 1; }
	golangci-lint run ./...

clean:
	rm -f *_test.txt _*.txt _origen.txt _destino.txt _bitacora.txt _entrada.txt _salida.txt _nums.txt _test_vacio.txt
