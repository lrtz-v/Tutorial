package test;

import display_element.CurrentConditionsDisplay;
import display_element.HeatIndexDisplay;
import subject.WeatherDate;

/**
 * @author lvtao03
 * @date 2021/1/19
 **/
public class WeatherStation {

    public static void main(String[] args) {
        WeatherDate weatherData = new WeatherDate();

        CurrentConditionsDisplay currentConditionsDisplay = new CurrentConditionsDisplay(weatherData);
        HeatIndexDisplay heatIndexDisplay = new HeatIndexDisplay(weatherData);

        weatherData.setMeasurements(80, 65, 30.4f);
        weatherData.setMeasurements(82, 70, 29.4f);
    }

}
