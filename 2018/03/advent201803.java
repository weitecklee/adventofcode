import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

class Solver {

  private static HashMap<String, Integer> suit = new HashMap<>();
  private static List<Claim> claims = new ArrayList<>();

  private static String makeCoord(int x, int y) {
    return x + "," + y;
  }

  private static class Claim {
    private int id;
    private int x;
    private int y;
    private int w;
    private int h;

    public Claim(int id, int x, int y, int w, int h) {
      this.id = id;
      this.x = x;
      this.y = y;
      this.w = w;
      this.h = h;
      makeClaim();
    }

    private void makeClaim() {
      for (int i = 0; i < w; i++) {
        for (int j = 0; j < h; j++) {
          String coord = makeCoord(x + i, y + j);
          suit.compute(coord, (k, oldValue) -> oldValue == null ? 1 : oldValue + 1);
        }
      }
    }

    public boolean checkClaim() {
      for (int i = 0; i < w; i++) {
        for (int j = 0; j < h; j++) {
          String coord = makeCoord(x + i, y + j);
          if (suit.get(coord) >= 2) {
            return false;
          }
        }
      }
      return true;
    }

    public int getId() {
      return id;
    }
  }

  private static void parseInput() {
    String fileName = "input.txt";

    Pattern pattern = Pattern.compile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)");

    try (BufferedReader br = new BufferedReader(new FileReader(fileName))) {
      String line;
      while ((line = br.readLine()) != null) {
        Matcher matcher = pattern.matcher(line);
        if (matcher.find()) {
          try {
            int id = Integer.parseInt(matcher.group(1));
            int x = Integer.parseInt(matcher.group(2));
            int y = Integer.parseInt(matcher.group(3));
            int w = Integer.parseInt(matcher.group(4));
            int h = Integer.parseInt(matcher.group(5));
            claims.add(new Claim(id, x, y, w, h));
          } catch (NumberFormatException e) {
            throw e;
          }
        }
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

  }

  private static int part1() {
    return (int) suit.values().stream().filter(n -> n >= 2).count();
  }

  private static int part2() {
    for (Claim claim: claims) {
      if (claim.checkClaim()) {
        return claim.getId();
      }
    }
    return -1;
  }

  public static void main(String[] args) {
    parseInput();
    System.out.println(part1());
    System.out.println(part2());


  }
}
