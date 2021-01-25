package state;

import gumball.GumballMachine;

/**
 * 已售空
 *
 * @author lvtao03
 * @date 2021/1/25
 **/
public class SoldOutState implements State {

    transient private GumballMachine gumballMachine;

    public SoldOutState(GumballMachine gumballMachine) {
        this.gumballMachine = gumballMachine;
    }

    @Override
    public void insertCoins() {
        System.out.println("Oos, sold out");
    }

    @Override
    public void ejectCoins() {
        System.out.println("Oos, sold out");
    }

    @Override
    public void turnCrank() {
        System.out.println("Oos, sold out");
    }

    @Override
    public void dispense() {
        System.out.println("Oos, sold out");
    }

    @Override
    public String toString() {
        return "sold out";
    }
}
