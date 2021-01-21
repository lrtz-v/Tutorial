package control;

import command.Command;
import command.NoCommand;

/**
 * @author lvtao03
 * @date 2021/1/21
 **/
public class RemoteControlWithUndo {

    Command[] onCommands;
    Command[] offCommands;
    Command undoCommand;

    public RemoteControlWithUndo() {
        onCommands = new Command[7];
        offCommands = new Command[7];

        Command noCommand = new NoCommand();
        for (int i = 0; i < onCommands.length; i++) {
            onCommands[i] = noCommand;
            offCommands[i] = noCommand;
        }

        undoCommand = noCommand;
    }

    public void setCommand(int slot, Command on, Command off) {
        onCommands[slot] = on;
        offCommands[slot] = off;
    }

    public void onButtonWaePushed(int slot) {
        onCommands[slot].execute();
        undoCommand = onCommands[slot];
    }

    public void offButtonWaePushed(int slot) {
        offCommands[slot].execute();
        undoCommand = offCommands[slot];
    }

    public void undoButtonPushed() {
        undoCommand.undo();
    }

    @Override
    public String toString() {
        StringBuilder str = new StringBuilder();
        str.append("\n----- Remote Control -----\n");
        for (int i = 0; i < onCommands.length; i++) {
            str.append("[slot ")
                .append(i)
                .append("] ")
                .append(onCommands[i].getClass().getName())
                .append("   ")
                .append(offCommands[i].getClass().getName())
                .append("\n");
        }
        return str.toString();
    }
}
