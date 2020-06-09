import test.Tester;

public class TranslateNumber {
    public static int translateNum(int num) {
        char[] array = String.valueOf(num).toCharArray();
        int m = 1, n = 1;
        for (int i = 1; i < array.length; i++) {
            int tmp = m;
            m = n;
            if (array[i - 1] == '1' || array[i - 1] == '2' && array[i] <= '5')
                n += tmp;
        }
        return n;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(TranslateNumber.class);
        tester.addTestCase(12258, 5);
        tester.runTestCases();
    }
}
