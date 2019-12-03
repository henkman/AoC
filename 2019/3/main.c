// gcc -Wall -s -O2 -o main main.c && cat input.txt | main
#include <stdio.h>
#include <stdlib.h>

typedef struct {
	int x, y;
} Point;

typedef struct {
	Point a, b;
} Line;

typedef struct {
	Line* part;
	size_t count, capacity;
} Wire;

static int point_manhattan_distance(Point* a, Point* b)
{
	return abs(b->x - a->x) + abs(b->y - a->y);
}

static int line_intersect(Line* a, Line* b, Point* i)
{
	float s1_x, s1_y, s2_x, s2_y;
	s1_x = a->b.x - a->a.x;
	s1_y = a->b.y - a->a.y;
	s2_x = b->b.x - b->a.x;
	s2_y = b->b.y - b->a.y;

	float s, t;
	s = (-s1_y * (a->a.x - b->a.x) + s1_x * (a->a.y - b->a.y)) / (-s2_x * s1_y + s1_x * s2_y);
	t = (s2_x * (a->a.y - b->a.y) - s2_y * (a->a.x - b->a.x)) / (-s2_x * s1_y + s1_x * s2_y);

	if (s >= 0 && s <= 1 && t >= 0 && t <= 1) {
		i->x = a->a.x + (t * s1_x);
		i->y = a->a.y + (t * s1_y);
		return 1;
	}
	return 0;
}

static void wire_init(Wire* wire, size_t capacity)
{
	wire->count = 0;
	wire->capacity = capacity;
	wire->part = malloc(sizeof(Wire) * wire->capacity);
}

static void wire_add_part(Wire* wire, Line* part)
{
	size_t nc = wire->count + 1;
	if (nc >= wire->capacity) {
		wire->capacity *= 2;
		wire->part = realloc(wire->part, sizeof(Wire) * wire->capacity);
	}
	wire->part[wire->count] = *part;
	wire->count = nc;
}

static void wire_read(Wire* wire, FILE* fd)
{
	Line part;
	part.a = (Point){ 0, 0 };
	char dir;
	int length;
	for (;;) {
		fscanf(stdin, "%c%d", &dir, &length);
		switch (dir) {
		case 'R':
			part.b.x = part.a.x + length;
			part.b.y = part.a.y;
			break;
		case 'L':
			part.b.x = part.a.x - length;
			part.b.y = part.a.y;
			break;
		case 'D':
			part.b.x = part.a.x;
			part.b.y = part.a.y + length;
			break;
		case 'U':
			part.b.x = part.a.x;
			part.b.y = part.a.y - length;
			break;
		}
		wire_add_part(wire, &part);
		int o = fgetc(stdin);
		if (o != ',') {
			break;
		}
		part.a = part.b;
	}
}

int main(int argc, char** argv)
{
	Wire first, second;
	wire_init(&first, 16);
	wire_init(&second, 16);
	wire_read(&first, stdin);
	wire_read(&second, stdin);

	if (0) {
		for (int i = 0; i < first.count; i++) {
			Line* part = &first.part[i];
			Point* a = &part->a;
			Point* b = &part->b;
			printf("%d,%d -> %d,%d\n", a->x, a->y, b->x, b->y);
		}
		printf("------\n");
		for (int i = 0; i < second.count; i++) {
			Line* part = &second.part[i];
			Point* a = &part->a;
			Point* b = &part->b;
			printf("%d,%d -> %d,%d\n", a->x, a->y, b->x, b->y);
		}
	}

	Point zero = { 0, 0 }, closest = { INT_MAX, INT_MAX };
	int closest_dist = INT_MAX;

	Point least_steps = { 0, 0 };
	int least_steps_amount = INT_MAX;

	int first_steps = 0, second_steps;
	for (int i = 0; i < first.count; i++) {
		Line* a = &first.part[i];
		second_steps = 0;
		for (int e = 0; e < second.count; e++) {
			Line* b = &second.part[e];
			Point intersect;
			if (line_intersect(a, b, &intersect)) {
				int dist = point_manhattan_distance(&intersect, &zero);
				if (dist < closest_dist) {
					closest = intersect;
					closest_dist = dist;
				}
				int steps = (first_steps + point_manhattan_distance(&a->a, &intersect))
					+ (second_steps + point_manhattan_distance(&b->a, &intersect));
				if (steps < least_steps_amount) {
					least_steps = intersect;
					least_steps_amount = steps;
				}
			}
			second_steps += point_manhattan_distance(&b->a, &b->b);
		}
		first_steps += point_manhattan_distance(&a->a, &a->b);
	}
	printf("first: %d (%d,%d)\n",
		closest_dist, closest.x, closest.y);
	printf("second: %d (%d,%d)\n",
		least_steps_amount, least_steps.x, least_steps.y);

	return 0;
}