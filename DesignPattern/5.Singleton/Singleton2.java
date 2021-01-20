import java.util.Objects;

/**
 * @author lvtao03
 * @date 2021/1/20
 **/
public class Singleton2 {

    private static Singleton2 instance;

    private Singleton2() {
    }

    /**
     * 只有第一次执行方法时，才真正需要同步
     *
     * @return Singleton
     */
    public static synchronized Singleton2 getInstance() {
        if (Objects.isNull(instance)) {
            instance = new Singleton2();
        }
        return instance;
    }
}
