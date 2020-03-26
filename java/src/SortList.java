import leetcode.ListNode;
import test.Tester;

public class SortList {
    private static ListNode FAKE_HEAD = new ListNode(-1);

    public static ListNode sortList(ListNode head) {
        if (head == null || head.next == null)
            return head;
        ListNode nextList = splitList(head);
        return merge(sortList(head), sortList(nextList));
    }

    private static ListNode splitList(ListNode head) {
        ListNode ptr = head.next;
        while (ptr != null && ptr.next != null) {
            head = head.next;
            ptr = ptr.next.next;
        }
        ListNode res = head.next;
        head.next = null;
        return res;
    }

    private static ListNode merge(ListNode lhs, ListNode rhs) {
        ListNode ptr = FAKE_HEAD;
        while (lhs != null && rhs != null) {
            if (lhs.val < rhs.val) {
                ptr = ptr.next = lhs;
                lhs = lhs.next;
            } else {
                ptr = ptr.next = rhs;
                rhs = rhs.next;
            }
        }
        if (lhs != null) {
            ptr.next = lhs;
        } else {
            ptr.next = rhs;
        }
        return FAKE_HEAD.next;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(SortList.class);
        tester.addTestCase(
                ListNode.newList(4, 2, 3, 1),
                ListNode.newList(1, 2, 3, 4)
        );
        tester.addTestCase(
                ListNode.newList(-1, 5, 3, 4, 0),
                ListNode.newList(-1, 0, 3, 4, 5)
        );
        tester.runTestCases();
    }
}