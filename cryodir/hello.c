#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>

int main() {
    for (int i = 1; i <= 10; i++) {
        system("clear");
        printf("%d\n", i);
        sleep(1);
    }

    return 0;
}