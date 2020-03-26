import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.Collections;
import java.util.List;

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

        System.out.println("inorder " + Arrays.toString(Arrays.copyOfRange(inorder, inorderStart, inorderEnd)));
        System.out.println("postorder " + Arrays.toString(Arrays.copyOfRange(postorder, postorderStart, postorderEnd)));

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

    @Test
    void test() {
        List<TestCase> testCases = Collections.singletonList(
                new TestCase(new int[]{9, 3, 15, 20, 7}, new int[]{9, 15, 7, 20, 3},
                        TreeNode.newTree(3, 9, 20, null, null, 15, 7))
        );
        for (TestCase testCase : testCases) {
            TreeNode actual = buildTree(testCase.inorder, testCase.postorder);
            Assertions.assertEquals(testCase.expected, actual);
        }
    }

    private static class TestCase {
        int[] inorder;
        int[] postorder;
        TreeNode expected;

        public TestCase(int[] inorder, int[] postorder, TreeNode expected) {
            this.inorder = inorder;
            this.postorder = postorder;
            this.expected = expected;
        }
    }
}
