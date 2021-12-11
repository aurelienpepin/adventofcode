import scala.collection.mutable.Stack

object Day10 extends App {
  def solve(line: String): Int = {
    val stack = Stack[Char]()

    for (c <- line.toCharArray) {
      if (isOpening(c))
        stack.push(c)
      else {
        if (stack.isEmpty || stack.head != matchingCharacter(c)) {
          return score(c)
        } else
          stack.pop()
      }
    }

    0
  }

  def solve2(line: String): Option[BigInt] = {
    val stack = Stack[Char]()

    for (c <- line.toCharArray) {
      if (isOpening(c))
        stack.push(c)
      else {
        if (stack.isEmpty || stack.head != matchingCharacter(c)) {
          return None // no point for corrupted lines
        } else
          stack.pop()
      }
    }

    var totalScore: BigInt = 0
    while (!stack.isEmpty) {
      totalScore = totalScore * 5 + score2(stack.head)
      stack.pop()
    }

    Some(totalScore)
  }

  val isOpening = (c) => Set('(', '[', '{', '<').contains(c)
  var matchingCharacter = (c: Char) =>
    c match {
      case ')' => '('
      case '}' => '{'
      case ']' => '['
      case '>' => '<'
    }
  val score = Map((')' -> 3), (']' -> 57), ('}' -> 1197), ('>' -> 25137))
  val score2 = Map(('(' -> 1), ('[' -> 2), ('{' -> 3), ('<' -> 4))
  val input = scala.io.Source.fromFile("inputs/day10").mkString.split('\n').map(_.strip)

  // println(input.map(solve(_)).sum)

  val res = input.map(solve2(_)).flatten.sorted
  println(res(res.length / 2))
}
