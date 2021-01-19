package fly;

/**
 * @author lvtao03
 * @date 2021/1/19
 **/
public class FlyRocketPowered implements FlyBehavior {

    @Override
    public void fly() {
        System.out.println("I'm flying with a rocket!");
    }
}
