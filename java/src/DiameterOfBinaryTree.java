import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

public class DiameterOfBinaryTree {
    static int diameterOfBinaryTree(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int leftDiameter = root.left == null ? 0 : diameterOfBinaryTree(root.left);
        int rightDiameter = root.right == null ? 0 : diameterOfBinaryTree(root.right);

        int leftDepth = root.left == null ? -1 : root.left.val;
        int rightDepth = root.right == null ? -1 : root.right.val;
        root.val = Math.max(leftDepth, rightDepth) + 1;

        return Math.max(Math.max(leftDiameter, rightDiameter), leftDepth + rightDepth + 2);
    }

    @Test
    void test() {
        List<TestCase<Integer, TreeNode>> testCases = Arrays.asList(
                new TestCase<>(3, TreeNode.newTree(1, 2, 3, 4, 5)),
                new TestCase<>(0, TreeNode.newTree(1)),
                new TestCase<>(0, null)
        );
        for (TestCase<Integer, TreeNode> testCase : testCases) {
            int actual = diameterOfBinaryTree(testCase.source);
            Assertions.assertEquals(testCase.expected, actual);
        }
    }
}
