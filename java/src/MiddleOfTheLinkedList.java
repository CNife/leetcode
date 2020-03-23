import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

public class MiddleOfTheLinkedList {
    static ListNode middleNode(ListNode head) {
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

    @Test
    void test() {
        List<TestCase<ListNode, ListNode>> testCases = Arrays.asList(
                new TestCase<>(ListNode.newList(3, 4, 5), ListNode.newList(1, 2, 3, 4, 5)),
                new TestCase<>(ListNode.newList(4, 5, 6), ListNode.newList(1, 2, 3, 4, 5, 6))
        );
        for (TestCase<ListNode, ListNode> testCase : testCases) {
            ListNode actual = middleNode(testCase.source);
            Assertions.assertEquals(testCase.expected, actual);
        }
    }
}
