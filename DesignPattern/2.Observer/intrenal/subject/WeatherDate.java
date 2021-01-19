package intrenal.subject;

import java.util.Observable;

/**
 * @author lvtao03
 * @date 2021/1/19
 **/
public class WeatherDate extends Observable {

    private float temperature;
    private float humidity;
    private float pressure;

    public WeatherDate() {
    }

    private void measurementsChanged() {
        setChanged();
        notifyObservers();
    }

    public void setMeasurements(float temperature, float humidity, float pressure) {
        this.temperature = temperature;
        this.humidity = humidity;
        this.pressure = pressure;
        measurementsChanged();
    }

    public float getTemperature() {
        return temperature;
    }

    public float getHumidity() {
        return humidity;
    }

    public float getPressure() {
        return pressure;
    }
}
