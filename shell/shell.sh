#!/bin/bash 

# ============Variable============

# define variable
name="name1"

# use variable
echo $name
echo ${name}

# update variable
name="name2"

# readonly variable
readonly name2="readonly_variable"
echo $name2
# name2="try_update_readonly_variable" # error
# echo ${name2}

# delete variable
unset name2

# ============String============

# use variable in string
echo "my name is ${name}"
# variable cannot be used in signle quota string
echo 'my name is ${name}'

# get string length 
var_length=
echo "string length: ${#name2}"

# substring
echo "sub string 1: ${name2:1:5}"
echo "sub string 2: ${name2::5}"

# string contact
echo "this is"" my name"
echo "this is my name"
echo "this" is "my name"
echo 'this is'' my nam'
echo 'this is my name'
echo 'this' is 'my name'

# ============Array============

# array
array_name=("li" "wang" "xiang" "zhang")
array_name[4]="test"
array_name[0]="lv"
echo "First element: ${array_name[0]}"
echo "First element length: ${#array_name[0]}"
echo "All elements: ${array_name[@]}"
echo "Array length: ${#array_name[@]}"


# ============Parameter============

echo "filename: $0"
echo "first parameter: $1"
echo "paramenter number: $#"
echo "all paramenter in one: $*"
echo "all paramenter in one: $@"
echo "pid: $$"
echo "last pid: $!"
echo "last command status: $?"

# ==========Calculate with number=============

echo `expr 2 + 2`
echo $[2+2]
echo $((2+2))
echo $((2-2))
echo $((2*2))
echo $((2/2))

val=3
echo $((val++))
echo $((val--))
echo $((++val))
echo $((--val))
echo $((!val))
echo $((~val))
echo $((val**val))
echo $((val<<2))
echo $((val>>2))

a=10
b=20

[ $a -eq $b ] && echo "$a equal $b"
[ $a -ne $b ] && echo "$a not equal $b"
[ $a -gt $b ] && echo "$a bigger than $b"
[ $a -ge $b ] && echo "$a bigger|equal than $b"
[ $a -lt $b ] && echo "$b bigger than $a"
[ $a -le $b ] && echo "$b bigger|equal than $a"

# ==========Calculate with string=============

a="abc"
b="efg"

[ $a = $b ] && echo "$a equal $b"
[ $a != $b ] && echo "$a not equal $b"
[ -z $a ] && echo "$a is empty"
[ -n $a ] && echo "$a is not empty"
[ $a ] && echo "$a is not empty"

# ==========Logic=============

a=10
b=20

[ !false ] && echo "true"
[ $a -lt 20 -o $b -gt 100 ] && echo "$a smaller than 20 OR $b bigger than 100"
[[ $a -lt 20 || $b -gt 100 ]] && echo "$a smaller than 20 OR $b bigger than 100"
[ $a -lt 20 -a $b -gt 100 ] && echo "$a smaller than 20 AND $b bigger than 100"
[[ $a -lt 20 && $b -gt 100 ]] && echo "$a smaller than 20 AND $b bigger than 100"

# ==========File=============

file=$0

[ -b $file ] && echo "$0 is block file"
[ -c $file ] && echo "$0 is char file"
[ -p $file ] && echo "$0 is pipe"
[ -d $file ] && echo "$0 is dir"
[ -f $file ] && echo "$0 is file"
[ -g $file ] && echo "$0 has SGID"
[ -k $file ] && echo "$0 has sticky bit"
[ -u $file ] && echo "$0 has SUID"
[ -r $file ] && echo "$0 is readable"
[ -w $file ] && echo "$0 is writable"
[ -x $file ] && echo "$0 is executable"
[ -s $file ] && echo "$0 is not empty"
[ -e $file ] && echo "$0 exists"

# ==========Command=============

echo `ls .`
echo $(ls /Users)
path=$(cd `dirname $0`; pwd)
echo $path


# =============Control==========

if [ -f .bash_profile ]; then
    echo "You have a .bash_profile"
else 
    echo "You have no .bash_profile"
fi

if [ "$(id -u)" = "0" ]; then
    echo "superuser"
else 
    echo "need to be superuser to continue"
fi

for i in word1 word2 word3; do
    echo "$i"
done

for i in ${array_name[@]}; do
    echo "$i"
done

count=0
for i in $(cat ~/.bash_profile); do
    count=$((count++))
    echo "Word $count ($i) contains $(echo -n $i | wc -c) characters"
done

for i in "$@"; do
    echo $i
done

i=0
while [[ $i<5 ]]
do 
    echo $i
    let "i++"
done

# echo "<CTRL-D> to exit"
# while read input
# do
#     echo "The input is $input"
# done

i=0
until [[ $i -gt 5 ]]
do
    echo Counter: $i
    ((i++))
done

# while true
# do
#     echo ""
# done

# for (( ; ; ))
# do
#     echo ""
# done


case $1 in
    a) echo "a"
    ;;
    b) echo "b"
    ;;
    *) echo "..."
    ;;
esac

f1 () {
    echo "function definition f1"
}

function f2 {
    echo "func definition f2"
}

f1
f2

f3 () {
    local name="f3"
    echo "In function $name"
}

f3


tmp=0
f4 () {
    echo "In function 4"
    tmp=1
    return 55
}

f4
echo $?
echo $tmp

# return value by send the value to stdout
f5 () {
    local func_result="some result"
    echo "$func_result"
}

func_result="$(f5)"
echo $func_result

# pass variable to function

f6 () {
    echo $@
}

f6 1 2 3 4 5 6
