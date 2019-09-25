#ifndef LEETCODE_C_TREE_H
#define LEETCODE_C_TREE_H

#include <stddef.h>

typedef struct TreeNode TreeNode;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
};

TreeNode *newTreeNode(int val);

#endif //LEETCODE_C_TREE_H
