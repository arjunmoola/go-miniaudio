.PHONY: build test clean prod

build: build/miniaudio.o
	go build -race -tags=dev .

build/miniaudio.o: src/miniaudio.c
	gcc -c -o build/miniaudio.o src/miniaudio.c

prod:
	go build .

test:
	go test -v -tags=dev .

clean:
	rm build/miniaudio.o
