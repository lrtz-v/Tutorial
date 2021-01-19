package duck;

import fly.FlyWithWings;
import quack.Quack;

/**
 * @author lvtao03
 * @date 2021/1/19
 **/
public class MallardDuck extends Duck {

    public MallardDuck() {
        quackBehavior = new Quack();
        flyBehavior = new FlyWithWings();
    }

    @Override
    public void display() {
        System.out.println("I'm a real Mallard duck.");
    }

}
