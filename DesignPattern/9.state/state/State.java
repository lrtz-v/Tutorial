package state;

/**
 * 状态接口，方法直接映射到糖果机上可能发生的动作
 *
 * 没有投入硬币 =》 投入硬币 =》发放硬币 =》 没有投入硬币
 *                  👇
 *                  已售空 =》 退回硬币
 * @author lvtao03
 * @date 2021/1/25
 **/
public interface State {

    /**
     * 投入硬币
     */
    void insertCoins();

    /**
     * 退换硬币
     */
    void ejectCoins();

    /**
     * 旋转手柄
     */
    void turnCrank();

    /**
     * 发放糖果
     */
    void dispense();

}
