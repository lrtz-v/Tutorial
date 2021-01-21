package command;

import iterm.Light;

/**
 * @author lvtao03
 * @date 2021/1/21
 **/
public class LightOffCommand implements Command {

    Light light;

    public LightOffCommand(Light light) {
        this.light = light;
    }

    @Override
    public void execute() {
        light.off();
    }

    @Override
    public void undo() {
        light.on();
    }
}
