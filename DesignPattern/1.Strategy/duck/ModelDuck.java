package duck;

import fly.FlyNoWay;
import quack.MuteQuack;
import quack.Quack;

/**
 * @author lvtao03
 * @date 2021/1/19
 **/
public class ModelDuck extends Duck {

    public ModelDuck() {
        flyBehavior = new FlyNoWay();
        quackBehavior = new MuteQuack();
    }

    @Override
    public void display() {
        System.out.println("I'm a model duck");
    }

}
