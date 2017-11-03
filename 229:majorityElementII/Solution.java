package majorityElementII;
import java.util.ArrayList;
import java.util.List;

public class Solution {
    public List<Integer> majorityElement(int[] nums) {
        // 最多有两个数大于总数的1/3
        // 摩尔投票法 Moore Voting   注：因为每一对不一样的数都会互相消去，最后留下来的candidate就是众数。
        List<Integer> result = new ArrayList<>();
        if(nums.length == 0) return result;
        int a = 1, b = 0, a_num = 0, b_num = 0;
        for(int n : nums){
            // 如果和某个候选数相等，将其计数器加1
            if(a == n){
                a_num++;
            } else if(b == n){
                b_num++;
            } else if(a_num == 0){
                // 如果 a和b的数量为0，则重置一次
                a = n;
                a_num = 1;
            } else if (b_num == 0){
                b = n;
                b_num = 1;
            } else {
                a_num--;
                b_num--;
            }
        }
        a_num = 0;
        b_num = 0;
        for(int n : nums){
            if(a == n) a_num++;
            else if(b == n) b_num++;
        }
        if(a_num > nums.length / 3) result.add(a);
        if(b_num > nums.length / 3) result.add(b);
        return result;
    }
}
