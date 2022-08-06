public class Ch072 {
    private static int staticVar1 = 10;
    private int var2;
    public Ch072(int val){
        var2 = val;
    }

    public static void main(String[] args){
        Ch072 obj = new Ch072(19);
        System.out.println(obj.var2);
    }
}