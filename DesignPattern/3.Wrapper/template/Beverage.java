package template;

/**
 * @author lrtz
 */
public abstract class Beverage {

    String description;

    public String getDescription() {
        return description;
    }

    /**
     * 计算费用
     *
     * @return cost
     */
    public abstract double cost();
}
