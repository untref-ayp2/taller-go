.PHONY: build test fmt clean

build:
	go build ./...

test:
	go test ./...

fmt:
	go fmt ./...

test-ejercicios:
	go test ./... 2>&1 | grep -E "^(ok|FAIL|---)" || true

clean:
	rm -f *_test.txt _*.txt _config.txt _origen.txt _destino.txt _bitacora.txt _entrada.txt _salida.txt _nums.txt _test_vacio.txt
