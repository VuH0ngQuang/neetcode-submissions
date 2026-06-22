class Solution {
    private boolean search(int[] arr, int num) {
        int left = 0;
        int right = arr.length - 1;
        while (left <= right) {
            var mid = left + (right - left) / 2;
            var tmp = arr[mid];
            if (tmp == num) return true;
            if (tmp < num) left = mid + 1;
            if (tmp > num) right = mid - 1;
        }

        return false;
    }
    public boolean searchMatrix(int[][] matrix, int target) {
        for (int i = 0; i < matrix.length; i++) {
            if (matrix[i][matrix[i].length - 1] >= target) {
                return search(matrix[i], target);
            }
        }

        return false;
    }
}
