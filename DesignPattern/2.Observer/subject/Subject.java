package subject;

import observer.Observer;

public interface Subject {

    /**
     * 注册观察者
     *
     * @param o 观察者
     */
    void registerObserver(Observer o);

    /**
     * 注销观察者
     *
     * @param o 观察者
     */
    void removeObserver(Observer o);

    /**
     * 通知所有观察者
     */
    void notifyObservers();
}