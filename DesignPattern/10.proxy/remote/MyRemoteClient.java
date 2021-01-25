package remote;

import java.rmi.Naming;

/**
 * @author lvtao03
 * @date 2021/1/25
 **/
public class MyRemoteClient {

    public void go() {
        try {
            MyRemote service = (MyRemote) Naming.lookup("rmi://127.0.0.1/RemoteHello");
            String s = service.sayHello();
            System.out.println(s);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public static void main(String[] args) {
        new MyRemoteClient().go();
    }

}
