#include <iostream>
#include <string>
#include "utils.h"

using namespace std;

string filename;

// загрузка
void load_all() {
    load_array(filename + ".array", arr);
    load_list(filename + ".list", list);
    load_dlist(filename + ".dlist", DList);
    load_queue(filename + ".queue", queue);
    load_stack(filename + ".stack", stack);
    load_tree(filename + ".tree", tree);
}

// сохранение
void save_all() {
    save_array(filename + ".array", arr);
    save_list(filename + ".list", list);
    save_dlist(filename + ".dblist", DList);
    save_queue(filename + ".queue", queue);
    save_stack(filename + ".stack", stack);
    save_tree(filename + ".tree", tree);
}

int main(int argc, char* argv[]) {
    string query;

    for (int i = 1; i < argc; ++i) {
        string arg = argv[i];
        if (arg == "--file" && i + 1 < argc) {
            filename = argv[++i];
        } else if (arg == "--query" && i + 1 < argc) {
            query = argv[++i];
        }
    }

    if (filename.empty()) {
        cout << "Укажите файл с помощью --file <имя_файла>\n";
        return 1;
    }

    load_all();

    if (!query.empty()) {
        execute_query(query);
        save_all();
    } else {
        cout << "Введите команду (или exit):\n";
        string line;
        while (getline(cin, line)) {
            if (line == "exit") break;
            execute_query(line);
            save_all();
        }
    }

    return 0;
}