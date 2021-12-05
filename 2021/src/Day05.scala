object Day05 extends App {
  case class Point(x: Int, y: Int)

  case class Line(p1: Point, p2: Point) {
    def allPoints(): Iterable[Point] = {
      if (p1.x == p2.x) { // horizontal
        ((p1.y min p2.y) to (p1.y max p2.y)).map(new Point(p1.x, _))
      } else if (p1.y == p2.y) { // vertical
        ((p1.x min p2.x) to (p1.x max p2.x)).map(new Point(_, p1.y))
      } else { // diagonal, not available
        Array.empty[Point]
      }
    }
  }

  def solve(lines: Array[Line]): Int = {
    val freqByPoint: Map[Point, Int] = lines.foldLeft(Map.empty)((freqs, line) => {
      line.allPoints().foldLeft(freqs)((coll, point) => coll + (point -> (coll.getOrElse(point, 0) + 1)))
    })

    freqByPoint.values.filter(_ >= 2).size
  }

  val toPoint: (String) => Point = input => {
    val xy = input.split(",")
    new Point(xy(0).toInt, xy(1).toInt)
  }

  val toLine: (String) => Line = input => {
    val points = input.split(" -> ")
    new Line(toPoint(points(0)), toPoint(points(1)))
  }

  val lines = scala.io.Source.fromFile("inputs/day05").mkString.split('\n').map(toLine)

  println(solve(lines))
}
