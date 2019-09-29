import org.jetbrains.annotations.Contract;
import org.jetbrains.annotations.NotNull;

import java.util.Objects;

class ListNode {
    int val;
    ListNode next;

    @Contract(pure = true)
    ListNode(int val) {
        this.val = val;
    }

    @Contract(pure = true)
    static ListNode newList(@NotNull int[] values) {
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

    @Contract(value = "null -> false", pure = true)
    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (!(o instanceof ListNode)) return false;
        ListNode listNode = (ListNode) o;
        return val == listNode.val;
    }

    @Override
    public int hashCode() {
        return Objects.hash(val);
    }
}