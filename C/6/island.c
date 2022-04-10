#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct island
{
    char *name;
    char *opens;
    char *closes;
    struct island *next;
} island;

void display(island *t)
{
    island *i = t;
    for (; i != NULL; i = i->next)
    {
        printf("Name: %s open: %s-%s\n", i->name, i->opens, i->closes);
    }
}

island *create(char *name, char *opens, char *closes)
{
    island *i = malloc(sizeof(island));
    i->name = strdup(name);
    i->opens = opens;
    i->closes = closes;
    i->next = NULL;
    return i;
}

void release(island *t)
{
    island *i;
    island *next;
    for (; i != NULL; i = i->next)
    {
        next = i->next;
        free(i->name);
        free(i);
    }
}

int main()
{
    island head = {};
    island *prev = &head;
    for (int i = 0; i < 5; i++)
    {
        char name[80];
        fgets(name, 80, stdin);
        island *tmp = create(name, "", "");
        prev->next = tmp;
        prev = tmp;
    }
    display(head.next);
    release(&head);
}
