#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>

void error(char *msg)
{
    fprintf(stderr, "%s\n", msg);
    exit(1);
}

int main()
{
    for (int i = 0; i < 3; i++)
    {
        pid_t pid = fork();
        if (pid == -1)
        {
            error("fork失败");
        }
        if (!pid)
        {
            execl("/bin/ls", "/bin/ls"
                             ".",
                  NULL);
            int pid_status;
            if (waitpid(pid, &pid_status, 0) == -1)
            {
                error("等待程序退出发生了错误");
            }
            if (WEXITSTATUS(pid_status)) {
                error("退出状态非零");
            }
        }
    }
    return 0;
}
