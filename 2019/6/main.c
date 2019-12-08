// gcc -Wall -s -O2 -o main main.c && cat input.txt | main
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <inttypes.h>

#define NAME_LEN 3

struct sObject;

typedef struct {
    struct sObject* object;
    size_t count, capacity;
} Objects;

typedef struct {
    size_t* index;
    size_t count, capacity;
} ObjectIndices;

typedef struct sObject {
    char name[NAME_LEN];
    size_t parent_index;
    ObjectIndices indices;
} Object;

static void objectindices_init(ObjectIndices* oidxs, size_t capacity)
{
    oidxs->count = 0;
    oidxs->capacity = capacity;
    oidxs->index = malloc(sizeof(size_t) * oidxs->capacity);
}

static void objectindices_add(ObjectIndices* oidxs, size_t index)
{
    size_t nc = oidxs->count + 1;
    if (nc >= oidxs->capacity) {
        oidxs->capacity *= 2;
        oidxs->index = realloc(oidxs->index, sizeof(size_t) * oidxs->capacity);
    }
    oidxs->index[oidxs->count] = index;
    oidxs->count = nc;
}

static void object_init(Object* obj, char name[NAME_LEN])
{
    obj->parent_index = 0;
    objectindices_init(&obj->indices, 2);
    memcpy(&obj->name[0], name, NAME_LEN);
}

static void objects_init(Objects* objs, size_t capacity)
{
    objs->count = 0;
    objs->capacity = capacity;
    objs->object = malloc(sizeof(Object) * objs->capacity);
}

static size_t objects_add(Objects* objs, char name[NAME_LEN])
{
    size_t nc = objs->count + 1;
    if (nc >= objs->capacity) {
        objs->capacity *= 2;
        objs->object = realloc(objs->object, sizeof(Object) * objs->capacity);
    }
    Object object;
    object_init(&object, name);
    size_t oc = objs->count;
    objs->object[objs->count] = object;
    objs->count = nc;
    return oc;
}

static int objects_indexbyname(Objects* objs, char name[NAME_LEN], size_t* index)
{
    for (size_t i = 0; i < objs->count; i++) {
        if (memcmp(objs->object[i].name, name, NAME_LEN) == 0) {
            *index = i;
            return 1;
        }
    }
    return 0;
}

int main(int argc, char** argv)
{
    Objects objs;
    objects_init(&objs, 16);

    size_t icom;
    {
        char name[3];
        for (;;) {
            size_t n = fread(&name[0], sizeof(char), NAME_LEN, stdin);
            if (n != NAME_LEN)
                break;
            size_t iobj;
            if (!objects_indexbyname(&objs, name, &iobj)) {
                iobj = objects_add(&objs, name);
            }
            if (fgetc(stdin) != ')')
                break;
            n = fread(&name[0], sizeof(char), NAME_LEN, stdin);
            if (n != NAME_LEN)
                break;
            size_t iorbiter;
            if (!objects_indexbyname(&objs, name, &iorbiter)) {
                iorbiter = objects_add(&objs, name);
            }
            objs.object[iorbiter].parent_index = iobj;
            Object* obj = &objs.object[iobj];
            ObjectIndices* oidxs = &obj->indices;
            objectindices_add(oidxs, iorbiter);
            if (fgetc(stdin) != '\n')
                break;
        }

        objects_indexbyname(&objs, "COM", &icom);
        objs.object[icom].parent_index = icom;
    }

    if (0) {
        for (int i = 0; i < objs.count; i++) {
            Object* obj = &objs.object[i];
            printf("%.*s", NAME_LEN, obj->name);

            if (i != obj->parent_index) {
                Object* parent = &objs.object[obj->parent_index];
                printf(" - %.*s\n", NAME_LEN, parent->name);
            } else {
                printf("\n");
            }

            ObjectIndices* oidxs = &obj->indices;
            for (int e = 0; e < oidxs->count; e++) {
                Object* orbiter = &objs.object[oidxs->index[e]];
                printf("\t%.*s\n", NAME_LEN, orbiter->name);
            }
        }
    }

    if (1) {
        int total = 0;
        for (size_t i = 0; i < objs.count; i++) {
            size_t index = i;
            while (index != icom) {
                total++;
                index = objs.object[index].parent_index;
            }
        }
        printf("first: %d\n", total);
    }

    if (1) {
        ObjectIndices pyou;
        {
            objectindices_init(&pyou, 2);
            size_t index;
            objects_indexbyname(&objs, "YOU", &index);
            while (index != icom) {
                index = objs.object[index].parent_index;
                objectindices_add(&pyou, index);
            }
        }

        {
            size_t index;
            objects_indexbyname(&objs, "SAN", &index);
            size_t sansteps = 0;
            index = objs.object[index].parent_index;
            while (index != icom) {
                for (size_t i = 0; i < pyou.count; i++) {
                    if (index == pyou.index[i]) {
                        printf("second: %" PRIuMAX "\n", i + sansteps);
                        goto end;
                    }
                }
                sansteps++;
                index = objs.object[index].parent_index;
            }
        }
    }

end:
    return 0;
}