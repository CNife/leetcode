import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

public class PartitionList {
    static ListNode partition(ListNode head, int x) {
        if (head == null || head.next == null) {
            return head;
        }

        ListNode tmpHead = new ListNode(-1);
        tmpHead.next = head;

        ListNode smallTail = tmpHead;
        while (smallTail.next != null && smallTail.next.val < x) {
            smallTail = smallTail.next;
        }
        if (smallTail.next == null) {
            return head;
        }

        ListNode bigHead = smallTail.next, curr = bigHead, next = curr.next;
        while (next != null) {
            if (next.val < x) {
                curr.next = next.next;
                next.next = bigHead;
                smallTail = smallTail.next = next;
                next = curr.next;
            } else {
                curr = next;
                next = next.next;
            }
        }
        return tmpHead.next;
    }

    @Test
    void test() {
        List<TestCase> testCases = Arrays.asList(
                new TestCase(ListNode.newList(1, 4, 3, 2, 5, 2), 3, ListNode.newList(1, 2, 2, 4, 3, 5)),
                new TestCase(ListNode.newList(7, 5, 3, 4, 6, 2), 5, ListNode.newList(3, 4, 2, 7, 5, 6))
        );
        for (TestCase testCase : testCases) {
            ListNode output = partition(testCase.head, testCase.x);
            Assertions.assertEquals(testCase.expected, output);
        }
    }

    private static class TestCase {
        ListNode head;
        int x;
        ListNode expected;

        public TestCase(ListNode head, int x, ListNode expected) {
            this.head = head;
            this.x = x;
            this.expected = expected;
        }
    }
}
