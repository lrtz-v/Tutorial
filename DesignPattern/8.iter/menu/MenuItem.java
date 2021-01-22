package menu;

/**
 * @author lvtao03
 * @date 2021/1/22
 **/
public class MenuItem {

    private String name;
    private String description;
    private Boolean vegetarian;
    private Float price;

    public void setDescription(String description) {
        this.description = description;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setPrice(Float price) {
        this.price = price;
    }

    public void setVegetarian(Boolean vegetarian) {
        this.vegetarian = vegetarian;
    }

    public String getName() {
        return name;
    }

    public String getDescription() {
        return description;
    }

    public Boolean getVegetarian() {
        return vegetarian;
    }

    public Float getPrice() {
        return price;
    }
}
