import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import java.util.Set;

class Solver {

  private static List<String> parseInput() {
    String fileName = "input.txt";
    List<String> messages = new ArrayList<>();

    try (BufferedReader br = new BufferedReader(new FileReader(fileName))) {
      String line;
      while ((line = br.readLine()) != null) {
        messages.add(line);
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

    return messages;

  }

  public static void main(String[] args) {
    List<String> messages = parseInput();
    Map<Integer, Map<Character, Integer>> charmap = new HashMap<>();

    for (String message: messages) {
      int pos = 0;
      for (char c: message.toCharArray()) {
        Map<Character, Integer> posmap = charmap.getOrDefault(pos, new HashMap<>());
        posmap.put(c, posmap.getOrDefault(c, 0) + 1);
        charmap.put(pos, posmap);
        pos++;
      }
    }

    StringBuilder part1 = new StringBuilder();
    StringBuilder part2 = new StringBuilder();

    for (Map<Character, Integer> posmap: charmap.values()) {
      Set<Map.Entry<Character, Integer>> entrySet = posmap.entrySet();
      Iterator<Map.Entry<Character, Integer>> iterator = entrySet.iterator();

      Map.Entry<Character, Integer> first = iterator.next();
      char mostChar = first.getKey();
      char leastChar = first.getKey();
      int mostVal = first.getValue();
      int leastVal = first.getValue();

      while (iterator.hasNext()) {
          Map.Entry<Character, Integer> entry = iterator.next();
          if (entry.getValue() > mostVal) {
            mostChar = entry.getKey();
            mostVal = entry.getValue();
          } else if (entry.getValue() < leastVal) {
            leastChar = entry.getKey();
            leastVal = entry.getValue();
          }
      }

      part1.append(mostChar);
      part2.append(leastChar);
    }

    System.out.println(part1.toString());
    System.out.println(part2.toString());
  }
}
