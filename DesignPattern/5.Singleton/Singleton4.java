import java.util.Objects;

/**
 * @author lvtao03
 * @date 2021/1/20
 **/
public class Singleton4 {

    private static volatile Singleton4 instance;

    private Singleton4() {
    }

    public static Singleton4 getInstance() {
        if (Objects.isNull(instance)) {
            synchronized (Singleton4.class) {
                if (Objects.isNull(instance)) {
                    instance = new Singleton4();
                }
            }
        }
        return instance;
    }

}
