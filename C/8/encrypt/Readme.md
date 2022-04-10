# 静态库

## Makefile

- 务必把存档命名为 libXXX.a，否则编译器找不到 它们

```Makefile
checksum.o: checksum.c checksum.h
	gcc -c checksum.c

encrypt.o: encrypt.c encrypt.h
	gcc -c encrypt.c

lib: encrypt.o checksum.o
	ar -rcs libhfsecurity.a encrypt.o checksum.o

test: test_code.c libhfsecurity.a
	gcc test_code.c -I . -L . -lhfsecurity -o test_code
```

## 创建 lib 存档

```bash
make lib
```

## 检查 lib

```bash
~ nm libhfsecurity.a

~ ar -t libhfsecurity.a

~ ar -x libhfsecurity.a encrypt.o
```

# lib 目录

- 把头文件保存在标准目录中： 例如/usr/local/include
- 在 include 语句中使用完整路径名：#include "/my_header_files/encrypt.h"
- 你可以告诉编译器去哪里找头文件：gcc -I/my_header_files test_code.c ... -o test_code
