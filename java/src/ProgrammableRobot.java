import test.Tester;

import java.util.Arrays;
import java.util.Objects;
import java.util.Set;
import java.util.stream.Collectors;

public class ProgrammableRobot {
    public static boolean robot(String commands, int[][] obstacles, int x, int y) {
        Game game = new Game(obstacles, x, y);
        char[] commandArray = commands.toCharArray();

        while (true) {
            for (char command : commandArray) {
                game.moveRobot(command);
                if (game.failed()) {
                    return false;
                } else if (game.succeeded()) {
                    return true;
                }
            }
        }
    }

    public static void main(String[] args) {
        Tester tester = new Tester(ProgrammableRobot.class);
        tester.addTestCase("URR", new int[][]{}, 3, 2, true);
        tester.addTestCase("URR", new int[][]{new int[]{2, 2}}, 3, 2, false);
        tester.addTestCase("URR", new int[][]{new int[]{4, 2}}, 3, 2, true);
        tester.runTestCases();
    }

    private static class Game {
        Point robot;
        Set<Point> obstacles;
        Point destination;

        Game(int[][] obstacles, int x, int y) {
            this.robot = new Point(0, 0);
            this.destination = new Point(x, y);
            this.obstacles = Arrays.stream(obstacles)
                    .map(obstacle -> new Point(obstacle[0], obstacle[1]))
                    .collect(Collectors.toSet());
        }

        boolean failed() {
            return obstacles.contains(robot) || robot.x > destination.x || robot.y > destination.y;
        }

        boolean succeeded() {
            return robot.equals(destination);
        }

        void moveRobot(char command) {
            switch (command) {
                case 'U':
                    robot.y++;
                    break;
                case 'R':
                    robot.x++;
                    break;
                default:
                    throw new IllegalArgumentException();
            }
        }
    }

    private static class Point {
        int x;
        int y;

        Point(int x, int y) {
            this.x = x;
            this.y = y;
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Point point = (Point) o;
            return x == point.x &&
                    y == point.y;
        }

        @Override
        public int hashCode() {
            return Objects.hash(x, y);
        }
    }
}
