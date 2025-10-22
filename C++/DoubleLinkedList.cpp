#include <iostream>

#include "DoubleLinkedList.h"

using namespace std;

void DoubleLinkedList::DPUSHFront(const string& value) {
    auto* newHead = new DListNode(value);

    if (head == nullptr) {
        head = tail = newHead;
        return;
    }

    newHead->next = head;
    head->prev = newHead;
    head = newHead;
}

void DoubleLinkedList::DPUSHBack(const string& value) {
    auto* newTail = new DListNode(value);

    if (head == nullptr) {
        head = tail = newTail;
        return;
    }

    tail->next = newTail;
    newTail->prev = tail;
    tail = newTail;
}

void DoubleLinkedList::DINSERTAfter(const string& value, const int key) {
    if (head == nullptr or key < 1) return;

    auto* iter = head;

    for (int i = 1; i < key; i++) {
        iter = iter->next;
        if (!iter) return;
    }

    auto* newNode = new DListNode(value);
    newNode->next = iter->next;
    newNode->prev = iter;

    if (iter->next != nullptr) {
        iter->next->prev = newNode;
    }
    else {
        tail = newNode;
    }

    iter->next = newNode;
}

void DoubleLinkedList::DINSERTBefore(const string& value, const int key) {
    if (head == nullptr or key < 1) return;

    if (key == 1) {DPUSHFront(value); return; }

    auto* iter = head;
    for (int i = 0; i < key - 1; i++) {
        iter = iter->next;
        if (!iter) return;
    }

    auto* newNode = new DListNode(value);

    newNode->prev = iter->prev;
    iter->prev->next = newNode;
    newNode->next = iter;
    iter->prev = newNode;
}

void DoubleLinkedList::DDELHead() {
    if (head == nullptr) return;

    auto* temp = head->next;
    temp->prev = nullptr;
    delete head;
    head = temp;
}

void DoubleLinkedList::DDELTail() {
    if (head == nullptr) return;

    auto* nodeToDelete = tail;

    tail->prev->next = nullptr;
    tail = tail->prev;
    delete nodeToDelete;
}

void DoubleLinkedList::DDELBefore(const int index) {
    if (head == nullptr || index < 1) return;

    if (index == 1) DDELHead();

    
    auto* iter = head;
    
    for (int i = 1; i < index - 1; i++) {
        iter = iter->next;
        if (!iter) return;
    }

    iter->prev->next = iter->next;
    iter->next->prev = iter->prev;
    delete iter;
}

void DoubleLinkedList::DDELAfter(const int index) {
    if (head == nullptr || index < 1) return;
    
    auto* iter = head;

    for (int i = 1; i < index; i++) {
        iter = iter->next;
        if (!iter) return;
    }

    iter->prev->next = iter->next;
    iter->next->prev = iter->prev;
}

void DoubleLinkedList::DDELValue(const string& value) {
    if (head == nullptr) return;

    if (head->value == value) DDELHead();
    if (tail->value == value) DDELTail();
    
    auto* iter = head;
    while (iter->value != value) {
        iter = iter->next;
        if (!iter) return;
    }

    iter->prev->next = iter->next;
    iter->next->prev = iter->prev;
}

int DoubleLinkedList::DGET(const string& value) {
    int index = 0;
    auto* iter = head;
    
    while (iter != nullptr) {
        if (iter->value == value) {
            return index;
        }
        iter = iter->next;
        index++;
    }
    
    return -1;
}

void DoubleLinkedList::printForward() {
    auto* iter = head;
    while (iter) {
        cout << iter->value << " ";
        iter = iter->next;
    }
    cout << endl;
}

void DoubleLinkedList::printBackward() {
    auto* iter = tail;
    while (iter) {
        cout << iter->value << " ";
        iter = iter->prev;
    }
    cout << endl;
}