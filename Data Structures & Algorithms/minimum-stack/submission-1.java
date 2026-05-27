class MinStack {

    Stack<Integer> min;
    Stack<Integer> stack;

    public MinStack() {
        min = new Stack<>();
        stack = new Stack<>();
    }
    
    public void push(int val) {
        if (stack.empty()) {
            min.push(val);
            stack.push(val);
        } else {
            if(min.peek() >= val) {
                min.push(val);
                stack.push(val);
            } else stack.push(val);
        }
    }
    
    public void pop() {
        int tmp = stack.pop();
        if (min.peek() == tmp) min.pop();
    }
    
    public int top() {
        return stack.peek();
    }
    
    public int getMin() {
        return min.peek();
    }
}