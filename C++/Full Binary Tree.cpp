#include "Full Binary Tree.h"

#include <iostream>
void Tree::TPUSH(const int value) {
    if (root == nullptr) {
        root = new TreeNode(value);
        return;
    }

    auto* iter = root;
    TreeNode* parent = nullptr;

    while (iter != nullptr) {
        parent = iter;
        if (value < iter->value) {
            iter = iter->left;
        }
        else if (value > iter->value) {
            iter = iter->right;
        }
        else {
            return;
        }
    }

    if (value < parent->value) {
        parent->left = new TreeNode(value);
    }
    else {
        parent->right = new TreeNode(value);
    }
}

bool Tree::TSEARCH(TreeNode* root, const int value) {
    auto* iter = root;

    while (iter != nullptr) {
        if (iter->value == value) {
            return true;
        }
        else if (iter->value <value) {
            iter = iter->left;
        }
        else {
            iter = iter->right;
        }
    }
    return false;
}

bool Tree::isFullBinaryTree(TreeNode* root) {
    if (root == nullptr) {
        return true;
    }

    if (root->left == nullptr && root->right == nullptr) {
        return true;
    }

    if (root->left != nullptr && root->right != nullptr) {
        return isFullBinaryTree(root->left) && isFullBinaryTree(root->right);
    }

    return false;
}

int Tree::heightTree(TreeNode* root) {
    if (!root) return 0;
    int lh = heightTree(root->left);
    int rh = heightTree(root->right);
    return 1 + (lh > rh ? lh : rh);
}

void Tree::printLevel(TreeNode* root, int level) {
    if (!root) return;
    if (level == 1)
        cout << root->value << " ";
    else {
        printLevel(root->left, level - 1);
        printLevel(root->right, level - 1);
    }
}

void Tree::print(TreeNode* root) {
    int h = heightTree(root);
    for (int i = 1; i <= h; ++i) {
        printLevel(root, i);
    }
}

void Tree::printInOrder(TreeNode* root) {
    if (!root) return;
    printInOrder(root->left);
    cout << root->value << " ";
    printInOrder(root->right);
}

void Tree::printPreOrder(TreeNode* root) {
    if (!root) return;
    cout << root->value << " ";
    printPreOrder(root->left);
    printPreOrder(root->right);
}

void Tree::printPostOrder(TreeNode* root) {
    if (!root) return;
    printPostOrder(root->left);
    printPostOrder(root->right);
    cout << root->value << " ";
}
