#ifndef STACK_H
#define STACK_H

#include <iostream>
#include<string>

using namespace std;

template<typename T>
struct StackNode {
    T value;
    StackNode<T>* next;

    explicit StackNode (const T& value) : value(value), next(nullptr) {}
};

template<typename T>
struct Stack {
    StackNode<T>* top;

    Stack() : top(nullptr) {}

    void SPUSH(const T& value) {
        auto* newNode = new StackNode(value);
        newNode->next = top;
        top = newNode;
    }

    T SPOP() {
        if (top == nullptr) {
            cout << "Стек пустой" << endl;
            return T();
        }
    
        auto* nodeToDelete = top;
        top = top->next;
        string result = nodeToDelete->value;
        delete nodeToDelete;
        return result;
    }

    void print() {
        if (top == nullptr) {
            cout << "Стек пустой" << endl;
            return;
        }
    
        auto* iter = top;
        while (iter != nullptr) {
            cout << iter->value << " ";
            iter = iter->next;
        }
        cout << endl;
    }
};

#endif

