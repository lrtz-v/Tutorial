# 编译过程

- 得到所有源文件的目标代码

```bash
gcc -c *.c
```

- 把目标文件链接起来

```bash
gcc *.o -o main
```

# 加快编译速度

## 部分文件修改

- 如果要修改其中一个文件，只需要重新编译这一个文件,然后重新链接程序即可

```bash
gcc -c modifyed_file.c
gcc *.o -o main
```

## 记不住修改了哪些文件

- make

```Makefile
encrypt.o: encrypt.c encrypt.h
	gcc -c encrypt.c

main.o: main.c encrypt.h
	gcc -c main.c

main: main.o encrypt.o
	gcc main.o encrypt.o -o main
```

## 自动化生成 makefile

- autoconf
