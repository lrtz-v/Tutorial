package observer;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public class QuackLogist implements DuckObserver {

    @Override
    public void update(QuackObservable duck) {
        System.out.println("QuackLogist: " + duck + " just quacked.");
    }
}
