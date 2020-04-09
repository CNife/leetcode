import leetcode.ListNode;
import test.Tester;

public class InsertionSortList {
    public static ListNode insertionSortList(ListNode head) {
        ListNode dummyHead = new ListNode(Integer.MIN_VALUE);
        ListNode node = head;
        while (node != null) {
            ListNode next = node.next;
            insertSortedList(dummyHead, node);
            node = next;
        }
        return dummyHead.next;
    }

    private static void insertSortedList(ListNode dummyHead, ListNode node) {
        ListNode prev = dummyHead;
        while (prev.next != null) {
            if (node.val < prev.next.val) {
                node.next = prev.next;
                prev.next = node;
                return;
            } else {
                prev = prev.next;
            }
        }
        prev.next = node;
        node.next = null;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(InsertionSortList.class);
        tester.addTestCase(
                ListNode.newList(4, 2, 1, 3),
                ListNode.newList(1, 2, 3, 4)
        );
        tester.addTestCase(
                ListNode.newList(-1, 5, 3, 4, 0),
                ListNode.newList(-1, 0, 3, 4, 5)
        );
        tester.runTestCases();
    }
}
