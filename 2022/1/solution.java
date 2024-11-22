import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.nio.file.*;
import java.util.ArrayList;
import java.util.List;
import java.util.PriorityQueue;

class Solver {
  private static List<Integer> parseInput() {
    String fileName = "input.txt";
    Path currentDirectory = Paths.get("").toAbsolutePath();
    Path filePath = Paths.get(currentDirectory.toString(), fileName);
    List<Integer> elves = new ArrayList<>();

    try (BufferedReader br = new BufferedReader(new FileReader(filePath.toString()))) {
      String line;
      int elf = 0;
      while ((line = br.readLine()) != null) {
        try {
          int number = Integer.parseInt(line.trim());
          elf += number;
        } catch (NumberFormatException e) {
          elves.add(elf);
          elf = 0;
        }
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

    return elves;
  }

  public static void main(String[] args) {
    List<Integer> input = parseInput();
    int max = 0;
    PriorityQueue<Integer> max3 = new PriorityQueue<>(3, null);

    for (int n: input) {
      if (n > max)
        max = n;
      max3.offer(n);
      if (max3.size() > 3)
        max3.poll();
    }

    System.out.println(max);

    int part2 = 0;
    while (!max3.isEmpty())
      part2 += max3.poll();

    System.out.println(part2);
  }
}
