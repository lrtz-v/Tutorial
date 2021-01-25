package remote;

import java.rmi.Naming;
import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;

/**
 * @author lvtao03
 * @date 2021/1/25
 **/
public class MyRemoteImpl extends UnicastRemoteObject implements MyRemote {

    public MyRemoteImpl() throws RemoteException {
    }

    @Override
    public String sayHello() throws RemoteException {
        return "Server says: `Hey`";
    }

    public static void main(String[] args) {
        try {
            MyRemote service = new MyRemoteImpl();
            Naming.rebind("RemoteHello", service);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
