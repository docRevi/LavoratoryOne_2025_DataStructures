#include <iostream>

#include "Stack.h"

void Stack::SPUSH(const string& value) {
    auto* newNode = new StackNode(value);
    newNode->next = top;
    top = newNode;
}

string Stack::SPOP() {
    if (top == nullptr) {
        cout << "Стек пустой" << endl;
        return "";
    }

    auto* nodeToDelete = top;
    top = top->next;
    string result = nodeToDelete->value;
    delete nodeToDelete;
    return result;
}

void Stack::print() {
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

