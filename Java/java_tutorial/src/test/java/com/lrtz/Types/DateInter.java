package com.lrtz.Types;

import java.util.Date;

/**
 * @author lvtao
 * @date 2020/9/29 23:14
 */
class DateInter extends Pair<Date> {

    // 父类的类型是Object，而子类的类型是Date，参数类型不一样，这如果是在普通的继承关系中，根本就不会是重写，而是重载
    @Override
    public void setValue(Date value) {
        super.setValue(value);
    }

    @Override
    public Date getValue() {
        return super.getValue();
    }
}


/*
* 使用javap -c className 反编译class文件
*
最后的两个方法，就是编译器自己生成的桥方法。可以看到桥方法的参数类型都是Object
子类中真正覆盖父类两个方法的就是这两个我们看不到的桥方法， 桥方法的内部实现，就只是去调用我们自己重写的那两个方法

class com.lrtz.Types.DateInter extends com.lrtz.Types.Pair<java.util.Date> {
  com.lrtz.Types.DateInter();
    Code:
       0: aload_0
       1: invokespecial #1                  // Method com/lrtz/Types/Pair."<init>":()V
       4: return

  public void setValue(java.util.Date);
    Code:
       0: aload_0
       1: aload_1
       2: invokespecial #2                  // Method com/lrtz/Types/Pair.setValue:(Ljava/lang/Object;)V
       5: return

  public java.util.Date getValue();
    Code:
       0: aload_0
       1: invokespecial #3                  // Method com/lrtz/Types/Pair.getValue:()Ljava/lang/Object;
       4: checkcast     #4                  // class java/util/Date
       7: areturn

  public void setValue(java.lang.Object);
    Code:
       0: aload_0
       1: aload_1
       2: checkcast     #4                  // class java/util/Date
       5: invokevirtual #5                  // Method setValue:(Ljava/util/Date;)V
       8: return

  public java.lang.Object getValue();
    Code:
       0: aload_0
       1: invokevirtual #6                  // Method getValue:()Ljava/util/Date;
       4: areturn
}


* */