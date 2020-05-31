import leetcode.TreeNode;
import test.Tester;

public class SymmetricTree {
    public static boolean isSymmetric(TreeNode root) {
        return fuck(root, root);
    }

    private static boolean fuck(TreeNode left, TreeNode right) {
        if (left == null)
            return right == null;
        if (right == null)
            return false;
        return left.val == right.val &&
                fuck(left.left, right.right) &&
                fuck(left.right, right.left);
    }

    public static void main(String[] args) {
        Tester tester = new Tester(SymmetricTree.class);
        tester.addTestCase(null, true);
        tester.addTestCase(
                TreeNode.newTree(1, 2, 2, 3, 4, 4, 3),
                true
        );
        tester.addTestCase(
                TreeNode.newTree(1, 2, 2, null, 3, null, 3),
                false
        );
        tester.runTestCases();
    }
}
