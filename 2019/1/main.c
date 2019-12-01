// gcc -Wall -s -O2 -o main main.c && cat input.txt | main
#include <stdlib.h>
#include <stdio.h>
#include <math.h>

static int required_fuel(int mass)
{
    return (int)(floorf(mass / 3) - 2);
}

int main(int argc, char** argv)
{
    if (0) {
        printf("%d\n", required_fuel(12));
        printf("%d\n", required_fuel(14));
        printf("%d\n", required_fuel(1969));
        printf("%d\n", required_fuel(100756));
    }

    char buf[256];
    int first = 0, second = 0;
    for (;;) {
        char* line = fgets(&buf[0], sizeof(line), stdin);
        if (line == NULL) {
            break;
        }
        int mass = atoi(line);
		int fuel = required_fuel(mass);
        first += fuel;

		int sumfuel = fuel;
		for(;;) {
			int nf = required_fuel(fuel);
			if(nf <= 0) {
				break;
			}
			sumfuel += nf;
			fuel = nf;
		}
		second += sumfuel;
    }
    printf("first: %d\n", first);
    printf("second: %d\n", second);
}