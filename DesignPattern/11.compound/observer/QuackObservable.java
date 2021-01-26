package observer;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public interface QuackObservable {

    void registerObserver(DuckObserver observer);

    void notifyObservers();

}
