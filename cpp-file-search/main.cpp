#include <iostream>
#include "FileSearcher.h"

int main() {
    std::string path, keyword;

    std::cout << "Enter directory path: ";
    std::getline(std::cin, path);

    std::cout << "Enter file name keyword: ";
    std::getline(std::cin, keyword);

    FileSearcher searcher;
    searcher.search(path, keyword);

    return 0;
}