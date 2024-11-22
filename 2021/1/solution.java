import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

class solution {

  private static List<Integer> parseInput() {
    String fileName = "input.txt";
    List<Integer> readings = new ArrayList<>();

    try (BufferedReader br = new BufferedReader(new FileReader(fileName))) {
      String line;
      while ((line = br.readLine()) != null) {
        readings.add(Integer.parseInt(line));
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

    return readings;
  }

  private static Integer part1(List<Integer> readings) {
    int res = 0;
    for (int i = 1; i < readings.size(); i++) {
      if (readings.get(i) > readings.get(i - 1)) {
        res++;
      }
    }
    return res;
  }

  private static Integer part2(List<Integer> readings) {
    int res = 0;
    int slidingSum = readings.get(0) + readings.get(1) + readings.get(2);

    for (int i = 3; i < readings.size(); i++) {
      int newSlidingSum = slidingSum + readings.get(i) - readings.get(i - 3);
      if (newSlidingSum > slidingSum) {
        res++;
      }
      slidingSum = newSlidingSum;
    }

    return res;
  }

  public static void main(String[] args) {
    List<Integer> readings = parseInput();

    System.out.println(part1(readings));
    System.out.println(part2(readings));
  }
}
