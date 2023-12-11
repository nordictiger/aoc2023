#include <fstream>
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

const string north = "|7FS";
const string east = "-J7S";
const string south = "|LJS";
const string west = "-FLS";

vector<string> load_data(string file_name)
{

    vector<string> pipe_map;

    ifstream file(file_name);
    if (!file.is_open())
    {
        cerr << "Error opening file" << endl;
        return vector<string>();
    }

    string line;
    while (getline(file, line))
    {
        pipe_map.push_back(line);
    }

    file.close();
    return pipe_map;
}

pair<int, int> find_start(vector<string> pipe_map)
{
    pair<int, int> start;
    start.first = 0;
    start.second = 0;
    for (int i = 0; i < pipe_map.size(); i++)
    {
        for (int j = 0; j < pipe_map[0].size(); j++)
        {
            if (pipe_map[i][j] == 'S')
            {
                start.first = i;
                start.second = j;
                return start;
            }
        }
    }
    return start;
}

void add_step(vector<pair<int, int>> &steps, int x, int y)
{
    pair<int, int> step;
    step.first = x;
    step.second = y;
    steps.push_back(step);
}

void add_north_step(vector<string> &pipe_map, vector<pair<int, int>> &steps, int x, int y)
{
    x--;
    if (x >= 0 && (north.find(pipe_map[x][y]) != std::string::npos))
    {
        add_step(steps, x, y);
    }
}

void add_east_step(vector<string> &pipe_map, vector<pair<int, int>> &steps, int x, int y)
{
    y++;
    if (y < pipe_map[0].size() && (east.find(pipe_map[x][y]) != std::string::npos))
    {
        add_step(steps, x, y);
    }
}

void add_south_step(vector<string> &pipe_map, vector<pair<int, int>> &steps, int x, int y)
{
    x++;
    if (x < pipe_map.size() && (south.find(pipe_map[x][y]) != std::string::npos))
    {
        add_step(steps, x, y);
    }
}

void add_west_step(vector<string> &pipe_map, vector<pair<int, int>> &steps, int x, int y)
{
    y--;
    if (y >= 0 && (west.find(pipe_map[x][y]) != std::string::npos))
    {
        add_step(steps, x, y);
    }
}

vector<pair<int, int>> find_valid_steps(vector<string> pipe_map, pair<int, int> start)
{
    vector<pair<int, int>> result;
    switch (pipe_map[start.first][start.second])
    {
    case '|':
        add_north_step(pipe_map, result, start.first, start.second);
        add_south_step(pipe_map, result, start.first, start.second);
        break;
    case '-':
        add_east_step(pipe_map, result, start.first, start.second);
        add_west_step(pipe_map, result, start.first, start.second);
        break;
    case 'L':
        add_north_step(pipe_map, result, start.first, start.second);
        add_east_step(pipe_map, result, start.first, start.second);
        break;
    case 'J':
        add_north_step(pipe_map, result, start.first, start.second);
        add_west_step(pipe_map, result, start.first, start.second);
        break;
    case '7':
        add_south_step(pipe_map, result, start.first, start.second);
        add_west_step(pipe_map, result, start.first, start.second);
        break;
    case 'F':
        add_south_step(pipe_map, result, start.first, start.second);
        add_east_step(pipe_map, result, start.first, start.second);
        break;
    case 'S':
        add_north_step(pipe_map, result, start.first, start.second);
        add_east_step(pipe_map, result, start.first, start.second);
        add_south_step(pipe_map, result, start.first, start.second);
        add_west_step(pipe_map, result, start.first, start.second);
        break;
    }
    return result;
}

int solve1(vector<string> pipe_map)
{
    auto start = find_start(pipe_map);
    auto valid_steps = find_valid_steps(pipe_map, start);
    if (valid_steps.size() != 2)
    {
        cout << "There should be exactly 2 valid directions in the loop." << endl;
        return -1;
    }
    auto current_position = valid_steps[0];
    auto previous_position = start;
    int steps = 1;

    while (pipe_map[current_position.first][current_position.second] != 'S')
    {
        auto valid_steps = find_valid_steps(pipe_map, current_position);
        if (valid_steps.size() != 2)
        {
            cout << "There should be exactly 2 valid directions in the loop." << endl;
            return -1;
        }
        if (valid_steps[0].first == previous_position.first && valid_steps[0].second == previous_position.second)
        {
            previous_position = current_position;
            current_position = valid_steps[1];
        }
        else
        {
            previous_position = current_position;
            current_position = valid_steps[0];
        }
        steps++;
        if (steps > 19600)
        {
            cout << "Too many steps." << endl;
            return -1;
        }
    }
    return steps - 1;
}

pair<vector<vector<int>>, int> create_loop_map(vector<string> pipe_map)
{
    int n = pipe_map.size();
    int m = pipe_map[0].size();
    vector<vector<int>> loop_map(n, std::vector<int>(m, 0));

    auto start = find_start(pipe_map);
    int map_value = 1;
    loop_map[start.first][start.second] = map_value;

    auto valid_steps = find_valid_steps(pipe_map, start);

    if (valid_steps.size() != 2)
    {
        cout << "There should be exactly 2 valid directions in the loop." << endl;
        return pair<vector<vector<int>>, int>(vector<vector<int>>(), 0);
    }
    auto current_position = valid_steps[0];
    auto previous_position = start;
    while (pipe_map[current_position.first][current_position.second] != 'S')
    {
        loop_map[current_position.first][current_position.second] = ++map_value;
        auto valid_steps = find_valid_steps(pipe_map, current_position);
        if (valid_steps.size() != 2)
        {
            cout << "There should be exactly 2 valid directions in the loop." << endl;
            return pair<vector<vector<int>>, int>(vector<vector<int>>(), 0);
        }
        if (valid_steps[0].first == previous_position.first && valid_steps[0].second == previous_position.second)
        {
            previous_position = current_position;
            current_position = valid_steps[1];
        }
        else
        {
            previous_position = current_position;
            current_position = valid_steps[0];
        }
    }
    return pair<vector<vector<int>>, int>(loop_map, map_value);
}

pair<int, int> get_next_tile(vector<vector<int>> &loop_map, int loop_size, int i, int j)
{
    int next_value = loop_map[i][j];
    if (loop_map[i][j] == loop_size)
    {
        next_value = 1;
    }
    else
    {
        next_value++;
    }
    for (int k = max(0, i - 1); k < min(i + 2, int(loop_map.size())); k++)
    {
        for (int l = max(0, j - 1); l < min(j + 2, int(loop_map[0].size())); l++)
        {
            if (loop_map[k][l] == next_value)
            {
                return pair<int, int>(k, l);
            }
        }
    }
    cout << "Next tile not found"
         << "i: " << i << " j: " << j << endl;
    return pair<int, int>(-1, -1);
}

pair<int, int> get_prev_tile(vector<vector<int>> &loop_map, int loop_size, int i, int j)
{
    int next_value = loop_map[i][j];
    if (loop_map[i][j] == 1)
    {
        next_value = loop_size;
    }
    else
    {
        next_value--;
    }
    for (int k = max(0, i - 1); k < min(i + 2, int(loop_map.size())); k++)
    {
        for (int l = max(0, j - 1); l < min(j + 2, int(loop_map[0].size())); l++)
        {
            if (loop_map[k][l] == next_value)
            {
                return pair<int, int>(k, l);
            }
        }
    }
    cout << "Previous tile not found "
         << "i: " << i << " j: " << j << endl;
    return pair<int, int>(-1, -1);
}

pair<bool, int> check_crossing(vector<vector<int>> &loop_map, int loop_size, int i, int j)
{
    int next_shift = 0;
    int k = i;
    int l = j;
    while (next_shift == 0)
    {
        auto next_tile = get_next_tile(loop_map, loop_size, k, l);
        k = next_tile.first;
        l = next_tile.second;
        next_shift = l - j;
    }

    int prev_shift = 0;
    int m = i;
    int n = j;
    while (prev_shift == 0)
    {
        auto prev_tile = get_prev_tile(loop_map, loop_size, m, n);
        m = prev_tile.first;
        n = prev_tile.second;
        prev_shift = n - j;
    }
    bool is_crossing = (next_shift + prev_shift) == 0;
    int skip_tiles = abs(k - m);
    return pair<bool, int>(is_crossing, skip_tiles);
}

bool check_enclosed(vector<vector<int>> &loop_map, int loop_size, int i, int j)
{
    int crossing_count = 0;

    for (int k = i + 1; k < loop_map.size(); k++)
    {
        if (loop_map[k][j] != 0)
        {
            // is it crossing?
            auto crossing_info = check_crossing(loop_map, loop_size, k, j);
            if (crossing_info.first)
            {
                crossing_count++;
            }
            k += crossing_info.second;
        }
    }
    if (crossing_count % 2 == 0)
    {
        return false;
    }
    return true;
}

int solve2(vector<vector<int>> loop_map, int loop_size)
{
    int enclosed_area = 0;

    for (int i = 0; i < loop_map.size(); i++)
    {
        for (int j = 0; j < loop_map[0].size(); j++)
        {
            if (loop_map[i][j] == 0)
            {
                if (check_enclosed(loop_map, loop_size, i, j))
                {
                    enclosed_area++;
                }
            }
        }
    }
    return enclosed_area;
}

int main()
{
    auto data = load_data("input.txt");
    auto loop_size = solve1(data);
    cout << "Puzzle 1 - steps to the farthest point: " << (loop_size / 2) + 1 << endl;

    auto loop_map = create_loop_map(data);
    cout << "Loop map with :" << loop_map.second << " elements" << endl;
    auto enclosed_area = solve2(loop_map.first, loop_map.second);
    cout << "Puzzle 2 - enclosed area: " << enclosed_area << endl;

    return 0;
}