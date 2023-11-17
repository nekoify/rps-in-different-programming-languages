var readline = require('readline-sync');

console.log("Welcome to Rock Paper Scissors written in Javascript!")

function rps() {
    var answer = readline.question("Please select Rock, Paper or Scissors: ").toLowerCase()
    var res = checkAnswer(answer)
    if (res == false) return
    var aiAnswer = Math.floor(Math.random() * 3) + 1;
    var aiMove = parseAnswer(aiAnswer)
    console.log(`AI has picked: ${aiMove}`)
    console.log(calcWinner(answer, aiMove))
}

function checkAnswer(ans) {
    if (!(ans === "rock" || ans === "paper" || ans === "scissors")) {
        console.log("Please enter a valid move!")
        rps()
        return false
    }
}

function parseAnswer(ans) {
    if (ans == 1) {
        return "rock"
    } else if (ans == 2) {
        return "paper"
    } else if (ans == 3) {
        return "scissors"
    }
}

function calcWinner(user, ai) {
    if (
        (user === 'rock' && ai === 'scissors') ||
        (user === 'paper' && ai === 'rock') ||
        (user === 'scissors' && ai === 'paper')
      ) {
        return 'You win!';
      } else if (
        (ai === 'rock' && user === 'scissors') ||
        (ai === 'paper' && user === 'rock') ||
        (ai === 'scissors' && user === 'paper')
      ) {
        return 'The AI wins!';
      } else {
        return "It's a tie!";
      }
}

rps()