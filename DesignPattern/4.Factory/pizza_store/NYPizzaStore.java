package pizza_store;

import pizza.ClamPizza;
import pizza.NyCheesePizza;
import pizza.PepperoniPizza;
import pizza.Pizza;
import pizza.VeggiePizza;

/**
 * @author lvtao03
 * @date 2021/1/20
 **/
public class NYPizzaStore extends PizzaStore {

    @Override
    public Pizza createPizza(String type) {
        Pizza pizza = null;
        if ("cheese".equals(type)) {
            pizza = new NyCheesePizza();
        } else if ("perpperoni".equals(type)) {
            pizza = new PepperoniPizza();
        } else if ("clam".equals(type)) {
            pizza = new ClamPizza();
        } else if ("veggie".equals(type)) {
            pizza = new VeggiePizza();
        }
        return pizza;
    }
}
