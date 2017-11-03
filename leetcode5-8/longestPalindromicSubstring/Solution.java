package longestPalindromicSubstring;

public class Solution {

    public String longestPalindrome(String s) {
        if(s == null || s.length() <= 1) {
            return s;
        }

        int len = s.length();
        int maxLen = 1;
        boolean [][] dp = new boolean[len][len];

        String longest = String.valueOf(s.charAt(0));
        for(int k = 0; k < s.length(); k++){
            for(int i = 0; i < len - k; i++){
                int j = i + k;
                // 首先会判断两者间隔小于或等于1的是否相等，或者两者的内部相邻的是否相等
		// 每个回文字符串都会经历从间隔到0再到1,之后只需判断内部相邻的是否相等
                if(s.charAt(i) == s.charAt(j) && (j - i <= 2 || dp[i + 1][j - 1])){
                    dp[i][j] = true;

                    if(j - i + 1 > maxLen){
                        maxLen = j - i + 1;
                        longest = s.substring(i, j + 1);
                    }
                }
            }
        }

        return longest;
    }
}
