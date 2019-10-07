import java.util.List;

public class MaxDepthOfNaryTree {
    private static class Node {
        int val;
        List<Node> children;

        Node(int val, List<Node> children) {
            this.val = val;
            this.children = children;
        }
    }

    private int maxDepth(Node root) {
        if (root == null)
            return 0;

        int max = 0;
        for (Node child : root.children)
            max = Math.max(max, maxDepth(child));
        return max + 1;
    }
}
