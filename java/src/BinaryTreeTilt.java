import leetcode.TreeNode;
import test.Tester;

public class BinaryTreeTilt {
    public static int findTilt(TreeNode root) {
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

    public static void main(String[] args) {
        Tester tester = new Tester(BinaryTreeTilt.class);
        tester.addTestCase(null, 0);
        tester.addTestCase(TreeNode.newTree(1, 2, 3), 1);
        tester.runTestCases();
    }
}
