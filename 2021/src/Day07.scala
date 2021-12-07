object Day07 extends App {
  def median(values: Array[Int]): Int = {
    values.sorted.apply(values.length / 2)
  }

  val fuel1 = (values: Array[Int], position: Int) => values.map(v => Math.abs(v - position)).sum
  val fuel2 = (values: Array[Int], position: Int) =>
    values.map(v => Math.abs(v - position) * (Math.abs(v - position) + 1)).sum / 2

  def solve(values: Array[Int]): Int = fuel1(values, median(values))

  def solve2(values: Array[Int], min: Int, max: Int): Int = {
    if (max < min) {
      Int.MaxValue
    } else {
      val mid = min + (max - min) / 2;

      val left = fuel2(values, mid - 1)
      val here = fuel2(values, mid)
      val right = fuel2(values, mid + 1)

      if (here < left && here < right)
        here
      else if (left < here)
        solve2(values, min, mid - 1)
      else
        solve2(values, mid + 1, max)
    }
  }

  val digits = scala.io.Source.fromFile("inputs/day07").mkString.strip.split(',').map(_.toInt)
  println(solve2(digits, digits.min, digits.max))
}
