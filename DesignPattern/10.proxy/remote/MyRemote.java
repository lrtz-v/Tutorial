package remote;

import java.rmi.Remote;
import java.rmi.RemoteException;

/**
 * @author lvtao03
 * @date 2021/1/25
 **/
public interface MyRemote extends Remote {

    String sayHello() throws RemoteException;
}
