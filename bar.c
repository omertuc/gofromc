#include "foo.h"
#include <stdio.h>

int main(int argc, char **argv) {
    GoInt x = Add(3, 5);
    printf("%d\n", x);
    return 0;
}
