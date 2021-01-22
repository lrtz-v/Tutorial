package iterator;

/**
 * @author lvtao03
 * @date 2021/1/22
 **/
public interface Iterator {

    /**
     * 是否存在下一个元素
     *
     * @return boolean
     */
    Boolean hasNext();

    /**
     * 下一个元素
     *
     * @return Object
     */
    Object next();
}
