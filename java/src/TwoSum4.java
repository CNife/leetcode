import leetcode.TreeNode;
import test.Tester;

import java.util.*;

public class TwoSum4 {
    public static boolean findTarget(TreeNode root, int k) {
        List<Integer> nums = bst2List(root);
        for (int i = 0; i < nums.size(); i++) {
            int num = nums.get(i);
            if (num * 2 <= k) {
                int found = Collections.binarySearch(nums, k - num);
                if (found > 0 && found != i) {
                    return true;
                }
            } else {
                break;
            }
        }
        return false;
    }

    private static List<Integer> bst2List(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        Deque<TreeNode> stack = new ArrayDeque<>();
        TreeNode curr = root;
        while (curr != null || !stack.isEmpty()) {
            while (curr != null) {
                stack.addLast(curr);
                curr = curr.left;
            }
            if (!stack.isEmpty()) {
                TreeNode top = stack.removeLast();
                result.add(top.val);
                curr = top.right;
            }
        }
        return result;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(TwoSum4.class);
        tester.addTestCase(
                TreeNode.newTree(5, 3, 6, 2, 4, null, 7), 9,
                true
        );
        tester.addTestCase(
                TreeNode.newTree(5, 3, 6, 2, 4, null, 7), 28,
                false
        );
        tester.addTestCase(
                TreeNode.newTree(2, 0, 3, -4, 1), -1,
                true
        );
        tester.addTestCase(
                TreeNode.newTree(2, null, 3), 6,
                false
        );
        tester.runTestCases();
    }
}
