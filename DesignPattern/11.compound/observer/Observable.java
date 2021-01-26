package observer;

import java.util.ArrayList;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public class Observable implements QuackObservable {

    ArrayList<DuckObserver> observables;
    QuackObservable duck;

    public Observable(QuackObservable duck) {
        this.duck = duck;
        observables = new ArrayList<>();
    }

    @Override
    public void registerObserver(DuckObserver observer) {
        observables.add(observer);
    }

    @Override
    public void notifyObservers() {
        for (DuckObserver observer: observables) {
            observer.update(duck);
        }
    }
}
