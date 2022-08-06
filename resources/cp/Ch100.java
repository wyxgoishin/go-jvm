public class Ch100 {
    public static void main(String[] args){
        try{
            new Ch100().test();
        }catch(Exception e){
            System.out.println(e.getMessage());
        }
    }

    public void test() throws Exception {
        // test1(); //ToDo: implement it
        // test2(); //ToDo: implement it
        throw new Exception();
    }

    public void test1(){
        int i = 10;
        try{
            int j = i / 0;
        }catch(Exception e){
            System.out.println(e.toString());
        }
    }

    public void test2(){
        try{
            subTest2();
        }catch(Exception e){
            System.out.println(e.toString());
        }
    }

    public int subTest2(){
        int[] arr = new int[10];
        return arr[11];
    }
}