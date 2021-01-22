package iterator;


import menu.MenuItem;

/**
 * @author lvtao03
 * @date 2021/1/22
 **/
public class DinerMenuIterator implements Iterator {

    private MenuItem[] items;
    private int position = 0;

    public DinerMenuIterator(MenuItem[] items) {
        this.items = items;
    }

    @Override
    public Object next() {
        MenuItem menuItem = items[position];
        position++;
        return menuItem;
    }

    @Override
    public Boolean hasNext() {
        return position < items.length && items[position] != null;
    }
}
