package test;

import duck.Duck;
import duck.MallardDuck;
import duck.ModelDuck;
import duckcall.DuckCall;
import fly.FlyRocketPowered;

/**
 * @author lvtao03
 * @date 2021/1/19
 **/
public class MiniDuckSimulator {

    public static void main(String[] args) {
        Duck mallard = new MallardDuck();
        mallard.performFly();
        mallard.performQuack();

        Duck model = new ModelDuck();
        model.performFly();
        model.setFlyBehavior(new FlyRocketPowered());
        model.performFly();

        DuckCall duckcall = new DuckCall();
        duckcall.setDuck(new MallardDuck());
        duckcall.performDuckCall();
        duckcall.setDuck(new ModelDuck());
        duckcall.performDuckCall();
    }
}
