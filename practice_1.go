package main

import ("fmt")

var p1Score, p2Score int
func main(){
  gameOver := false
  board := [9]int{0,0,0,0,0,0,0,0,0}
  turnNumber := 1

  for gameOver != true{
    presentBoard(board)
    player := 0
    if turnNumber % 2 == 1{
      fmt.Println("Player 1's turn")
      player = 1
    } else {
      fmt.Println("Player 2's turn")
      player = 2
    }

    currentMove := askForPlay()
      // Quit with '9'
      if currentMove == 9 {
        return
      }
    board = executePlayerMove(currentMove, player, board)

    result := checkForWin(board)
    if result > 0{
      fmt.Printf("Player %d wins!\n\n", result)
      fmt.Println("Score: you", p1Score, "me", p2Score)
        fmt.Println("\nLet's play again.")
      gameOver = true
    } else if turnNumber == 9 {
      // Tie game example: 0 2 1 3 4 7 5 8 6 
      fmt.Printf("Tie game!\n\n")
      fmt.Println("\nLet's play again.")
      gameOver = true
    } else {
      turnNumber++
    }
  }

}

func askForPlay() int{
  fmt.Println("Select a move")
  var moveInt int
  fmt.Scan(&moveInt)
  // fmt.Println("moveInt is", moveInt)
  return moveInt
}

func executePlayerMove(moveInt int, player int, b [9]int) [9]int {

  // Check for occupied spaces
  if b[moveInt] != 0 {
    fmt.Println("Please pick an unoccupied space.")
    moveInt = askForPlay()
    b = executePlayerMove(moveInt, player, b)
  } else {
    if player == 1 {
      b[moveInt] = 1
    } else if player == 2 {
      b[moveInt] = 10
    }
  }

  // Check for out-of-bounds
  for moveInt > 9 {
      fmt.Println("Please enter a number under 10.")
      moveInt = askForPlay()
  }

  if player == 1{
    b[moveInt] = 1
  }else if player == 2{
    b[moveInt] = 10
  }
  return b
}

func presentBoard(b [9]int) {
  for i, v := range b {
    if v == 0{
      // empty space. Display number
      fmt.Printf("%d", i)
    } else if v == 1{
      fmt.Printf("😂")
    } else if v == 10{
      fmt.Printf("😱")
    }
    // And now the decorations
    if i> 0 && (i+1) % 3 == 0{
      fmt.Printf("\n")
    } else {
      fmt.Printf(" || ")
    }
  }
}

func checkForWin(b [9]int) int {
  // re-calculate sums Array
  sums := [8] int {0,0,0,0,0,0,0,0}
  // for _, v := range b[0:2] { sums[7] += v }
  // for _, v := range b[3:5] { sums[6] += v }
  // for _, v := range b[6:8] { sums[5] += v }

  sums[0] = b[2]+b[4]+b[6]
  sums[1] = b[0]+b[3]+b[6]
  sums[2] = b[1]+b[4]+b[7]
  sums[3] = b[2]+b[5]+b[8]
  sums[4] = b[0]+b[4]+b[8]
  sums[5] = b[6]+b[7]+b[8]
  sums[6] = b[3]+b[4]+b[5]
  sums[7] = b[0]+b[1]+b[2]
  for _, v := range sums {
    if v == 3{
      fmt.Println("player2 win!")
      p2Score++
      return 1
    } else if v == 30{
      fmt.Println("player1 win!")
      p1Score++
      return 2
    }
  }
  return 0
}

