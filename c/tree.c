#include "tree.h"

#include <stdlib.h>

TreeNode *newTreeNode(int val) {
    TreeNode *node = malloc(sizeof(TreeNode));
    if (node) {
        node->val = val;
        node->left = node->right = NULL;
    }
    return node;
}
