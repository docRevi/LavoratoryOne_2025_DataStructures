#ifndef ARRAY_H
#define ARRAY_H

#include <iostream>
#include <string>

using namespace std;

template<typename T>
struct Array {
    T *data;
    int size;
    int capacity;

    Array() : data(nullptr), size(0), capacity(0) {}

    void reserve(int new_capacity) {
        if (new_capacity <= capacity) return;
        T* new_data = new T[new_capacity];
        for (int i = 0; i < size; ++i) {
            new_data[i] = data[i];
        }
        delete[] data;
        data = new_data;
        capacity = new_capacity;
    }

    void Aadd(const int index, const T& value) {
        if (index < 0 or index > size) return;
        if (index == size) {
            reserve(size + 1);
        }
        data[index] = value;
    }

    void Apush_back(const T& value) {
        int newSize = size + 1;
        auto* newArray = new T[newSize];

        for (int i = 0; i < size; i++) {
            newArray[i] = data[i];
        }

        newArray[size] = value;

        delete[] data;

        data = newArray;
        size = newSize;
        }
    
    T Aget(const int index) {
        if (index < 0 or index >= size) return T();
        const T& elem = data[index];
        return elem;
    }

    void Adel(const int index){
        data[index] = "";
    }
    
    void Areplace(const int index, const T& value) {
        if (index < 0 or index >= size) return;
        data[index] = value;
    }
    
    int Asize() {
        return size;
    }
    
    void Aprint() {
        for (int i = 0; i < size; i++) {
            cout << data[i] << " " ;
        }
        cout << endl;
    }
};

#endif
