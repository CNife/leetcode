import test.Tester;

public class DecodeString {
    public static String decodeString(String src) {
        StringBuilder builder = new StringBuilder();
        decode(src, 0, builder);
        return builder.toString();
    }

    private static int decode(String src, int index, StringBuilder builder) {
        int count = 0;
        for (; index < src.length(); index++) {
            char c = src.charAt(index);
            if (c <= '9' && c >= '0') {
                count = count * 10 + c - '0';
            } else if (c == '[') {
                StringBuilder tmpBuilder = new StringBuilder();
                int next = decode(src, index + 1, tmpBuilder);
                for (int i = 0; i < count; i++) {
                    builder.append(tmpBuilder);
                }
                count = 0;
                index = next;
            } else if (c == ']') {
                return index;
            } else {
                builder.append(c);
            }
        }
        return index + 1;
    }

    public static void main(String[] args) {
        Tester tester = new Tester(DecodeString.class);
        tester.addTestCase("3[a]2[bc]", "aaabcbc");
        tester.addTestCase("3[a2[c]]", "accaccacc");
        tester.addTestCase("2[abc]3[cd]ef", "abcabccdcdcdef");
        tester.runTestCases();
    }
}
