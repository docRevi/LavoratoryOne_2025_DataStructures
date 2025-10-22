#ifndef ARRAY_H
#define ARRAY_H

#include <string>

using namespace std;

struct Array {
    string *data;
    int size;
    int capacity;

    Array() : data(nullptr), size(0), capacity(0) {}

    void reserve(int new_capacity);
    void Aadd(int index, const string &value);
    void Apush_back(const string &value);
    string Aget(int index);
    void Adel(int index);
    void Areplace (int index, const string &value);
    int Asize();
    void Aprint();
};

#endif
