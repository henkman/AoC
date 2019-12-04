// gcc -Wall -s -O2 -o main main.c && cat input.txt | main
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define DIGITS 6

static void number_get_digits(int number, char digits[DIGITS])
{
	int i = 0;
	do {
		char digit = (char)(number % 10);
		digits[i] = digit;
		number = number / 10;
		i++;
	} while (number > 0);
}

static int is_valid_first(char digits[DIGITS])
{
	int has_double = 0;
	for (int i = 1; i < DIGITS; i++) {
		char last = digits[i - 1];
		char cur = digits[i];
		if (cur > last) {
			return 0;
		}
		if (cur == last) {
			has_double = 1;
		}
	}
	return has_double;
}

static int is_valid_second(char digits[DIGITS])
{
	int has_double = 0, multiple = 0;
	for (int i = 1; i <= DIGITS; i++) {
		char last = digits[i - 1];
		char cur = digits[i];
		if (cur > last) {
			return 0;
		}
		if (cur == last) {
			multiple++;
		} else {
			if (multiple == 1) {
				has_double = 1;
			}
			multiple = 0;
		}
	}
	return has_double;
}

int main(int argc, char** argv)
{
	int begin, end;
	fscanf(stdin, "%d-%d", &begin, &end);

	char digits[DIGITS];

	if (0) {
		printf("%d,%d\n", begin, end);

		number_get_digits(111111, digits);
		printf("%d -> %d\n", 111111, is_valid_first(digits));
		number_get_digits(223450, digits);
		printf("%d -> %d\n", 223450, is_valid_first(digits));
		number_get_digits(123789, digits);
		printf("%d -> %d\n", 123789, is_valid_first(digits));

		number_get_digits(112233, digits);
		printf("%d -> %d\n", 112233, is_valid_second(digits));
		number_get_digits(123444, digits);
		printf("%d -> %d\n", 123444, is_valid_second(digits));
		number_get_digits(111122, digits);
		printf("%d -> %d\n", 111122, is_valid_second(digits));

		return 0;
	}

	int first = 0, second = 0;
	for (int i = begin; i <= end; i++) {
		if (i % 10 == 0)
			continue;
		number_get_digits(i, digits);
		first += is_valid_first(digits);
		second += is_valid_second(digits);
	}
	printf("first: %d\nsecond: %d\n", first, second);

	return 0;
}