public class DeepestLeavesSum {

    public static void main(String[] args) {
        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.left.left = new TreeNode(4);
        root.left.left.left = new TreeNode(7);
        root.left.right = new TreeNode(5);
        root.right = new TreeNode(3);
        root.right.right = new TreeNode(6);
        root.right.right.right = new TreeNode(8);

        int output = new DeepestLeavesSum().deepestLeavesSum(root);
        System.out.println(output);
        assert output == 15;
    }

    private DepthInfo fuck(TreeNode root) {
        if (root == null) {
            return DepthInfo.EmptyNode;
        } else if (root.left == null && root.right == null) {
            return new DepthInfo(1, root.val);
        }

        DepthInfo left = fuck(root.left);
        DepthInfo right = fuck(root.right);
        if (left.depth == right.depth) {
            return new DepthInfo(left.depth + 1, left.sum + right.sum);
        } else if (left.depth > right.depth) {
            return new DepthInfo(left.depth + 1, left.sum);
        } else {
            return new DepthInfo(right.depth + 1, right.sum);
        }
    }

    public int deepestLeavesSum(TreeNode root) {
        return fuck(root).sum;
    }

    private static class DepthInfo {
        final static DepthInfo EmptyNode = new DepthInfo(0, 0);
        final int depth;
        final int sum;

        DepthInfo(int depth, int sum) {
            this.depth = depth;
            this.sum = sum;
        }
    }
}