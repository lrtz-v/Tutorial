package beverage;

/**
 * @author lvtao03
 * @date 2021/1/20
 **/
public class Decaf extends Beverage {

    public Decaf() {
        description = "低咖啡因";
    }

    @Override
    public double cost() {
        return 1.05f;
    }
}
