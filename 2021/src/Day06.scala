import scala.collection.mutable

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

  def memoize[I, O](f: I => O): I => O = new mutable.HashMap[I, O]() {
    override def apply(key: I) = getOrElseUpdate(key, f(key))
  }

  val fibo7_9: Int => BigInt = memoize { (n) =>
    if (n < 0)
      0
    else if (n < 9)
      1
    else
      fibo7_9(n - 7) + fibo7_9(n - 9)
  }

  def solve(timers: Array[Int], days: Int): Int = {
    (0 until days)
      .foldLeft(timers)((ts, _) => nextCycle(ts))
      .length
  }

  def solve2(timers: Array[Int], days: Int): BigInt = {
    timers.map(t => fibo7_9(days + 8 - t)).reduce(_ + _)
  }

  val digits = scala.io.Source.fromFile("inputs/day06").mkString.strip.split(',').map(_.toInt)
  println(solve2(digits, 256))
}
