using System.Text;
using System.Text.RegularExpressions;

partial class Program
{
    static List<Tuple<string, List<int>>> LoadData(string filename)
    {
        List<Tuple<string, List<int>>> result = new();
        string line_pattern = @"(.*) (.*)";
        string number_pattern = @"(\d+)";

        try
        {
            using StreamReader reader = new(filename);

            string? line;
            while ((line = reader.ReadLine()) != null)
            {
                var matches = Regex.Matches(line, line_pattern);
                List<int> numbers = new();
                var number_matches = Regex.Matches(matches[0].Groups[2].Value, number_pattern);
                foreach (var number in number_matches.Cast<Match>())
                {
                    numbers.Add(int.Parse(number.Value));
                }

                Tuple<string, List<int>> line_result = new(matches[0].Groups[1].Value, numbers);
                result.Add(line_result);
            }
        }

        catch (Exception ex)
        {
            Console.WriteLine("An error occurred: " + ex.Message);
            Environment.Exit(1);
        }
        return result;
    }

    static int SolveLine(List<List<(int, int)>> solutions, Stack<(int, int)> path, string line, int line_start, List<int> numbers, int depth)
    {
        // solve the end of recursion
        if (depth == numbers.Count)
        {
            // We are at the end of the numbers
            if (path.Count != numbers.Count)
            {
                Console.WriteLine("Path count and numbers count do not match at the end of numbers!");
                Console.WriteLine($"Path count: {path.Count}");
                Console.WriteLine($"Numbers count: {numbers.Count}");
                Console.WriteLine($"Line: {line}");
                Environment.Exit(1);
            }
            if (line_start < line.Length && (line[(line_start)..]).Contains('#'))
            {
                // We have used all numbers but there are still #s in the line
                return 0;
            }
            // We got the solution
            var solution = new List<(int, int)>(path);
            solution.Reverse();
            solutions.Add(solution);
            return 1;
        }
        // we are not at the end of the numbers
        int arrangements_in_level = 0;
        string pattern_string = $"([\\?#]{{{numbers[depth]}}})([\\?\\.]|$)";
        do
        {
            var match = Regex.Match(line[line_start..], pattern_string);
            if (!match.Success)
            {
                // no more matches, we do not need to follow this branch
                break;
            }
            var match_index = match.Groups[1].Index;
            var match_length = match.Groups[1].Length;
            // did we skip any #
            if (line.Substring(line_start, match_index).Contains('#')) {
                // we did, so we are not following this branch
                break;
            }

            // There was match and we are not done with all numbers
            // Recurse with the rest of the line and the rest of the numbers
            line_start += match_index;
            path.Push((line_start, match_length));
            int result = SolveLine(solutions, path, line, Math.Min(line_start + match_length + 1, line.Length), numbers, depth + 1);
            path.Pop();
            arrangements_in_level += result;

            if (line_start < line.Length && line[line_start] != '#')
            {
                line_start++;
            }
            else
            {
                break;
            }
        } while (true);

        return arrangements_in_level;
    }

    static void WritePaths(List<List<(int, int)>> solutions, int length)
    {
        foreach (var s in solutions)
        {
            var line = new string('.', length);
            var line_array = line.ToCharArray();
            foreach (var pair in s)
            {
                // Console.Write($"({pair.Item1}, {pair.Item2})");
                for (int j = pair.Item1; j < pair.Item1 + pair.Item2; j++)
                {
                    line_array[j] = '#';
                }
            }
            // Console.WriteLine();

            var output = new string(line_array);
            Console.WriteLine(output);
        }
    }

    static void SolvePuzzle1(List<Tuple<string, List<int>>> data)
    {
        Console.WriteLine("Solve Puzzle 1");
        int sum = 0;
        foreach (var item in data)
        {
            Stack<(int, int)> path = new();
            List<List<(int, int)>> solutions = new();
            int result = SolveLine(solutions, path, item.Item1, 0, item.Item2, 0);
            /*
            Console.WriteLine("------------------------");
            Console.Write($"{item.Item1} - ");
            foreach (var number in item.Item2)
            {
                Console.Write($"{number} ");
            }
            Console.WriteLine();
            WritePaths(solutions, item.Item1.Length);
            Console.WriteLine($"Result: {result}");
            */
            sum += result;
        }
        Console.WriteLine($"Sum: {sum}");
    }


    static List<Tuple<string, List<int>>> PrepareData2(List<Tuple<string, List<int>>> data)
    {
        List<Tuple<string, List<int>>> result = new();
        foreach (var item in data)
        {
            string new_line = $"{item.Item1}?{item.Item1}?{item.Item1}?{item.Item1}?{item.Item1}";
            List<int> new_numbers = new();
            for (int i = 0; i < 5; i++)
            {
                new_numbers.AddRange(item.Item2);
            }

            int sum = 0;
            foreach (var number in item.Item2)
            {
                sum += number;
            }
            result.Add(new Tuple<string, List<int>>(new_line, new_numbers));
        }
        return result;
    }

    static void SolvePuzzle2(List<Tuple<string, List<int>>> data)
    {
        var data2 = PrepareData2(data);
        Console.WriteLine("Solve Puzzle 2");
        int sum = 0;
        foreach (var item in data2)
        {
            Stack<(int, int)> path = new();
            List<List<(int, int)>> solutions = new();
            int result = SolveLine(solutions, path, item.Item1, 0, item.Item2, 0);
            Console.WriteLine(item.Item1);
            Console.WriteLine($"Result: {result}");
            sum += result;
        }
        Console.WriteLine($"Sum: {sum}");
    }

    static void Main()
    {
        string filePath = "input.txt";
        var data = LoadData(filePath);
        Console.WriteLine($"Loaded {data.Count} lines of data");
        SolvePuzzle1(data);
        SolvePuzzle2(data);

    }
}
