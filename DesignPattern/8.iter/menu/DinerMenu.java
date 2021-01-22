package menu;

import iterator.DinerMenuIterator;
import iterator.Iterator;

/**
 * @author lvtao03
 * @date 2021/1/22
 **/
public class DinerMenu {

    static final int MAX_ITEMS = 6;
    private int numberOffItems = 0;
    private MenuItem[] menuItems;

    public DinerMenu() {
        menuItems = new MenuItem[MAX_ITEMS];
        addItem("", "", Boolean.TRUE, 2.99f);
        addItem("", "", Boolean.FALSE, 2.99f);
        addItem("", "", Boolean.FALSE, 3.29f);
        addItem("", "", Boolean.FALSE, 3.05f);
    }

    public void addItem(String name, String description, Boolean vegetarian, Float price) {
        MenuItem item = new MenuItem();
        item.setName(name);
        item.setDescription(description);
        item.setVegetarian(vegetarian);
        item.setPrice(price);
    }

    public Iterator createIterator() {
        return new DinerMenuIterator(menuItems);
    }

    public MenuItem[] getMenuItems() {
        return menuItems;
    }

}
