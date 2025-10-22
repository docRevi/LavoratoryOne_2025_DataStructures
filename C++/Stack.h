#ifndef STACK_H
#define STACK_H

using namespace std;

struct StackNode {
    string value;
    StackNode* next;

    explicit StackNode (const string& value) : value(value), next(nullptr) {}
};

struct Stack {
    StackNode* top;

    Stack() : top(nullptr) {}

    void SPUSH(const string& value);
    string SPOP();
    void print();
};

#endif

