package adapter;

import goose.Goose;
import observer.DuckObserver;
import observer.Observable;
import quack.QuackAble;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public class GooseAdapter implements QuackAble {

    Goose goose;
    Observable observable;

    public GooseAdapter(Goose goose) {
        this.goose = goose;
        observable = new Observable(this);
    }

    @Override
    public void quack() {
        goose.honk();
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
