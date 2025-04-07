import java.util.Random;

public class GoGame {
    private static Random random = new Random();

    public static void main(String[] args) {
        int numberToGuess = random.nextInt(100);
        System.out.println("I'm thinking of a number between 0 and 99. Can you guess it?");
        
        int userGuess;
        do {
            userGuess = inputUserGuess(numberToGuess);
            
            if (userGuess == numberToGuess) {
                System.out.println("Congratulations! You guessed the correct number.");
                break;
            } else if (userGuess < numberToGuess) {
                System.out.println("Too low, try again.");
            } else {
                System.out.println("Too high, try again.");
            }
        } while (true);
    }

    private static int inputUserGuess(int targetNumber) {
        Scanner scanner = new Scanner(System.in);
        int guess;

        do {
            System.out.print("Enter your guess: ");
            guess = scanner.nextInt();
            
            if (guess < 0 || guess > 99) {
                System.out.println("Invalid input. Try again.");
            } else if (guess == targetNumber) {
                break;
            }
        } while (true);

        return guess;
    }
}
