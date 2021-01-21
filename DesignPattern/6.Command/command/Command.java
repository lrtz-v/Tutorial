package command;

/**
 * @author lvtao03
 * @date 2021/1/20
 **/
public interface Command {

    /**
     * 执行命令
     */
    void execute();

    /**
     * 撤销执行
     */
    void undo();

}
