class Solution {
    public int[] productExceptSelf(int[] nums) {
        int[] arr1 = new int[nums.length];
        int[] arr2 = new int[nums.length];
        int[] ans = new int[nums.length];
        int len = nums.length;
        int temp1 = 1;
        int temp2 = 1;

        for(int i = 0; i < len; i++) {
            if (i != 0) {
                arr1[i] = temp1 * nums[i - 1];
                temp1 = arr1[i];
                arr2[len - i - 1] = temp2 * nums[len - i];
                temp2 = arr2[len - i -1];
            } else {
                arr1[i] = 1;
                arr2[len - i - 1] = 1;
            }
        }

        for(int i = 0; i < len; i++) {
            ans[i] = arr1[i] * arr2[i];
        }

        return ans;
    }
}  
