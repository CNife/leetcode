import leetcode.TreeNode;
import test.Tester;

public class DiameterOfBinaryTree {
    public static int diameterOfBinaryTree(TreeNode root) {
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

    public static void main(String[] args) {
        Tester tester = new Tester(DiameterOfBinaryTree.class);
        tester.addTestCase(TreeNode.newTree(1, 2, 3, 4, 5), 3);
        tester.addTestCase(TreeNode.newTree(1), 0);
        tester.addTestCase(null, 0);
        tester.runTestCases();
    }
}
