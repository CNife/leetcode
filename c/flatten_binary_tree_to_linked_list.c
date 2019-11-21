#include "tree.h"

#include <assert.h>

TreeNode *do_flatten(TreeNode *root) {
    assert(root);
    TreeNode *left = root->left,
            *right = root->right,
            *left_end = left ? do_flatten(left) : NULL,
            *right_end = right ? do_flatten(right) : NULL;
    if (left_end) {
        root->left = NULL;
        root->right = left;
        root = left_end;
    }
    if (right_end) {
        root->left = NULL;
        root->right = right;
        root = right_end;
    }
    return root;
}

void flatten(TreeNode *root) {
    if (root)
        do_flatten(root);
}