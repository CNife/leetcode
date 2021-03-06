import leetcode.TreeNode;
import test.Tester;

import java.util.ArrayList;
import java.util.List;
import java.util.Objects;

public class LeafSimilarTree {
    public static boolean leafSimilar(TreeNode root1, TreeNode root2) {
        if (Objects.equals(root1, root2)) {
            return true;
        }
        List<Integer> list1 = new ArrayList<>();
        walkTree(root1, list1);
        List<Integer> list2 = new ArrayList<>();
        walkTree(root2, list2);
        return list1.equals(list2);
    }

    private static void walkTree(TreeNode root, List<Integer> list) {
        if (root.left == null && root.right == null) {
            list.add(root.val);
            return;
        }
        if (root.left != null) {
            walkTree(root.left, list);
        }
        if (root.right != null) {
            walkTree(root.right, list);
        }
    }

    public static void main(String[] args) {
        Tester tester = new Tester(LeafSimilarTree.class);
        tester.addTestCase(new TreeNode(1), new TreeNode(2), false);
        tester.runTestCases();
    }
}
