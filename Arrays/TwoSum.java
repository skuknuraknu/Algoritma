package Arrays;

import java.util.Arrays;
import java.util.HashMap;

public class TwoSum {
    public static void main( String[] args) {
        int[] nums = new int[]{ -5, 6, 7, 2, -1, 5, 2 };
        twoSum(nums, 7);
    }

    // O(n) solution
    private static void twoSum( int[] nums, int target ){
        // Menggunakan hashmap untuk menyimpan nilai yang sudah di cek.
        HashMap<Integer, Boolean> hashNum = new HashMap<Integer, Boolean>();
        log(nums);
        log("Target: "+ target);

        // Looping sebanyak jumlah elemen pada array .
        for ( int e : nums ) {
            /*** PotentialSum: hasil pengurangan target dengan elemen array
             * yang dimana jika hasil pengurangan tersebut ada pada hashmap
             * maka akan mengembalikan nilai tersebut.
             * */
            final int potentialSum = target - e;
            if ( hashNum.containsKey(potentialSum)) {
                log(String.format("%s, %s", potentialSum, e));
                return;
            }
            // Menyimpan nilai elemen array ke dalam hashmap jika nilai tersebut belum ada.
            hashNum.put(e, true);
        }
        // Jika target tidak ditemukan.
        log("[]");
        return;
    }
    private static void log(Object o){
        if ( o instanceof int[] ) {
            System.out.println(Arrays.toString((int[]) o));
        } else {
            System.out.println(o);
        }
    }
}
