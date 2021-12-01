object Day01 extends App {
  def solve(depths: Array[Int]): Int = {
    depths.zipAll(depths.drop(1), Int.MaxValue, Int.MinValue)
      .filter((a, b) => a < b)
      .length
  }

  val input = scala.io.Source.fromFile("inputs/day01")
  var integers = (try input.mkString finally input.close())
    .split('\n')
    .map(_.toInt)

  println(solve(integers))
}
