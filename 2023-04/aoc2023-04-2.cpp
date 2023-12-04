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

int get_intersections(set<string> winning_numbers, set<string> selected_numbers) {
    vector<string> intersection;
    set_intersection(winning_numbers.begin(), winning_numbers.end(), 
                        selected_numbers.begin(), selected_numbers.end(),
                        std::back_inserter(intersection));
    return intersection.size();
}

void process_cards(vector<tuple<int, int, int>>& cards_info) {
    for (int i = 0; i < cards_info.size(); i++) {
        for (int j = i+1; j < i + get<1>(cards_info[i]) + 1; j++) {
            get<2>(cards_info[j]) += get<2>(cards_info[i]);
        }
    }
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
    int index = 0;
    vector<tuple<int, int, int>> cards_info;
    while (getline(file, line))
    {
        auto number_sets = parse_line_set(line);
        
        int number_of_intersections = get_intersections(number_sets.first, number_sets.second);

        tuple<int, int, int> card_info = make_tuple(index, number_of_intersections, 1);
        cards_info.push_back(card_info);
        index++;

    }

    process_cards(cards_info);

    int sum = 0;
    for (int i = 0; i < cards_info.size(); i++) {
        sum += get<2>(cards_info[i]);
    }

    cout << "Sum is: " << sum << endl;
    file.close();
    return 0;
}