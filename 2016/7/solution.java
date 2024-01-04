import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;
import java.util.regex.*;

class Solver {
  private static String patternString = "\\[(.*?)\\]";
  private static Pattern pattern = Pattern.compile(patternString);
  private static Pattern patternABBA = Pattern.compile("(\\w)(\\w)\\2\\1");

  private static class Address {
    private String address;
    private String hypernet;
    private String supernet;

    public Address(String address) {
      this.address = address;
      Matcher matcher = pattern.matcher(address);
      List<String> hypernets = new ArrayList<>();
      while (matcher.find()) {
        hypernets.add(matcher.group(1));
      }
      this.hypernet = String.join(" ", hypernets);
      this.supernet = address.replaceAll(patternString, " ");
    }

    public boolean checkPart1() {
      return checkABBA(supernet) && !checkABBA(hypernet);
    }

    public boolean checkPart2() {
      for (int i = 0; i < supernet.length() - 2; i++) {
        if (supernet.charAt(i) == supernet.charAt(i + 2) && supernet.charAt(i) != supernet.charAt(i + 1)) {
          StringBuilder sb = new StringBuilder();
          sb.append(supernet.charAt(i + 1))
            .append(supernet.charAt(i))
            .append(supernet.charAt(i + 1));
          Pattern patternBAB = Pattern.compile(sb.toString());
          Matcher matcher2 = patternBAB.matcher(hypernet);
          if (matcher2.find()) {
            return true;
          }
        }
      }
      return false;
    }
  }

  private static boolean checkABBA(String str) {
    Matcher matcher = patternABBA.matcher(str);
    while (matcher.find()) {
      String match = matcher.group();
      if (match.charAt(0) != match.charAt(1)) {
        return true;
      }
    }
    return false;
  }

  private static List<Address> parseInput() {
    String fileName = "input.txt";
    List<Address> puzzleInput = new ArrayList<>();

    try {
        List<String> lines = Files.readAllLines(Paths.get(fileName));
        for (String line : lines) {
            puzzleInput.add(new Address(line));
        }
    } catch (IOException e) {
        e.printStackTrace();
    }

    return puzzleInput;

  }

  public static void main(String[] args) {
    List<Address> addresses = parseInput();
    int part1 = 0;
    int part2 = 0;
    for (Address address: addresses) {
      if (address.checkPart1()) {
        part1++;
      }
      if (address.checkPart2()) {
        part2++;
      }
    }
    System.out.println(part1);
    System.out.println(part2);
  }
}
