#include <stdio.h>

int main() {
    char name[50];

    printf("Insert Your Name: ");
    fgets(name, sizeof(name), stdin);
    printf("\nHello, %s", name);


    return 0;
}