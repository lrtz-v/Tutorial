package pizza;

import java.util.ArrayList;

/**
 * @author lvtao03
 * @date 2021/1/20
 **/
public abstract class Pizza {

    String name;
    String dough;
    String sauce;

    ArrayList<String> toppings = new ArrayList<>();

    public void prepare() {
        System.out.println("Preparing " + name);
        System.out.println("Adding topping...");
        for (String s : toppings) {
            System.out.println("   " + s);
        }
    }

    public void bake() {
        System.out.println("Bake");
    }

    public void cut() {
        System.out.println("Cut");
    }

    public void box() {
        System.out.println("Box");
    }

    public String getName() {
        return name;
    }
}
