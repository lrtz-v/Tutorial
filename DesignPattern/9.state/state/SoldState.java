package state;

import gumball.GumballMachine;

/**
 * 发放糖果
 *
 * @author lvtao03
 * @date 2021/1/25
 **/
public class SoldState implements State {

    private GumballMachine gumballMachine;

    public SoldState(GumballMachine gumballMachine) {
        this.gumballMachine = gumballMachine;
    }

    @Override
    public void insertCoins() {
        System.out.println("Please wait");
    }

    @Override
    public void ejectCoins() {
        System.out.println("you already turned crank");
    }

    @Override
    public void turnCrank() {
        System.out.println("you can't turn crank twice");
    }

    @Override
    public void dispense() {
        gumballMachine.releaseBall();
        if (gumballMachine.getCount() > 0) {
            gumballMachine.setState(gumballMachine.getNoCoinState());
        } else {
            System.out.println("Oops, out of gumballs.");
            gumballMachine.setState(gumballMachine.getSoldOutState());
        }
    }
}
