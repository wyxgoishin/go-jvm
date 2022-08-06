import java.util.Random;
public class Ch081 {
//     private static Random rand = new Random(12345);

    public static void main(String[] args){
        System.out.println("Hello world");
        for(String arg : args){
            System.out.println(arg);
        }
        int[] nums = new int[]{1, 2, -3, -40, 2, 121, -2, 2, 2, 3, -3, 3, 3, 3, 90, -32, -2};
        quickSort(nums, 0, nums.length - 1);
        printArray(nums);
    }

    public static void printArray(int[] nums){
        for(int i = 0, j = nums.length; i < j; ++i){
            System.out.println(nums[i]);
        }
    }

    public static void quickSort(int[] nums, int left, int right){
        if(left < right){
            int[] pivots = partition(nums, left, right);
            quickSort(nums, left, pivots[0] - 1);
            quickSort(nums, pivots[0] + 1, right);
        }
    }

    public static int[] partition(int[] nums, int left, int right){
//         int randLoc = rand.nextInt(right - left + 1) + left;
//         swap(nums, left, randLoc);
        int[] pivots = new int[]{right, right};
        int val = nums[right];
        for(int i = right - 1; i >= left; --i){
            if(nums[i] > val){
                swap(nums, pivots[0] - 1, i);
                swap(nums, pivots[1], pivots[0] - 1);
                --pivots[0];
                --pivots[1];
            }else if(nums[i] == val){
                swap(nums, pivots[0] - 1, i);
                --pivots[0];
            }
        }
        return pivots;
    }

    public static void swap(int[] nums, int left, int right){
        if(left != right){
            nums[left] ^= nums[right];
            nums[right] ^= nums[left];
            nums[left] ^= nums[right];
        }
    }
}