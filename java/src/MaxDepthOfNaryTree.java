import java.util.List;

public class MaxDepthOfNaryTree {
    public static int maxDepth(Node root) {
        if (root == null)
            return 0;

        int max = 0;
        for (Node child : root.children)
            max = Math.max(max, maxDepth(child));
        return max + 1;
    }

    private static class Node {
        int val;
        List<Node> children;

        Node(int val, List<Node> children) {
            this.val = val;
            this.children = children;
        }
    }
}
