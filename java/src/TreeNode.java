import java.util.ArrayDeque;
import java.util.Deque;
import java.util.Objects;

class TreeNode {
    int val;
    TreeNode left = null;
    TreeNode right = null;

    TreeNode(int val) {
        this.val = val;
    }

    static TreeNode newTree(Integer... nums) {
        int len = nums.length;
        if (len == 0) {
            return null;
        }

        TreeNode root = new TreeNode(nums[0]);
        Deque<TreeNode> queue = new ArrayDeque<>();
        queue.addLast(root);

        for (int i = 1; i < len; ) {
            TreeNode node = queue.removeFirst();
            Integer leftVal = nums[i++];
            if (leftVal != null) {
                node.left = new TreeNode(leftVal);
                queue.addLast(node.left);
            }
            if (i < len) {
                Integer rightVal = nums[i++];
                if (rightVal != null) {
                    node.right = new TreeNode(rightVal);
                    queue.addLast(node.right);
                }
            }
        }
        return root;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        TreeNode treeNode = (TreeNode) o;
        return val == treeNode.val &&
                Objects.equals(left, treeNode.left) &&
                Objects.equals(right, treeNode.right);
    }

    @Override
    public int hashCode() {
        return Objects.hash(val, left, right);
    }

    @Override
    public String toString() {
        return "TreeNode(" + val + "){" + left + "," + right + "}";
    }
}
