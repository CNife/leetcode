import leetcode.TreeNode;
import test.Tester;

import java.util.*;

public class BinaryTreeRightSideView {
    public static List<Integer> rightSideView(TreeNode root) {
        if (root == null)
            return Collections.emptyList();

        List<Integer> result = new ArrayList<>();
        Deque<TreeNode> queue = new ArrayDeque<>();
        queue.addLast(root);
        while (!queue.isEmpty()) {
            result.add(queue.getLast().val);
            int count = queue.size();
            for (int i = 0; i < count; i++) {
                TreeNode node = queue.removeFirst();
                if (node.left != null)
                    queue.addLast(node.left);
                if (node.right != null)
                    queue.addLast(node.right);
            }
        }
        return result;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(BinaryTreeRightSideView.class);
        tester.addTestCase(
                TreeNode.newTree(1, 2, 3, null, 5, null, 4),
                Arrays.asList(1, 3, 4)
        );
        tester.runTestCases();
    }
}
