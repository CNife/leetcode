import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

public class DeepDarkFraction {
    public int[] fraction(int[] cont) {
        int[] result = new int[]{
                cont[cont.length - 1],
                1
        };
        for (int i = cont.length - 1; i > 0; i--) {
            int tmp = result[0];
            result[0] = cont[i - 1] * result[0] + result[1];
            result[1] = tmp;
        }
        return result;
    }

    @Test
    void test() {
        int[][] input = new int[][]{
                new int[]{3, 2, 0, 2},
                new int[]{0, 0, 3}
        };
        int[][] expected = new int[][]{
                new int[]{13, 4},
                new int[]{3, 1}
        };

        for (int i = 0; i < input.length; i++) {
            assertArrayEquals(expected[i], fraction(input[i]));
        }
    }
}