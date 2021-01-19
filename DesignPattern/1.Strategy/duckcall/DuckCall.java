package duckcall;

import duck.Duck;

/**
 * @author lvtao03
 * @date 2021/1/19
 **/
public class DuckCall {

    Duck duck;

    public DuckCall() {
    }

    public void setDuck(Duck d) {
        duck = d;
    }

    public void performDuckCall() {
        duck.performQuack();
    }

}
