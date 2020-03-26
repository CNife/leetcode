import leetcode.TreeNode;
import test.Tester;

public class DeepestLeavesSum {
    public static int deepestLeavesSum(TreeNode root) {
        return fuck(root).sum;
    }

    private static DepthInfo fuck(TreeNode root) {
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

    public static void main(String[] args) {
        Tester tester = new Tester(DeepestLeavesSum.class);
        tester.addTestCase(
                TreeNode.newTree(1, 2, 3, 4, 5, null, 6, 7, null, null, null, null, 8),
                15
        );
        tester.runTestCases();
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