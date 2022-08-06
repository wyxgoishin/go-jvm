import java.util.List;
import java.util.ArrayList;
public class Ch095 {
    public static void main(String[] args){
       List<Integer> list = new ArrayList<>();
       Integer i1 = 10;
       Integer i2 = 10;
       System.out.println(i1 == i2);
       System.out.println(i1.equals(i2));
       list.add(i1);
       list.add(100);
       System.out.println(list);
    }
}