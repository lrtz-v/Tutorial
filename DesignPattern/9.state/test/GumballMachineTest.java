package test;

import gumball.GumballMachine;
import java.lang.management.GarbageCollectorMXBean;

/**
 * @author lvtao03
 * @date 2021/1/25
 **/
public class GumballMachineTest {

    public static void main(String[] args) {
        GumballMachine gumballMachine = new GumballMachine(5);
        System.out.println(gumballMachine);

        gumballMachine.insertCoin();
        gumballMachine.turnCrank();
        System.out.println(gumballMachine);

        gumballMachine.insertCoin();
        gumballMachine.turnCrank();
        gumballMachine.insertCoin();
        gumballMachine.turnCrank();
        System.out.println(gumballMachine);
    }

}
