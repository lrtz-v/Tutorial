#include <stdio.h>
#include <string.h>

char tracks[][80] = {
    "I left my heart in Harvard Med School",
    "Newark, Newark - a wonderful town",
    "Dancing with a Dork",
    "From here to maternity",
    "The girl from Iwo Jima",
};

void go_south_east(int *lat, int *lon)
{
    *lat -= 1;
    *lon += 1;
}

void fortune_cookie(char msg[])
{
    printf("Message reads: %s\n", msg);
    printf("msg occupies %i bytes\n", sizeof(msg));
}

void find_track(char search_for[])
{
    int i;
    for (i = 0; i < 5; i++)
    {
        if (strstr(tracks[i], search_for))
            printf("Track %i: '%s'\n", i, tracks[i]);
    }
}
int main()
{
    int latitude = 32;
    int longitude = -64;
    go_south_east(&latitude, &longitude);
    printf("停！当前位置: [%i, %i]\n", latitude, longitude);
    char quote[] = "Cookies make you fat";
    fortune_cookie(quote);

    char saerch_for[80];
    printf("Search for: ");
    fgets(saerch_for, 80, stdin);
    saerch_for[strlen(saerch_for) - 1] = '\0';
    find_track(saerch_for);

    return 0;
}
