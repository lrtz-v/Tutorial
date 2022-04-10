checksum.o: checksum.c checksum.h
	gcc -c checksum.c

encrypt.o: encrypt.c encrypt.h
	gcc -c encrypt.c

lib: encrypt.o checksum.o
	ar -rcs libhfsecurity.a encrypt.o checksum.o

test: test_code.c libhfsecurity.a
	gcc test_code.c -I . -L . -lhfsecurity -o test_code
