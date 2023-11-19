import random
print("Welcome to Rock Paper Scissors written in Python!")


def rps():
   answer = input("Please select Rock, Paper or Scissors: ").lower()
   res = checkAnswer(answer)
   if res == False:
       return
   aiAnswer = random.randint(1, 3)
   aiMove = parseAnswer(aiAnswer)
   print(f"Ai has picked {aiMove}")
   print(findWinner(answer, aiMove))

def checkAnswer(ans):
    if not (ans == "rock" or ans == "paper" or ans == "scissors"):
        print("Please select a valid option!")
        rps()
        return False

def parseAnswer(ans):
    if ans == 1:
        return "rock"
    elif ans == 2:
        return "paper"
    elif ans == 3:
        return "scissors"

def findWinner(user, ai):
    if (
        (user == 'rock' and ai == 'scissors') or
        (user == 'paper' and ai == 'rock') or
        (user == 'scissors' and ai == 'paper')
    ):
        return 'You win!'
    elif (
        (ai == 'rock' and user == 'scissors') or
        (ai == 'paper' and user == 'rock') or
        (ai == 'scissors' and user == 'paper')
    ):
        return 'The AI wins!'
    else:
        return "It's a tie!"

rps()