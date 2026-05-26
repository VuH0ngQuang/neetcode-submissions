class Solution {

    boolean rowColumn(char[][] board, int n) {
        HashSet<Character> rowSet = new HashSet<>();
        HashSet<Character> colSet = new HashSet<>();

        for (int i = 0; i < 9; i++) {
            char row = board[n][i];
            char col = board[i][n];

            if (row != '.') {
                if (rowSet.contains(row)) return false;
                rowSet.add(row);
            }

            if (col != '.') {
                if (colSet.contains(col)) return false;
                colSet.add(col);
            }
        }

        return true;
    }

    boolean square(char[][] board, int rowStart, int colStart) {
        HashSet<Character> set = new HashSet<>();

        for (int r = rowStart; r < rowStart + 3; r++) {
            for (int c = colStart; c < colStart + 3; c++) {
                char ch = board[r][c];

                if (ch != '.') {
                    if (set.contains(ch)) return false;
                    set.add(ch);
                }
            }
        }

        return true;
    }

    public boolean isValidSudoku(char[][] board) {

        // Check rows and columns
        for (int i = 0; i < 9; i++) {
            if (!rowColumn(board, i)) {
                return false;
            }
        }

        // Check 3x3 squares
        for (int r = 0; r < 9; r += 3) {
            for (int c = 0; c < 9; c += 3) {
                if (!square(board, r, c)) {
                    return false;
                }
            }
        }

        return true;
    }
}