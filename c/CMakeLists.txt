cmake_minimum_required(VERSION 3.14)
project(leetcode_c C)

set(CMAKE_C_STANDARD 99)

add_executable(main main.c tree.h list.h)

add_library(tree SHARED tree.c tree.h)
add_library(list SHARED list.c list.h)
target_link_libraries(main tree list)