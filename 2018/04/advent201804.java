import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

class Solver {

  private static HashMap<Integer, Guard> guards = new HashMap<>();
  private static Pattern pattern = Pattern.compile("\\[(.+)\\] (.+)");
  private static Pattern pattern2 = Pattern.compile("\\d+");
  private static Pattern pattern3 = Pattern.compile("\\d+:(\\d+)");

  private static class Guard {
    private int id;
    private List<String[]> records = new ArrayList<>();
    private int totalSleep = 0;
    private int mostAsleep = 0;
    private HashMap<Integer, Integer> minutes = new HashMap<>();

    public Guard(int id) {
      this.id = id;
    }

    public void addRecord(String[] record) {
      records.add(record);
    }

    public void analyze() {
      int sleepTime = 0;
      for (String[] record: records) {
        Matcher matcher = pattern3.matcher(record[0]);
        matcher.find();
        int minuteStamp = Integer.parseInt(matcher.group(1));
        if (record[1].equals("wakes up")) {
          totalSleep += minuteStamp - sleepTime;
          for (int i = sleepTime; i < minuteStamp; i++) {
            minutes.compute(i, (k, v) -> v == null ? 1 : v + 1);
          }
        } else if (record[1].equals("falls asleep")) {
          sleepTime = minuteStamp;
        }
      }
      int maxSleep = 0;
      for (Map.Entry<Integer, Integer> entry: minutes.entrySet()) {
        int minute = entry.getKey();
        int sleep = entry.getValue();
        if (sleep > maxSleep) {
          maxSleep = sleep;
          mostAsleep = minute;
        }
      }
    }

    public int getTotalSleep() {
      return totalSleep;
    }

    public int getMostAsleep() {
      return mostAsleep;
    }

    public int getMostAsleepTime() {
      if (!minutes.containsKey(mostAsleep)) {
        return 0;
      }
      return minutes.get(mostAsleep);
    }

    public int getId() {
      return id;
    }

  }

  private static void parseInput() {
    String fileName = "input.txt";
    List<String> logs = new ArrayList<>();

    try (BufferedReader br = new BufferedReader(new FileReader(fileName))) {
      String line;
      while ((line = br.readLine()) != null) {
        logs.add(line);
      }
    } catch (IOException e) {
      e.printStackTrace();
    }

    Collections.sort(logs);
    Guard curr = null;

    for (String log: logs) {
      Matcher matcher = pattern.matcher(log);
      if (matcher.find()) {
        String timestamp = matcher.group(1);
        String action = matcher.group(2);
        Matcher matcher2 = pattern2.matcher(action);
        if (matcher2.find()) {
          int id = Integer.parseInt(matcher2.group());
          if (guards.containsKey(id)) {
            curr = guards.get(id);
          } else {
            curr = new Guard(id);
            guards.put(id, curr);
          }
          action = "begins shift";
        }
        String[] record = {timestamp, action};
        curr.addRecord(record);
      }
    }

  }

  public static void main(String[] args) {
    parseInput();

    for (Guard guard: guards.values()) {
      guard.analyze();
    }

    Optional<Guard> part1 = guards.values().stream().sorted((a, b) -> Integer.compare(b.getTotalSleep(), a.getTotalSleep())).findFirst();
    if (part1.isPresent()) {
      Guard guard1 = part1.get();
      System.out.println(guard1.getId() * guard1.getMostAsleep());
    } else {
      System.out.println("Error with part 1");
    }

    Optional<Guard> part2 = guards.values().stream().sorted((a, b) -> Integer.compare(b.getMostAsleepTime(), a.getMostAsleepTime())).findFirst();
    if (part2.isPresent()) {
      Guard guard2 = part2.get();
      System.out.println(guard2.getId() * guard2.getMostAsleep());
    } else {
      System.out.println("Error with part 2");
    }
  }
}
