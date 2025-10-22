#ifndef QUERY_H
#define QUERY_H
#include <string>

using namespace std;

struct QueueNode {
    string value;
    QueueNode* next;
    explicit QueueNode(const string& value) : value(value), next(nullptr) {}
};

struct Queue {
    QueueNode* head;
    QueueNode* tail;

    Queue() : head(nullptr), tail(nullptr) {}

    void QPUSH(const string& value);
    string QPOP();
    void print();
};


#endif
