package control;

import command.Command;
import command.NoCommand;

/**
 * @author lvtao03
 * @date 2021/1/20
 **/
public class RemoteControl {

    Command[] onCommands;
    Command[] offCommands;

    public RemoteControl() {
        onCommands = new Command[7];
        offCommands = new Command[7];

        Command noCommand = new NoCommand();
    }

}
