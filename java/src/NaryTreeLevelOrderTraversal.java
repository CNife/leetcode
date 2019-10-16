import java.util.ArrayList;
import java.util.List;

public class NaryTreeLevelOrderTraversal {
    private static class Node {
        int val;
        List<Node> children;
    }

    public List<List<Integer>> levelOrder(Node root) {
        if (root == null)
            return new ArrayList<>();

        List<List<Integer>> res = new ArrayList<>();
        List<Node> levelNode = new ArrayList<>();
        levelNode.add(root);
        while (!levelNode.isEmpty()) {
            List<Integer> levelInt = new ArrayList<>();
            List<Node> nextLevelNode = new ArrayList<>();
            for (Node node : levelNode) {
                if (node != null) {
                    levelInt.add(node.val);
                    nextLevelNode.addAll(node.children);
                }
            }
            res.add(levelInt);
            levelNode = nextLevelNode;
        }
        return res;
    }
}
