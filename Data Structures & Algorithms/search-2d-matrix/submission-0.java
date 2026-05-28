class Solution {
    boolean binarySeach(int[] arr, int target) {
        int mid;
        int left = 0;
        int right = arr.length - 1;
        while (left <= right) {
            mid = left + (right - left) / 2;
            if (arr[mid] == target) {
                return true;
            } else if (arr[mid] < target) {
                left = mid + 1;
            } else right = mid - 1;
        }
        return false;
    }

    public boolean searchMatrix(int[][] matrix, int target) {
        for(int i = 0; i < matrix.length; i++) {
            if (binarySeach(matrix[i], target)) return true;
        }
        return false;
    }
}
