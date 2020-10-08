package com.lrtz.locks;

import java.util.concurrent.locks.ReentrantLock;

/**
 * @author lvtao
 * @date 2020/10/8 13:39
 *
 * synchronized 和 ReentrantLock 的区别
 *
 * 1.两者都是可重入锁
 *      “可重入锁” 指的是自己可以再次获取自己的内部锁
 *      同一个线程每次获取锁，锁的计数器都自增 1，所以要等到锁的计数器下降为 0 时才能释放锁
 * 2.synchronized 依赖于 JVM 而 ReentrantLock 依赖于 API
 *      synchronized 是依赖于 JVM 实现的
 *      ReentrantLock 是 JDK 层面实现的
 * 3.ReentrantLock 比 synchronized 增加了一些高级功能
 *      等待可中断
 *          ReentrantLock提供了一种能够中断等待锁的线程的机制，通过 lock.lockInterruptibly() 来实现这个机制。也就是说正在等待的线程可以选择放弃等待，改为处理其他事情
 *      可实现公平锁
 *          ReentrantLock可以指定是公平锁还是非公平锁。而synchronized只能是非公平锁
 *      可实现选择性通知（锁可以绑定多个条件）
 *          synchronized关键字与wait()和notify()/notifyAll()方法相结合可以实现等待/通知机制。
 *          ReentrantLock类当然也可以实现，但是需要借助于Condition接口与newCondition()方法
 *
 */
public class ReentrantLockDemo {

    public void method() {
        ReentrantLock lock = new ReentrantLock();
        lock.lock();
        try {
            //
        } catch (Exception e) {
            //
        } finally {
            lock.unlock();
        }
    }
}
