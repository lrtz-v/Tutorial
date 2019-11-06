
#include <stdio.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <netinet/in.h>


int make_socket(uint16_t port)
{
    int sock;
    struct sockaddr_in name;

    /* Init IPV4 socket. */
    sock = socket(PF_INET, SOCK_STREAM, 0);
    if (sock < 0) {
        perror ("socket");
        exit (EXIT_FAILURE);
    }

    /* Set port & ip & addr. */
    name.sin_family = AF_INET; /* IPV4 */
    name.sin_port = htons (port);  /* port */
    name.sin_addr.s_addr = htonl(INADDR_ANY); /* address */

    /* Bind addr */
    if (bind(sock, (struct sockaddr *) &name, sizeof(name)) < 0) {
        perror ("bind");
        exit (EXIT_FAILURE);
    }
    return sock
}
