package display_element;

import observer.Observer;
import subject.Subject;

/**
 * @author lvtao03
 * @date 2021/1/19
 **/
public class CurrentConditionsDisplay implements Observer, DisplayElement {

    private float temperature;
    private float humidity;
    private Subject weatherData;

    public CurrentConditionsDisplay(Subject weatherData) {
        this.weatherData = weatherData;
        weatherData.registerObserver(this);
    }

    @Override
    public void update(float temperature, float humidity, float pressure) {
        this.temperature = temperature;
        this.humidity = humidity;
        display();
    }

    @Override
    public void display() {
        System.out.println(
            "Current conditions: " +
                temperature + "F degrees and " +
                humidity + "% humidity");
    }
}
