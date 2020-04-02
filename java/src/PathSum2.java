import leetcode.TreeNode;
import test.Tester;

import java.util.*;
import java.util.stream.Collectors;

public class PathSum2 {
    public static List<List<Integer>> pathSum(TreeNode root, int sum) {
        if (root == null) {
            return Collections.emptyList();
        }

        Deque<TreeNode> stack = new ArrayDeque<>();
        List<List<Integer>> result = new ArrayList<>();
        backtrack(root, sum, stack, result);
        return result;
    }

    private static void backtrack(
            TreeNode current,
            int remaining,
            Deque<TreeNode> stack,
            List<List<Integer>> result
    ) {
        stack.addLast(current);
        if (current.left == null && current.right == null) {
            if (current.val == remaining) {
                List<Integer> oneResult = stack.stream()
                        .map(node -> node.val)
                        .collect(Collectors.toList());
                result.add(oneResult);
            }
        } else {
            if (current.left != null) {
                backtrack(current.left, remaining - current.val, stack, result);
            }
            if (current.right != null) {
                backtrack(current.right, remaining - current.val, stack, result);
            }
        }
        stack.removeLast();
    }

    public static void main(String[] args) {
        @SuppressWarnings("unchecked")
        Tester tester = new Tester(PathSum2.class,
                result -> new HashSet<>((List<List<Integer>>) result));
        tester.addTestCase(
                TreeNode.newTree(5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1),
                22,
                Arrays.asList(
                        Arrays.asList(5, 4, 11, 2),
                        Arrays.asList(5, 8, 4, 5)
                )
        );
        tester.addTestCase(
                TreeNode.newTree(),
                1,
                Collections.emptyList()
        );
        tester.runTestCases();
    }
}
