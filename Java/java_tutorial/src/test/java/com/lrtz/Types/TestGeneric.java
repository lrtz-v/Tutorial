package com.lrtz.Types;

import org.junit.Test;

import java.lang.reflect.InvocationTargetException;
import java.util.ArrayList;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

/**
 * @author lvtao
 * @date 2020/9/29 22:40
 *
 *
 *
 * 泛型一般有三种使用方式:泛型类、泛型接口、泛型方法
 *
 * 经常用T、E、K、V等形式的参数来表示泛型参数
 * T：代表一般的任何类
 * E：代表 Element 的意思，或者 Exception 异常的意思
 * K：代表 Key 的意思
 * V：代表 Value 的意思，通常与 K 一起配合使用
 * N: Number
 *
 */


public class TestGeneric {

    // 范型的类型擦除
    @Test
    public void arrayListTypes() throws NoSuchMethodException, InvocationTargetException, IllegalAccessException {
        ArrayList<String> list1 = new ArrayList<String>();
        list1.add("abc");

        ArrayList<Integer> list2 = new ArrayList<Integer>();
        list2.add(123);

        // 原始类型相等
        System.out.println(list1.getClass() == list2.getClass());


        // 不能用类型参数替换基本类型。就比如，没有ArrayList<double>，只有ArrayList<Double>。
        // 因为当类型擦除后，ArrayList的原始类型变为Object，但是Object类型不能存储double值，只能引用Double的值
//        ArrayList<double> l = new ArrayList<>();
        ArrayList<Double> listDouble = new ArrayList<>();


        // 通过反射添加其它类型元素
        ArrayList<Integer> list = new ArrayList<Integer>();
        list.add(1);  //这样调用 add 方法只能存储整形，因为泛型类型的实例为 Integer
        list.getClass().getMethod("add", Object.class).invoke(list, "asd");

        //  不能使用foreach
        for (int i = 0; i < list.size(); i++) {
            System.out.println(list.get(i));
        }
    }

    // 类型擦除后保留的原始类型
    //在不指定泛型的情况下，泛型变量的类型为该方法中的几种类型的同一父类的最小级，直到Object
    //在指定泛型的情况下，该方法的几种类型必须是该泛型的实例的类型或者其子类
    @Test
    public void testGeneric() {
        /**不指定泛型的时候*/
        int i = TestGeneric.add(1, 2); //这两个参数都是Integer，所以T为Integer类型
        Number f = TestGeneric.add(1, 1.2); //这两个参数一个是Integer，以风格是Float，所以取同一父类的最小级，为Number
        Object o = TestGeneric.add(1, "asd"); //这两个参数一个是Integer，以风格是Float，所以取同一父类的最小级，为Object

        /**指定泛型的时候*/
        int a = TestGeneric.<Integer>add(1, 2); //指定了Integer，所以只能为Integer类型或者其子类
//        int b = TestGeneric.<Integer>add(1, 2.2); //编译错误，指定了Integer，不能为Float
        Number c = TestGeneric.<Number>add(1, 2.2); //指定为Number，所以可以为Integer和Float

        ArrayList list = new ArrayList(); // 同 ArrayList list = new ArrayList<String>()
        list.add(1);
        list.add("121");
        list.add(new Date());
        Integer num = (Integer) list.get(0);
        System.out.println(num);
        String str = (String) list.get(1);
        System.out.println(str);
        Date date = (Date) list.get(2);
        System.out.println(date);

        System.out.println(list);

    }

    public static <T> T add(T x, T y) {
        return y;
    }


    // 先检查，再编译以及编译的对象和引用传递问题
    @Test
    public void testTypeCHeckInGeneric() {
        ArrayList<String> list1 = new ArrayList();
        list1.add("1"); //编译通过

//        list1.add(1); //编译错误, 说明这就是在编译之前的检查，因为如果是在编译之后检查，类型擦除后，原始类型为Object，是应该允许任意引用类型添加的
        // new ArrayList()只是在内存中开辟了一个存储空间，可以存储任何类型对象，而真正设计类型检查的是它的引用

        String str1 = list1.get(0); //返回类型就是String

        ArrayList list2 = new ArrayList<String>();
        list2.add("1"); //编译通过
        list2.add(1); //编译通过
        Object object = list2.get(0); //返回类型就是Object

        new ArrayList<String>().add("11"); //编译通过
//        new ArrayList<String>().add(22); //编译错误


//        String str2 = new ArrayList<String>().get(0); //返回类型就是String
//        通过get方法查询数据， 在return之前，会根据泛型变量进行强转
//        (E) elementData[index]

    }

    // 类型擦除与多态的冲突和解决方法
    @Test
    public void testExtends() {
        DateInter dateInter = new DateInter();
        dateInter.setValue(new Date());
//        dateInter.setValue(new Object()); //编译错误

        Map m = new HashMap<>();
    }


}
