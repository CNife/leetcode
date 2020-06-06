import test.Tester;

import java.util.HashMap;

public class LongestConsecutiveSequence {
    public static int longestConsecutive(int[] nums) {
        HashMap<Integer, Integer> ds = new HashMap<>(nums.length);
        for (int num : nums) {
            ds.put(num, num + 1);
        }

        int result = 0;
        for (int num : nums) {
            int len = find(num, ds);
            result = Math.max(result, len - num);
        }
        return result;
    }

    private static int find(int x, HashMap<Integer, Integer> ds) {
        Integer p = ds.get(x);
        if (p == null) {
            return x;
        } else {
            int q = find(p, ds);
            ds.put(x, q);
            return q;
        }
    }

    public static void main(String[] args) {
        Tester tester = new Tester(LongestConsecutiveSequence.class);
        tester.addTestCase(new int[]{100, 4, 200, 1, 3, 2}, 4);
        tester.runTestCases();
    }
}
