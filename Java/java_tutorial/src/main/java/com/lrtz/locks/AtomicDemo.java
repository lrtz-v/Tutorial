package com.lrtz.locks;

import java.util.concurrent.atomic.AtomicInteger;

/**
 * @author lvtao
 * @date 2020/10/8 14:54
 */
public class AtomicDemo {

    private AtomicInteger val = new AtomicInteger();

    public static void main(String[] args) {
        // 初始化实例
        AtomicDemo atomicDemo = new AtomicDemo();

        for (int i = 0; i < 10; ++i)
        {
            new Thread(atomicDemo::increase).start();
        }

        // 让主线程休眠5秒，保证前面起的10个异步线程都执行完毕
        try {
            Thread.sleep(3000);
        } catch (InterruptedException e)
        {
            e.printStackTrace();
        }
        System.out.println(atomicDemo.getVal().toString());
    }

    private void increase()
    {
        for (int i = 0; i < 1000; ++i)
        {
            this.val.incrementAndGet();
        }
    }

    private AtomicInteger getVal()
    {
        return this.val;
    }
}
