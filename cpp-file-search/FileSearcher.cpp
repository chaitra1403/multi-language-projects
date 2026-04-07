#include "FileSearcher.h"
#include <iostream>
#include <filesystem>

namespace fs = std::filesystem;

void FileSearcher::search(const std::string& path, const std::string& keyword) {
    for (const auto& entry : fs::recursive_directory_iterator(path)) {
        if (entry.is_regular_file()) {
            std::string filename = entry.path().filename().string();

            if (filename.find(keyword) != std::string::npos) {
                std::cout << "Found: " << entry.path() << std::endl;
                std::cout << "Size: " << entry.file_size() << " bytes\n";
                std::cout << "--------------------------\n";
            }
        }
    }
}