import leetcode.ListNode;
import test.Tester;


public class PartitionList {
    public static ListNode partition(ListNode head, int x) {
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

    public static void main(String[] args) {
        Tester tester = new Tester(PartitionList.class);
        tester.addTestCase(
                ListNode.newList(1, 4, 3, 2, 5, 2), 3,
                ListNode.newList(1, 2, 2, 4, 3, 5)
        );
        tester.addTestCase(
                ListNode.newList(7, 5, 3, 4, 6, 2), 5,
                ListNode.newList(3, 4, 2, 7, 5, 6)
        );
        tester.runTestCases();
    }
}
