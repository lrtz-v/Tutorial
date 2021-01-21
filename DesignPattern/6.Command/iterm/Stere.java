package iterm;

/**
 * @author lvtao03
 * @date 2021/1/21
 **/
public class Stere {

    private int volume;

    public void on() {
        System.out.println("Stere on");
    }

    public void off() {
        System.out.println("Stere off");
    }

    public void setCD() {
        System.out.println("Stere set CD");
    }

    public void setVolume(int volume) {
        this.volume = volume;
        System.out.println("Stere set volume");
    }

}
