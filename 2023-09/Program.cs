using System;
using System.Dynamic;
using System.Text.RegularExpressions;

class Program
{
    static List<List<int>> LoadData(string filename)
    {
        List<List<int>> result = new List<List<int>>();
        string pattern = @"(-?\d+)";
        try
        {
            using StreamReader reader = new StreamReader(filename);

            string? line;
            while ((line = reader.ReadLine()) != null)
            {
                MatchCollection matches = Regex.Matches(line, pattern);
                List<int> lineNumbers = new List<int>();
                foreach (Match match in matches.Cast<Match>())
                {
                    lineNumbers.Add(int.Parse(match.Value));
                }
                result.Add(lineNumbers);
            }
        }

        catch (Exception ex)
        {
            Console.WriteLine("An error occurred: " + ex.Message);
            Environment.Exit(1);
        }
        return result;
    }

    static List<List<int>> GetDerivedValues(List<int> numbers)
    {
        List<List<int>> derivedValues = new()
        {
            numbers
        };

        int index = 0;
        while (true)
        {
            if (derivedValues[index].Count <= 1)
            {
                Console.WriteLine("Not enough values to extrapolate");
                Environment.Exit(1);
            }
            List<int> newValues = new();
            bool allZeros = true;
            for (int i = 1; i < derivedValues[index].Count; i++)
            {
                int newValue = derivedValues[index][i] - derivedValues[index][i - 1];
                if (newValue != 0)
                {
                    allZeros = false;
                }
                newValues.Add(newValue);
            }
            derivedValues.Add(newValues);
            index++;
            if (allZeros)
            {
                break;
            }
        }
        return derivedValues;
    }
    static int GetExtraPolatedValueEnd(List<int> numbers)
    {
        List<List<int>> derivedValues = GetDerivedValues(numbers);
        int newNumber = 0;
        for (int i = derivedValues.Count - 1; i >= 0; i--)
        {
            newNumber += derivedValues[i][^1];
        }
        return newNumber;
    }
    static int GetExtraPolatedValueStart(List<int> numbers)
    {
        List<List<int>> derivedValues = GetDerivedValues(numbers);
        int newNumber = 0;
        for (int i = derivedValues.Count - 1; i >= 0; i--)
        {
            newNumber = derivedValues[i][0] - newNumber;
        }
        return newNumber;
    }

    static void SolvePuzzle1(List<List<int>> data)
    {
        Console.WriteLine("Solve Puzzle 1");
        int sum = 0;
        foreach (var line in data)
        {
            var value = GetExtraPolatedValueEnd(line);
            sum += value;
        }
        Console.WriteLine($"Sum: {sum}");
    }


    static void SolvePuzzle2(List<List<int>> data)
    {
        Console.WriteLine("Solve Puzzle 2");
        int sum = 0;
        foreach (var line in data)
        {
            var value = GetExtraPolatedValueStart(line);
            sum += value;
        }
        Console.WriteLine($"Sum: {sum}");
    }

    static void Main()
    {
        string filePath = "input.txt";

        var data = LoadData(filePath);
        SolvePuzzle1(data);
        SolvePuzzle2(data);
    }
}
