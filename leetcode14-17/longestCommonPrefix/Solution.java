package longestCommonPrefix;

public class Solution {
    public String longestCommonPrefix(String[] strs) {
        // 首先求出各字符串的最小长度
        if (strs.length < 1) {
            return "";
        } else {
            int min_len = strs[0].length();
            for (String str: strs) {
                if (str.length() < min_len) {
                    min_len = str.length();
                }
            }

            if (min_len == 0) {
                return "";
            } else {
                for (int i = 0; i < min_len; i++) {
                    char temp = strs[0].charAt(i);
                    int j  = 1;
                    for (; j < strs.length; j++) {
                        if (temp != strs[j].charAt(i)) {
                            break;
                        }
                    }

                    if (j < strs.length) {
                        return strs[0].substring(0, i);
                    }

                }

                return strs[0].substring(0, min_len);
            }

        }
    }
}
