package duck;

import observer.DuckObserver;
import observer.Observable;
import quack.QuackAble;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public class DuckCall implements QuackAble {

    Observable observable;

    public DuckCall() {
        observable = new Observable(this);
    }

    @Override
    public void quack() {
        System.out.println("Kwak");
        notifyObservers();
    }

    @Override
    public void registerObserver(DuckObserver observer) {
        observable.registerObserver(observer);
    }

    @Override
    public void notifyObservers() {
        observable.notifyObservers();
    }
}
