package com.company;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class RobotTest {

    @Test
    public void testInput1() {
        Robot robot = new Robot();
        assertEquals("0 1 NORTH", robot.execute("0 0 NORTH A"));
    }
    @Test
    public void testInput2() {
        Robot robot = new Robot();
        assertEquals("9 4 WEST", robot.execute("7 3 NORTH RAALAL"));
    }

    @Test
    public void testInput3() {
        Robot robot = new Robot();
        assertEquals("0 0 WEST", robot.execute("0 0 NORTH L"));
    }

    @Test
    public void testInput4() {
        Robot robot = new Robot();
        assertEquals("0 0 EAST", robot.execute("0 0 NORTH R"));
    }

    @Test
    public void testInput5() {
        Robot robot = new Robot();
        assertEquals("0 0 NORTH", robot.execute("0 0 NORTH AAAALAAAALAAAALAAAAL"));
    }
}