# lem-in

## Objectives

The **lem-in** project is a Go program designed to simulate a digital ant farm. The goal is to find the quickest way to move `n` ants from a starting point (`##start`) to an endpoint (`##end`) across a colony consisting of rooms and tunnels. The project emphasizes pathfinding efficiency, error handling, and adherence to coding best practices.

## How It Works

1. **Ant Farm Simulation**:
    - The colony is represented by rooms connected through tunnels.
    - Ants start in the `##start` room and must reach the `##end` room.

2. **Pathfinding Requirements**:
    - Minimize the number of moves required for all ants to reach the `##end`.
    - Avoid traffic jams and ensure only one ant occupies a room at a time (except `##start` and `##end`).
    - Use the shortest path while considering traffic optimization.

3. **Input Description**:
    - Input is read from a file passed as an argument.
    - File structure:
        - `number_of_ants`: The total number of ants.
        - `the_rooms`: A list of rooms with coordinates (e.g., `RoomName x y`).
        - `the_links`: A list of tunnels connecting rooms (e.g., `RoomA-RoomB`).

4. **Output Format**:
    - The program outputs the file content followed by the moves of ants in the following format:
      ```
      Lx-y Lz-w Lr-o ...
      ```
      - `x, z, r`: Ant numbers.
      - `y, w, o`: Room names.

## Instructions

1. **Rooms**:
    - Defined as `RoomName coord_x coord_y` (e.g., `Room1 10 20`).
    - Names must not start with `L` or `#` and must not contain spaces.

2. **Tunnels**:
    - Represented as `RoomA-RoomB`.
    - Each tunnel connects exactly two rooms.

3. **Rules**:
    - Each room can only contain one ant at a time (except `##start` and `##end`).
    - Tunnels can only be used once per turn.
    - Input validation is crucial:
        - Detect missing `##start` or `##end`.
        - Handle invalid formats, duplicate rooms, or links to unknown rooms.

4. **Error Handling**:
    - Return an error message for invalid inputs:
        ```
        ERROR: invalid data format
        ```
  

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/jesee-kuya/lem-in.git
    cd lem-in
    ```

2. Compile and run the program with a file as input:

   ```bash
   go run . [INPUT-FILE]


### Examples

#### Example 1
```bash
$ go run . test0.txt
3
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5

L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0
```

#### Example 2
```bash
$ go run . test1.txt
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1

L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3
L3-1
```

## Implementation Details

- **Programming Language**: Go
- **Coding Standards**: Follow Go best practices.
- **Testing**: Include unit tests to ensure the program handles edge cases and invalid inputs correctly.

## Allowed Packages

Only Go's standard library is permitted.


## Contributors

This project exists thanks to all the people who contribute:

- [Jesee Kuya](https://github.com/jesee-kuya)

- [Joel Adero](https://github.com/Murzuqisah)

- [John Odhiambo](https://github.com/johneliud)


# Licence
This project is licenced under [MIT](https://github.com/jesee-kuya/lem-in/blob/main/LICENCE)