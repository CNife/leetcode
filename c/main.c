#include "tree.h"

//typedef struct TreeNode TreeNode;

TreeNode *lowestCommonAncestor(TreeNode *root, TreeNode *p, TreeNode *q) {
    TreeNode *l = root->left ? lowestCommonAncestor(root->left, p, q) : NULL;
    TreeNode *r = root->right ? lowestCommonAncestor(root->right, p, q) : NULL;
    return root == p || root == q || (l && r) ? root : l ? l : r;
}