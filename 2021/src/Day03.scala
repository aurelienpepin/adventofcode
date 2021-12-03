object Day03 extends App {
  val toInt: (Array[Char]) => Int = word => Integer.parseInt(word.mkString, 2)

  def mostCommonBit(word: Array[Char]): Char = {
    word.groupBy(identity).toSeq.sortWith(_._1 > _._1).maxBy(_._2.size)._1
  }

  def leastCommonBit(word: Array[Char]): Char = {
    word.groupBy(identity).toSeq.sortWith(_._1 < _._1).minBy(_._2.size)._1
  }

  def rating(words: Array[Array[Char]])(bit: (Array[Char]) => Char): Array[Char] = {
    (0 to words(0).length - 1)
      .foldLeft(words)((ws, i) => ws.filter(_(i) == bit(ws.map(_(i)))))(0)
  }

  def solve(words: Array[Array[Char]]): Int = {
    val gamma = words.transpose.map(mostCommonBit)
    val epsilon = gamma.map('1' - _).map(_.toChar)

    toInt(gamma) * toInt(epsilon)
  }

  def solve2(words: Array[Array[Char]]): Int = {
    val oxygen = rating(words)(mostCommonBit)
    val co2 = rating(words)(leastCommonBit)

    toInt(oxygen) * toInt(co2)
  }

  val words = scala.io.Source
    .fromFile("inputs/day03")
    .mkString
    .split('\n')
    .map(_.toCharArray)

  println(solve2(words))
}
