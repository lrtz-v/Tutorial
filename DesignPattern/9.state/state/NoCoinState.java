package state;

import gumball.GumballMachine;

/**
 * 没有投入硬币
 *
 * @author lvtao03
 * @date 2021/1/25
 **/
public class NoCoinState implements State {

    transient private GumballMachine gumballMachine;

    public NoCoinState(GumballMachine gumballMachine) {
        this.gumballMachine = gumballMachine;
    }

    @Override
    public void insertCoins() {
        System.out.println("insert coins");
        gumballMachine.setState(gumballMachine.getHasCoinState());
    }

    @Override
    public void ejectCoins() {
        System.out.println("you haven't inserted coins");
    }

    @Override
    public void turnCrank() {
        System.out.println("you turned, but no coins");
    }

    @Override
    public void dispense() {
        System.out.println("insert coins first");
    }

    @Override
    public String toString() {
        return "Has no coin";
    }
}
