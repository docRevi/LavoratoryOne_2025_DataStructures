#ifndef LINKEDLIST_H
#define LINKEDLIST_H

#include <string>

using namespace std;

struct ListNode {
    string value;
    ListNode* next;

    explicit ListNode (const string& value) : value(value), next(nullptr) {}
};

struct LinkedList {
    ListNode* head;
    ListNode* tail;

    LinkedList() : head(nullptr), tail(nullptr) {}

    void LPUSHBack(const string& value);
    void LPUSHFront(const string& value);
    void LINSERTAfter(const string& value, int index);
    void LINSERTBefore(const string& value, int index);

    void LDELHead();
    void LDELTail();
    void LDELAfter(int index);
    void LDELBefore(int index);
    void LDEL(const string& value);
    void print();
    int LGET(const string& value);

};

#endif
