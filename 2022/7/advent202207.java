import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

class Solver {

  private static Directory root;
  private static int limit = 100000;
  private static int totalSpace = 70000000;
  private static int unusedSpace = 30000000;
  private static int spaceToDelete;
  private static int spaceNeeded;

  public static class Directory {
    private Directory parent;
    private HashMap<String, Directory> subDirectories;
    private int size;
    private int totalSize;

    public Directory(Directory parent) {
      this.parent = parent;
      this.subDirectories = new HashMap<>();
      this.size = 0;
      this.totalSize = 0;
    }

    public void addSubDirectory(String name) {
      Directory subDirectory = new Directory(this);
      subDirectories.put(name, subDirectory);
    }

    public void addFile(int size) {
      this.size += size;
    }

    public int getSize() {
      return size;
    }

    public Directory getParent() {
      return parent;
    }

    public Directory getSubDirectory(String name) {
      return subDirectories.get(name);
    }

    public int getTotalSize() {
      if (totalSize != 0) {
        return totalSize;
      }
      int total = size;
      for (Directory subDir: subDirectories.values()) {
          total += subDir.getTotalSize();
      }
      totalSize = total;
      return totalSize;
    }

  }

  private static void parseInput() {
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

    Pattern pattern = Pattern.compile("\\$ cd (.+)");

    root = new Directory(null);
    root.addSubDirectory("/");
    Directory curr = root;
    for (String line: lines) {
      Matcher matcher = pattern.matcher(line);
      if (matcher.find()) {
        String path = matcher.group(1);
        if (path.equals("..")) {
          curr = curr.getParent();
        } else {
          curr = curr.getSubDirectory(path);
        }
      } else {
        if (line.charAt(0) == '$') {
          continue;
        }
        String[] parts = line.split(" ");
        try {
          int size = Integer.parseInt(parts[0]);
          curr.addFile(size);
        } catch (NumberFormatException e) {
          curr.addSubDirectory(parts[1]);
        }
      }
    }
  }

  private static int part1(Directory dir) {
    int sum = 0;
    for (Directory subdir: dir.subDirectories.values()) {
      int size = subdir.getTotalSize();
      if (size <= limit) {
        sum += size;
      }
      sum += part1(subdir);
    }
    return sum;
  }

  private static int part2() {
    spaceNeeded = unusedSpace - totalSpace + root.getTotalSize();
    spaceToDelete = totalSpace;
    recur2(root);
    return spaceToDelete;
  }

  private static void recur2(Directory dir) {
    for (Directory subdir: dir.subDirectories.values()) {
      int size = subdir.getTotalSize();
      if (size >= spaceNeeded && size <= spaceToDelete) {
        spaceToDelete = size;
      }
      recur2(subdir);
    }
  }

  public static void main(String[] args) {
    parseInput();
    System.out.println(part1(root));
    System.out.println(part2());

  }
}
