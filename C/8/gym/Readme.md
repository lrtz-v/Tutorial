# 动态库

- 动态库在运行时链接程序
- 在 Windows 中，动态库通常叫动态链接库，后缀名是.dll;在 Linux 和 Unix 上，它们叫共享目标文件，后缀名.so;而在 Mac 上，它们就叫动态库，后缀名.dylib

## Makefile

- -fPIC, 它告诉 gcc 你想创建位置无关代码
- gcc -shared hfcal.o -o libhfcal.dylib[dll, dll.a, so, dylib]

```Makefile
hfcal.o:
	gcc -I./includes -fPIC -c hfcal.c -o hfcal.o
share:
	gcc -shared hfcal.o -o libs/libhfcal.dylib
elliptical.o:
	gcc -I./includes -c elliptical.c -o elliptical.o
elliptical:
	gcc elliptical.o -I . -L./libs -lhfcal -o elliptical
```
