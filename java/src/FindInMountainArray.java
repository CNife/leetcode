import test.Tester;

import java.util.Arrays;

public class FindInMountainArray {
    public static int findInMountainArray(int target, MountainArray array) {
        int size = array.length();
        int top = findMountainTop(array, 0, size - 1);
        int result = findAscendingArray(array, 0, top, target);
        return result == -1 ? findDescendingArray(array, top + 1, size - 1, target) : result;
    }

    private static int findMountainTop(MountainArray array, int l, int r) {
        while (l < r) {
            int mid = l + (r - l) / 2;
            if (array.get(mid) < array.get(mid + 1))
                l = mid + 1;
            else
                r = mid;
        }
        return l;
    }

    private static int findAscendingArray(MountainArray array, int l, int r, int target) {
        while (l < r) {
            int mid = l + (r - l) / 2;
            if (array.get(mid) < target)
                l = mid + 1;
            else
                r = mid;
        }
        if (array.get(l) == target)
            return l;
        return -1;
    }

    private static int findDescendingArray(MountainArray array, int l, int r, int target) {
        while (l < r) {
            int mid = l + (r - l) / 2;
            if (array.get(mid) > target)
                l = mid + 1;
            else
                r = mid;
        }
        if (array.get(l) == target)
            return l;
        return -1;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(FindInMountainArray.class);
        tester.addTestCase(
                3, new MountainArray(1, 2, 3, 4, 5, 3, 1),
                2
        );
        tester.addTestCase(
                3, new MountainArray(0, 1, 2, 4, 2, 1),
                -1
        );
        tester.runTestCases();
    }
}

class MountainArray {
    private final int[] array;
    private int counter;

    public MountainArray(int... array) {
        this.array = array;
        this.counter = 0;
    }

    int get(int index) {
        if (++counter > 100)
            throw new IllegalStateException("call MountainArray.get over 100 times");
        return array[index];
    }

    int length() {
        return array.length;
    }

    @Override
    public String toString() {
        return "MountainArray" + Arrays.toString(array);
    }
}