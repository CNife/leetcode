import test.Tester;

public class DeepDarkFraction {
    public static int[] fraction(int[] cont) {
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

    public static void main(String[] args) {
        Tester tester = new Tester(DeepDarkFraction.class);
        tester.addTestCase(
                new int[]{3, 2, 0, 2},
                new int[]{13, 4}
        );
        tester.addTestCase(
                new int[]{0, 0, 3},
                new int[]{3, 1}
        );
        tester.runTestCases();
    }
}