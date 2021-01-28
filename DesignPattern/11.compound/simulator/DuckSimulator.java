package simulator;

import adapter.GooseAdapter;
import counter.QuackCounter;
import duck.DuckCall;
import duck.MallarDuck;
import duck.RedheadDuck;
import duck.RubberDuck;
import factory.AbstractDuckFactory;
import factory.CountingDuckFactory;
import flock.Flock;
import goose.Goose;
import observer.QuackLogist;
import quack.QuackAble;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public class DuckSimulator {

    public static void main(String[] args) {
        DuckSimulator simulator = new DuckSimulator();
        AbstractDuckFactory duckFactory = new CountingDuckFactory();
//        simulator.simulator();
        simulator.simulator(duckFactory);
    }

    /**
     * 鸭子模拟器 1、适配器 2、装饰器-实现quack计数
     */
    private void simulator() {
//        QuackAble mallardDuck = new MallarDuck();
//        QuackAble redheadDuck = new RedheadDuck();
//        QuackAble duckCall = new DuckCall();
//        QuackAble rubberDuck = new RubberDuck();
//        QuackAble gooseDuck = new GooseAdapter(new Goose());

        QuackAble mallardDuck = new QuackCounter(new MallarDuck());
        QuackAble redheadDuck = new QuackCounter(new RedheadDuck());
        QuackAble duckCall = new QuackCounter(new DuckCall());
        QuackAble rubberDuck = new QuackCounter(new RubberDuck());
        QuackAble gooseDuck = new GooseAdapter(new Goose());

        System.out.println("\n Duck simulator");

        simulator(mallardDuck);
        simulator(redheadDuck);
        simulator(duckCall);
        simulator(rubberDuck);
        simulator(gooseDuck);

        System.out.println("quack count: " + QuackCounter.getQuackCount());
    }

    /**
     * 鸭子模拟器: 工厂 + 适配器 + 组合, 管理一群鸭子
     */
    private void simulator(AbstractDuckFactory duckFactory) {
        QuackAble mallardDuck = duckFactory.createMallarDuck();
        QuackAble redheadDuck = duckFactory.createRedheadDuck();
        QuackAble duckCall = duckFactory.createDuckCall();
        QuackAble rubberDuck = duckFactory.createRubberDuck();
        QuackAble gooseDuck = new GooseAdapter(new Goose());

        Flock flock = new Flock();
        flock.add(mallardDuck);
        flock.add(redheadDuck);
        flock.add(duckCall);
        flock.add(rubberDuck);
        flock.add(gooseDuck);

        QuackLogist quackLogist = new QuackLogist();
        flock.registerObserver(quackLogist);
        System.out.println("\n Duck simulator with factory");
        simulator(flock);
        System.out.println("quack count: " + QuackCounter.getQuackCount());
    }

    private void simulator(QuackAble duck) {
        duck.quack();
    }
}
