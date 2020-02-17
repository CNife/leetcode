import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

public class SwapNodesInPairs {
    static ListNode swapPairs(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }

        ListNode newHead = head.next;
        head.next = newHead.next;
        newHead.next = head;

        ListNode prev = head;
        while (prev.next != null && prev.next.next != null) {
            ListNode next = prev.next, nextNext = next.next;
            next.next = nextNext.next;
            nextNext.next = next;
            prev.next = nextNext;
            prev = next;
        }
        return newHead;
    }

    @Test
    void test() {
        TestCase[] cases = new TestCase[]{
                new TestCase(new int[]{1, 2, 3, 4}, new int[]{2, 1, 4, 3}),
                new TestCase(new int[]{1, 2, 3}, new int[]{2, 1, 3}),
                new TestCase(new int[]{}, new int[]{})
        };
        for (TestCase testCase : cases) {
            Assertions.assertEquals(testCase.expected, swapPairs(testCase.input));
        }
    }

    private static class TestCase {
        ListNode input;
        ListNode expected;

        TestCase(int[] input, int[] expected) {
            this.input = ListNode.newList(input);
            this.expected = ListNode.newList(expected);
        }
    }
}
