import leetcode.TreeNode;
import test.Tester;

import java.util.ArrayList;
import java.util.List;

public class MinimumAbsoluteDifferenceInBST {
    public static int getMinimumDifference(TreeNode root) {
        List<TreeNode> stack = new ArrayList<>();
        TreeNode node = root;
        Integer prevValue = null;
        int result = Integer.MAX_VALUE;
        while (node != null || !stack.isEmpty()) {
            while (node != null) {
                stack.add(node);
                node = node.left;
            }
            node = stack.get(stack.size() - 1);
            stack.remove(stack.size() - 1);

            if (prevValue != null) {
                result = Math.min(result, node.val - prevValue);
            }
            prevValue = node.val;

            node = node.right;
        }
        return result;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(MinimumAbsoluteDifferenceInBST.class);
        tester.addTestCase(TreeNode.newTree(1, null, 2, null, 3), 1);
        tester.addTestCase(TreeNode.newTree(5, 4, 7), 1);
        tester.runTestCases();
    }
}
