#include <stdio.h>

int main(void) 
{
    int ten = 10;
    int minu_ten = -10;

    char bigA = 'A';

    printf("%p, %p", &ten, &minu_ten);

    printf("%p", &bigA);

    return 0;
}