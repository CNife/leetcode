import java.util.ArrayList;

public class EncodeDecodeTinyUrl {
    private ArrayList<String> data = new ArrayList<>();

    private String encode(String longUrl) {
        String shortUrl = Integer.toString(data.size(), 36);
        data.add(longUrl);
        return shortUrl;
    }

    private String decode(String shortUrl) {
        int index = Integer.parseInt(shortUrl, 36);
        return data.get(index);
    }
}
