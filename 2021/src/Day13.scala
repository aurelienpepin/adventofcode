object Day13 extends App {

  def solve(points: Set[(Int, Int)], verticalFold: Int, horizontalFold: Int): Int = {
    points.map {
      case (x, y) if x > verticalFold   => (verticalFold - (x - verticalFold), y)
      case (x, y) if y > horizontalFold => (x, horizontalFold - (y - horizontalFold))
      case (x, y)                       => (x, y)
    }.size
  }

  def solve2(points: Set[(Int, Int)], instructions: Array[(Int, String)]): Unit = {
    val endPoints = instructions.foldLeft(points)((ps, instruction) =>
      val largestX = ps.maxBy(_._1)._1
      val largestY = ps.maxBy(_._2)._2

      instruction match {
        case (d, "x") if d >= largestX / 2 =>
          ps.map {
            case (x, y) if x > d => (d - (x - d), y)
            case (x, y)          => (x, y)
          }
        case (d, "x") =>
          ps.map {
            case (x, y) if x < d  => (d - (d - x), y)
            case (x, y)           => (x - d - 1, y)
          }
        case (d, "y") if d >= largestY / 2 =>
          ps.map {
            case (x, y) if y > d => (x, d - (y - d))
            case (x, y) => (x, y)
          }
        case (d, "y") =>
          ps.map {
            case (x, y) if y < d => (x, d - (d - y))
            case (x, y) => (x, y - d - 1)
          }
        case _ => throw new IllegalArgumentException
      }
    )

    for (i <- 0 to endPoints.maxBy(_._2)._2) {
      for (j <- 0 to endPoints.maxBy(_._1)._1) {
        if (endPoints.contains((j, i))) {
          print("*")
        } else {
          print(" ")
        }
      }
      println("")
    }
  }

  val input = scala.io.Source
    .fromFile("inputs/day13")
    .mkString
    .split("\n\n")

  val points = input(0).strip.split('\n').map(_.split(',')).map { case Array(x, y) => (x.toInt, y.toInt) }.toSet
  val instructions =
    input(1).strip.split('\n').map(_.split(' ').last.split('=')).map { case Array(c, d) => (d.toInt, c) }.toArray

  // println(solve(points, 655, Int.MaxValue))
  solve2(points, instructions)
}
