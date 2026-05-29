class Solution {
    public int maxProfit(int[] prices) {
        int i = 0;
        int j;
        int ans = 0;
        int min = prices[0];
        int max = prices[0];
        for (j = 1; j < prices.length; j++) {
            if (min >= prices[j]) {
                ans = Math.max(ans, max - min);
                min = prices[j];
                max = 0;
            } else max = Math.max(max, prices[j]);
        }
        ans = Math.max(ans, max - min);
        return ans;
    }
}
