package reverseInteger;

public class Solution {
    public int reverse(int x) {
        // min:-2147483648  max:2147483647
        //  10 => 1
        String result = "";
        String st_x = String.valueOf(x);
        if(x < 0){
            result = "-";
            st_x = String.valueOf(x*-1);

        }

        for (int i = st_x.length() - 1; i >= 0; i--){
            result += String.valueOf(st_x.charAt(i));
        }
        if (result.charAt(0) == '-'){
            if (result.length() <= 11){
                if (Long.parseLong(result) < -2147483648){
                    return 0;
                } else {
                    return Integer.parseInt(result);
                }
            }
        } else {
            if (Long.parseLong(result) > 2147483647){
                return 0;
            } else {
                return Integer.parseInt(result);
            }
        }

        return 0;
    }
}
