class Solution {
    public int[] dailyTemperatures(int[] temperatures) {
        Stack<Integer> temp = new Stack<>();
        Stack<Integer> index = new Stack<>();

        int len = temperatures.length;
        int[] ans = new int[len];
        int tmp, tmpIndex;

        temp.push(temperatures[0]);
        index.push(0);

        for(int i = 1; i < len; i++) {
            while(!temp.empty() && temp.peek() < temperatures[i]) {
                tmp = temp.pop();
                tmpIndex = index.pop();
                ans[tmpIndex] = i - tmpIndex;
            }
            temp.push(temperatures[i]);
            index.push(i);
        }
        return ans;
    }
}
