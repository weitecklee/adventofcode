import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

class Solver {
  private static List<String> parseInput() {
    String fileName = "input.txt";
    List<String> lines = new ArrayList<>();

    try (BufferedReader br = new BufferedReader(new FileReader(fileName))) {
      String line;
      while ((line = br.readLine()) != null) {
        lines.add(line);
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

    return lines;
  }

  private static HashMap<Character, Integer> makeMap(String line) {
    HashMap<Character, Integer> chars = new HashMap<>();
    for (Character c: line.toCharArray()) {
      if (chars.containsKey(c)) {
        chars.put(c, chars.get(c) + 1);
      } else {
        chars.put(c, 1);
      }
    }
    return chars;
  }

  private static int part1(List<String> lines) {
    int two = 0;
    int three = 0;
    for (String line: lines) {
      HashMap<Character, Integer> chars = makeMap(line);
      if (chars.containsValue(2)) {
        two++;
      }
      if (chars.containsValue(3)) {
        three++;
      }
    }
    return two * three;
  }

  private static String part2(List<String> lines) {
    for (int i = 0; i < lines.size(); i++) {
      char[] line1 = lines.get(i).toCharArray();
      for (int j = i + 1; j < lines.size(); j++) {
        char[] line2 = lines.get(j).toCharArray();
        int diff = 0;
        for (int k = 0; k < line1.length; k++) {
          if (line1[k] != line2[k]) {
            diff++;
            if (diff > 1) {
              break;
            }
          }
        }
        if (diff == 1) {
          StringBuilder res = new StringBuilder(line1.length - 1);
          for (int k = 0; k < line1.length; k++) {
            if (line1[k] == line2[k]) {
              res.append(line1[k]);
            }
          }
          return res.toString();
        }
      }
    }
    return "";
  }
  public static void main(String[] args) {
    List<String> lines = parseInput();
    System.out.println(part1(lines));
    System.out.println(part2(lines));

  }
}
