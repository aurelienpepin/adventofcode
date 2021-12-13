object Day13 extends App {

  def solve(points: Set[(Int, Int)], verticalFold: Int, horizontalFold: Int): Int = {
    points.map {
      case (x, y) if x > verticalFold   => (verticalFold - (x - verticalFold), y)
      case (x, y) if y > horizontalFold => (x, horizontalFold - (y - horizontalFold))
      case (x, y)                       => (x, y)
    }.size
  }

  val input = scala.io.Source
    .fromFile("inputs/day13")
    .mkString
    .split("\n\n")

  val points = input(0).strip.split('\n').map(_.split(',')).map { case Array(x, y) => (x.toInt, y.toInt) }.toSet
  val instructions = input(1).strip.split('\n').map(_.split(' ').last.split('=')).map { case Array(c, d) => (c, d) }

  println(solve(points, 655, Int.MaxValue))
}
