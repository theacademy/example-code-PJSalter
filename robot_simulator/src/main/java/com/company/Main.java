package com.company;

public class Main {
    public static void main(String[] args) {

        Robot robot = new Robot();

        // Example test cases
        System.out.println(robot.execute("0 0 NORTH A")); // Output: 0 1 NORTH
        System.out.println(robot.execute("7 3 NORTH RAALAL")); // Output: 9 4 WEST
        System.out.println(robot.execute("0 0 NORTH L")); // Output: 0 0 WEST
        System.out.println(robot.execute("0 0 NORTH R")); // Output: 0 0 EAST
        System.out.println(robot.execute("0 0 NORTH AAAALAAAALAAAALAAAAL")); // Output: 0 0 NORTH

    }
}