import leetcode.ListNode;
import test.Tester;

public class SwapNodesInPairs {
    public static ListNode swapPairs(ListNode head) {
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

    public static void main(String[] args) {
        Tester tester = new Tester(SwapNodesInPairs.class);
        tester.addTestCase(
                ListNode.newList(1, 2, 3, 4),
                ListNode.newList(2, 1, 4, 3)
        );
        tester.addTestCase(
                ListNode.newList(1, 2, 3),
                ListNode.newList(2, 1, 3)
        );
        tester.addTestCase(
                ListNode.newList(),
                ListNode.newList()
        );
        tester.runTestCases();
    }
}
