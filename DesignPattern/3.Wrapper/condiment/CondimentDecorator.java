package condiment;

import beverage.Beverage;

/**
 * @author lrtz
 */
public abstract class CondimentDecorator extends Beverage {

    /**
     * 返回描述
     *
     * @return description
     */
    @Override
    public abstract String getDescription();
}
