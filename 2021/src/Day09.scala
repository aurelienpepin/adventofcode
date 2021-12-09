import scala.collection.mutable

object Day09 extends App {
  class Grid(var grid: Array[Array[Int]]) {
    def isLowPoint(x: Int, y: Int): Boolean = {
      val above = if (x > 0) grid(x - 1)(y) else Int.MaxValue
      val below = if (x < grid.length - 1) grid(x + 1)(y) else Int.MaxValue
      val left = if (y > 0) grid(x)(y - 1) else Int.MaxValue
      val right = if (y < grid(x).length - 1) grid(x)(y + 1) else Int.MaxValue

      Array(above, below, left, right).filter(_ <= grid(x)(y)).isEmpty
    }

    def riskLevel(i: Int, j: Int): Int = if (isLowPoint(i, j)) 1 + grid(i)(j) else 0

    def allPoints(): Array[(Int, Int)] = {
      (0 until grid.length)
        .map(x => (x, (0 until grid(x).length).toArray))
        .flatMap({ case (v, r) => r.map((v, _)) })
        .toArray
    }

    def neighbours(point: (Int, Int)): Array[(Int, Int)] = {
      val get: (Int, Int) => Option[(Int, Int)] = (x, y) =>
        if (x < 0 || x == grid.length || y < 0 || y == grid(x).length) Option.empty else Some((x, y))

      val x = point._1
      val y = point._2
      Array((x - 1, y), (x + 1, y), (x, y + 1), (x, y - 1)).map { case (x, y) => get(x, y) }.flatten
    }

    def getBasinSize(x: Int, y: Int): Int = {
      if (!isLowPoint(x, y))
        return 1

      var queue = mutable.Queue((x, y))
      var visited = mutable.Set[(Int, Int)]()
      var size = 0

      while (!queue.isEmpty) {
        val point = queue.dequeue
        if (!visited.contains(point)) {
          visited.add(point)

          neighbours(point).filterNot(p => grid(p._1)(p._2) == 9).foreach(queue.enqueue(_))
          size += 1
        }
      }

      size
    }
  }

  def solve(grid: Grid): Int = grid.allPoints().map { case (x, y) => grid.riskLevel(x, y) }.sum

  def solve2(grid: Grid): Int =
    grid.allPoints().map { case (x, y) => grid.getBasinSize(x, y) }.sortBy(identity).reverse.take(3).product

  val input = scala.io.Source.fromFile("inputs/day09").mkString.split('\n').map(_.strip.map(_.asDigit).toArray)
  var grid = new Grid(input)

  println(solve2(grid))
}
