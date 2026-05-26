class Solution {
    public int maxArea(int[] heights) {
        int max = 0;
        int amt;
        int left = 0;
        int right = heights.length - 1;
        while (left < right) {
            amt = Math.min(heights[left],heights[right]) * (right - left);
            if (amt > max) max = amt;
            if(heights[left] > heights[right]) {
                right--;
            } else left++;
        }

        return max;
    }
}
