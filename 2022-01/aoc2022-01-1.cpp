#include <fstream>
#include <iostream>
#include <string>

using namespace std;

int main()
{
    ifstream file("input.txt");
    if (!file.is_open())
    {
        cerr << "Error opening file" << endl;
        return 1;
    }

    string line;
    int max_sum = 0;
    int running_sum = 0;
    while (getline(file, line))
    {
        if (line.empty())
        {
            if (max_sum < running_sum)
            {
                max_sum = running_sum;
                running_sum = 0;
            }
            running_sum = 0;
            continue;
        }
        try
        {
            int num = stoi(line);
            running_sum += num;
        }
        catch (invalid_argument &e)
        {
            cerr << "Invalid argument: " << e.what() << endl;
            return 1;
        }
        catch (out_of_range &e)
        {
            cerr << "Out of range: " << e.what() << endl;
            return 1;
        }
    }
    if (max_sum < running_sum)
    {
        max_sum = running_sum;
    }

    cout << "Maximum is: " << max_sum << endl;
    file.close();
    return 0;
}