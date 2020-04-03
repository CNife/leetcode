import java.util.ArrayDeque;
import java.util.Deque;

public class PopulatingNextRightPointersInEachNode2 {
    public static Node connect(Node root) {
        if (root == null) {
            return null;
        }

        Deque<Node> queue = new ArrayDeque<>();
        queue.addLast(root);
        Node dummyHead = new Node(-1);
        while (!queue.isEmpty()) {
            int levelCount = queue.size();
            Node node = dummyHead;
            for (int i = 0; i < levelCount; i++) {
                node.next = queue.removeFirst();
                node = node.next;
                if (node.left != null) {
                    queue.addLast(node.left);
                }
                if (node.right != null) {
                    queue.addLast(node.right);
                }
            }
            node.next = null;
        }
        return root;
    }

    public static void main(String[] args) {
        Node root = new Node(1);
        root.left = new Node(2);
        root.right = new Node(3);
        root.left.left = new Node(4);
        root.left.right = new Node(5);
        root.right.right = new Node(6);

        Node output = connect(root);
        System.out.println(output);
    }

    private static class Node {
        int val;
        Node left;
        Node right;
        Node next;

        Node(int val) {
            this.val = val;
        }

        @Override
        public String toString() {
            return "Node{" +
                    "val=" + val +
                    ", left=" + left +
                    ", right=" + right +
                    ", next=" + next +
                    '}';
        }
    }
}
