import leetcode.TreeNode;
import test.Tester;

import java.util.ArrayList;

public class MinDistanceBetweenBstNodes {
    public static int minDiffInBST(TreeNode root) {
        ArrayList<Integer> values = inorder(root);
        if (values.size() <= 1) {
            return 0;
        }

        int result = Integer.MAX_VALUE;
        for (int i = 0; i < values.size() - 1; i++) {
            int diff = Math.abs(values.get(i) - values.get(i + 1));
            result = Math.min(result, diff);
        }
        return result;
    }

    private static ArrayList<Integer> inorder(TreeNode root) {
        ArrayList<Integer> result = new ArrayList<>();
        inorderRecursive(root, result);
        return result;
    }

    private static void inorderRecursive(TreeNode node, ArrayList<Integer> values) {
        if (node != null) {
            inorderRecursive(node.left, values);
            values.add(node.val);
            inorderRecursive(node.right, values);
        }
    }

    public static void main(String[] args) {
        Tester tester = new Tester(MinDistanceBetweenBstNodes.class);
        tester.addTestCase(null, 0);
        tester.addTestCase(new TreeNode(1), 0);
        tester.addTestCase(TreeNode.newTree(4, 2, 6, 1, 3), 1);
        tester.runTestCases();
    }
}
