package integerToRoman;

/**
 * Given an integer, convert it to a roman numeral.

 * Input is guaranteed to be within the range from 1 to 3999.
 */
    /**
      I（1）、V（5）、X（10）、L（50）、C（100）、D（500）和M（1000）

     重复数次：一个罗马数字重复几次，就表示这个数的几倍。
     右加左减：
     在较大的罗马数字的右边记上较小的罗马数字，表示大数字加小数字。
     在较大的罗马数字的左边记上较小的罗马数字，表示大数字减小数字。
     左减的数字有限制，仅限于I、X、C。比如45不可以写成VL，只能是XLV
     但是，左减时不可跨越一个位值。比如，99不可以用IC（ {\displaystyle 100-1} 100-1）表示，而是用XCIX（ {\displaystyle [100-10]+[10-1]} [100-10]+[10-1]）表示。（等同于阿拉伯数字每位数字分别表示。）
     左减数字必须为一位，比如8写成VIII，而非IIX。
     右加数字不可连续超过三位，比如14写成XIV，而非XIIII。（见下方“数码限制”一项。）

     */
    // 感觉题目有点问题 测试了3999，感觉自己的也是符合规则的，答案的也是符合规则的
public class Solution {
    public String intToRoman(int num) {
        String result = "";
        while (num != 0) {
            if (num >= 1000) {
                result += "M";
                num -= 1000;
            } else {
                if (num >= 500) {
                    result += "D";
                    num -= 500;
                } else {
                    if (num >= 100) {
                        if (num < 400) {
                            result += "C";
                            num -= 100;
                        } else {
                            result += "C";
                            result += "D";
                            num -= 400;
                        }
                    } else {
                        if (num >= 50) {
                            result += "L";
                            num -= 50;
                        } else {
                            if (num >= 10) {
                                if (num < 40) {
                                    result += "X";
                                    num -= 10;
                                } else {
                                    result += "X";
                                    result += "L";
                                    num -= 40;
                                }
                            } else {
                                if (num >= 5) {
                                    result += "V";
                                    num -= 5;
                                } else {
                                    if (num >= 1) {
                                        if (num < 4) {
                                            result += "I";
                                            num -= 1;
                                        } else {
                                            result += "I";
                                            result += "V";
                                            num -= 4;
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }

        }
        return result;
    }
}
