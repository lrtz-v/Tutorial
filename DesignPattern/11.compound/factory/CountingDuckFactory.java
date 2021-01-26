package factory;

import counter.QuackCounter;
import duck.DuckCall;
import duck.MallarDuck;
import duck.RedheadDuck;
import duck.RubberDuck;
import quack.QuackAble;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public class CountingDuckFactory extends AbstractDuckFactory {

    @Override
    public QuackAble createMallarDuck() {
        return new QuackCounter(new MallarDuck());
    }

    @Override
    public QuackAble createRedheadDuck() {
        return new QuackCounter(new RedheadDuck());
    }

    @Override
    public QuackAble createDuckCall() {
        return new QuackCounter(new DuckCall());
    }

    @Override
    public QuackAble createRubberDuck() {
        return new QuackCounter(new RubberDuck());
    }
}
