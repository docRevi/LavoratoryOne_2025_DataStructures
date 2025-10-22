#include <iostream>

#include "LinkedList.h"

using namespace std;

void LinkedList::LPUSHBack(const string &value) {
    auto* newNode = new ListNode(value);
    if (head == nullptr) {
        head = newNode;
        tail = newNode;
    }
    tail->next = newNode;
    tail = newNode;
}

void LinkedList::LPUSHFront(const string& value) {
    auto* newNode = new ListNode(value);
    if (head == nullptr) {
        head = newNode;
        tail = newNode;
    }
    newNode->next = head;
    head = newNode;
}

void LinkedList::LINSERTAfter(const string& value, const int index) {
    auto* newNode = new ListNode(value);
    auto* iter = head;
    for (int i = 1; i < index; i++ ) {
        iter = iter->next;
    }
    if (!iter) return;
    newNode->next = iter->next;
    iter->next = newNode;
}

void LinkedList::LINSERTBefore(const string& value, const int index) {
    if (head == nullptr) return;

    if (index == 1) LPUSHFront(value);

    auto* iter = head;
    for (int i = 1; i < index - 1; i++ ) {
        iter = iter->next;
    }

    if (!iter) return;

    auto* newNode = new ListNode(value);

    newNode->next = iter->next;
    iter->next = newNode;
}

void LinkedList::LDELHead() {
    if (head == nullptr) return;

    auto* nodeToDelete = head;
    head = head->next;
    delete nodeToDelete;
}


void LinkedList::LDELTail() {
    if (head == nullptr) return;

    if (head == tail) {
        delete head;
        head = tail = nullptr;
        return;
    }

    auto* iter = head;
    while (iter->next != tail) {
        iter = iter->next;
    }

    auto* nodeToDelete = tail;
    tail = iter;
    tail->next = nullptr;
    delete nodeToDelete;
}

void LinkedList::LDELAfter(const int index) {
    if (head == nullptr) return;

    auto* iter = head;
    for (int i = 1; i < index; i++ ) {
        iter = iter->next;
    }
    if (!iter) return;

    auto* nodeToDelete = iter->next;
    iter->next = nodeToDelete->next;

    delete nodeToDelete;
}

void LinkedList::LDELBefore(const int index) {
    if (head == nullptr or index <= 1) return;

    if (index == 2) {
        auto* nodeToDelete = head;
        head = head->next;
        delete nodeToDelete;
        return;
    }

    auto* iter = head;

    for (int i = 1; i < index - 2; i++ ) {
        iter = iter->next;
    }

    if (!iter) return;

    auto* nodeToDelete = iter->next;
    iter->next = nodeToDelete->next;

    delete nodeToDelete;
}

void LinkedList::LDEL(const string& value) {
    if (head == nullptr) return;

    if (head->value == value) LDELHead();
    if (tail->value == value) LDELTail();

    auto* iter = head;

    while (iter->next->value != value) {
        iter = iter->next;
    }

    auto* nodeToDelete = iter;
    iter->next = nodeToDelete->next;

    delete nodeToDelete;
}

void LinkedList::print() {
    auto* current = head;
    while (current != nullptr) {
        cout << current->value << " ";
        current = current->next;
    }
    cout << endl;
}

int LinkedList::LGET(const string& value) {
    int index = 1;
    auto* current = head;

    while (current->next!= nullptr) {
        if (current->value == value) {
            return index;
        }
        current = current->next;
        index++;
    }

    return -1;
}