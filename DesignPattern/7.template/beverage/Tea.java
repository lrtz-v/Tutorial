package beverage;

import template.CaffeineBeverage;

/**
 * @author lvtao03
 * @date 2021/1/21
 **/
public class Tea extends CaffeineBeverage {

    @Override
    public void brew() {
        System.out.println("Steeping the tea");
    }

    @Override
    public void addCondiments() {
        System.out.println("Add lemon");
    }
}
