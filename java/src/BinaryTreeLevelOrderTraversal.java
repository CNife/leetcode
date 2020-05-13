import leetcode.TreeNode;
import test.Tester;

import java.util.*;

public class BinaryTreeLevelOrderTraversal {
    public static List<List<Integer>> levelOrder(TreeNode root) {
        if (root == null) {
            return Collections.emptyList();
        }

        Deque<TreeNode> queue = new ArrayDeque<>();
        queue.addLast(root);

        List<List<Integer>> result = new ArrayList<>();
        while (!queue.isEmpty()) {
            List<Integer> level = new ArrayList<>();
            int levelCount = queue.size();
            for (int i = 0; i < levelCount; i++) {
                TreeNode node = queue.removeFirst();
                level.add(node.val);
                if (node.left != null) {
                    queue.addLast(node.left);
                }
                if (node.right != null) {
                    queue.addLast(node.right);
                }
            }
            result.add(level);
        }
        return result;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(BinaryTreeLevelOrderTraversal.class);
        tester.addTestCase(
                null,
                Collections.emptyList()
        );
        tester.addTestCase(
                TreeNode.newTree(3, 9, 20, null, null, 15, 7),
                Arrays.asList(
                        Collections.singletonList(3),
                        Arrays.asList(9, 20),
                        Arrays.asList(15, 7)
                )
        );
        tester.runTestCases();
    }
}
