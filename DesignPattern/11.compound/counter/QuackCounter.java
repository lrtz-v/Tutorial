package counter;

import observer.DuckObserver;
import observer.Observable;
import quack.QuackAble;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public class QuackCounter implements QuackAble {

    QuackAble duck;
    static int quackCount;

    public QuackCounter(QuackAble duck) {
        this.duck = duck;
    }

    @Override
    public void quack() {
        duck.quack();
        quackCount++;
    }

    public static int getQuackCount() {
        return quackCount;
    }

    @Override
    public void registerObserver(DuckObserver observer) {
        duck.registerObserver(observer);
    }

    @Override
    public void notifyObservers() {
        //
    }
}
