import leetcode.TreeNode;
import test.Tester;

public class ConstructBinaryTreeFromPreorderAndInorderTraversal {
    public static TreeNode buildTree(int[] preorder, int[] inorder) {
        if (preorder == null || inorder == null) {
            return null;
        }
        return buildTree(preorder, 0, preorder.length,
                inorder, 0, inorder.length);
    }

    private static TreeNode buildTree(int[] preorder, int preStart, int preEnd,
                                      int[] inorder, int inStart, int inEnd) {
        if (preStart >= preEnd) {
            return null;
        }

        int rootVal = preorder[preStart];

        int rootIndex = inStart;
        while (rootIndex < inEnd) {
            if (inorder[rootIndex] == rootVal) {
                break;
            } else {
                rootIndex++;
            }
        }

        int leftCount = rootIndex - inStart;
        TreeNode root = new TreeNode(rootVal);
        root.left = buildTree(
                preorder, preStart + 1, preStart + leftCount + 1,
                inorder, inStart, rootIndex
        );
        root.right = buildTree(
                preorder, preStart + leftCount + 1, preEnd,
                inorder, rootIndex + 1, inEnd
        );
        return root;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(
                ConstructBinaryTreeFromPreorderAndInorderTraversal.class
        );
        tester.addTestCase(
                null, null,
                null
        );
        tester.addTestCase(
                new int[]{}, new int[]{},
                null
        );
        tester.addTestCase(
                new int[]{3, 9, 20, 15, 7}, new int[]{9, 3, 15, 20, 7},
                TreeNode.newTree(3, 9, 20, null, null, 15, 7)
        );
        tester.runTestCases();
    }
}
