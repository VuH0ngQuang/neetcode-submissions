class Solution {

    public int search(int[] nums, int target) {
        int left = 0;
        int right = nums.length - 1;

        while (left <= right) {
            var mid = left + (right - left) / 2 ;
            var tmp = nums[mid];
            if (tmp == target) {
                return mid;
            } else if (tmp < target) {
                left = mid+1;
            } else {
                right = mid-1;
            }
        }

        return -1;
    }
}