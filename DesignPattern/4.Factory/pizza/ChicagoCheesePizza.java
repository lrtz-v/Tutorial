package pizza;

/**
 * @author lvtao03
 * @date 2021/1/20
 **/
public class ChicagoCheesePizza extends Pizza {

    public ChicagoCheesePizza() {
        name = "ChicagoCheesePizza";
        dough = "";
        sauce = "";
        toppings.add("Shredded Mozzarella Cheese");
    }

    @Override
    public void cut() {
        System.out.println("Cutting into square slices");
    }
}
