package state;

import gumball.GumballMachine;
import java.util.Random;

/**
 * 有硬币
 *
 * @author lvtao03
 * @date 2021/1/25
 **/
public class HasCoinState implements State {

    private GumballMachine gumballMachine;
    private Random randomWinner = new Random(System.currentTimeMillis());

    public HasCoinState(GumballMachine gumballMachine) {
        this.gumballMachine = gumballMachine;
    }

    @Override
    public void insertCoins() {
        System.out.println("can't insert another coin");
    }

    @Override
    public void ejectCoins() {
        System.out.println("coin returned");
        gumballMachine.setState(gumballMachine.getNoCoinState());
    }

    @Override
    public void turnCrank() {
        System.out.println("You turned...");
        int winner = randomWinner.nextInt(10);
        if ((winner == 0) && (gumballMachine.getCount() > 1)) {
            gumballMachine.setState(gumballMachine.getWinnerState());
        } else {
            gumballMachine.setState(gumballMachine.getSoldState());
        }
    }

    @Override
    public void dispense() {
        System.out.println("no gumball dispensed");
    }
}
