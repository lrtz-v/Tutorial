import java.util.Objects;

/**
 * @author lvtao03
 * @date 2021/1/20
 **/
public class Singleton {

    private static Singleton instance;

    private Singleton() {
    }

    /**
     * 只适用于单线程
     *
     * @return Singleton
     */
    public static Singleton getInstance() {
        if (Objects.isNull(instance)) {
            instance = new Singleton();
        }
        return instance;
    }

}
