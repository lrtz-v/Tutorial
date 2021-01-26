package quack;

import observer.QuackObservable;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public interface QuackAble extends QuackObservable {

    void quack();
}
