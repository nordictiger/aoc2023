#include <fstream>
#include <iostream>
#include <string>
#include <unordered_map>
#include <regex>
#include <algorithm>

using namespace std;

int main()
{
    unordered_map<std::string, int> numberMap = {
        {"1", 1}, {"one", 1},   {"eno", 1},
        {"2", 2}, {"two", 2},   {"owt", 2},
        {"3", 3}, {"three", 3}, {"eerht", 3},
        {"4", 4}, {"four", 4},  {"ruof", 4},
        {"5", 5}, {"five", 5},  {"evif", 5},
        {"6", 6}, {"six", 6},   {"xis", 6},
        {"7", 7}, {"seven", 7}, {"neves", 7},
        {"8", 8}, {"eight", 8}, {"thgie", 8},
        {"9", 9}, {"nine", 9},  {"enin", 9}
    };

    regex start_pattern(R"((?:1|one|2|two|3|three|4|four|5|five|6|six|7|seven|8|eight|9|nine))");
    regex end_pattern(R"((?:1|eno|2|owt|3|eerht|4|ruof|5|evif|6|xis|7|neves|8|thgie|9|enin))");

    ifstream file("input.txt");
    if (!file.is_open())
    {
        cerr << "Error opening file" << endl;
        return 1;
    }

    string line;
    int sum = 0;
    while (getline(file, line))
    {
        cout << "Line is: " << line << endl;
        int code = 0;
        smatch match;

        if (regex_search(line, match, start_pattern)) {
            code = numberMap[match.str()]*10;
        } else {
            cerr << "No match found" << endl;
            return -1;
        }

        reverse(line.begin(), line.end());
        if (regex_search(line, match, end_pattern)) {
            code += numberMap[match.str()];
        } else {
            cerr << "No match found" << endl;
            return -1;
        }

        cout << "Code is: " << code << endl;
        sum += code;
    }
    cout << "Sum is: " << sum << endl;
    file.close();
    return 0;
}