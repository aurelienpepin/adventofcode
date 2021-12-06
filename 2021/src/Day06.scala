object Day06 extends App {
  def nextCycle(timers: Array[Int]): Array[Int] = {
    timers
      .flatMap(timer => {
        if (timer == 0) {
          Array(6, 8)
        } else {
          Array(timer - 1)
        }
      })
  }

  def solve(timers: Array[Int], days: Int): Int = {
    (0 until days)
      .foldLeft(timers)((ts, _) => nextCycle(ts))
      .length
  }

  val digits = scala.io.Source.fromFile("inputs/day06").mkString.strip.split(',').map(_.toInt)
  println(solve(digits, 80))
}
