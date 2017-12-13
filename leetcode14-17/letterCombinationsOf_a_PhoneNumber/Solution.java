package letterCombinationsOf_a_PhoneNumber;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

/**
 * Given a digit string, return all possible letter combinations that the number could represent.

 A mapping of digit to letters (just like on the telephone buttons) is given below.

 Input:Digit string "23"
 Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 Note:
 Although the above answer is in lexicographical order, your answer could be in any order you want.
 */

public class Solution {
    private static final HashMap<Character, String> map = new HashMap<>();
    public Solution() {
        map.put('2', "abc");
        map.put('3', "def");
        map.put('4', "ghi");
        map.put('5', "jkl");
        map.put('6', "mno");
        map.put('7', "pqrs");
        map.put('8', "tuv");
        map.put('9', "wxyz");
    }

    public List<String> letterCombinations(String digits) {
        if (digits.length() == 0) {
            return new ArrayList<>();
        } else {
            List<String> list = new ArrayList<>();
            return getResult(list, digits, 1, new StringBuffer());
        }
    }

    private List<String> getResult(List<String> list, String digits, int n, StringBuffer temp) {
        if (n == digits.length()) {
            for (char ch: map.get(digits.charAt(n-1)).toCharArray()) {
                temp.append(ch);
                list.add(temp.toString());
                temp.deleteCharAt(temp.length() - 1);
            }
        } else {
            for (char ch : map.get(digits.charAt(n - 1)).toCharArray()) {
                temp.append(ch);
                getResult(list, digits, n+1, temp);
                temp.deleteCharAt(temp.length() - 1);
            }

        }
        return list;

    }
}
