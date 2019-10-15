#pragma once

#include <boost/core/noncopyable.hpp>

struct TreeNode : private boost::noncopyable {
    int val;
    TreeNode *left;
    TreeNode *right;

    explicit TreeNode(int val) : val(val), left(nullptr), right(nullptr) {}
};