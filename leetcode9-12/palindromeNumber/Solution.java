package palindromeNumber;

public class Solution {
    public boolean isPalindrome(int x) {

        if (x < 0) {
            return false;
        }
        int x2 = x;
        int i = 1;
        while (true){
            x2 /= 10;
            if (x2 == 0) {
                break;
            }
            i *= 10;
        }

        while (x != 0 && i != 0) {
            if (x % 10 != (x/i) % 10) {
                return false;
            } else {
                x = x/10;
                i /= 100;
            }
        }
        return true;
    }

}
