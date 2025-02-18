package Arrays;

import java.util.*;

public class Tournament {
    public static void main( String args[] ) {
        HashMap<String, List<List<String>>> competition = new HashMap<>();

        // Generate data
        List<List<String>> nestedList = new ArrayList<>();
        nestedList.add(Arrays.asList("html", "c#"));
        nestedList.add(Arrays.asList("go", "html"));
        nestedList.add(Arrays.asList("python", "go"));
        competition.put("competitions", nestedList);
        int[] results = new int[]{ 1, 0, 1 };

        log("Tournament Start!");
        tournamentWinner( competition, results );
    }
    private static void tournamentWinner( HashMap<String, List<List<String>>> competitions, int[] results ) {
        final int HOME_TEAM_WINS = 1;
        HashMap<String, Integer> currentBestTeam = new HashMap<>();
        List<List<String>> listCompetition = competitions.get("competitions");


        int highestScore = 0;
        String winnerName = "";

        for (int e = 0; e < results.length; e++) {
            List<String> match = listCompetition.get(e);
            String winner = results[e] == HOME_TEAM_WINS ? match.get(0) : match.get(1);
            int count = currentBestTeam.getOrDefault(winner, 0) + 3;

            currentBestTeam.put(winner, count);

            if ( count > highestScore ) {
                highestScore = count;
                winnerName = winner;
            }
        }
        log( winnerName );
    }
    private static void log(Object o){
        System.out.println(o);
    }
}
