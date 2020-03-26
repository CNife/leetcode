import leetcode.TreeNode;
import test.Tester;

public class ConstructBinaryTreeFromInorderAndPostorderTraversal {
    public static TreeNode buildTree(int[] inorder, int[] postorder) {
        if (inorder == null || postorder == null) {
            return null;
        }
        return buildTreeRecursively(inorder, 0, inorder.length, postorder, 0, postorder.length);
    }

    private static TreeNode buildTreeRecursively(
            int[] inorder, int inorderStart, int inorderEnd,
            int[] postorder, int postorderStart, int postorderEnd
    ) {
        if (inorderStart >= inorderEnd || postorderStart >= postorderEnd) {
            return null;
        }

        int rootVal = postorder[postorderEnd - 1];
        int rootOffset = 0;
        for (int i = inorderStart; i < inorderEnd; i++) {
            if (inorder[i] == rootVal) {
                rootOffset = i - inorderStart;
                break;
            }
        }

        TreeNode root = new TreeNode(rootVal);
        root.left = buildTreeRecursively(inorder, inorderStart, inorderStart + rootOffset, postorder, postorderStart, postorderStart + rootOffset);
        root.right = buildTreeRecursively(inorder, inorderStart + rootOffset + 1, inorderEnd, postorder, postorderStart + rootOffset, postorderEnd - 1);
        return root;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(ConstructBinaryTreeFromInorderAndPostorderTraversal.class);
        tester.addTestCase(
                new int[]{9, 3, 15, 20, 7},
                new int[]{9, 15, 7, 20, 3},
                TreeNode.newTree(3, 9, 20, null, null, 15, 7)
        );
        tester.runTestCases();
    }
}
