class Solution {
    public int[] topKFrequent(int[] nums, int k) {

        HashMap<Integer, Integer> map = new HashMap<>();

        for (int num : nums) {
            map.put(num, map.getOrDefault(num, 0) + 1);
        }

        List<Integer> sortedKey = map.keySet()
                .stream()
                .sorted((a, b) -> map.get(b) - map.get(a))
                .limit(k)
                .collect(Collectors.toList());

        int[] ans = new int[k];

        for (int i = 0; i < k; i++) {
            ans[i] = sortedKey.get(i);
        }

        return ans;
    }
}