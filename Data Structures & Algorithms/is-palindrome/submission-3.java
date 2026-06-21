class Solution {
    public boolean isPalindrome(String s) {
        char[] str = s.toCharArray();
        int left = 0;
        int right = s.length() - 1;

        while (left < right) {
            while (left < right && !Character.isLetterOrDigit(str[left])) {
                left++;
            }
            while (left < right && !Character.isLetterOrDigit(str[right])) {
                right--;
            }
            if (Character.toLowerCase(str[left]) != Character.toLowerCase(str[right])) return false;
            left++;
            right--;
        }

        return true;
    }
}
