import leetcode.TreeNode;
import test.Tester;

import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Deque;
import java.util.List;

public class SumRoot2LeafNumbers {
    public static int sumNumbers(TreeNode root) {
        if (root == null) {
            return 0;
        }

        Deque<Integer> stack = new ArrayDeque<>();
        List<Integer> results = new ArrayList<>();
        backtrack(root, stack, results);
        return results.stream()
                .mapToInt(Integer::intValue)
                .sum();
    }

    private static void backtrack(TreeNode node, Deque<Integer> stack, List<Integer> results) {
        stack.addLast(node.val);
        if (node.left == null && node.right == null) {
            results.add(stack.stream().reduce(0, (acc, digit) -> acc * 10 + digit));
        } else {
            if (node.left != null)
                backtrack(node.left, stack, results);
            if (node.right != null)
                backtrack(node.right, stack, results);
        }
        stack.removeLast();
    }

    public static void main(String[] args) {
        Tester tester = new Tester(SumRoot2LeafNumbers.class);
        tester.addTestCase(
                TreeNode.newTree(1, 2, 3),
                25
        );
        tester.addTestCase(
                TreeNode.newTree(4, 9, 0, 5, 1),
                1026
        );
        tester.runTestCases();
    }
}
