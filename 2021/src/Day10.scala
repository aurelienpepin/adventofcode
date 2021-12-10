import scala.collection.mutable.Stack

object Day10 extends App {
  def solve(line: String): Int = {
    // println("------")
    val stack = Stack[Char]()

    for (c <- line.toCharArray) {
      if (isOpening(c))
        stack.push(c)
      else {
        if (stack.isEmpty || stack.head != matchingCharacter(c)) {
          // println(c + " " + stack)
          // println("illegal " + c)
          return score(c)
        } else
          stack.pop()
      }

      // println(c + " " + stack)
    }

    0
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
  val input = scala.io.Source.fromFile("inputs/day10").mkString.split('\n').map(_.strip)

  println(input.map(solve(_)).sum)
}
