package pizza_store;

import pizza.Pizza;

/**
 * @author lvtao03
 * @date 2021/1/20
 **/
public abstract class PizzaStore {

    public PizzaStore() {
    }

    /**
     * 创建pizza
     *
     * @param type pizza类型
     * @return pizza
     */
    abstract Pizza createPizza(String type);

    /**
     * 点单
     *
     * @param type pizza类型
     * @return pizza
     */
    public Pizza orderPizza(String type) {
        Pizza pizza = createPizza(type);

        pizza.prepare();
        pizza.bake();
        pizza.cut();
        pizza.box();

        return pizza;
    }
}
