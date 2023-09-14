import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

class Solver {

  private static HashMap<Character, Integer> priority = new HashMap<>();

  private static List<String> parseInput() {
    String fileName = "input.txt";
    List<String> rucksacks = new ArrayList<>();

    try (BufferedReader br = new BufferedReader(new FileReader(fileName))) {
      String line;
      while ((line = br.readLine()) != null) {
        rucksacks.add(line);
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

    return rucksacks;
  }

  private static Set<Character> makeCharSet(String str) {
    Set<Character> charSet = new HashSet<>();
    char[] charArray = str.toCharArray();
    for (char c: charArray) {
      charSet.add(c);
    }
    return charSet;
  }

  private static int part1(List<String> rucksacks) {
    int res = 0;

    for (String rucksack: rucksacks) {
      Set<Character> set1 = new HashSet<>();
      Set<Character> set2 = new HashSet<>();
      int len = rucksack.length() / 2;
      for (int i = 0; i < len; i++) {
        set1.add(rucksack.charAt(i));
        set2.add(rucksack.charAt(i + len));
      }
      set1.retainAll(set2);
      for (char c: set1) {
        res += priority.get(c);
      }
    }

    return res;
  }

  private static int part2(List<String> rucksacks) {
    int res = 0;

    for (int i = 0; i < rucksacks.size(); i += 3) {
      Set<Character> set1 = makeCharSet(rucksacks.get(i));
      Set<Character> set2 = makeCharSet(rucksacks.get(i + 1));
      Set<Character> set3 = makeCharSet(rucksacks.get(i + 2));
      set1.retainAll(set2);
      set1.retainAll(set3);
      for (char c: set1) {
        res += priority.get(c);
      }
    }

    return res;
  }
  public static void main(String[] args) {
    List<String> rucksacks = parseInput();
    String prio = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

    for (int i = 0; i < prio.length(); i++) {
      char c = prio.charAt(i);
      priority.put(c, i + 1);
    }

    System.out.println(part1(rucksacks));
    System.out.println(part2(rucksacks));

  }
}
