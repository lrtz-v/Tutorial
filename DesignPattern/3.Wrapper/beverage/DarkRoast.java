package beverage;

/**
 * @author lvtao03
 * @date 2021/1/20
 **/
public class DarkRoast extends Beverage {

    public DarkRoast() {
        description = "深焙";
    }

    @Override
    public double cost() {
        return .99;
    }
}
