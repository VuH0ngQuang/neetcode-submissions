class Solution {
    public int evalRPN(String[] tokens) {
        Stack<Integer> stack = new Stack<>();

        for (String token : tokens) {
            if (isOp(token)) {
                int tmp2 = stack.pop();
                int tmp1 = stack.pop();

                if (token.equals("+")) {
                    stack.push(tmp1 + tmp2);
                } else if (token.equals("-")) {
                    stack.push(tmp1 - tmp2);
                } else if (token.equals("*")) {
                    stack.push(tmp1 * tmp2);
                } else {
                    stack.push(tmp1 / tmp2);
                }
            } else {
                stack.push(Integer.parseInt(token));
            }
        }

        return stack.pop();
    }

    private boolean isOp(String c) {
        return c.equals("+") || c.equals("-") || c.equals("*") || c.equals("/");
    }
}