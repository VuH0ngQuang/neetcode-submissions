class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        char[] str = s.toCharArray();
        for (int i = 0; i < str.length; i++) {
            if (open(str[i])) {
                stack.push(str[i]);
            } else {
                if (stack.isEmpty()) return false;
                var tmp = stack.pop();
                if (str[i] == ')' && tmp != '(') return false;
                if (str[i] == ']' && tmp != '[') return false;
                if (str[i] == '}' && tmp != '{') return false;
            }
        }
        
        return stack.isEmpty();
    }

    private boolean open(char c) {
        return c == '(' || c == '[' || c == '{'; 
    }
}
