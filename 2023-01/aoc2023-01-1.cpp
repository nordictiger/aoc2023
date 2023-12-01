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
    int sum = 0;
    while (getline(file, line))
    {
        string code(2, '0');
        int i = 0;
        while (!isdigit(line[i]) && i < line.length() - 1)
        {
            i++;
        }
        int j = line.length() - 1;
        while (!isdigit(line[j]) && j > i)
        {
            j--;
        }
        code[0] = line[i];
        code[1] = line[j];
        try
        {
            int num = stoi(code);
            sum += num;
        }
        catch (invalid_argument &e)
        {
            cerr << "Invalid argument: " << e.what() << endl;
            return 1;
        }
    }
    cout << "Sum is: " << sum << endl;
    file.close();
    return 0;
}