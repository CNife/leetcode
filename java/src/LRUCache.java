import java.util.HashMap;

public class LRUCache {
    private HashMap<Integer, Node> map;
    private Node head;
    private Node tail;
    private int capacity;

    public LRUCache(int capacity) {
        this.capacity = capacity;
        map = new HashMap<>(capacity);
        head = null;
        tail = null;
    }

    public int get(int key) {
        Node node = map.get(key);
        if (node == null)
            return -1;
        moveToHead(node);
        return node.value;
    }

    public void put(int key, int value) {
        Node node = map.get(key);
        if (node != null) {
            node.value = value;
            moveToHead(node);
        } else {
            if (map.size() >= capacity) {
                Node removed = removeTail();
                map.remove(removed.key);
            }
            node = new Node(key, value);
            map.put(key, node);
            insertAsHead(node);
        }
    }

    private void moveToHead(Node node) {
        if (node.prev == null)
            return;
        node.prev.next = node.next;
        if (node.next != null)
            node.next.prev = node.prev;
        else
            tail = node.prev;
        node.prev = null;
        node.next = head;
        head.prev = node;
        head = node;
    }

    private void insertAsHead(Node node) {
        if (head == null) {
            head = tail = node;
        } else {
            node.next = head;
            head.prev = node;
            head = node;
        }
    }

    private Node removeTail() {
        Node res = tail;
        if (head == tail) {
            head = tail = null;
        } else {
            tail = tail.prev;
            tail.next = null;
        }
        return res;
    }

    private static class Node {
        int key;
        int value;
        Node prev;
        Node next;

        Node(int key, int value) {
            this.key = key;
            this.value = value;
            this.prev = null;
            this.next = null;
        }
    }
}
