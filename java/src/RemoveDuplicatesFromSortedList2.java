import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;

public class RemoveDuplicatesFromSortedList2 {
    static ListNode deleteDuplicates(ListNode head) {
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

    @Test
    void test() {
        List<TestCase<ListNode, ListNode>> testCases = Arrays.asList(
                new TestCase<>(ListNode.newList(1, 2, 5), ListNode.newList(1, 2, 3, 3, 4, 4, 5)),
                new TestCase<>(ListNode.newList(2, 3), ListNode.newList(1, 1, 1, 2, 3))
        );
        for (TestCase<ListNode, ListNode> testCase : testCases) {
            ListNode output = deleteDuplicates(testCase.source);
            Assertions.assertEquals(testCase.expected, output);
        }
    }
}
