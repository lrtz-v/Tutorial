package test;

import command.LightOffCommand;
import command.LightOnCommand;
import command.StereOffCommand;
import command.StereOnWithCDCommand;
import control.RemoteControl;
import control.RemoteControlWithUndo;
import iterm.Light;
import iterm.Stere;

/**
 * @author lvtao03
 * @date 2021/1/21
 **/
public class RemoteLoader {

    public static void main(String[] args) {
        RemoteControl control = new RemoteControl();

        Light light = new Light();
        Stere stere = new Stere();

        LightOnCommand lightOnCommand = new LightOnCommand(light);
        LightOffCommand lightOffCommand = new LightOffCommand(light);
        StereOnWithCDCommand stereOnWithCDCommand = new StereOnWithCDCommand(stere);
        StereOffCommand stereOffCommand = new StereOffCommand(stere);

        control.setCommand(0, lightOnCommand, lightOffCommand);
        control.setCommand(1, stereOnWithCDCommand, stereOffCommand);
        System.out.println(control.toString());

        control.onButtonWaePushed(0);
        control.offButtonWaePushed(0);
        control.onButtonWaePushed(1);
        control.offButtonWaePushed(1);

        RemoteControlWithUndo undoControl = new RemoteControlWithUndo();

        undoControl.setCommand(0, lightOnCommand, lightOffCommand);
        undoControl.setCommand(1, stereOnWithCDCommand, stereOffCommand);
        System.out.println(undoControl.toString());

        undoControl.onButtonWaePushed(0);
        undoControl.offButtonWaePushed(0);
        undoControl.undoButtonPushed();
        undoControl.onButtonWaePushed(1);
        undoControl.offButtonWaePushed(1);
        undoControl.undoButtonPushed();
    }

}
