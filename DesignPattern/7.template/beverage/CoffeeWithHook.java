package beverage;

import template.CaffeineBeverageWithHook;

/**
 * @author lvtao03
 * @date 2021/1/21
 **/
public class CoffeeWithHook extends CaffeineBeverageWithHook {

    private Boolean wantsCondiments;

    @Override
    public void brew() {
        System.out.println("Dripping coffee through filter");
    }

    @Override
    public void addCondiments() {
        System.out.println("Add milk");
    }

    @Override
    public Boolean customerWantsCondiments() {
        return wantsCondiments;
    }

    public void setWantsCondiments(Boolean wants) {
        wantsCondiments = wants;
    }
}
