import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;

class Solver {
  private static List<Integer> parseInput() {
    String fileName = "input.txt";
    List<Integer> ops = new ArrayList<>();

    try (BufferedReader br = new BufferedReader(new FileReader(fileName))) {
      String line;
      while ((line = br.readLine()) != null) {
        try {
          int op = Integer.parseInt(line);
          ops.add(op);
        } catch (NumberFormatException e) {
          throw e;
        }
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

    return ops;
  }

  private static int part1(List<Integer> ops) {
    int res = 0;
    for (int op: ops) {
      res += op;
    }
    return res;
  }

  private static int part2(List<Integer> ops) {
    int freq = 0;
    HashSet<Integer> set = new HashSet<>();
    set.add(freq);
    while (true) {
      for (int op: ops) {
        freq += op;
        if (set.contains(freq)) {
          return freq;
        }
        set.add(freq);
      }
    }
  }

  public static void main(String[] args) {
    List<Integer> ops = parseInput();
    System.out.println(part1(ops));
    System.out.println(part2(ops));

  }
}
