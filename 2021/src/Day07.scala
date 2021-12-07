object Day07 extends App {
  def median(values: Array[Int]): Int = {
    values.sorted.apply(values.length / 2)
  }

  def solve(values: Array[Int]): Int = {
    val med = median(values)
    values.map(x => Math.abs(x - med)).sum
  }

  val digits = scala.io.Source.fromFile("inputs/day07").mkString.strip.split(',').map(_.toInt)
  println(solve(digits))
}
