package gumball;

import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;
import remote.GumballMachineRemote;
import state.HasCoinState;
import state.NoCoinState;
import state.SoldOutState;
import state.SoldState;
import state.State;
import state.WinnerState;

/**
 * @author lvtao03
 * @date 2021/1/25
 **/
public class GumballMachine extends UnicastRemoteObject implements GumballMachineRemote {

    private State soldState;
    private State soldOutState;
    private State hasCoinState;
    private State noCoinState;
    private State winnerState;
    private String location;

    /**
     * 当前状态，初始为已售完
     */
    private State state;

    /**
     * 糖果数量，初始为0
     */
    private int count = 0;

    public GumballMachine(String location, int numberGumballs) throws RemoteException {
        this.location = location;
        this.winnerState = new WinnerState(this);
        this.soldState = new SoldState(this);
        this.soldOutState = new SoldOutState(this);
        this.hasCoinState = new HasCoinState(this);
        this.noCoinState = new NoCoinState(this);

        this.count = numberGumballs;
        if (numberGumballs > 0) {
            this.state = noCoinState;
        } else {
            this.state = soldState;
        }
    }

    public void refill(int numberGumballs) {
        this.count = count;
        if (numberGumballs > 0) {
            this.state = noCoinState;
        } else {
            this.state = soldState;
        }
    }

    public State getNoCoinState() {
        return noCoinState;
    }

    public State getHasCoinState() {
        return hasCoinState;
    }

    public State getSoldState() {
        return soldState;
    }

    public State getSoldOutState() {
        return soldOutState;
    }

    /**
     * 更新状态
     *
     * @param state
     */
    public void setState(State state) {
        this.state = state;
    }

    @Override
    public State getState() throws RemoteException {
        return state;
    }

    public void insertCoin() {
        state.insertCoins();
    }

    public void ejectCoin() {
        state.ejectCoins();
    }

    public void turnCrank() {
        state.turnCrank();
        state.dispense();
    }

    public void releaseBall() {
        System.out.println("A gumball comes rolling out slot...");
        if (count != 0) {
            count--;
        }
    }

    @Override
    public int getCount() {
        return count;
    }

    public State getWinnerState() {
        return winnerState;
    }

    @Override
    public String getLocation() {
        return location;
    }
}
