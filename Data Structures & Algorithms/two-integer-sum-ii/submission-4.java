class Solution {
    public int[] twoSum(int[] numbers, int target) {
        int left = 0;
        int right = numbers.length - 1;
        while (left < right) {
            var tmp = numbers[right] + numbers[left];
            if (tmp == target) {
                return new int[]{left+1, right+1};
            }
            if (tmp < target) {
                left++;
            } else {
                right--;
            }
        }

        return new int[]{0,0};
    }
}
