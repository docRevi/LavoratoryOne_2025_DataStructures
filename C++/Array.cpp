#include "Array.h"
#include <iostream>
using namespace std;

void Array::reserve(int new_capacity) {
    if (new_capacity <= capacity) return;
    string* new_data = new string[new_capacity];
    for (int i = 0; i < size; ++i) {
        new_data[i] = data[i];
    }
    delete[] data;
    data = new_data;
    capacity = new_capacity;
}

void Array::Aadd(const int index, const string& value) {
    if (index < 0 or index > size) return;
    if (index == size) {
        reserve(size + 1);
    }
    data[index] = value;
}

void Array::Apush_back(const string& value) {
    int newSize = size + 1;
    auto* newArray = new string[newSize];

    for (int i = 0; i < size; i++) {
        newArray[i] = data[i];
    }

    newArray[size] = value;

    delete[] data;

    data = newArray;
    size = newSize;
}

string Array::Aget(const int index) {
    if (index < 0 or index >= size) return "";
    const string& elem = data[index];
    return elem;
}

void Array::Adel(const int index){
    data[index] = "";
}

void Array::Areplace(const int index, const string& value) {
    if (index < 0 or index >= size) return;
    data[index] = value;
}

int Array::Asize() {
    return size;
}

void Array::Aprint() {
    for (int i = 0; i < size; i++) {
        cout << data[i] << " " ;
    }
    cout << endl;
}