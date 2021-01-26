package person;

import invocation.NonOwnerInvocationHandler;
import invocation.OwnerInvocationHandler;
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Proxy;

/**
 * @author lvtao03
 * @date 2021/1/26
 **/
public class PersonBeanImpl implements PersonBean {

    private String name;
    private String gender;
    private String interests;
    private int rating;
    private int ratingCount = 0;

    public PersonBeanImpl() {
    }


    @Override
    public String getName() {
        return name;
    }

    @Override
    public String getGender() {
        return gender;
    }

    @Override
    public String getInterests() {
        return interests;
    }

    @Override
    public int getHotOrNotRating() {
        if (ratingCount == 0) {
            return 0;
        }
        return rating / ratingCount;
    }

    @Override
    public void setName(String name) {
        this.name = name;
    }

    @Override
    public void setGender(String gender) {
        this.gender = gender;
    }

    @Override
    public void setInterests(String interests) {
        this.interests = interests;
    }

    @Override
    public void setHotOrNotRating(int rating) {
        this.rating = rating;
        ratingCount++;
    }

    public PersonBean getOwnerProxy(PersonBean personBean) {
        return (PersonBean) Proxy.newProxyInstance(
            personBean.getClass().getClassLoader(),
            personBean.getClass().getInterfaces(),
            (InvocationHandler) new OwnerInvocationHandler(personBean)
        );
    }

    public PersonBean getNonOwnerProxy(PersonBean personBean) {
        return (PersonBean) Proxy.newProxyInstance(
            personBean.getClass().getClassLoader(),
            personBean.getClass().getInterfaces(),
            (InvocationHandler) new NonOwnerInvocationHandler(personBean)
        );
    }
}
