import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.util.*;

public class UniqueBinarySearchTrees2 {
    // 回溯算法，因为官方没有给 TreeNode 重载 hashCode() 方法而去重失败
    //
    // static List<TreeNode> generateTrees(int n) {
    //     Set<TreeNode> result = new HashSet<>();
    //     if (n < 1) {
    //         return new ArrayList<>();
    //     }
    //
    //     boolean[] used = new boolean[n + 1];
    //     for (int i = 1; i <= n; i++) {
    //         TreeNode tree = new TreeNode(i);
    //         used[i] = true;
    //         backtrack(1, tree, used, result);
    //         used[i] = false;
    //     }
    //     return new ArrayList<>(result);
    // }
    //
    // private static void backtrack(int count, TreeNode tree, boolean[] used, Set<TreeNode> result) {
    //     if (count >= used.length - 1) {
    //         result.add(deepCloneTree(tree));
    //         return;
    //     }
    //
    //     for (int i = 1; i < used.length; i++) {
    //         if (!used[i]) {
    //             TreeNode parent = findParentToInsert(tree, i);
    //             if (i < parent.val)
    //                 parent.left = new TreeNode(i);
    //             else
    //                 parent.right = new TreeNode(i);
    //             used[i] = true;
    //
    //             backtrack(count + 1, tree, used, result);
    //
    //             if (i < parent.val)
    //                 parent.left = null;
    //             else
    //                 parent.right = null;
    //             used[i] = false;
    //         }
    //     }
    // }
    //
    // private static TreeNode deepCloneTree(TreeNode tree) {
    //     if (tree == null) {
    //         return null;
    //     } else {
    //         TreeNode cloned = new TreeNode(tree.val);
    //         cloned.left = deepCloneTree(tree.left);
    //         cloned.right = deepCloneTree(tree.right);
    //         return cloned;
    //     }
    // }
    //
    // private static TreeNode findParentToInsert(TreeNode tree, int value) {
    //     TreeNode parent = tree, next = tree;
    //     while (next != null) {
    //         parent = next;
    //         next = value < parent.val ? parent.left : parent.right;
    //     }
    //     return parent;
    // }

    static List<TreeNode> generateTrees(int n) {
        if (n <= 0)
            return new LinkedList<>();
        HashMap<Pair, LinkedList<TreeNode>> memo = new HashMap<>();
        return generateTrees(1, n, memo);
    }

    private static LinkedList<TreeNode> generateTrees(int start, int end, HashMap<Pair, LinkedList<TreeNode>> memo) {
        LinkedList<TreeNode> result = new LinkedList<>();
        if (start > end) {
            result.add(null);
            return result;
        }

        Pair pair = new Pair(start, end);
        LinkedList<TreeNode> cached = memo.get(pair);
        if (cached != null) {
            return deepClone(cached);
        }

        for (int i = start; i <= end; i++) {
            LinkedList<TreeNode> leftTrees = generateTrees(start, i - 1, memo);
            LinkedList<TreeNode> rightTrees = generateTrees(i + 1, end, memo);
            for (TreeNode left : leftTrees) {
                for (TreeNode right : rightTrees) {
                    TreeNode root = new TreeNode(i);
                    root.left = left;
                    root.right = right;
                    result.add(root);
                }
            }
        }
        memo.put(pair, result);
        return result;
    }

    private static LinkedList<TreeNode> deepClone(LinkedList<TreeNode> src) {
        LinkedList<TreeNode> result = new LinkedList<>();
        for (TreeNode tree : src) {
            result.add(deepCloneTree(tree));
        }
        return result;
    }

    private static TreeNode deepCloneTree(TreeNode tree) {
        if (tree == null) {
            return null;
        } else {
            TreeNode cloned = new TreeNode(tree.val);
            cloned.left = deepCloneTree(tree.left);
            cloned.right = deepCloneTree(tree.right);
            return cloned;
        }
    }

    @Test
    void test() {
        List<TestCase<List<TreeNode>, Integer>> testCases = Arrays.asList(
                new TestCase<>(
                        Arrays.asList(
                                TreeNode.newTree(1, null, 3, 2),
                                TreeNode.newTree(3, 2, null, 1),
                                TreeNode.newTree(3, 1, null, null, 2),
                                TreeNode.newTree(2, 1, 3),
                                TreeNode.newTree(1, null, 2, null, 3)),
                        3),
                new TestCase<>(
                        Collections.singletonList(
                                TreeNode.newTree(1)),
                        1
                )
        );
        for (TestCase<List<TreeNode>, Integer> testCase : testCases) {
            List<TreeNode> actual = generateTrees(testCase.source);
            Assertions.assertEquals(new HashSet<>(testCase.expected), new HashSet<>(actual));
        }
    }

    private static class Pair {
        int start;
        int end;

        public Pair(int start, int end) {
            this.start = start;
            this.end = end;
        }

        @Override
        public boolean equals(Object o) {
            if (this == o) return true;
            if (o == null || getClass() != o.getClass()) return false;
            Pair pair = (Pair) o;
            return start == pair.start &&
                    end == pair.end;
        }

        @Override
        public int hashCode() {
            return Objects.hash(start, end);
        }
    }
}
