package observer;

/**
 * @author lvtao03
 * @date 2021/1/19
 **/
public interface Observer {

    /**
     * 更新观察者数据
     *
     * @param temp     温度
     * @param humidity 湿度
     * @param pressure 气压
     */
    void update(float temp, float humidity, float pressure);
}
