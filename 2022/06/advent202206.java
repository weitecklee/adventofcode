import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.HashMap;

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

  private static int findMarker(String signal, int distinctChars) {
    HashMap<Character, Integer> letters = new HashMap<>();
    int n = distinctChars;
    int i = 0;
    while (i < n) {
      Character c = signal.charAt(i);
      if (letters.containsKey(c)) {
        n = Math.max(n, letters.get(c) + distinctChars + 1);
      }
      letters.put(c, i);
      i++;
    }
    return i;
  }

  public static void main(String[] args) {
    String signal = parseInput();
    System.out.println(findMarker(signal, 4));
    System.out.println(findMarker(signal, 14));

  }
}
