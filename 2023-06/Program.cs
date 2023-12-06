using System;
using System.Text.RegularExpressions;

class Program
{
    static Tuple<List<int>, List<int>> LoadData(string filename)
    {
        List<int> times = new();
        List<int> distances = new();
        string pattern = @"(\d+)";
        try
        {
            using StreamReader reader = new StreamReader(filename);

            string line = reader.ReadLine()!;
            MatchCollection matches = Regex.Matches(line, pattern);
            foreach (Match match in matches.Cast<Match>())
            {
                times.Add(int.Parse(match.Value));
            }

            line = reader.ReadLine()!;
            matches = Regex.Matches(line, pattern);
            foreach (Match match in matches.Cast<Match>())
            {
                distances.Add(int.Parse(match.Value));
            }
        }
        catch (Exception ex)
        {
            Console.WriteLine("An error occurred: " + ex.Message);
            Environment.Exit(1);
        }
        Tuple<List<int>, List<int>> result = new Tuple<List<int>, List<int>>(times, distances);
        return result;
    }

    static long GetRaceResult(long time, long distance)
    {
        long firstTime = 0;
        for (long i = 1; i < time; i++)
        {
            if ((i*(time-i) > distance)) {
                firstTime = i;
                break;
            }
        }
        long secondTime = 0;
        for (long i = time - 1; i > 0; i--)
        {
            if ((i*(time-i) > distance)) {
                secondTime = i;
                break;
            }
        }
        return (secondTime - firstTime) + 1;
    }

    static void SolvePuzzle1(List<int> times, List<int> distances)
    {
        long product = 1;
        Console.WriteLine("Solve Puzzle 1");
        for (int i = 0; i < times.Count; i++)
        {
            product *= GetRaceResult(times[i], distances[i]);
        }
        Console.WriteLine($"Product: {product}");
    }

    static void SolvePuzzle2()
    {
        Console.WriteLine("Solve Puzzle 2");
        long result = GetRaceResult(44899691, 277113618901768);
        Console.WriteLine($"Number of ways: {result}");
    }

    static void Main()
    {
        string filePath = "input.txt";

        var data = LoadData(filePath);
        SolvePuzzle1(data.Item1, data.Item2);
        SolvePuzzle2();
    }
}
