import leetcode.TreeNode;

import java.util.ArrayList;

public class IncreasingOrderSearchTree {
    public static TreeNode increasingBST(TreeNode root) {
        if (root == null)
            return null;

        ArrayList<Integer> values = toList(root);
        TreeNode newRoot = new TreeNode(values.get(0));
        TreeNode ptr = newRoot;
        for (int i = 1; i < values.size(); i++)
            ptr = ptr.right = new TreeNode(values.get(i));
        return newRoot;
    }

    private static ArrayList<Integer> toList(TreeNode root) {
        ArrayList<Integer> list = new ArrayList<>();
        inOrder(root, list);
        return list;
    }

    private static void inOrder(TreeNode node, ArrayList<Integer> list) {
        if (node.left != null) {
            inOrder(node.left, list);
        }
        list.add(node.val);
        if (node.right != null) {
            inOrder(node.right, list);
        }
    }
}
