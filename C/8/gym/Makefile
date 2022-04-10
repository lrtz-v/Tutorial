hfcal.o:
	gcc -I./includes -fPIC -c hfcal.c -o hfcal.o
share:
	gcc -shared hfcal.o -o libs/libhfcal.dylib
elliptical.o:
	gcc -I./includes -c elliptical.c -o elliptical.o
elliptical:
	gcc elliptical.o -I . -L./libs -lhfcal -o elliptical
