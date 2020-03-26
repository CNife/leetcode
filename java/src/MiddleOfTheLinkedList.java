import leetcode.ListNode;
import test.Tester;

public class MiddleOfTheLinkedList {
    public static ListNode middleNode(ListNode head) {
        ListNode fast = head, slow = head;
        while (fast != null) {
            fast = fast.next;
            if (fast != null) {
                slow = slow.next;
                fast = fast.next;
            }
        }
        return slow;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(MiddleOfTheLinkedList.class);
        tester.addTestCase(
                ListNode.newList(1, 2, 3, 4, 5),
                ListNode.newList(3, 4, 5)
        );
        tester.addTestCase(
                ListNode.newList(1, 2, 3, 4, 5, 6),
                ListNode.newList(4, 5, 6)
        );
        tester.runTestCases();
    }
}
