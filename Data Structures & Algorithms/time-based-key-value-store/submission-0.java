class TimeMap {
    
    HashMap<String, List<Pair<Integer, String>>> map;
    
    public TimeMap() {
        map = new HashMap<>();
    }
    
    public void set(String key, String value, int timestamp) {
        map.putIfAbsent(key, new ArrayList<>());
        map.get(key).add(new Pair<>(timestamp, value)); 
    }
    
    public String get(String key, int timestamp) {
        List<Pair<Integer, String>> tmp = map.get(key);

        if (tmp == null) {
            return "";
        }

        int left = 0;
        int right = tmp.size() - 1;
        String result = "";

        while (left <= right) {
            int mid = left + (right - left) / 2;

            int time = tmp.get(mid).getKey();

            if (time <= timestamp) {
                result = tmp.get(mid).getValue();
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return result;
    }
}

class Pair<K, V> {
    private K key;
    private V value;

    Pair(K key, V value) {
        this.key = key;
        this.value = value;
    }

    public K getKey() {
        return key;
    }

    public V getValue() {
        return value;
    }

    public void setKey(K key) {
        this.key = key;
    }

    public void setValue(V value) {
        this.value = value;
    }
}
