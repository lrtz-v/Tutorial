package com.lrtz;

import org.junit.Test;

import java.util.Date;
import java.util.concurrent.*;

/**
 * @author lvtao
 * @date 2020/10/8 14:22
 */
public class TestExecutors {

    private static final int CORE_POOL_SIZE = 5;
    private static final int MAX_POOL_SIZE = 10;
    private static final int QUEUE_CAPACITY = 100;
    private static final Long KEEP_ALIVE_TIME = 1L;

    @Test
    public void testExecutors() {
        ExecutorService P1 = Executors.newFixedThreadPool(1);
        ExecutorService P2 = Executors.newCachedThreadPool();
        ExecutorService P3 = Executors.newScheduledThreadPool(1);
        ExecutorService P4 = Executors.newSingleThreadExecutor();
        ExecutorService P5 = Executors.newSingleThreadScheduledExecutor();
        ExecutorService P6 = Executors.newWorkStealingPool();
    }

    public class MyRunnable implements Runnable {

        private String command;

        public MyRunnable(String s) {
            this.command = s;
        }

        @Override
        public void run() {
            System.out.println(Thread.currentThread().getName() + " Start. Time = " + new Date());
            processCommand();
            System.out.println(Thread.currentThread().getName() + " End. Time = " + new Date());
        }

        private void processCommand() {
            try {
                Thread.sleep(5000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }

        @Override
        public String toString() {
            return this.command;
        }
    }

    @Test
    public void testThreadPoolExecutor() {
        //使用阿里巴巴推荐的创建线程池的方式
        //通过ThreadPoolExecutor构造函数自定义参数创建
        ThreadPoolExecutor executor = new ThreadPoolExecutor(
                CORE_POOL_SIZE,
                MAX_POOL_SIZE,
                KEEP_ALIVE_TIME,
                TimeUnit.SECONDS,
                new ArrayBlockingQueue<>(QUEUE_CAPACITY),
                new ThreadPoolExecutor.CallerRunsPolicy());

        for (int i = 0; i < 10; i++) {
            //创建WorkerThread对象（WorkerThread类实现了Runnable 接口）
            Runnable worker = new MyRunnable("" + i);
            //执行Runnable
            executor.execute(worker);
        }
        //终止线程池
        executor.shutdown();
        while (!executor.isTerminated()) {

        }
        System.out.println("Finished all threads");
    }

}
