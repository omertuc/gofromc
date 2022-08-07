all: main

main: foo.a bar.c
	gcc bar.c foo.a -o main

foo.a: foo.go
	go build -buildmode=c-archive foo.go

clean:
	rm -f main *.o
