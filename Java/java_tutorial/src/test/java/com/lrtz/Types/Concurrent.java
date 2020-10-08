package com.lrtz.Types;

import org.junit.Test;

import java.util.Hashtable;
import java.util.concurrent.*;

/**
 * @author lvtao
 * @date 2020/10/2 20:55
 */
public class Concurrent {

    @Test
    public void testHashTable() {
        Hashtable t = new Hashtable();
        t.put(1, 1); // 该方法为synchronized方法
//        触发rehash
    }

    @Test
    public void testConcurrentHashMap() {
        ConcurrentHashMap m = new ConcurrentHashMap();
        m.put(1, 1);
//        不冲突时，使用cas
//        检查是否在扩容中，帮助扩容
//        synchronized 锁节点
    }

    @Test
    public void testCopyOnWriteArrayList() {
        CopyOnWriteArrayList l = new CopyOnWriteArrayList();
        l.add(1);
//        获取ReentrantLock lock
//        copy array实现扩容
//        set new array
    }

    @Test
    public void testCopyOnWriteArraySet() {
        CopyOnWriteArraySet s = new CopyOnWriteArraySet();
        s.add(1);
//        ReentrantLock lock 加锁
//        copy array实现扩容
//        set new array
    }

    /**
     * 非阻塞队列
     */
    @Test
    public void testConcurrentLinkedQueue() {
        ConcurrentLinkedQueue q = new ConcurrentLinkedQueue();
        q.add(1);
    }

    @Test
    public void testBlockQueue() {
        ArrayBlockingQueue q1 = new ArrayBlockingQueue(1);
        LinkedBlockingQueue q2 = new LinkedBlockingQueue(1);
        PriorityBlockingQueue q3 = new PriorityBlockingQueue(1);
    }

}
