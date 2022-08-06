public class Ch092 {
    public static void main(String[] args){
        int[] a = new int[]{1, 2};
        int[] b = new int[2];
        System.arraycopy(a, 0, b, 0, 2);
        System.out.println(b[1]);
    }
}