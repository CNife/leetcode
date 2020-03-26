import leetcode.ListNode;
import test.Tester;

public class ReverseLinkedList2 {
    public static ListNode reverseBetween(ListNode head, int m, int n) {
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

    public static void main(String[] args) {
        Tester tester = new Tester(ReverseLinkedList2.class);
        tester.addTestCase(
                ListNode.newList(1, 2, 3, 4, 5), 2, 4,
                ListNode.newList(1, 4, 3, 2, 5)
        );
        tester.addTestCase(
                ListNode.newList(1, 2, 3), 1, 3,
                ListNode.newList(3, 2, 1)
        );
        tester.runTestCases();
    }
}
