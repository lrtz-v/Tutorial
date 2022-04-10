#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef struct
{
    int width;
    int height;
} rectangle;

int compare_scores(const void *score_a, const void *score_b)
{
    int a = *(int *)score_a;
    int b = *(int *)score_b;
    return a - b;
}

int compare_scores_desc(const void *score_a, const void *score_b)
{
    int a = *(int *)score_a;
    int b = *(int *)score_b;
    return b - a;
}

int compare_areas(const void *a, const void *b)
{
    rectangle *rectangle_a = (rectangle *)a;
    rectangle *rectangle_b = (rectangle *)b;
    return rectangle_a->height * rectangle_a->width - rectangle_b->height * rectangle_b->width;
}

int compare_areas_desc(const void *a, const void *b)
{
    rectangle *rectangle_a = (rectangle *)a;
    rectangle *rectangle_b = (rectangle *)b;
    return rectangle_b->height * rectangle_b->width - rectangle_a->height * rectangle_a->width;
}

int compare_names(const void *a, const void *b)
{
    char **str_a = (char **)a;
    char **str_b = (char **)b;
    return strcmp(*str_a, *str_b);
}

int compare_names_desc(const void *a, const void *b)
{
    char **str_a = (char **)a;
    char **str_b = (char **)b;
    return strcmp(*str_b, *str_a);
}

int main()
{
    char *names[] = {
        "AB",
        "BA",
        "DF",
        "ER",
        "XX",
        "DF",
        "FG"};
    qsort(names, 7, sizeof(char *), compare_names_desc);
    for (int i = 0; i < 7; i++)
    {
        printf("%s\n", names[i]);
    }
    return 0;
}










