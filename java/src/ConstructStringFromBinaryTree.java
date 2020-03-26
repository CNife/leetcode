import leetcode.TreeNode;
import test.Tester;

public class ConstructStringFromBinaryTree {
    public static String tree2str(TreeNode t) {
        if (t == null) {
            return "";
        }

        StringBuilder sb = new StringBuilder();
        sb.append(t.val);
        if (t.left != null || t.right != null) {
            sb.append('(');
            sb.append(tree2str(t.left));
            sb.append(')');
        }
        if (t.right != null) {
            sb.append('(');
            sb.append(tree2str(t.right));
            sb.append(')');
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        Tester tester = new Tester(ConstructStringFromBinaryTree.class);
        tester.addTestCase(
                TreeNode.newTree(1, 2, 3, 4),
                "1(2(4))(3)"
        );
        tester.addTestCase(
                TreeNode.newTree(1, 2, 3, null, 4),
                "1(2()(4))(3)"
        );
        tester.runTestCases();
    }
}
