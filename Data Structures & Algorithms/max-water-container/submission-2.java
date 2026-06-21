class Solution {
    public int maxArea(int[] heights) {
        int left = 0;
        int right = heights.length - 1;
        int ans = 0;
        while (left < right) {
            var tmp = Math.min(heights[left],heights[right]) * (right - left);
            ans = Math.max(ans, tmp);
            if (heights[left] < heights[right]) {
                left++;
            } else {
                right--;
            }
        }

        return ans;
    }
}
