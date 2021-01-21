package beverage;

import template.CaffeineBeverage;

/**
 * @author lvtao03
 * @date 2021/1/21
 **/
public class Coffee extends CaffeineBeverage {

    @Override
    public void brew() {
        System.out.println("Dripping coffee through filter");
    }

    @Override
    public void addCondiments() {
        System.out.println("Add milk");
    }
}
