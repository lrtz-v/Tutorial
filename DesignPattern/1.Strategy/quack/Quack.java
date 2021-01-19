package quack;

/**
 * @author lvtao03
 * @date 2021/1/19
 **/
public class Quack implements QuackBehavior{

    @Override
    public void quack() {
        System.out.println("Quack");
    }
}
