import leetcode.ListNode;
import test.Tester;

public class RemoveDuplicatesFromSortedList2 {
    public static ListNode deleteDuplicates(ListNode head) {
        ListNode tmpHead = new ListNode(-1);
        tmpHead.next = head;
        ListNode prev = tmpHead;
        while (prev.next != null) {
            ListNode curr = prev.next, next = curr.next;
            while (next != null && curr.val == next.val) {
                next = next.next;
            }
            if (next != curr.next) {
                prev.next = next;
            } else {
                prev = prev.next;
            }
        }
        return tmpHead.next;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(RemoveDuplicatesFromSortedList2.class);
        tester.addTestCase(
                ListNode.newList(1, 2, 3, 3, 4, 4, 5),
                ListNode.newList(1, 2, 5)
        );
        tester.addTestCase(
                ListNode.newList(1, 1, 1, 2, 3),
                ListNode.newList(2, 3)
        );
        tester.runTestCases();
    }
}
