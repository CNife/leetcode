#include "tree.h"

int sumOfLeftLeaves(TreeNode *root) {
    if (!root)
        return 0;

    int leftSum = 0;
    TreeNode *left = root->left;
    if (left)
        leftSum = (left->left || left->right) ? sumOfLeftLeaves(left) : left->val;
    int rightSum = sumOfLeftLeaves(root->right);
    return leftSum + rightSum;
}