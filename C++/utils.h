#ifndef UTILS_H
#define UTILS_H

#include "Array.h"
#include "LinkedList.h"
#include "DoubleLinkedList.h"
#include "Stack.h"
#include "Queue.h"
#include "Full Binary Tree.h"

extern Array arr;
extern LinkedList list;
extern DoubleLinkedList DList;
extern Stack stack;
extern Queue queue;
extern Tree tree;

void save_array(const string& filename, const Array& arr);
void load_array(const string& filename, Array& arr);

void save_list(const string& filename, const LinkedList& list);
void load_list(const string& filename, LinkedList& list);

void save_dlist(const string& filename, const DoubleLinkedList& dlist);
void load_dlist(const string& filename, DoubleLinkedList& dlist);

void save_queue(const string& filename, const Queue& q);
void load_queue(const string& filename, Queue& q);

void save_stack(const string& filename, const Stack& db);
void load_stack(const string& filename, Stack& db);

void save_tree(const string& filename, const Tree& t);
void load_tree(const string& filename, Tree& t);

void execute_query(const string& query);

#endif //UTILS_H
