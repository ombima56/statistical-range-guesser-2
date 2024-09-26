# Guess-It-2

## Description

`guess-it-2` is a Go-based application designed to perform a sequence of numbers as input and predict a range where the next number in the sequence is expected to fall. Each number provided represents a point on a graph where the x-axis is the line number (starting from 0) and the y-axis is the number itself.

markdown


- **student/calculation/calculation.go**: Contains the statistical functions for calculating the linearRegression, average, variance, standard deviation.
- **student/main.go**: The main application file that reads input, processes it, and calculates the range prediction based on the provided data.
- **student/script.sh**: A shell script to run the Go program.

## Installation

1. Clone the repository:
```bash
   git clone https://learn.zone01kisumu.ke/git/hiombima/guess-it-2.git
```
2. Navigate to the student directory:

```bash
    cd guess-it-1/student
```
## How it works

1. **Input:** The program receives numbers as standard input, one per line.
 - Example
```go
    189
    113
    121
    114
    145
    110
```

2. **Graph Representation:**
    - The x-axis represents the line number (0, 1, 2, 3, ...) and the y-axis represents the input values.
    - For each new number, the program calculates a predicted range for the next number based on previous values.


## Usage

- Ensure you have Go installed. If not, download and install it from the Go website.
- Navigate to the student directory where main.go is located.
- Download the zip file provided, the program will output the calculated range based on the provided data.
- Run the program using the provided shell script:

### How to test this function:
To test the student program, these commands should be ran to have the
dependencies needed and to start the webpage on the port 3000:

```console
docker compose up
```

or

```sh
npm install
```
after installing you can run in the terminal:
```sh
npm start
```

After opening your browser of preference in the port
[3000](http://localhost:3000/), if you try clicking on any of the `Test Data`
buttons, you will notice that in the Dev Tool/ Console there is a message which
tells you that you need another guesser besides the student.

Adding a guesser is simple. You need to add in the URL a guesser, in other
words, the name of one of the files present in the `ai/` folder:

```console
?guesser=<name_of_guesser>
```

For example:

```console
?guesser=big-range
```

After that, choose data 4 and 5 set to test. After that you can click
`Quick` to skip the waiting and be presented with the results.

Since the website uses big data sets, we advise you to clear the displays
clicking on the `Clean` button after each test.

## Contribution

Follow these steps if you would like to contribute to this project or if you've discovered a bug.

1. **Fork the repository:** Click the "Fork" button at the top-right corner of the repository page.
2. **Clone your fork:** Clone the forked repository to your local machine using:

```bas
git clone https://github.com/<your-username>/<your-forked-repo>.git
```
3. **Create a new branch:** Create a branch for your feature or bug fix.

```bash
git checkout -b feature/your-feature-name
```
4. **Make your changes:** Implement your feature or bug fix, ensuring code quality.
5. **Run tests:** Make sure your changes do not break existing functionality. Write tests where applicable.
6. **Commit your changes:** Follow good commit message practices (e.g., `git commit -m "Add feature X`").
7. **Push your branch:** Push your changes to your fork.
8. **Create a Pull Request:** Go to the original repository, and click "Compare & pull request." Describe your changes in detail.

## License
This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.