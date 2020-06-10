import test.Tester;

public class PalindromeNumber {
    public static boolean isPalindrome(int x) {
        if (x < 0)
            return false;

        char[] array = String.valueOf(x).toCharArray();
        for (int i = 0; i < array.length; i++) {
            if (array[i] != array[array.length - 1 - i])
                return false;
        }
        return true;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(PalindromeNumber.class);
        tester.addTestCase(121, true);
        tester.addTestCase(-121, false);
        tester.addTestCase(10, false);
        tester.runTestCases();
    }
}
