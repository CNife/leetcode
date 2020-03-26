import leetcode.TreeNode;
import test.Tester;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class FindModeInBinarySearchTree {
    public static int[] findMode(TreeNode root) {
        if (root == null) {
            return new int[]{};
        }

        List<TreeNode> stack = new ArrayList<>();
        TreeNode top = root;

        Map<Integer, Integer> map = new HashMap<>();

        while (!stack.isEmpty() || top != null) {
            while (top != null) {
                stack.add(top);
                top = top.left;
            }
            if (!stack.isEmpty()) {
                top = stack.get(stack.size() - 1);
                map.put(top.val, map.getOrDefault(top.val, 0) + 1);
                stack.remove(stack.size() - 1);
                top = top.right;
            }
        }

        int maxCount = 0;
        List<Integer> result = new ArrayList<>();
        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            if (maxCount < entry.getValue()) {
                result.clear();
                result.add(entry.getKey());
                maxCount = entry.getValue();
            } else if (maxCount == entry.getValue()) {
                result.add(entry.getKey());
            }
        }
        return result.stream().mapToInt(i -> i).toArray();
    }

    public static void main(String[] args) {
        Tester tester = new Tester(FindModeInBinarySearchTree.class);
        tester.addTestCase(
                TreeNode.newTree(1, null, 2, null, 2),
                new int[]{2}
        );
        tester.runTestCases();
    }
}
