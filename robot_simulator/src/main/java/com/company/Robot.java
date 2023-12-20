package com.company;

public class Robot {
    public String execute(String param) {
        String[] vars = param.split(" ");
        int counter = 0;
        String d = vars[2];
        int x = Integer.parseInt(vars[0]);
        int y = Integer.parseInt(vars[1]);
        String directions = vars[3];

        for (int i = 0; i < directions.length(); i++) {
            if (param.charAt(0) == '\n') {
                return null;
            } else if (directions.charAt(i) == 'L') {
                counter--;
            } else if (directions.charAt(i) == 'R') {
                counter++;
            }

            if (counter >= 0) {
                switch (counter % 4) {
                    case 0 -> d = "NORTH";
                    case 1 -> d = "EAST";
                    case 2 -> d = "SOUTH";
                    case 3 -> d = "WEST";
                }
            }

            if (counter < 0) {
                d = switch (counter % 4) {
                    case 0 -> "NORTH";
                    case -1 -> "WEST";
                    case -2 -> "SOUTH";
                    case -3 -> "EAST";
                    default -> d;
                };
            }

            if (directions.charAt(i) == 'A') {
                switch (d) {
                    case "NORTH" -> y++;
                    case "SOUTH" -> y--;
                    case "EAST" -> x++;
                    case "WEST" -> x--;
                }
            }
        }

        return x + " " + y + " " + d;
    }
}
