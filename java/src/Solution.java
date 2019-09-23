import java.util.ArrayList;
import java.util.List;

class Solution {
    private String[] LETTER_TABLE = {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};

    public List<String> letterCombinations(String digits) {
        List<String> res = new ArrayList<>();
        doCombination(res, new StringBuilder(), digits, 0);
        return res;
    }

    private void doCombination(List<String> res, StringBuilder sb, String digits, int i) {
        if (i < digits.length()) {
            char letter = digits.charAt(i);
            for (char next : LETTER_TABLE[letter - '0'].toCharArray()) {
                StringBuilder nsb = new StringBuilder(sb);
                nsb.append(next);
                doCombination(res, nsb, digits, i + 1);
            }
        } else if (!digits.isEmpty()) {
            res.add(sb.toString());
        }
    }
}