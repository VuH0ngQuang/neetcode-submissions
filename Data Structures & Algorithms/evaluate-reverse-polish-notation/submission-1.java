class Solution {
    public int evalRPN(String[] tokens) {
        Stack<Integer> stack = new Stack<>();
        int a,b,ans;
        for (int i = 0; i < tokens.length; i++) {
            if (tokens[i].equals("+") ||
                tokens[i].equals("-") ||
                tokens[i].equals("*") ||
                tokens[i].equals("/")
            ) {
                b = stack.pop();
                a = stack.pop();
                switch(tokens[i]) {
                    case "+":
                        stack.push(a+b);
                        break;
                    case "-":
                        stack.push(a-b);
                        break;
                    case "*":
                        stack.push(a*b);
                        break;
                    case "/":
                        stack.push(a/b);
                        break;
                }
            } else stack.push(Integer.parseInt(tokens[i]));
        }
        return stack.pop();
    }
}
