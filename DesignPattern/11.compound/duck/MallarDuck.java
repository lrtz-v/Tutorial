package duck;

import observer.DuckObserver;
import observer.Observable;
import quack.QuackAble;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public class MallarDuck implements QuackAble {

    Observable observable;

    public MallarDuck() {
        observable = new Observable(this);
    }

    @Override
    public void quack() {
        System.out.println("Quack");
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

    @Override
    public String toString() {
        return "MallarDuck";
    }
}
