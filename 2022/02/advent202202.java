import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

class Solver {

  private static int[] scores = {6, 3, 0}; // win, draw, loss
  private static int[] rpsValues = {1, 2, 3};

  public static class Round {
    private String p1;
    private String p2;

    public Round(String[] players) {
      this.p1 = players[0];
      this.p2 = players[1];
    }

    public int result1() {
      int res = 0;
      switch (p2) {
        case "X":
          res += rpsValues[0];
          switch (p1) {
            case "A":
              res += scores[1];
              break;
            case "B":
              res += scores[2];
              break;
            case "C":
              res += scores[0];
              break;
          }
          break;
        case "Y":
          res += rpsValues[1];
          switch (p1) {
            case "A":
              res += scores[0];
              break;
            case "B":
              res += scores[1];
              break;
            case "C":
              res += scores[2];
              break;
          }
          break;
        case "Z":
          res += rpsValues[2];
          switch (p1) {
            case "A":
              res += scores[2];
              break;
            case "B":
              res += scores[0];
              break;
            case "C":
              res += scores[1];
              break;
          }
          break;
      }
      return res;
    }

    public int result2() {
      int res = 0;
      switch (p2) {
        case "X":
          res += scores[2];
          switch (p1) {
            case "A":
              res += rpsValues[2];
              break;
            case "B":
              res += rpsValues[0];
              break;
            case "C":
              res += rpsValues[1];
              break;
          }
          break;
        case "Y":
          res += scores[1];
          switch (p1) {
            case "A":
              res += rpsValues[0];
              break;
            case "B":
              res += rpsValues[1];
              break;
            case "C":
              res += rpsValues[2];
              break;
          }
          break;
        case "Z":
          res += scores[0];
          switch (p1) {
            case "A":
              res += rpsValues[1];
              break;
            case "B":
              res += rpsValues[2];
              break;
            case "C":
              res += rpsValues[0];
              break;
          }
          break;
      }
      return res;
    }

  }

  private static List<Round> parseInput() {
    String fileName = "input.txt";
    List<Round> rounds = new ArrayList<>();

    try (BufferedReader br = new BufferedReader(new FileReader(fileName))) {
      String line;
      while ((line = br.readLine()) != null) {
        String[] players = line.split(" ");
        rounds.add(new Round(players));
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

    return rounds;
  }

  public static void main(String[] args) {
    List<Round> rounds = parseInput();
    int part1 = 0;
    int part2 = 0;

    for (Round round: rounds) {
      part1 += round.result1();
      part2 += round.result2();
    }

    System.out.println(part1);
    System.out.println(part2);

  }
}
