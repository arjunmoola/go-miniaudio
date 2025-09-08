.PHONY: build test clean

build: build/miniaudio.o
	go build .

build/miniaudio.o: src/miniaudio.c
	gcc -c -o build/miniaudio.o src/miniaudio.c

test:
	go test .

clean:
	rm build/miniaudio.o
