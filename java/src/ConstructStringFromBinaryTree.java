import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

public class ConstructStringFromBinaryTree {
    static String tree2str(TreeNode t) {
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

    @Test
    void test() {
        List<TestCase<String, TreeNode>> testCases = Arrays.asList(
                new TestCase<>("1(2(4))(3)", TreeNode.newTree(1, 2, 3, 4)),
                new TestCase<>("1(2()(4))(3)", TreeNode.newTree(1, 2, 3, null, 4))
        );
        for (TestCase<String, TreeNode> testCase : testCases) {
            Assertions.assertEquals(testCase.expected, tree2str(testCase.source));
        }
    }
}
