public class Ch06 {
    public static int staticVar;
    public int instanceVar;
    public static void main(String[] args){
        int x = 12345;
        Ch06 obj = new Ch06();
        obj.staticVar = x;
        x = obj.staticVar;
        obj.instanceVar = x;
        x = obj.instanceVar;
        Object o = obj;
        if (o instanceof Ch06) {
            obj = (Ch06) o;
            System.out.println(obj.instanceVar);
        }
    }
}