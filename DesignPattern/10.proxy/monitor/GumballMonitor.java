package monitor;

import gumball.GumballMachine;
import java.rmi.RemoteException;
import remote.GumballMachineRemote;

/**
 * @author lvtao03
 * @date 2021/1/25
 **/
public class GumballMonitor {

    private GumballMachineRemote machine;

    public GumballMonitor(GumballMachineRemote machine) {
        this.machine = machine;
    }

    public void report() {
        try {
            System.out.println("Gamball Machine: " + machine.getLocation());
            System.out.println("Current inventory: " + machine.getCount() + " gumballs");
            System.out.println("Current state: " + machine.getState());
        } catch (RemoteException e) {
            e.printStackTrace();
        }
    }
}
