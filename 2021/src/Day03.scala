object Day03 extends App {
  def mostCommonBit(word: Array[Char]): Char = {
    word.groupBy(identity).toSeq.sortWith(_._1 > _._1).maxBy(_._2.size)._1
  }

  def leastCommonBit(word: Array[Char]): Char = {
    word.groupBy(identity).toSeq.sortWith(_._1 < _._1).minBy(_._2.size)._1
  }

  def rating(
      words: Array[Array[Char]]
  )(criteria: (Array[Char]) => Char): Array[Char] = {
    (0 to words(0).length - 1)
      .foldLeft(words)((words, i) =>
        words.filter(_(i) == criteria(words.map(_(i))))
      )(0)
  }

  def solve(words: Array[Array[Char]]): Int = {
    var gamma = words.transpose.map(mostCommonBit).mkString
    var epsilon = gamma.toCharArray.map('1' - _).mkString

    Integer.parseInt(gamma, 2) * Integer.parseInt(epsilon, 2)
  }

  def solve2(words: Array[Array[Char]]): Int = {
    val oxygen = Integer.parseInt(rating(words)(mostCommonBit).mkString, 2)
    val co2 = Integer.parseInt(rating(words)(leastCommonBit).mkString, 2)

    return oxygen * co2
  }

  val words = scala.io.Source
    .fromFile("inputs/day03")
    .mkString
    .split('\n')
    .map(_.toCharArray)

  println(solve2(words))
}
