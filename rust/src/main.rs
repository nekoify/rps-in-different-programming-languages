use std::io;
use rand::Rng;
fn main() {
    println!("Welcome to Rock Paper Scissors written in Rust!");
    rps();
}

fn rps() {
    let mut rng = rand::thread_rng();
    println!("Please select Rock, Paper or Scissors:");
    let mut user_input = String::new();
    io::stdin().read_line(&mut user_input).unwrap();
    let answer = user_input.to_lowercase().trim().to_owned();
    let res = check_answer(answer.clone());
    if res == false {return}
    let ai_answer = rng.gen_range(1..4);
    let ai_move = parse_answer(ai_answer);
    println!("Ai has picked {ai_move}"); 
    println!("{}", find_winner(&answer, ai_move))
}

fn check_answer(ans: String) -> bool {
    if !(ans == "rock" || ans == "paper" || ans == "scissors") {
    println!("Please select a valid move!");
    rps();
    return false
    }
    return true
}

fn parse_answer(ans: i32) -> &'static str {
    if ans == 1 {
        return "rock"
    } else if ans == 2 {
        return "paper"
    } else if ans == 3 {
        return "scissors"
    }
    return "?"
}

fn find_winner(user: &str, ai: &str) -> &'static str {
    if (user == "rock" && ai == "scissors") ||
        (user == "paper" && ai == "rock") ||
        (user == "scissors" && ai == "paper") {
        return "You win!";
    } else if (ai == "rock" && user == "scissors") ||
        (ai == "paper" && user == "rock") ||
        (ai == "scissors" && user == "paper") {
        return "The AI wins!";
    } else {
        return "It's a tie!";
    }
}