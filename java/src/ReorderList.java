import leetcode.ListNode;
import test.Tester;

public class ReorderList {
    public static ListNode reorderList(ListNode head) {
        if (head == null) {
            return null;
        }

        ListNode right = split(head);
        right = reverse(right);
        insert(head, right);
        return head;
    }

    private static ListNode split(ListNode head) {
        ListNode fast = head, slow = head;
        while (fast != null) {
            fast = fast.next;
            if (fast != null) {
                fast = fast.next;
                slow = slow.next;
            }
        }
        ListNode rightHead = slow.next;
        slow.next = null;
        return rightHead;
    }

    private static ListNode reverse(ListNode head) {
        ListNode prev = null, curr = head;
        while (curr != null) {
            ListNode next = curr.next;
            curr.next = prev;
            prev = curr;
            curr = next;
        }
        return prev;
    }

    private static void insert(ListNode left, ListNode right) {
        while (right != null) {
            ListNode leftNext = left.next, rightNext = right.next;
            left.next = right;
            left = right.next = leftNext;
            right = rightNext;
        }
    }

    public static void main(String[] args) {
        Tester tester = new Tester(ReorderList.class);
        tester.addTestCase(
                ListNode.newList(1, 2, 3, 4),
                ListNode.newList(1, 4, 2, 3)
        );
        tester.addTestCase(
                ListNode.newList(1, 2, 3, 4, 5),
                ListNode.newList(1, 5, 2, 4, 3)
        );
        tester.addTestCase(
                ListNode.newList(),
                ListNode.newList()
        );
        tester.runTestCases();
    }
}
