public class Ch07 implements Runnable {
    public static void main(String[] args){
        Runnable obj = new Ch07();
        ((Ch07) obj).test();
    }
    public void test() {
        Ch07.staticMethod();
        Ch07 demo = new Ch07();
        demo.instanceMethod();
        super.equals(null);
        this.run();
        ((Runnable) demo).run();
    }
    public static void staticMethod(){
        System.out.println(0);
    }
    private void instanceMethod() {
        System.out.println(1);
    }
    public void run(){}
}