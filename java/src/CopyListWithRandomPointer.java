public class CopyListWithRandomPointer {
    public static Node copyRandomList(Node head) {
        if (head == null) {
            return null;
        }

        Node next = head;
        while (next != null) {
            Node copied = new Node(next.val);
            copied.next = next.next;
            next.next = copied;
            next = copied.next;
        }

        next = head;
        Node copiedNext = head.next;
        while (copiedNext != null) {
            if (next.random != null) {
                copiedNext.random = next.random.next;
            }
            next = copiedNext.next;
            if (next == null) {
                break;
            } else {
                copiedNext = next.next;
            }
        }

        next = head;
        Node copiedHead = copiedNext = head.next;
        while (copiedNext != null && copiedNext.next != null) {
            next = next.next = next.next.next;
            copiedNext = copiedNext.next = copiedNext.next.next;
        }
        next.next = null;
        return copiedHead;
    }

    public static void main(String[] args) {
        Node node = new Node(1);
        node.next = new Node(2);
        node.random = node.next;
        node.next.random = node.next;

        Node copied = CopyListWithRandomPointer.copyRandomList(node);
        System.out.println(node);
        System.out.println(copied);
    }

    private static class Node {
        int val;
        Node next = null;
        Node random = null;

        public Node(int val) {
            this.val = val;
        }

        @Override
        public String toString() {
            StringBuilder sb = new StringBuilder("Node{val=");
            sb.append(val);
            sb.append(", random=");
            if (random == null) {
                sb.append("null");
            } else {
                sb.append(random.val);
            }
            sb.append(", next=");
            sb.append(next);
            sb.append("}");
            return sb.toString();
        }
    }
}
