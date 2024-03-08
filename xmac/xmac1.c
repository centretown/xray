#include <stdio.h>

#define FOR_LIST_OF_VARIABLES(DO) \
    DO(JOY_FLAG_0, joy_x, 0) \
    DO(JOY_FLAG_1, joy_y, 1) \
    DO(JOY_FLAG_2, joy_z, 2) \
    DO(JOY_FLAG_3, joy_t, 3) \
    DO(JOY_FLAG_4, joy_n, 4)

#define DEFINE_FLAG_ENUMERATION(id, name, val, ...) id=val*val,
enum my_id_list_type {
    FOR_LIST_OF_VARIABLES( DEFINE_FLAG_ENUMERATION )
};

#define DEFINE_NAME_VAR(id, name, ...) float name = id;
FOR_LIST_OF_VARIABLES(DEFINE_NAME_VAR)

void print_variables(void)
{
#define PRINT_NAME_AND_VALUE(id, name, ...) printf(#name " = %f\n", name);
    FOR_LIST_OF_VARIABLES(PRINT_NAME_AND_VALUE)
}

int main(void)
{
    print_variables();
}
