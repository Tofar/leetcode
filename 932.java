class Solution {
    public int[] beautifulArray(int N) {
        List<Integer> array = new ArrayList<>();
        for (int i = 1; i <= N; i += 2) {
            array.add(i);
        }
        for (int i = 2; i <= N; i += 2) {
            array.add(i);
        }

        help(array.subList(0, (N + 1) / 2));
        help(array.subList((N + 1) / 2, N));
        int[] result = new int[N];
        int i = 0;
        for (Integer n : array) {
            result[i] = n;
            i++;
        }
        return result;
    }

    private void help(List<Integer> array) {
        if (array == null || array.size() <= 2) {
            return;
        }
        
        int N = array.size();
        for (int i = 1; i <= N/2; i++) {
            moveToTail(array, i);
        }
        help(array.subList(0, (N + 1) / 2));
        help(array.subList((N + 1) / 2, N));
    }

    static void moveToTail(List<Integer> array, int i) {
        for (; i < array.size() - 1; i++) {
            Integer temp = array.get(i);
            array.set(i, array.get(i + 1));
            array.set(i + 1, temp);
        }
    }
}
