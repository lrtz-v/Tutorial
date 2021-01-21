package command;

import iterm.Stere;

/**
 * @author lvtao03
 * @date 2021/1/21
 **/
public class StereOnWithCDCommand implements Command {

    Stere stere;

    public StereOnWithCDCommand(Stere stere) {
        this.stere = stere;
    }

    @Override
    public void execute() {
        stere.on();
        stere.setCD();
        stere.setVolume(11);
    }

    @Override
    public void undo() {
        stere.off();
    }
}
