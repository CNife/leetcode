import test.Tester;

public class SpiralOrder {
    public static int[] spiralOrder(int[][] matrix) {
        if (matrix == null || matrix.length <= 0 || matrix[0].length <= 0)
            return new int[]{};

        int[] result = new int[matrix.length * matrix[0].length];
        int i = 0;

        int up = 0, down = matrix.length;
        int left = 0, right = matrix[0].length;
        while (down - up >= 2 && right - left >= 2) {
            for (int c = left; c < right; c++)
                result[i++] = matrix[up][c];
            for (int r = up + 1; r < down - 1; r++)
                result[i++] = matrix[r][right - 1];
            for (int c = right - 1; c >= left; c--)
                result[i++] = matrix[down - 1][c];
            for (int r = down - 2; r >= up + 1; r--)
                result[i++] = matrix[r][left];

            up++;
            down--;
            left++;
            right--;
        }

        if (down - up == 1)
            for (int j = left; j < right; j++)
                result[i++] = matrix[up][j];
        else if (right - left == 1)
            for (int j = up; j < down; j++)
                result[i++] = matrix[j][left];

        return result;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(SpiralOrder.class);
        tester.addTestCase(null, new int[]{});
        tester.addTestCase(new int[][]{}, new int[]{});
        tester.addTestCase(new int[][]{new int[]{}}, new int[]{});
        tester.addTestCase(
                new int[][]{
                        new int[]{1, 2, 3},
                        new int[]{4, 5, 6},
                        new int[]{7, 8, 9}
                },
                new int[]{1, 2, 3, 6, 9, 8, 7, 4, 5}
        );
        tester.addTestCase(
                new int[][]{
                        new int[]{1, 2, 3, 4},
                        new int[]{5, 6, 7, 8},
                        new int[]{9, 10, 11, 12},
                },
                new int[]{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
        );
        tester.runTestCases();
    }
}
