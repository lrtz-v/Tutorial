package com.lrtz.Types;

/**
 * @author lvtao
 * @date 2020/9/29 23:45
 */

// 泛型类中的静态方法和静态变量不可以使用泛型类所声明的泛型类型参数
// 因为泛型类中的泛型参数的实例化是在定义对象的时候指定的，而静态变量和静态方法不需要使用对象来调用。对象都没有创建，如何确定这个泛型参数是何种类型，所以当然是错误的
public class Test2<T> {
//    public static T one;   //编译错误
//
//    public static T show(T one) { //编译错误
//        return null;
//    }
}
