package invocation;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import person.PersonBean;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public class OwnerInvocationHandler implements InvocationHandler {

    private PersonBean personBean;

    public OwnerInvocationHandler(PersonBean personBean) {
        this.personBean = personBean;
    }

    @Override
    public Object invoke(Object proxy, Method method, Object[] args) {
        try {
            if (method.getName().startsWith("get")) {
                return method.invoke(personBean, args);
            } else if ("setHotOrNotRating".equals(method.getName())) {
                throw new IllegalAccessException();
            } else if (method.getName().startsWith("set")) {
                return method.invoke(personBean, args);
            }
        } catch (IllegalAccessException | InvocationTargetException e) {
            e.printStackTrace();
        }
        return null;
    }

}
