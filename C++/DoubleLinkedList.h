#ifndef DOUBLELINKEDLIST_H
#define DOUBLELINKEDLIST_H

#include <string>

using namespace std;

struct DListNode {
    string value;
    DListNode* next;
    DListNode* prev;

    explicit DListNode(const string& value) : value(value), next(nullptr), prev(nullptr) {}
};

struct DoubleLinkedList {
    DListNode* head;
    DListNode* tail;

    DoubleLinkedList() : head(nullptr), tail(nullptr) {}

    void DPUSHFront(const string& value);
    void DPUSHBack(const string& value);
    void DINSERTAfter(const string& value, int key);
    void DINSERTBefore(const string& value, int key);

    void DDELHead();
    void DDELTail();
    void DDELBefore(int index);
    void DDELAfter(int index);
    void DDELValue(const string& value);

    int DGET(const string& value);
    void printForward();
    void printBackward();

};

#endif
