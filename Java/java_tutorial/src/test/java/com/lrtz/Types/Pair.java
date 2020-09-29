package com.lrtz.Types;

/**
 * @author lvtao
 * @date 2020/9/29 22:41
 */

// T 是一个无限定的类型变量，所以类型擦出后用Object替换
// 如果类型变量有限定，那么原始类型就用第一个边界的类型变量类替换

class Pair<T> {
    private T value;

    public T getValue() {
        return value;
    }

    public void setValue(T value) {
        this.value = value;
    }
}