package remote;

import java.rmi.Remote;
import java.rmi.RemoteException;
import state.State;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public interface GumballMachineRemote extends Remote {

    int getCount() throws RemoteException;

    String getLocation() throws RemoteException;

    State getState() throws RemoteException;

}
