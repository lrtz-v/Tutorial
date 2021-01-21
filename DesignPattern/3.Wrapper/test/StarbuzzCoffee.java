package test;

import template.Beverage;
import template.DarkRoast;
import template.Espresso;
import condiment.Mocha;

/**
 * @author lrtz
 */
public class StarbuzzCoffee {

    public static void main(String[] args) {
        Beverage beverage = new Espresso();
        System.out.println(beverage.getDescription() + "$" + beverage.cost());

        Beverage beverage2 = new DarkRoast();
        beverage2 = new Mocha(beverage2);
        beverage2 = new Mocha(beverage2);
        System.out.println(beverage2.getDescription() + "$" + beverage2.cost());
    }
}