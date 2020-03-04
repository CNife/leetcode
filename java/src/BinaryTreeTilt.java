import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

public class BinaryTreeTilt {
    static int findTilt(TreeNode root) {
        if (root == null) {
            return 0;
        }
        return findTiltRecursively(root);
    }

    private static int findTiltRecursively(TreeNode root) {
        int leftTilt = 0, leftSum = 0;
        if (root.left != null) {
            leftTilt = findTiltRecursively(root.left);
            leftSum = root.left.val;
        }

        int rightTilt = 0, rightSum = 0;
        if (root.right != null) {
            rightTilt = findTiltRecursively(root.right);
            rightSum = root.right.val;
        }

        root.val += leftSum + rightSum;
        return Math.abs(leftSum - rightSum) + leftTilt + rightTilt;
    }

    @Test
    void test() {
        List<TestCase<Integer, TreeNode>> testCases = Arrays.asList(
                new TestCase<>(0, null),
                new TestCase<>(1, TreeNode.newTree(1, 2, 3))
        );
        for (TestCase<Integer, TreeNode> testCase : testCases) {
            Assertions.assertEquals(testCase.expected, findTilt(testCase.source));
        }
    }
}
