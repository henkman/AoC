// gcc -Wall -s -O2 -o main main.c && cat input.txt | main
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef enum {
	ADD = 1,
	MUL = 2,
	INPUT = 3,
	OUTPUT = 4,
	JUMP_IF_TRUE = 5,
	JUMP_IF_FALSE = 6,
	LESS_THAN = 7,
	EQUALS = 8,
	HALT = 99
} Opcode;

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

static void parse_instruction(int instruction, Opcode* opcode,
	int* first, int* second, int* third)
{
	*opcode = instruction % 100;
	*first = (instruction / 100) % 10;
	*second = (instruction / 1000) % 10;
	*third = (instruction / 10000) % 10;
}

static int memory_get_value(Memory* mem, int value, int mode)
{
	if (mode) {
		return value;
	}
	return mem->ptr[value];
}

static void execute(Memory* mem, FILE* out, int input)
{
	int ip = 0;
	for (;;) {
		Opcode opcode;
		int fm, sm, tm;
		parse_instruction(mem->ptr[ip], &opcode, &fm, &sm, &tm);
		switch (opcode) {
		case ADD: {
			int first = memory_get_value(mem, mem->ptr[ip + 1], fm);
			int second = memory_get_value(mem, mem->ptr[ip + 2], sm);
			int third = mem->ptr[ip + 3];
			mem->ptr[third] = first + second;
			ip += 4;
		} break;
		case MUL: {
			int first = memory_get_value(mem, mem->ptr[ip + 1], fm);
			int second = memory_get_value(mem, mem->ptr[ip + 2], sm);
			int third = mem->ptr[ip + 3];
			mem->ptr[third] = first * second;
			ip += 4;
		} break;
		case INPUT: {
			int first = mem->ptr[ip + 1];
			mem->ptr[first] = input;
			ip += 2;
		} break;
		case OUTPUT: {
			int first = mem->ptr[ip + 1];
			fprintf(out, "OUTPUT: %d\n", mem->ptr[first]);
			ip += 2;
		} break;
		case JUMP_IF_TRUE: {
			int first = memory_get_value(mem, mem->ptr[ip + 1], fm);
			int second = memory_get_value(mem, mem->ptr[ip + 2], sm);
			if (first) {
				ip = second;
			} else {
				ip += 3;
			}
		} break;
		case JUMP_IF_FALSE: {
			int first = memory_get_value(mem, mem->ptr[ip + 1], fm);
			int second = memory_get_value(mem, mem->ptr[ip + 2], sm);
			if (!first) {
				ip = second;
			} else {
				ip += 3;
			}
		} break;
		case LESS_THAN: {
			int first = memory_get_value(mem, mem->ptr[ip + 1], fm);
			int second = memory_get_value(mem, mem->ptr[ip + 2], sm);
			int third = mem->ptr[ip + 3];
			mem->ptr[third] = first < second;
			ip += 4;
		} break;
		case EQUALS: {
			int first = memory_get_value(mem, mem->ptr[ip + 1], fm);
			int second = memory_get_value(mem, mem->ptr[ip + 2], sm);
			int third = mem->ptr[ip + 3];
			mem->ptr[third] = first == second;
			ip += 4;
		} break;
		case HALT:
			return;
		default:
			fprintf(out, "invalid opcode %d at ip %d\n", opcode, ip);
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
	execute(&program, stdout, 1);

	printf("==================\n");

	memory_copy(&initial, &program);
	execute(&program, stdout, 5);

	return 0;
}