#include "Queue.h"

#include <iostream>

void Queue::QPUSH(const string& value) {
    auto* newQueueNode = new QueueNode(value);

    if (head == nullptr) {
        head = tail =  newQueueNode;
        return;
    }
    tail->next = newQueueNode;
    tail = newQueueNode;
}

string Queue::QPOP() {
    if (head == nullptr) {
        cout << "Очередь пустая" << endl;
        return "";
    }

    auto* nodeToDelete = head;
    head = head->next;
    if (nodeToDelete == tail) tail = nullptr;
    string result = nodeToDelete->value;
    delete nodeToDelete;
    return result;
}

void Queue::print() {
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
