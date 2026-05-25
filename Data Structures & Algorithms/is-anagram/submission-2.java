class Solution {
    public boolean isAnagram(String s, String t) {
        int lengths = s.length();
        int lengtht = t.length();
        char ch;
        int val;
        if (lengtht != lengths) return false;
        HashMap<Character, Integer> map = new HashMap<>();
        for(int i = 0; i < lengths; i++) {
            ch = s.charAt(i);
            map.computeIfPresent(ch, (k,v) -> v + 1);
            map.putIfAbsent(ch,1);
        }
        for(int i = 0; i < lengths; i++) {
            ch = t.charAt(i);
            if (map.containsKey(ch)) {
                val = map.get(ch);
                if (val > 0) {
                    map.put(ch, val - 1);
                } else return false;
            } else return false;
        }
        return true;
    }
}
