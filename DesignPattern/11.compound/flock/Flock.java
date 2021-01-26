package flock;

import java.util.ArrayList;
import observer.DuckObserver;
import quack.QuackAble;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public class Flock implements QuackAble {

    ArrayList<QuackAble> quackers;

    public Flock() {
        quackers = new ArrayList<>();
    }

    public void add(QuackAble quacker) {
        quackers.add(quacker);
    }

    @Override
    public void quack() {
        for (QuackAble quacker : quackers) {
            quacker.quack();
        }
    }

    @Override
    public void registerObserver(DuckObserver observer) {
        for (QuackAble quacker : quackers) {
            quacker.registerObserver(observer);
        }
    }

    @Override
    public void notifyObservers() {
        //
    }
}
