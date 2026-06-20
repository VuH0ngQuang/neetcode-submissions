class Solution {
    public boolean isAnagram(String s, String t) {
        char[] str = s.toCharArray();
        char[] st = t.toCharArray();
        HashMap<Character,Integer> set = new HashMap<>();

        for (int i = 0; i < str.length; i++) {
            set.computeIfPresent(str[i], (k,v) -> v+1);
            set.computeIfAbsent(str[i], v -> 1);
        }

        for (int i =0; i < st.length; i++) {
            if (!set.containsKey(st[i])) {
                return false;
            }
            set.computeIfPresent(st[i], (k,v) -> v-1);
            if (set.get(st[i]) == 0) set.remove(st[i]);
        }

        return set.isEmpty();
    }
}
