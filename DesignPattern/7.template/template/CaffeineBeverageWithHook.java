package template;

/**
 * @author lvtao03
 * @date 2021/1/21
 **/
public abstract class CaffeineBeverageWithHook {

    /**
     * 定义算法模版 准备饮料
     */
    public void prepareRecipe() {
        boilWater();
        brew();
        pourInCpu();
        if (customerWantsCondiments()) {
            addCondiments();
        }
    }

    /**
     * 醒茶/咖啡
     */
    public abstract void brew();

    /**
     * 加配料
     */
    public abstract void addCondiments();

    /**
     * 烧水
     */
    public void boilWater() {
        System.out.println("Boiling water");
    }

    /**
     * 倒入杯中
     */
    public void pourInCpu() {
        System.out.println("Pouring");
    }

    public Boolean customerWantsCondiments() {
        return true;
    }
}
