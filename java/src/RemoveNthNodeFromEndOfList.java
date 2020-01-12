import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class RemoveNthNodeFromEndOfList {
    public static ListNode removeNthFromEnd(ListNode head, int n) {
        if (n < 1) {
            return head;
        }

        ListNode ahead = getNthNode(head, n);
        if (ahead == null) {
            return head.next;
        }

        ListNode behind = head;
        while (ahead.next != null) {
            ahead = ahead.next;
            behind = behind.next;
        }

        behind.next = behind.next.next;
        return head;
    }

    private static ListNode getNthNode(ListNode head, int n) {
        ListNode node = head;
        for (int i = 0; i < n; i++) {
            if (node != null) {
                node = node.next;
            } else {
                break;
            }
        }
        return node;
    }

    @Test
    void test1() {
        new TestCase(new int[]{1, 2, 3, 4, 5}, 2, new int[]{1, 2, 3, 5}).runTest();
    }

    @Test
    void test2() {
        new TestCase(new int[]{1, 2}, 1, new int[]{1}).runTest();
    }

    private static class TestCase {
        ListNode origin;
        ListNode expected;
        int n;

        TestCase(int[] origin, int n, int[] expected) {
            this.origin = ListNode.newList(origin);
            this.n = n;
            this.expected = ListNode.newList(expected);
        }

        void runTest() {
            ListNode output = removeNthFromEnd(origin, n);
            assertEquals(expected, output);
        }
    }
}