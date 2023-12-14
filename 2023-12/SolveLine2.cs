using System.Text.RegularExpressions;

partial class Program
{
    static string GetKey(string line, List<int> numbers)
    {
        return $"{line} {string.Join(',', numbers)}";
    }
    static long SolveLine2(string line, List<int> numbers)
    {

        function_calls++;
        // Do we have result in cache?
        string original_line = line;
        long cached_result = cache.GetValueOrDefault(GetKey(original_line, numbers), -1);
        if (cached_result != -1)
        {
            cache_hits++;
            return cached_result;
        }
        // solve the end of recursion
        if (numbers.Count == 0)
        {
            if (line.Contains('#'))
            {
                // We have used all numbers but there are still #s in the line
                return 0;
            }
            return 1;
        }
        // we are not at the end of the numbers
        long arrangements_in_level = 0;
        string pattern_string = $"([\\?#]{{{numbers[0]}}})([\\?\\.]|$)";
        do
        {
            var match = Regex.Match(line, pattern_string);
            if (!match.Success)
            {
                // no more matches, we do not need to follow this branch
                break;
            }
            var match_index = match.Groups[1].Index;
            var match_length = match.Groups[1].Length;
            // did we skip any #
            if (line[..match_index].Contains('#')) {
                // we did, so we are not following this branch
                break;
            }

            // There was match and we are not done with all numbers
            // Recurse with the rest of the line and the rest of the numbers
            line = line[match_index..];
            var result = SolveLine2(line[Math.Min(match_length + 1, line.Length)..], numbers.GetRange(1, numbers.Count - 1));
            arrangements_in_level += result;

            if (line[0] != '#')
            {
                line = line[1..];
            }
            else
            {
                break;
            }
        } while (true);
        cache.Add(GetKey(original_line, numbers), arrangements_in_level);
        return arrangements_in_level;
    }
}