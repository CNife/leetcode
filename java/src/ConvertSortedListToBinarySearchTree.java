import leetcode.ListNode;
import leetcode.TreeNode;

import java.util.ArrayList;

public class ConvertSortedListToBinarySearchTree {
    public static TreeNode sortedListToBST(ListNode head) {
        ArrayList<Integer> nums = linkedList2Array(head);
        return buildTreeFromArray(nums, 0, nums.size());
    }

    private static ArrayList<Integer> linkedList2Array(ListNode head) {
        ArrayList<Integer> result = new ArrayList<>();
        while (head != null) {
            result.add(head.val);
            head = head.next;
        }
        return result;
    }

    private static TreeNode buildTreeFromArray(ArrayList<Integer> nums, int start, int end) {
        if (start >= end) {
            return null;
        }

        int middle = start + (end - start) / 2;
        TreeNode root = new TreeNode(nums.get(middle));
        root.left = buildTreeFromArray(nums, start, middle);
        root.right = buildTreeFromArray(nums, middle + 1, end);
        return root;
    }

    public static void main(String[] args) {
        ListNode list = ListNode.newList(-10, -3, 0, 5, 9);
        TreeNode actual = sortedListToBST(list);
        System.out.println(actual);
    }
}
