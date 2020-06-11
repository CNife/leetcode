import test.Tester;

import java.util.ArrayDeque;
import java.util.Deque;

public class DailyTemperatures {
    public static int[] dailyTemperatures(int[] temps) {
        if (temps == null || temps.length == 0) {
            return new int[]{};
        }

        Deque<Integer> stack = new ArrayDeque<>();
        int[] result = new int[temps.length];
        for (int i = 0; i < temps.length; i++) {
            while (!stack.isEmpty() && temps[stack.getLast()] < temps[i]) {
                int top = stack.removeLast();
                result[top] = i - top;
            }
            stack.addLast(i);
        }
        return result;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(DailyTemperatures.class);
        tester.addTestCase(
                new int[]{73, 74, 75, 71, 69, 72, 76, 73},
                new int[]{1, 1, 4, 2, 1, 1, 0, 0}
        );
        tester.runTestCases();
    }
}
