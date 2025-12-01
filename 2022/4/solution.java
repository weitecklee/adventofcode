import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

class Solver {

  public static class Pair {
    private boolean overlap1;
    private boolean overlap2;

    public Pair(int p1n1, int p1n2, int p2n1, int p2n2) {
      this.overlap1 = (p1n1 >= p2n1 && p1n1 <= p2n2 && p1n2 >= p2n1 && p1n2 <= p2n2) || (p2n1 >= p1n1 && p2n1 <= p1n2 && p2n2 >= p1n1 && p2n2 <= p1n2);
      this.overlap2 = (p1n1 >= p2n1 && p1n1 <= p2n2) || (p1n2 >= p2n1 && p1n2 <= p2n2) || (p2n1 >= p1n1 && p2n1 <= p1n2) || (p2n2 >= p1n1 && p2n2 <= p1n2);
    }

    public boolean isOverlap1() {
      return overlap1;
    }
    public boolean isOverlap2() {
      return overlap2;
    }

  }

  private static List<Pair> parseInput() {
    String fileName = "input.txt";
    List<Pair> pairs = new ArrayList<>();

    String re = "(\\d+)-(\\d+),(\\d+)-(\\d+)";
    Pattern pattern = Pattern.compile(re);


    try (BufferedReader br = new BufferedReader(new FileReader(fileName))) {
      String line;
      while ((line = br.readLine()) != null) {
        Matcher matcher = pattern.matcher(line);
        if (matcher.find()) {
          int p1n1 = Integer.parseInt(matcher.group(1));
          int p1n2 = Integer.parseInt(matcher.group(2));
          int p2n1 = Integer.parseInt(matcher.group(3));
          int p2n2 = Integer.parseInt(matcher.group(4));

          pairs.add(new Pair(p1n1, p1n2, p2n1, p2n2));
        }
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

    return pairs;
  }

  private static int part1(List<Pair> pairs) {
    int res = 0;
    for (Pair pair: pairs) {
      if (pair.isOverlap1()) {
        res++;
      }
    }
    return res;
  }

  private static int part2(List<Pair> pairs) {
    int res = 0;
    for (Pair pair: pairs) {
      if (pair.isOverlap2()) {
        res++;
      }
    }
    return res;
  }

  public static void main(String[] args) {
    List<Pair> pairs = parseInput();
    System.out.println(part1(pairs));
    System.out.println(part2(pairs));

  }
}
