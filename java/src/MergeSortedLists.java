import leetcode.ListNode;
import test.Tester;

import java.util.Comparator;
import java.util.PriorityQueue;
import java.util.Queue;

public class MergeSortedLists {
    public static ListNode mergeKLists(ListNode[] lists) {
        if (lists == null || lists.length == 0)
            return null;

        Queue<ListNode> heap = new PriorityQueue<>(
                lists.length,
                Comparator.comparingInt((ListNode lhs) -> lhs.val)
        );
        for (ListNode list : lists)
            if (list != null)
                heap.add(list);


        ListNode dummyHead = new ListNode(-1), node = dummyHead;
        while (!heap.isEmpty()) {
            ListNode nextNode = heap.remove();
            if (nextNode.next != null)
                heap.add(nextNode.next);
            node = node.next = nextNode;
        }
        return dummyHead.next;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(MergeSortedLists.class);
        tester.addTestCase(
                new ListNode[]{
                        ListNode.newList(1, 4, 5),
                        ListNode.newList(1, 3, 4),
                        ListNode.newList(2, 6)
                },
                ListNode.newList(1, 1, 2, 3, 4, 4, 5, 6)
        );
        tester.addTestCase(
                null,
                null
        );
        tester.addTestCase(
                new ListNode[]{},
                null
        );
        tester.addTestCase(
                new ListNode[]{null},
                null
        );
        tester.runTestCases();
    }
}
