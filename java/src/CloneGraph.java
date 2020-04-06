import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

public class CloneGraph {
    public static Node cloneGraph(Node node) {
        return node == null ? null : cloneNode(node, new HashMap<>());
    }

    private static Node cloneNode(Node node, HashMap<Integer, Node> clonedNodes) {
        Node cloned = new Node(node.val);
        clonedNodes.put(node.val, cloned);
        for (Node neighbor : node.neighbors) {
            Node newNeighbor = clonedNodes.get(neighbor.val);
            if (newNeighbor == null) {
                newNeighbor = cloneNode(neighbor, clonedNodes);
            }
            cloned.neighbors.add(newNeighbor);
        }
        return cloned;
    }

    public static void main(String[] args) {
        List<Node> nodes = Node.createNodes(4);
        Node.link(nodes, 1, 2);
        Node.link(nodes, 2, 3);
        Node.link(nodes, 3, 4);
        Node.link(nodes, 4, 1);
        Node node = cloneGraph(nodes.get(0));
        System.out.println(node);
    }

    private static class Node {
        int val;
        ArrayList<Node> neighbors;

        Node(int val) {
            this.val = val;
            this.neighbors = new ArrayList<>();
        }

        Node(int val, ArrayList<Node> neighbors) {
            this.val = val;
            this.neighbors = neighbors;
        }

        static List<Node> createNodes(int count) {
            List<Node> nodes = new ArrayList<>(count);
            for (int i = 1; i <= count; i++) {
                nodes.add(new Node(i));
            }
            return nodes;
        }

        static void link(List<Node> nodes, int lhs, int rhs) {
            nodes.get(lhs - 1).neighbors.add(nodes.get(rhs - 1));
            nodes.get(rhs - 1).neighbors.add(nodes.get(lhs - 1));
        }
    }
}
