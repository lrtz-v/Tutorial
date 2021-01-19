package fly;

/**
 * @author lvtao03
 * @date 2021/1/19
 **/
public class FlyNoWay implements FlyBehavior{

    @Override
    public void fly() {
        System.out.println("I can't fly");
    }
}
