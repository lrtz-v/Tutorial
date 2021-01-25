package test;

import java.rmi.Naming;
import java.rmi.RemoteException;
import remote.GumballMachineRemote;

/**
 * @author lvtao03
 * @date 2021/1/25
 **/
public class GumballMonitorTest {

    public static void main(String[] args) throws RemoteException {

        GumballMachineRemote machine;
        try {
            machine = (GumballMachineRemote) Naming.lookup("rmi://127.0.0.1/gumball");
            int count = machine.getCount();
            System.out.println(count);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

}
