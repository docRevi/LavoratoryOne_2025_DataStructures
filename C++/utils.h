#ifndef UTILS_H
#define UTILS_H

#include "Array.h"
#include "LinkedList.h"
#include "DoubleLinkedList.h"
#include "Stack.h"
#include "Queue.h"
#include "Full Binary Tree.h"
#include <fstream>
#include <iostream>
#include <string>
#include <vector>
#include <functional>
#include <map>
#include <sstream>

extern Array<string> arr;
extern LinkedList<string> list;
extern DoubleLinkedList<string> DList;
extern Stack<string> stack;
extern Queue<string> queue;
extern Tree tree;

// массив
inline void save_array(const string& filename, const Array<string>& arr) {
    ofstream out(filename);
    for (int i = 0; i < arr.size; ++i)
        out << arr.data[i] << " ";
}
inline void load_array(const string& filename, Array<string>& arr) {
    ifstream in(filename);
    string x;
    while(getline(in, x)){
        if (!x.empty())
            arr.Apush_back(x);
    }
}

//односвязанный список
inline void save_list(const string& filename, const LinkedList<string>& list) {
    ofstream out(filename);
    for (auto iter = list.head; iter; iter = iter->next)
        out << iter->value << " ";
}
inline void load_list(const string& filename, LinkedList<string>& list) {
    ifstream in(filename);
    string x;
    while(getline(in, x)){
        if (!x.empty())
            list.LPUSHBack(x);
    }
}

// двусвязный список
inline void save_dlist(const string& filename, const DoubleLinkedList<string>& dlist) {
    ofstream out(filename);
    for (auto iter = dlist.head; iter; iter = iter->next)
        out << iter->value << " ";
}
inline void load_dlist(const string& filename, DoubleLinkedList<string>& dlist) {
    ifstream in(filename);
    string x;
    while(getline(in, x)){
        if (!x.empty())
            dlist.DPUSHBack(x);
    }
}

//очередь
inline void save_queue(const string& filename, const Queue<string>& q) {
    ofstream out(filename);
    for (auto iter = q.head; iter; iter = iter->next)
        out << iter->value << " ";
}
inline void load_queue(const string& filename, Queue<string>& q) {
    ifstream in(filename);
    string x;
    while(getline(in, x)){
        if (!x.empty())
            q.QPUSH(x);
    }
}

// стек
inline void save_stack(const string& filename, const Stack<string>& db) {
    ofstream out(filename);
    for (auto iter = db.top; iter; iter = iter->next)
        out << iter->value << " ";
}
inline void load_stack(const string& filename, Stack<string>& db) {
    ifstream in(filename);
    string x;
    vector<string> vals;
    while(getline(in, x)){
        if (!x.empty())
            vals.push_back(x);
    }
    for (auto iter = vals.rbegin(); iter != vals.rend(); ++iter)
        db.SPUSH(*iter);
}

// дерево
inline void save_tree_helper(TreeNode* node, ofstream& out) {
    if(!node) return;
    out << node->value << " ";
    save_tree_helper(node->left, out);
    save_tree_helper(node->right, out);
}

inline void save_tree(const string& filename, const Tree& t) {
    ofstream out(filename);
    save_tree_helper(t.root, out);
}

inline void load_tree(const string& filename, Tree& t) {
    ifstream in(filename);
    string x;
    while(getline(in, x)){
        if (!x.empty())
            t.TPUSH(stoi(x));
    }
}

void execute_query(const string& query);

#endif //UTILS_H
