import leetcode.ListNode;
import test.Tester;

import java.util.ArrayList;
import java.util.Collections;

public class AddTwoNumbers2 {
    public static ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ArrayList<Integer> lhs = list2Array(l1);
        ArrayList<Integer> rhs = list2Array(l2);
        ArrayList<Integer> result = lhs.size() < rhs.size() ? addArrays(rhs, lhs) : addArrays(lhs, rhs);
        return array2List(result);
    }

    private static ArrayList<Integer> list2Array(ListNode head) {
        ArrayList<Integer> result = new ArrayList<>();
        while (head != null) {
            result.add(head.val);
            head = head.next;
        }
        return result;
    }

    private static ArrayList<Integer> addArrays(ArrayList<Integer> lhs, ArrayList<Integer> rhs) {
        boolean carry = false;
        ArrayList<Integer> result = new ArrayList<>();
        for (
                int i = lhs.size() - 1, j = rhs.size() - 1;
                i >= 0;
                i--
        ) {
            int digit = lhs.get(i);
            if (j >= 0)
                digit += rhs.get(j--);
            if (carry)
                digit++;
            if (digit > 9) {
                digit -= 10;
                carry = true;
            } else
                carry = false;
            result.add(digit);
        }
        if (carry)
            result.add(1);
        Collections.reverse(result);
        return result;
    }

    private static ListNode array2List(ArrayList<Integer> array) {
        ListNode dummyHead = new ListNode(-1);
        ListNode node = dummyHead;
        for (int i : array) {
            node.next = new ListNode(i);
            node = node.next;
        }
        return dummyHead.next;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(AddTwoNumbers2.class);
        tester.addTestCase(
                ListNode.newList(7, 2, 4, 3),
                ListNode.newList(5, 6, 4),
                ListNode.newList(7, 8, 0, 7)
        );
        tester.addTestCase(
                ListNode.newList(5),
                ListNode.newList(5),
                ListNode.newList(1, 0)
        );
        tester.runTestCases();
    }
}
