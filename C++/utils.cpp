#include "utils.h"
#include <fstream>
#include <iostream>
#include <string>
#include <vector>
#include <functional>
#include <map>
#include <sstream>

Array arr;
LinkedList list;
DoubleLinkedList DList;
Stack stack;
Queue queue;
Tree tree;

using namespace std;

// массив
void save_array(const string& filename, const Array& arr) {
    ofstream out(filename);
    for (int i = 0; i < arr.size; ++i)
        out << arr.data[i] << " ";
}
void load_array(const string& filename, Array& arr) {
    ifstream in(filename);
    string x;
    while(getline(in, x)){
        if (!x.empty())
            arr.Apush_back(x);
    }
}

void save_list(const string& filename, const LinkedList& list) {
    ofstream out(filename);
    for (auto iter = list.head; iter; iter = iter->next)
        out << iter->value << " ";
}
void load_list(const string& filename, LinkedList& list) {
    ifstream in(filename);
    string x;
    while(getline(in, x)){
        if (!x.empty())
            list.LPUSHBack(x);
    }
}

// двусвязный список
void save_dlist(const string& filename, const DoubleLinkedList& dlist) {
    ofstream out(filename);
    for (auto iter = dlist.head; iter; iter = iter->next)
        out << iter->value << " ";
}
void load_dlist(const string& filename, DoubleLinkedList& dlist) {
    ifstream in(filename);
    string x;
    while(getline(in, x)){
        if (!x.empty())
            dlist.DPUSHBack(x);
    }
}

//очередь
void save_queue(const string& filename, const Queue& q) {
    ofstream out(filename);
    for (auto iter = q.head; iter; iter = iter->next)
        out << iter->value << " ";
}
void load_queue(const string& filename, Queue& q) {
    ifstream in(filename);
    string x;
    while(getline(in, x)){
        if (!x.empty())
            q.QPUSH(x);
    }
}

// стек
void save_stack(const string& filename, const Stack& db) {
    ofstream out(filename);
    for (auto iter = db.top; iter; iter = iter->next)
        out << iter->value << " ";
}
void load_stack(const string& filename, Stack& db) {
    ifstream in(filename);
    string x;
    vector<string> vals;
    while(getline(in, x)){
        if (!x.empty())
            vals.push_back(x);
    }
    for (auto iter = vals.rbegin(); iter != vals.rend(); ++iter)
        db.SPUSH(*iter);
}

// дерево
void save_tree_helper(TreeNode* node, ofstream& out) {
    if(!node) return;
    out << node->value << " ";
    save_tree_helper(node->left, out);
    save_tree_helper(node->right, out);
}
void save_tree(const string& filename, const Tree& t) {
    ofstream out(filename);
    save_tree_helper(t.root, out);
}
void load_tree(const string& filename, Tree& t) {
    ifstream in(filename);
    string x;
    while(getline(in, x)){
        if (!x.empty())
            t.TPUSH(stoi(x));
    }
}

map<string, function<void(const string&)>> one_arg = {
    // Массив
    {"APUSHB", [](const string& val){ arr.Apush_back(val); }},
    {"ADEL",   [](const string& idx){
        int i = stoi(idx);
        arr.Adel(i);
    }},
    {"AGET",   [](const string& idx){
        int i = stoi(idx);
        string result = arr.Aget(i);
        if (!result.empty())
            cout << result << endl;
        else
            cout << "Индекс вне диапазона" << endl;
    }},
    {"ASIZE",  [](const string&){
        cout << "Размер массива: " << arr.Asize() << endl;
    }},

    // Односвязный список
    {"LPUSHF", [](const string& val){ list.LPUSHFront(val); }},
    {"LPUSHB", [](const string& val){ list.LPUSHBack(val); }},
    {"LDEL",   [](const string& val){ list.LDEL(val); }},
    {"LDELH",  [](const string&){ list.LDELHead(); }},
    {"LDELT",  [](const string&){ list.LDELTail(); }},
    {"LDELA",  [](const string& idx){
        int i = stoi(idx);
        list.LDELAfter(i);
    }},
    {"LDELB",  [](const string& idx){
        int i = stoi(idx);
        list.LDELBefore(i);
    }},
    {"LGET",   [](const string& val){
        int result = list.LGET(val);
        if (result != -1)
            cout << "Найдено на позиции: " << result << endl;
        else
            cout << "Элемент не найден" << endl;
    }},

    // Двусвязный список
    {"DPUSHF", [](const string& val){ DList.DPUSHFront(val); }},
    {"DPUSHB", [](const string& val){ DList.DPUSHBack(val); }},
    {"DDEL",   [](const string& val){ DList.DDELValue(val); }},
    {"DDELH",  [](const string&){ DList.DDELHead(); }},
    {"DDELT",  [](const string&){ DList.DDELTail(); }},
    {"DDELA",  [](const string& idx){
        int i = stoi(idx);
        DList.DDELAfter(i);
    }},
    {"DDELB",  [](const string& idx){
        int i = stoi(idx);
        DList.DDELBefore(i);
    }},
    {"DGET",   [](const string& val){
        int result = DList.DGET(val);
        if (result != -1)
            cout << "Найдено на позиции: " << result << endl;
        else
            cout << "Элемент не найден" << endl;
    }},

    // Очередь
    {"QPUSH",  [](const string& val){ queue.QPUSH(val); }},
    {"QPOP",   [](const string&){
        string result = queue.QPOP();
        if (!result.empty())
            cout << "Извлечено: " << result << endl;
    }},

    // Стек
    {"SPUSH",  [](const string& val){ stack.SPUSH(val); }},
    {"SPOP",   [](const string&){
        string result = stack.SPOP();
        if (!result.empty())
            cout << "Извлечено: " << result << endl;
    }},

    // Дерево
    {"TPUSH",  [](const string& val){ tree.TPUSH(stoi(val)); }},
    {"TSEARCH", [](const string& val){
        bool found = tree.TSEARCH(tree.root, stoi(val));
        cout << (found ? "Найдено" : "Не найдено") << endl;
    }},
    {"ISFULL", [](const string&){
        bool isFull = tree.isFullBinaryTree(tree.root);
        cout << (isFull ? "Полное бинарное дерево" : "Не полное бинарное дерево") << endl;
    }},

    // Вывод структур
    {"PRINT", [](const string& target){
        if (target == "list") list.print();
        else if (target == "dlist") DList.printForward();
        else if (target == "array") arr.Aprint();
        else if (target == "queue") queue.print();
        else if (target == "stack") stack.print();
        else if (target == "tree") tree.print(tree.root);
        else cout << "Неизвестная структура: " << target << endl;
    }},
};

map<string, function<void(const string&, const string&)>> two_args = {
    // Односвязный список
    {"LINSERTA", [](const string& idx, const string& val){
        int i = stoi(idx);
        list.LINSERTAfter(val, i);
    }},
    {"LINSERTB", [](const string& idx, const string& val){
        int i = stoi(idx);
        list.LINSERTBefore(val, i);
    }},

    // Двусвязный список
    {"DINSERTA", [](const string& idx, const string& val){
        int i = stoi(idx);
        DList.DINSERTAfter(val, i);
    }},
    {"DINSERTB", [](const string& idx, const string& val){
        int i = stoi(idx);
        DList.DINSERTBefore(val, i);
    }},

    // Массив
    {"AADD", [](const string& idx, const string& val){
        int i = stoi(idx);
        arr.Aadd(i, val);
    }},
    {"AREPLACE", [](const string& idx, const string& val){
        int i = stoi(idx);
        arr.Areplace(i, val);
    }},
};

void execute_query(const string& query) {
    stringstream ss(query);
    string cmd;
    ss >> cmd;

    if (one_arg.count(cmd)) {
        string x;
        if (ss >> x) {
            one_arg[cmd](x);
        } else {
            cout << "Ошибка: команда " << cmd << " требует 1 аргумент\n";
        }
    }
    else if (two_args.count(cmd)) {
        string a, b;
        if (ss >> a >> b) {
            two_args[cmd](a, b);
        } else {
            cout << "Ошибка: команда " << cmd << " требует 2 аргумента\n";
        }
    }
    else {
        cout << "Неизвестная команда: " << cmd << "\n";
        cout << "Доступные команды:\n";
        cout << "Односвязный список: LPUSHF, LPUSHB, LDEL, LDELH, LDELT, LDELA, LDELB, LGET, LINSERTA, LINSERTB\n";
        cout << "Двусвязный список: DPUSHF, DPUSHB, DDEL, DDELH, DDELT, DDELA, DDELB, DGET, DINSERTA, DINSERTB\n";
        cout << "Массив: APUSHB, ADEL, AGET, ASIZE, AADD, AREPLACE\n";
        cout << "Очередь: QPUSH, QPOP\n";
        cout << "Стек: SPUSH, SPOP\n";
        cout << "Дерево: TPUSH, TSEARCH, ISFULL\n";
        cout << "Вывод: PRINT <структура>\n";
    }
}