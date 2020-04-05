import leetcode.TreeNode;
import test.Tester;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class SecondMinimumNodeInBinaryTree {
    public static int findSecondMinimumValue(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        tree2List(root, result);
        Collections.sort(result);

        if (result.size() > 0) {
            int min = result.get(0);
            for (int i = 1; i < result.size(); i++) {
                if (result.get(i) != min)
                    return result.get(i);
            }
        }
        return -1;
    }

    private static void tree2List(TreeNode node, List<Integer> result) {
        if (node != null) {
            result.add(node.val);
            tree2List(node.left, result);
            tree2List(node.right, result);
        }
    }

    public static void main(String[] args) {
        Tester tester = new Tester(SecondMinimumNodeInBinaryTree.class);
        tester.addTestCase(
                TreeNode.newTree(2, 2, 5, null, null, 5, 7),
                5
        );
        tester.addTestCase(
                TreeNode.newTree(2, 2, 2),
                -1
        );
        tester.runTestCases();
    }
}
