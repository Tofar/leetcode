package zigZagConversion;

public class Solution {
    public String convert(String s, int numRows) {
        String[] resultArray = new String[numRows];
        for (int i = 0; i < numRows; i++){
            resultArray[i] = "";
        }
        if (numRows <= 1){
            return s;
        } else {
            if (numRows == 2){
                for (int i = 0; i < s.length(); i++){
                    resultArray[i % 2] += String.valueOf(s.charAt(i));
                }
            } else {
                int rowNum = 1;
                for (int i = 0; i < s.length(); i++){
                    if (i%(2*numRows - 2) < numRows){
                        if (i%(2*numRows - 2) == 0){
                            rowNum--;
                        } else {
                            rowNum++;
                        }
                        resultArray[rowNum] += String.valueOf(s.charAt(i));

                    } else {
                        rowNum--;
                        resultArray[rowNum] += String.valueOf(s.charAt(i));

                    }
                }
            }
        }

        String result = "";
        for (String everyString : resultArray){
            result += everyString;
        }
        return result;
    }
}
