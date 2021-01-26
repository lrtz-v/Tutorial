package invocation;

import java.lang.reflect.Method;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public interface InvocationHandler {

    Object invoke(Object proxy, Method method, Object[] args);

}
