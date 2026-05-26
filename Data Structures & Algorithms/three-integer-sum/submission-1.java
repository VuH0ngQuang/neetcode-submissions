class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        int len = nums.length;
        List<List<Integer>> ans = new ArrayList<>();
        int left, right, sum;
        for(int i = 0; i < len - 2; i++) {
            while(i > 0 && nums[i] == nums[i-1] && i < len - 2) i++;
            left = i+1;
            right = len -1;
            while (left < right) {
                sum = nums[left]+nums[right]+nums[i];
                if (sum == 0) {
                    ans.add(new ArrayList<>(List.of(nums[i],nums[left],nums[right])));
                    left++;
                    right--;
                    while(left < right && nums[left] == nums[left-1]) left++;
                    while(left < right && nums[right] == nums[right+1]) right--;
                } else if (sum > 0) {
                    right--;
                } else left++;
            }
        }

        return ans;
    }
}
