package state;

import gumball.GumballMachine;

/**
 * 赢家
 *
 * @author lvtao03
 * @date 2021/1/25
 **/
public class WinnerState implements State {

    transient private GumballMachine gumballMachine;

    public WinnerState(GumballMachine gumballMachine) {
        this.gumballMachine = gumballMachine;
    }

    @Override
    public void insertCoins() {
        System.out.println("you win");
    }

    @Override
    public void ejectCoins() {
        System.out.println("you win");
    }

    @Override
    public void turnCrank() {
        System.out.println("you win");
    }

    @Override
    public void dispense() {
        System.out.println("you win, you got two gumballs for your coin");
        gumballMachine.releaseBall();
        if (gumballMachine.getCount() == 0) {
            gumballMachine.setState(gumballMachine.getSoldOutState());
        } else {
            gumballMachine.releaseBall();
            if (gumballMachine.getCount() > 0) {
                gumballMachine.setState(gumballMachine.getNoCoinState());
            } else {
                gumballMachine.setState(gumballMachine.getSoldOutState());
            }
        }
    }

    @Override
    public String toString() {
        return "winner";
    }
}
