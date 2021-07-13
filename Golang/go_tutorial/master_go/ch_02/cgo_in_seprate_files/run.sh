#/bin/bash

gcc -c callClib/*.c
ar rs callC.a *.o
rm *.o
go build ./callC.go
rm *.a
./callC
rm ./callC
