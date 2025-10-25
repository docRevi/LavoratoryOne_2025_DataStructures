#ifndef QUERY_H
#define QUERY_H

#include <iostream>
#include <string>

using namespace std;

template<typename T>
struct QueueNode {
    T value;
    QueueNode<T>* next;
    explicit QueueNode(const T& value) : value(value), next(nullptr) {}
};

template<typename T>
struct Queue {
    QueueNode<T>* head;
    QueueNode<T>* tail;

    Queue() : head(nullptr), tail(nullptr) {}

    void QPUSH(const T& value) {
        auto* newQueueNode = new QueueNode(value);
    
        if (head == nullptr) {
            head = tail =  newQueueNode;
            return;
        }
        tail->next = newQueueNode;
        tail = newQueueNode;
    }
    
    T QPOP() {
        if (head == nullptr) {
            cout << "Очередь пустая" << endl;
            return T();
        }
    
        auto* nodeToDelete = head;
        head = head->next;
        if (nodeToDelete == tail) tail = nullptr;
        T result = nodeToDelete->value;
        delete nodeToDelete;
        return result;
    }
    
    void print() {
        if (head == nullptr) {
            cout << "Очередь пустая" << endl;
            return;
        }
    
        auto* iter =  head;
        while (iter) {
            cout << iter->value << " ";
            iter = iter->next;
        }
        cout << endl;
    }
    
};


#endif
