// gcc -Wall -s -O2 -o main main.c && cat input.txt | main
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
	int* ptr;
	size_t size, capacity;
} Memory;

static void memory_init(Memory* mem)
{
	mem->size = 0;
	mem->capacity = 16;
	mem->ptr = malloc(sizeof(int) * mem->capacity);
}

static void memory_append(Memory* mem, int value)
{
	size_t newsize = mem->size + 1;
	if (newsize >= mem->capacity) {
		mem->capacity *= 2;
		mem->ptr = realloc(mem->ptr, sizeof(int) * mem->capacity);
	}
	mem->ptr[mem->size] = value;
	mem->size = newsize;
}

static void memory_copy(Memory* src, Memory* dst)
{
	if (dst->capacity < src->capacity) {
		dst->capacity = src->capacity;
		dst->ptr = realloc(dst->ptr, sizeof(int) * dst->capacity);
	}
	dst->size = src->size;
	memcpy(dst->ptr, src->ptr, sizeof(int) * dst->size);
}

static void execute(Memory* mem, int noun, int verb)
{
	mem->ptr[1] = noun;
	mem->ptr[2] = verb;

	int ip = 0;
	for (;;) {
		int opcode = mem->ptr[ip];
		switch (opcode) {
		case 1: // add
			mem->ptr[mem->ptr[ip + 3]] = mem->ptr[mem->ptr[ip + 1]] + mem->ptr[mem->ptr[ip + 2]];
			ip += 4;
			break;
		case 2: // mul
			mem->ptr[mem->ptr[ip + 3]] = mem->ptr[mem->ptr[ip + 1]] * mem->ptr[mem->ptr[ip + 2]];
			ip += 4;
			break;
		case 99: // halt
			return;
		default:
			printf("invalid opcode %d at ip %d\n", opcode, ip);
			return;
		}
	}
}

int main(int argc, char** argv)
{
	Memory initial, program;
	memory_init(&initial);
	memory_init(&program);

	{ // read in initial
		int c;
		for (;;) {
			fscanf(stdin, "%d", &c);
			memory_append(&initial, c);
			int o = fgetc(stdin);
			if (o != ',') {
				break;
			}
		}
	}

	if (0) {
		for (int i = 0; i < initial.size; i++) {
			printf("%d: %d\n", i, initial.ptr[i]);
		}
	}

	memory_copy(&initial, &program);
	execute(&program, 12, 2);
	printf("first: %d\n", program.ptr[0]);

	for (int noun = 0; noun <= 99; noun++) {
		for (int verb = 0; verb <= 99; verb++) {
			memory_copy(&initial, &program);
			execute(&program, noun, verb);
			if (program.ptr[0] == 19690720) {
				printf("second: noun=%d, verb=%d, answer=%d\n", noun, verb, 100 * noun + verb);
				goto end;
			}
		}
	}
end:

	return 0;
}