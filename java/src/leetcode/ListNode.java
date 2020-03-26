package leetcode;

import java.util.Objects;

public class ListNode {
    public int val;
    public ListNode next;

    public ListNode(int val) {
        this.val = val;
    }

    public static ListNode newList(int... values) {
        if (values == null) {
            return null;
        }
        ListNode head = null, prev = null;
        for (int value : values) {
            ListNode next = new ListNode(value);
            if (prev == null) {
                head = prev = next;
            } else {
                prev.next = next;
                prev = next;
            }
        }
        return head;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        ListNode listNode = (ListNode) o;
        return val == listNode.val &&
                Objects.equals(next, listNode.next);
    }

    @Override
    public int hashCode() {
        return Objects.hash(val, next);
    }

    @Override
    public String toString() {
        return val + "->" + next;
    }
}
