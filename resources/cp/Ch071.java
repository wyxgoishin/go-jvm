import java.util.Map;
import java.util.HashMap;
public class Ch071 {
    public static void main(String[] args){
        int x = solution(10);
        System.out.println(x);
    }
    private static int solution(int x){
        if(x <= 1){
            return 1;
        }
        return solution(x - 1) + solution(x - 2);
    }
}