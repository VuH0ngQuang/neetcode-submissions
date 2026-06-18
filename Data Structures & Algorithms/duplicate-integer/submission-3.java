class Solution {
    public boolean hasDuplicate(int[] nums) {
        Stack<Integer> st = new Stack<>();

        for (int i =0; i < nums.length; i++) {
            if (st.contains(nums[i])) {
                return true;
            } else {
                st.push(nums[i]);
            }
        }

        return false;
    }
}