.PHONY: build test clean prod

build: build/miniaudio.o
	go build -tags=dev .

build/miniaudio.o: src/miniaudio.c
	gcc -c -o build/miniaudio.o src/miniaudio.c

prod:
	go build .

test:
	go test -tags=dev .

clean:
	rm build/miniaudio.o
