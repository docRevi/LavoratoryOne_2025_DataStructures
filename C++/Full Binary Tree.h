#ifndef FULL_BINARY_TREE_H
#define FULL_BINARY_TREE_H

#include <string>

using namespace std;

struct TreeNode {
    int value;
    TreeNode* left;
    TreeNode* right;

    explicit TreeNode(int value) : value(value), left(nullptr), right(nullptr) {}
};

struct Tree {
    TreeNode* root;

    Tree(): root(nullptr) {}

    void TPUSH(int value);
    bool TSEARCH(TreeNode* root, int value);
    bool isFullBinaryTree(TreeNode* root);
    int heightTree(TreeNode* root);
    void printLevel(TreeNode* root, int level);
    void print(TreeNode* root);
    void printInOrder(TreeNode* root);
    void printPreOrder(TreeNode* root);
    void printPostOrder(TreeNode* root);

};

void printTreeHorizontal(TreeNode* root, const string& prefix = "", bool isLeft = true);


#endif
