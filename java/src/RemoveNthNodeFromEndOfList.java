import leetcode.ListNode;
import test.Tester;

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

    public static void main(String[] args) {
        Tester tester = new Tester(RemoveNthNodeFromEndOfList.class);
        tester.addTestCase(
                ListNode.newList(1, 2, 3, 4, 5), 2,
                ListNode.newList(1, 2, 3, 5)
        );
        tester.addTestCase(
                ListNode.newList(1, 2), 1,
                ListNode.newList(1)
        );
        tester.runTestCases();
    }
}