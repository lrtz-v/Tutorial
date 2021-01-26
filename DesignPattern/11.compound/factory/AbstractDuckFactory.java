package factory;

import quack.QuackAble;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public abstract class AbstractDuckFactory {

    public abstract QuackAble createMallarDuck();

    public abstract QuackAble createRedheadDuck();

    public abstract QuackAble createDuckCall();

    public abstract QuackAble createRubberDuck();

}
