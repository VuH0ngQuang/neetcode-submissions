class Solution {
    public int findMin(int[] nums) {
        int len = nums.length;
        int left = 0;
        int right = len - 1;
        int min = Math.min(nums[left], nums[right]);

        int mid;

        while(left <= right) {
            mid = left + (right - left) / 2;
            min = Math.min(min, nums[mid]);

            if((nums[mid]) > nums[right]) {
                left = mid + 1;
            } else right = mid - 1;
        }

        return min;
    }
}
