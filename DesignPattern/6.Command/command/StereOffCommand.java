package command;

import iterm.Stere;

/**
 * @author lvtao03
 * @date 2021/1/21
 **/
public class StereOffCommand implements Command {

    Stere stere;

    public StereOffCommand(Stere stere) {
        this.stere = stere;
    }

    @Override
    public void execute() {
        stere.off();
    }

    @Override
    public void undo() {
        stere.on();
        stere.setCD();
        stere.setVolume(11);
    }
}
