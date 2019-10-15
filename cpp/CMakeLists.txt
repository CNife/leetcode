cmake_minimum_required(VERSION 3.13)
project(cpp)

set(CMAKE_CXX_STANDARD 17)

add_library(tree SHARED tree.cpp tree.hpp)
add_executable(test main.cpp tree.hpp)
target_link_libraries(test tree)