object Day03 extends App {
  def solve(words: Array[Array[Char]]): Int = {
    var gamma = words.transpose
      .map(_.groupBy(identity).maxBy(_._2.size)._1)
      .mkString

    var epsilon = gamma.toCharArray.map('1' - _).mkString
    Integer.parseInt(gamma, 2) * Integer.parseInt(epsilon, 2)
  }

  val words = scala.io.Source
    .fromFile("inputs/day03")
    .mkString
    .split('\n')
    .map(_.toCharArray)

  println(solve(words))
}
