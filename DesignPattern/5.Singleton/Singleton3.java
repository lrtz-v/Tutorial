/**
 * @author lvtao03
 * @date 2021/1/20
 **/
public class Singleton3 {

    private static Singleton3 instance = new Singleton3();

    private Singleton3() {
    }

    /**
     * 加载这个类时就会创建单利实例
     *
     * @return Singleton
     */
    public static Singleton3 getInstance() {
        return instance;
    }

}
