#include "utils.h"
#include <fstream>
#include <iostream>
#include <string>
#include <vector>
#include <functional>
#include <map>
#include <sstream>

Array<string> arr;
LinkedList<string> list;
DoubleLinkedList<string> DList;
Stack<string> stack;
Queue<string> queue;
Tree tree;

using namespace std;

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