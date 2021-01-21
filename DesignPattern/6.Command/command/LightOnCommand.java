package command;

import iterm.Light;

/**
 * @author lvtao03
 * @date 2021/1/21
 **/
public class LightOnCommand implements Command {

    Light light;

    public LightOnCommand(Light light) {
        this.light = light;
    }

    @Override
    public void execute() {
        light.on();
    }

    @Override
    public void undo() {
        light.off();
    }
}
