import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

public class MinimumAbsoluteDifferenceInBST {
    static int getMinimumDifference(TreeNode root) {
        List<TreeNode> stack = new ArrayList<>();
        TreeNode node = root;
        Integer prevValue = null;
        int result = Integer.MAX_VALUE;
        while (node != null || !stack.isEmpty()) {
            while (node != null) {
                stack.add(node);
                node = node.left;
            }
            node = stack.get(stack.size() - 1);
            stack.remove(stack.size() - 1);

            if (prevValue != null) {
                result = Math.min(result, node.val - prevValue);
            }
            prevValue = node.val;

            node = node.right;
        }
        return result;
    }

    @Test
    void test() {
        TreeNode root = new TreeNode(1);
        root.right = new TreeNode(2);
        root.right.right = new TreeNode(3);
        Assertions.assertEquals(1, getMinimumDifference(root));

        root = new TreeNode(5);
        root.left = new TreeNode(4);
        root.right = new TreeNode(7);
        Assertions.assertEquals(1, getMinimumDifference(root));
    }
}
