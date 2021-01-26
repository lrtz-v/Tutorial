package factory;

import duck.DuckCall;
import duck.MallarDuck;
import duck.RedheadDuck;
import duck.RubberDuck;
import quack.QuackAble;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public class DuckFactory extends AbstractDuckFactory {

    @Override
    public QuackAble createMallarDuck() {
        return new MallarDuck();
    }

    @Override
    public QuackAble createRedheadDuck() {
        return new RedheadDuck();
    }

    @Override
    public QuackAble createDuckCall() {
        return new DuckCall();
    }

    @Override
    public QuackAble createRubberDuck() {
        return new RubberDuck();
    }
}
