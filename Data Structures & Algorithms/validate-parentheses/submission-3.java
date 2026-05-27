class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        char[] str = s.toCharArray();
        char open;
        for (int i = 0; i < s.length(); i++) {
            if (str[i] == '(' ||
                str[i] == '{' ||
                str[i] == '[') stack.push(str[i]);

            if (str[i] == ')' ||
                str[i] == '}' ||
                str[i] == ']') {
                    if (stack.isEmpty()) return false;
                    open = stack.pop();
                    switch(str[i]) {
                        case ')':
                            if(open != '(') return false;
                            break;
                        case '}':
                            if(open != '{') return false;
                            break;
                        case ']':
                            if(open != '[') return false;
                            break;
                    }
            }
        }
        if (stack.isEmpty()) return true;
        return false;
    }
}
