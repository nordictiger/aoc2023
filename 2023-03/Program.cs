using System;
using System.IO;

class Program
{
    static void Main()
    {
        string filePath = "input.txt";

        try
        {
            // Check if the file exists
            if (File.Exists(filePath))
            {
                // Read and process each line of the file
                foreach (string line in File.ReadLines(filePath))
                {
                    // Process the line (for now, just print it)
                    Console.WriteLine(line);
                }
            }
            else
            {
                Console.WriteLine($"File not found: {filePath}");
            }
        }
        catch (Exception ex)
        {
            // Handle any errors that might occur
            Console.WriteLine($"An error occurred: {ex.Message}");
        }
    }
}
