package Arrays;

import java.util.Arrays;

public class SortedSquare {
    public static void main( String[] args ) {
        int[] nums = new int[]{ 1 , 2 , 3 , 4 };
        log( firstSolution(nums) );
        log( secSolution(nums) );
    }
    // 🪴 O(n) time / O(n) space
    private static int[] secSolution( int[] nums ) {
        // generate zero array with the same length as nums
        int[] sortedSquare = new int[nums.length];
        int smallerValIdx = 0; // Track the smaller value 🦜
        int largerValIdx = sortedSquare.length - 1; // Track the larger value 🦜

        for (int i = sortedSquare.length - 1; i >= 0;  i--) { // ✅ Looping from the end of the array

            int smallerVal = Math.abs( nums[smallerValIdx] ); // 🌸 Absolute value
            int largerVal = Math.abs( nums[largerValIdx] ); // 🌸 Absolute value

            if ( smallerVal > largerVal ) {
                sortedSquare[i] = smallerVal * smallerVal;
                smallerValIdx += 1; // 🌱 Move to the next smaller value, ex: 1 => 2, 2 => 3
            } else {
                sortedSquare[i] = largerVal * largerVal;
                largerValIdx -= 1; // 🌱 Move to the next larger value, ex: 4 => 3, 3 => 2
            }
        }
        return sortedSquare;
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
