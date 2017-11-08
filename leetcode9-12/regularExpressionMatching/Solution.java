package regularExpressionMatching;

public class Solution {
    public boolean isMatch(String s, String p) {
        return helper(s,p,0,0);
    }

    // 不用处理 \
    private boolean helper(String s, String p, int i, int j) {
        if(j == p.length()) {
            return i == s.length();
        }

        if(j == p.length()-1 || p.charAt(j+1) != '*') {
            // 到达s 的末尾或者 不匹配
            if(i == s.length() || s.charAt(i) != p.charAt(j) && p.charAt(j) != '.') {
                return false;
            } else {
                // 此时匹配继续下一步
                return helper(s,p,i+1,j+1);
            }
        }

        // *功能
        while(i < s.length() && (p.charAt(j) == '.' || s.charAt(i) == p.charAt(j))) {
            if(helper(s, p, i,j+2)) {
                return true;
            }
            i++;
        }
        return helper(s, p, i,j+2);
    }

// 有问题的暴力求解
//    public boolean isMatch(String s, String p) {
//        if (p.equals("")) {
//            if (s.equals("")) {
//                return true;
//            } else {
//                return false;
//            }
//        }
//
//        int subscript = 0;
//        boolean isTransMeaning = false;
//        for (int i = 0; i < s.length(); i++) {
//            if (subscript < p.length()) {
//                switch (p.charAt(subscript)) {
//                    case '.':
//                        if (isTransMeaning) {
//                            if (s.charAt(i) != '.') {
//                                return false;
//                            }
//                        }
//                        isTransMeaning = false;
//                        subscript++;
//                        break;
//                    case '*':
//                        if (!isTransMeaning && p.charAt(subscript - 1) == '.') {
//                            return matchZeroOrMore(s.substring(i-1), p.substring(subscript + 1), p.charAt(subscript - 1), true);
//
//                        } else {
//                            return matchZeroOrMore(s.substring(i-1), p.substring(subscript + 1), p.charAt(subscript - 1), false);
//
//                        }
////                    break;
//                    case '\\':
//
//                        isTransMeaning = true;
//                        subscript++;
//                        break;
//                    default:
//                        if (s.charAt(i) != p.charAt(subscript)) {
//                            if (subscript < p.length() - 1 && p.charAt(subscript + 1) == '*') {
//                                return matchZeroOrMore(s.substring(i), p.substring(subscript + 2), p.charAt(subscript), false);
//                            } else {
//                                return false;
//                            }
//                        }
//                        isTransMeaning = false;
//                        subscript++;
//                        break;
//                }
//            } else {
//                return false;
//            }
//        }
//        if (subscript < p.length()) {
//            if (p.length() - subscript == 1) {
//                if (p.charAt(subscript) == '*') {
//                    return true;
//                } else {
//                    return false;
//                }
//            } else {
//                if (p.charAt(subscript + 1) == '*') {
//                    return isMatch("", p.substring(subscript + 2));
//                } else {
//                    return false;
//                }
//            }
//        } else {
//            return true;
//        }
//    }
//
//    public boolean matchZeroOrMore(String s, String p, char currentChar, boolean any){
//        if (any) {
//            for (int i = 0; i < s.length() - 1; i++) {
//
//                if (isMatch(s.substring(i + 1), p)) {
//                    return true;
//                }
//            }
//
//            if (p.equals("")) {
//                return true;
//            } else {
//                return false;
//            }
//
//        } else {
//            if (s.charAt(0) != currentChar) {
//                return isMatch(s, p);
//            } else {
//                for (int i = 0; i < s.length(); i++) {
//                    if (s.charAt(i) == currentChar) {
//                        if (isMatch(s.substring(i + 1), p)) {
//                            return true;
//                        }
//                    } else {
//                        return false;
//                    }
//                }
//                if (p.equals("")) {
//                    return true;
//                } else {
//                    return false;
//                }
//            }
//        }
//
//    }
}
