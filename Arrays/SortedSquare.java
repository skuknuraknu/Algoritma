package Arrays;

import java.util.Arrays;

public class SortedSquare {
    public static void main( String[] args ) {
        int[] nums = new int[]{ -5, -3, -2, 0, 1, 4, 5 };
        log( firstSolution(nums) );
    }

    // 🪴 O(nLogn) time / O(n) space
    private static int[] firstSolution( int[] nums ) {
        // generate zero array with the same length as nums
        int[] sortedSquare = new int[nums.length];

        for (int i = 0; i < sortedSquare.length; i++) {
            int value = nums[i];
            sortedSquare[i] = value * value;
        }
        Arrays.sort(sortedSquare); // this dude here make it O(nLogn) time complexity 🌸
        return sortedSquare;
    }
    private static void log(Object o){
        if ( o instanceof int[] ) {
            System.out.println(Arrays.toString((int[]) o));
        } else {
            System.out.println(o);
        }
    }
}
