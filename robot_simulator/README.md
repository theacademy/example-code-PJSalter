# 🧩 robot-simulator 🧩

## 🧑‍💻 Coding Example - Simulate a robot 🦾

The objective is to create a program that emulates the movements of a robot on a hypothetical infinite grid. The robot can perform three types of movements:

```
- Turn right
- Turn left
- Advance
```

### Robot Behavior:

- The robot is positioned on a hypothetical infinite grid, facing specific cardinal directions (north, east, south, or west) at a set of `{x,y}` coordinates. For example, `{3,8}`, where the coordinates increase towards the north and east directions.

- The program includes a `Robot` class with a method `execute` that accepts a series of instructions. This method calculates the new position and direction of the robot after executing these instructions.

## Example

```
For instance, the letter-string "RAALAL" translates to:

- Turn right
- Advance twice
- Turn left
- Advance once
- Turn left yet again
```

If a robot starts at `{7, 3}` facing north, executing this sequence of instructions should position it at `{9, 4}` facing west.

## Inputs and Outputs

The `execute` method of the `Robot` class takes a `String` argument in the format `X Y BEARING COMMANDS`. It returns a one-liner `String` with the format `X Y BEARING`.

For instance:
- `7 3 NORTH RAALAL` as the input should return `9 4 WEST`.

## Tests

### JUnit 5 Tests

Tests are written in JUnit 5 to validate the `execute` method functionality. To execute these tests:

1. Create a test class for the `RobotTest` class.
2. Use assertions to validate the expected output against the actual output from the `execute` method for various test cases.

```java
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class RobotTest {
    @Test
    public void testInput2() {
        Robot robot = new Robot();
        assertEquals("9 4 WEST", robot.execute("7 3 NORTH RAALAL"));
    }
}
```
