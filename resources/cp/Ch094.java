public class Ch094 implements Cloneable{
    private int val;
    public Ch094(int val){
        this.val = val;
    }
    public static void main(String[] args){
        Ch094 obj1 = new Ch094(1);
        try{
            Ch094 obj2 = (Ch094) obj1.clone();
            System.out.println(obj2.val);
        }catch(Exception e){
            System.out.println(e);
        }
    }
}