import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

class SortList {
    private static ListNode FAKE_HEAD = new ListNode(-1);

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

    private ListNode sortList(ListNode head) {
        if (head == null || head.next == null)
            return head;
        ListNode nextList = splitList(head);
        return merge(sortList(head), sortList(nextList));
    }

    static class SortListTest {
        @Test
        void test() {
            List<ListNode> cases = makeList(new int[][]{new int[]{4, 2, 3, 1}, new int[]{-1, 5, 3, 4, 0}});
            List<ListNode> expected = makeList(new int[][]{new int[]{1, 2, 3, 4}, new int[]{-1, 0, 3, 4, 5}});
            for (int i = 0; i < cases.size(); i++) {
                ListNode c = cases.get(i), e = expected.get(i);
                Assertions.assertEquals(new SortList().sortList(c), e);
            }
        }

//        @Test
//        void testListLength() {
//            int[] src = {1, 2};
//            ListNode head = ListNode.newList(src);
//            Assertions.assertEquals(new SortList().listLength(head), src.length);
//            Assertions.assertNotNull(head);
//        }

        private List<ListNode> makeList(int[][] src) {
            return Arrays.stream(src)
                    .map(ListNode::newList)
                    .collect(Collectors.toList());
        }
    }
}