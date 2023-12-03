using System;
using System.ComponentModel.DataAnnotations;
using System.Globalization;
using System.IO;

class Program
{
    static List<List<char>> matrix = new();
    static int partsSum = 0;
    static int gearSum = 0;
    static void LoadMatrix(string filename)
    {
        try
        {
            foreach (string line in File.ReadLines(filename))
            {
                List<char> charLine = new List<char>(line);
                matrix.Add(charLine);
            }
        }
        catch (Exception ex)
        {
            Console.WriteLine($"An error occurred: {ex.Message}");
        }
    }
    static bool IsPartNumber(int i_start, int j_start, int j_end)
    {
        for (int i = Math.Max(0, i_start - 1); i <= Math.Min(i_start + 1, matrix.Count - 1); i++)
        {
            for (int j = Math.Max(0, j_start - 1); j <= Math.Min(j_end + 1, matrix[i].Count - 1); j++)
            {
                if (!(char.IsDigit(matrix[i][j]) || (matrix[i][j] == '.')))
                {
                    return true;
                }
            }
        }
        return false;
    }
    static void ProcessNumber(int i_start, int j_start, int j_end)
    {
        if (IsPartNumber(i_start, j_start, j_end))
        {
            List<char> numberRange = matrix[i_start].GetRange(j_start, j_end - j_start + 1);
            string extractedNumber = new string(numberRange.ToArray());
            partsSum += int.Parse(extractedNumber);
        }

    }
    static void SolvePuzzle1()
    {
        Console.WriteLine("Solve Puzzle 1");
        bool startNumber = false;
        int i_start = 0;
        int j_start = 0;
        int j_end = 0;
        for (int i = 0; i < matrix.Count; i++)
        {
            for (int j = 0; j < matrix[i].Count; j++)
            {
                if (char.IsDigit(matrix[i][j]))
                {
                    if (!startNumber)
                    {
                        i_start = i;
                        j_start = j;
                        startNumber = true;
                    }
                }
                else
                {
                    if (startNumber)
                    {
                        startNumber = false;
                        j_end = j - 1;
                        ProcessNumber(i_start, j_start, j_end);
                    }
                }
            }
            if (startNumber)
            {
                startNumber = false;
                j_end = matrix[i].Count - 1;
                ProcessNumber(i_start, j_start, j_end);
            }
        }
        Console.WriteLine($"Sum of all part numbers: {partsSum}");
    }

    static Tuple<int, int, int, int> GetFullNumber(int i_num, int j_num)
    {
        int j_start = j_num;
        int j_end = j_num;
        while (j_start > 0 && char.IsDigit(matrix[i_num][j_start-1]))
        {
            j_start--;
        }
        while (j_end < matrix[i_num].Count - 1 && char.IsDigit(matrix[i_num][j_end + 1]))
        {
            j_end++;
        }

        List<char> numberRange = matrix[i_num].GetRange(j_start, j_end - j_start + 1);
        string extractedNumber = new string(numberRange.ToArray());
        // Console.WriteLine($"Number at {i_num}, {j_num} is {extractedNumber}");
        partsSum += int.Parse(extractedNumber);
        return new Tuple<int, int, int,int>(i_num,j_start, j_end, int.Parse(extractedNumber));
    }
    static void ProcessGear(int i_gear, int j_gear)
    {
        HashSet<Tuple<int, int, int, int>> gearNumbers = new();

        for (int i = Math.Max(0, i_gear - 1); i <= Math.Min(i_gear + 1, matrix.Count - 1); i++)
        {
            for (int j = Math.Max(0, j_gear - 1); j <= Math.Min(j_gear + 1, matrix[i].Count - 1); j++)
            {
                if (char.IsDigit(matrix[i][j]))
                {
                    Tuple<int, int, int, int> number;    
                    number = GetFullNumber(i,j);
                    gearNumbers.Add(number);
                }
            }
        }
        if (gearNumbers.Count == 2)
        {
            // Console.WriteLine($"Gear at {i_gear}, {j_gear}");
            // Console.WriteLine($"Gear ratio: {gearNumbers.ElementAt(0).Item4} : {gearNumbers.ElementAt(1).Item4}");
            gearSum += gearNumbers.ElementAt(0).Item4 * gearNumbers.ElementAt(1).Item4;
        }
    }
    static void SolvePuzzle2()
    {
        Console.WriteLine("Solve Puzzle 2");
        for (int i = 0; i < matrix.Count; i++)
        {
            for (int j = 0; j < matrix[i].Count; j++)
            {
                if (matrix[i][j] == '*')
                {
                    ProcessGear(i, j);
                }
            }
        }
        Console.WriteLine($"Sum of all gear ratios: {gearSum}");        
    }


    static void Main()
    {
        string filePath = "input.txt";

        LoadMatrix(filePath);
        SolvePuzzle1();
        SolvePuzzle2();
    }
}
