package medianOfTwoSortedArrays;

public class Solution {
    public double findMedianSortedArrays(int[] A, int[] B) {
        if ((A == null) && (B == null)) {
            return 0.0;
        }
        int lenA = A.length;
        int lenB = B.length;
        int len = lenA + lenB;


        int indexA = 0, indexB = 0, indexC = 0;
        int[] C = new int[len];

        while (indexA < lenA && indexB < lenB) {
            if (A[indexA] < B[indexB]) {
                C[indexC++] = A[indexA++];
            } else {
                C[indexC++] = B[indexB++];
            }
        }
        // 只有A有元素
        while (indexA < lenA) {
            C[indexC++] = A[indexA++];
        }
        // 只有B有元素
        while (indexB < lenB) {
            C[indexC++] = B[indexB++];
        }


        int indexOfM1 = (len - 1) / 2;
        int indexOfM2 = len / 2;
        if (len % 2 == 0) {
            return (C[indexOfM1] + C[indexOfM2]) / 2.0;
        } else {
            return C[indexOfM2];
        }
    }
}