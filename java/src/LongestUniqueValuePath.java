import leetcode.TreeNode;
import test.Tester;

public class LongestUniqueValuePath {
    public static int longestUnivaluePath(TreeNode root) {
        Int result = new Int();
        int _i = arrowLength(root, result);
        return result.value;
    }

    private static int arrowLength(TreeNode root, Int result) {
        if (root == null) {
            return 0;
        }

        int left = arrowLength(root.left, result);
        int right = arrowLength(root.right, result);
        int leftPath = 0;
        if (root.left != null && root.val == root.left.val) {
            leftPath += left + 1;
        }
        int rightPath = 0;
        if (root.right != null && root.val == root.right.val) {
            rightPath += right + 1;
        }
        result.value = Math.max(result.value, leftPath + rightPath);
        return Math.max(leftPath, rightPath);
    }

    public static void main(String[] args) {
        Tester tester = new Tester(LongestUniqueValuePath.class);
        tester.addTestCase(
                TreeNode.newTree(5, 4, 5, 1, 1, null, 5),
                2
        );
        tester.addTestCase(
                TreeNode.newTree(1, 4, 5, 4, 4, null, 5),
                2
        );
        tester.addTestCase(
                TreeNode.newTree(),
                0
        );
        tester.addTestCase(
                TreeNode.newTree(1),
                0
        );
        tester.runTestCases();
    }

    private static class Int {
        int value = 0;
    }
}
