#ifndef LINKEDLIST_H
#define LINKEDLIST_H

#include <iostream>
#include <string>

using namespace std;

template<typename T>
struct ListNode {
    T value;
    ListNode<T>* next;

    explicit ListNode (const T& value) : value(value), next(nullptr) {}
};

template<typename T>
struct LinkedList {
    ListNode<T>* head;
    ListNode<T>* tail;

    LinkedList() : head(nullptr), tail(nullptr) {}

    void LPUSHBack(const T &value) {
        auto* newNode = new ListNode(value);
        if (head == nullptr) {
            head = newNode;
            tail = newNode;
        }
        tail->next = newNode;
        tail = newNode;
    }
    
    void LPUSHFront(const T& value) {
        auto* newNode = new ListNode(value);
        if (head == nullptr) {
            head = newNode;
            tail = newNode;
        }
        newNode->next = head;
        head = newNode;
    }
    
    void LINSERTAfter(const T& value, const int index) {
        auto* newNode = new ListNode(value);
        auto* iter = head;
        for (int i = 1; i < index; i++ ) {
            iter = iter->next;
        }
        if (!iter) return;
        newNode->next = iter->next;
        iter->next = newNode;
    }
    
    void LINSERTBefore(const T& value, const int index) {
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
    
    void LDELHead() {
        if (head == nullptr) return;
    
        auto* nodeToDelete = head;
        head = head->next;
        delete nodeToDelete;
    }
    
    void LDELTail() {
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
    
    void LDELAfter(const int index) {
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
    
    void LDELBefore(const int index) {
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
    
    void LDEL(const T& value) {
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
    
    void print() {
        auto* current = head;
        while (current != nullptr) {
            cout << current->value << " ";
            current = current->next;
        }
        cout << endl;
    }
    
    int LGET(const T& value) {
        int index = 1;
        auto* current = head;
    
        while (current->next!= nullptr) {
            if (current->value == value) {
                return index;
            }
            current = current->next;
            index++;
        }
    
        return -1 ;
    }

};

#endif
