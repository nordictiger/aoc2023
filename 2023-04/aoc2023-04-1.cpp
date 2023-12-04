#include <fstream>
#include <iostream>
#include <string>
#include <set>
#include <vector>
#include <algorithm>
#include <regex>

using namespace std;

pair<set<string>,set<string>> parse_line_set(const string& line) {
    regex number_or_bar("(?:\\d+:)|(\\d+|\\|)");

    sregex_iterator currentMatch(line.begin(), line.end(), number_or_bar);
    sregex_iterator lastMatch;
    smatch match;
    set<string> winning_numbers;
    set<string> selected_numbers;

    while (currentMatch != lastMatch) {
        match = *currentMatch;
        if (match[1].str() == "|") {
            break;
        }
        winning_numbers.insert(match[1].str());
        currentMatch++;
    }

    while (currentMatch != lastMatch) {
        match = *currentMatch;
        selected_numbers.insert(match[1].str());
        currentMatch++;
    }

    pair<set<string>,set<string>> result;
    result.first = winning_numbers;
    result.second = selected_numbers;
    return result;
}

int main()
{
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
        auto number_sets = parse_line_set(line);
        
        vector<string> intersection;
        set_intersection(number_sets.first.begin(), number_sets.first.end(), 
                            number_sets.second.begin(), number_sets.second.end(),
                            std::back_inserter(intersection));

        if (intersection.size() > 0) {
            sum += 1 << (intersection.size() - 1);
        }
    }
    cout << "Sum is: " << sum << endl;
    file.close();
    return 0;
}