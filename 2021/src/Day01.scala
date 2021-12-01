object Day01 extends App {
  def solve(depths: Array[Int]): Int = {
    depths.zipAll(depths.drop(1), Int.MaxValue, Int.MinValue)
      .filter((a, b) => a < b)
      .length
  }

  def solve2(depths: Array[Int]): Int = {
    var currentSum: Int = depths.take(3).sum
    var increasing: Int = 0

    for (i <- 3 to depths.length - 1) {
      val newSum = currentSum + depths(i) - depths(i - 3)
      if (newSum > currentSum)
        increasing += 1

      currentSum = newSum
    }

    increasing
  }

  val input = scala.io.Source.fromFile("inputs/day01")
  var integers = (try input.mkString finally input.close())
    .split('\n')
    .map(_.toInt)

  // println(solve(integers))
  println(solve2(integers))
}
