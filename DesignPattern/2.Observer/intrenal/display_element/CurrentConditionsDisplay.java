package intrenal.display_element;

import intrenal.subject.WeatherDate;
import java.util.Observable;
import java.util.Observer;

/**
 * @author lvtao03
 * @date 2021/1/19
 **/
public class CurrentConditionsDisplay implements Observer, DisplayElement {

    Observable observable;
    private float temperature;
    private float humidity;
    private float pressure;

    public CurrentConditionsDisplay(Observable observable) {
        this.observable = observable;
        observable.addObserver(this);
    }

    @Override
    public void update(Observable obs, Object arg) {
        if (obs instanceof WeatherDate) {
            WeatherDate weatherDate = (WeatherDate) obs;
            this.temperature = weatherDate.getTemperature();
            this.humidity = weatherDate.getHumidity();
            this.pressure = weatherDate.getPressure();
            display();
        }
    }

    @Override
    public void display() {
        System.out.println(
            "Current conditions: " +
                temperature + "F degrees and " +
                humidity + "% humidity");
    }
}
