package Arrays;

import java.util.Arrays;
import java.util.HashMap;

public class Subsequence {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 7, 5, 10, 8};
        int[] sub = {3, 7, 11};
        log(isSubsequence(arr, sub));
    }
    public static boolean isSubsequence(int[] arr, int[] sub) {
        final int n = arr.length;
        int seqIdx  = 0;
        for ( int x = 0; x < n; x++ ) {
            if ( seqIdx == sub.length ) {
                return true;
            }
            if ( arr[x] == sub[seqIdx] ){
                seqIdx++;
            }
        }
        return false;
    }
    private static void log(Object o){
        if ( o instanceof int[] ) {
            System.out.println(Arrays.toString((int[]) o));
        } else {
            System.out.println(o);
        }
    }
}
