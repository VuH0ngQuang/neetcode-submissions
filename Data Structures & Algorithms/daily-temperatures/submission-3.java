class Solution {
    public int[] dailyTemperatures(int[] temperatures) {
        Stack<Integer> stack = new Stack<>();
        stack.push(0);
        int[] ans = new int[temperatures.length];

        for(int i = 1; i < temperatures.length; i++) {
            var temp = temperatures[i]; 
            while (!stack.isEmpty() && temp > temperatures[stack.peek()]) {
                var prev = stack.pop();
                ans[prev] = i - prev;
            }
            stack.push(i);
        }

        return ans;
    }
}