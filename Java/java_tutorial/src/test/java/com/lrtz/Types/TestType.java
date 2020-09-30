package com.lrtz.Types;

import org.junit.Test;

/**
 * @author lvtao
 * @date 2020/9/30 17:03
 */
public class TestType {

    @Test
    public void testType() {

        // 通过valueOf方法可以理解以下结果

        Integer i1 = 100;
        Integer i2 = 100;
        Integer i3 = 200;
        Integer i4 = 200;

        System.out.println(i1 == i2);
        System.out.println(i3 == i4);

        Double d1 = 100.0;
        Double d2 = 100.0;
        Double d3 = 200.0;
        Double d4 = 200.0;

        System.out.println(d1==d2);
        System.out.println(d3==d4);

        Boolean b1 = false;
        Boolean b2 = false;
        Boolean b3 = true;
        Boolean b4 = true;

        System.out.println(b1==b2);
        System.out.println(b3==b4);


        // 装箱：自动根据数值创建对应的 Integer对象， Integer i = 10;
        // 拆箱：自动将包装器类型转换为基本数据类型，  int n = i;

        // 当 "=="运算符的两个操作数都是 包装器类型的引用，则是比较指向的是否是同一个对象，
        // 而如果其中有一个操作数是表达式（即包含算术运算）则比较的是数值（即会触发自动拆箱的过程）
        // 另外，对于包装器类型，equals方法并不会进行类型转换
        Integer a = 1;
        Integer b = 2;
        Integer c = 3;
        Integer d = 3;
        Integer e = 321;
        Integer f = 321;
        Long g = 3L;
        Long h = 2L;

        System.out.println(c==d);  // true
        System.out.println(e==f);  // false
        System.out.println(c==(a+b));  // true，  a+b包含了算术运算，因此会触发自动拆箱过程
        System.out.println(c.equals(a+b));  // true，自动拆箱-->自动装箱
        System.out.println(g==(a+b));  // true
        System.out.println(g.equals(a+b));  // false  equals 不会自动转换类型
        System.out.println(g.equals(a+h));  // true
    }

}
