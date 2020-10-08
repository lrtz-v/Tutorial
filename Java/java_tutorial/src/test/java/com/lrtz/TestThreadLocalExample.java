package com.lrtz;

import org.junit.Test;

/**
 * @author lvtao
 * @date 2020/10/8 14:16
 */
public class TestThreadLocalExample {

    @Test
    public void testThreadLocal() {
        ThreadLocal t1 = new ThreadLocal();
        ThreadLocal t2 = new ThreadLocal();
        t1.set("1");
        System.out.println(t1.get());
        System.out.println(t2.get());
    }

    @Test
    public void testInheritableThreadLocal() {
        ThreadLocal<String> ThreadLocal = new ThreadLocal<>();
        ThreadLocal.set("父类数据:threadLocal");

        ThreadLocal<String> inheritableThreadLocal = new InheritableThreadLocal<>();
        inheritableThreadLocal.set("父类数据:inheritableThreadLocal");

        new Thread(new Runnable() {
            @Override
            public void run() {
                System.out.println("子线程获取父类`ThreadLocal`数据：" + ThreadLocal.get());
                System.out.println("子线程获取父类inheritableThreadLocal数据：" + inheritableThreadLocal.get());
            }
        }).start();
    }
}
