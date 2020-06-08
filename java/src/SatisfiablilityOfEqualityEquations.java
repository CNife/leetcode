import test.Tester;

public class SatisfiablilityOfEqualityEquations {
    public static boolean equationsPossible(String[] equations) {
        if (equations == null || equations.length == 0)
            return true;

        int[] us = new int[26];
        for (int i = 0; i < 26; i++)
            us[i] = i;

        for (String equation : equations)
            if (equation.charAt(1) == '=')
                union(equation.charAt(0) - 'a', equation.charAt(3) - 'a', us);

        for (String equation : equations)
            if (equation.charAt(1) == '!')
                if (find(equation.charAt(0) - 'a', us) == find(equation.charAt(3) - 'a', us))
                    return false;

        return true;
    }

    private static int find(int i, int[] us) {
        if (us[i] != i && us[i] != -1)
            us[i] = find(us[i], us);
        return us[i];
    }

    private static void union(int lhs, int rhs, int[] us) {
        us[find(lhs, us)] = us[find(rhs, us)];
    }

    public static void main(String[] args) {
        Tester tester = new Tester(SatisfiablilityOfEqualityEquations.class);
        tester.addTestCase(new String[]{"e==d", "e==a", "f!=d", "b!=c", "a==b"}, true);
        tester.addTestCase(new String[]{"c==c", "f!=a", "f==b", "b==c"}, true);
        tester.addTestCase(new String[]{"a!=a"}, false);
        tester.addTestCase(new String[0], true);
        tester.addTestCase(null, true);
        tester.addTestCase(new String[]{"a==b", "b!=a"}, false);
        tester.addTestCase(new String[]{"b==a", "a==b"}, true);
        tester.addTestCase(new String[]{"a==b", "b==c", "a==c"}, true);
        tester.addTestCase(new String[]{"a==b", "b!=c", "c==a"}, false);
        tester.addTestCase(new String[]{"c==c", "b==d", "x!=z"}, true);
        tester.runTestCases();
    }

}
