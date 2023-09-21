import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

class Solver {

  private static String parseInput() {
    String fileName = "input.txt";
    String line = "";

    try (BufferedReader br = new BufferedReader(new FileReader(fileName))) {
      line = br.readLine();
    } catch (IOException e) {
      e.printStackTrace();
    }

    return line;
  }

  private static int reactPolymer(String polymer) {
    StringBuilder res = new StringBuilder();
    int n = 0;
    for (char c: polymer.toCharArray()) {
      res.append(c);
      while ((n = res.length()) > 1 && oppositePolarity(res.charAt(n - 1), res.charAt(n - 2))) {
        res.delete(n - 2, n);
      }
    }
    return res.length();
  }

  private static boolean oppositePolarity(char a, char b) {
    return Character.toLowerCase(a) == Character.toLowerCase(b) && a != b;
  }

  private static int part1(String polymer) {
    return reactPolymer(polymer);
  }

  private static int part2(String polymer) {
    int res = polymer.length();
    for (int i = (int) 'a'; i <= (int) 'z'; i++) {
      char remove = (char) i;
      StringBuilder check = new StringBuilder();
      for (char c: polymer.toCharArray()) {
        if (Character.toLowerCase(c) != remove) {
          check.append(c);
        }
      }
      int n = reactPolymer(check.toString());
      if (n < res) {
        res = n;
      }
    }
    return res;

  }

  public static void main(String[] args) {
    String polymer = parseInput();
    System.out.println(part1(polymer));
    System.out.println(part2(polymer));
  }
}
