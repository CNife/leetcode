import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

public class ReverseLinkedList2 {
    static ListNode reverseBetween(ListNode head, int m, int n) {
        if (head == null || m <= 0 || m >= n) {
            return head;
        }

        ListNode prev;
        if (m == 1) {
            prev = new ListNode(-1);
            prev.next = head;
        } else {
            prev = head;
            for (int i = 2; i < m; i++) {
                prev = prev.next;
            }
        }

        ListNode node = prev.next;
        for (int i = m; i < n; i++) {
            ListNode next = node.next;
            node.next = next.next;
            next.next = prev.next;
            prev.next = next;
        }

        return m == 1 ? prev.next : head;
    }

    @Test
    void test() {
        List<TestCase> testCases = Arrays.asList(
                new TestCase(ListNode.newList(1, 2, 3, 4, 5), 2, 4, ListNode.newList(1, 4, 3, 2, 5)),
                new TestCase(ListNode.newList(1, 2, 3), 1, 3, ListNode.newList(3, 2, 1))
        );
        for (TestCase testCase : testCases) {
            ListNode actual = reverseBetween(testCase.head, testCase.m, testCase.n);
            Assertions.assertEquals(testCase.expected, actual);
        }
    }

    private static class TestCase {
        ListNode head;
        int m;
        int n;
        ListNode expected;

        public TestCase(ListNode head, int m, int n, ListNode expected) {
            this.head = head;
            this.m = m;
            this.n = n;
            this.expected = expected;
        }
    }
}
