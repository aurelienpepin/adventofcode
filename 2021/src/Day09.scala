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

    def getBasinSize(x: Int, y: Int): Int = 0
  }

  def solve(grid: Grid): Int = grid.allPoints().map { case (x, y) => grid.riskLevel(x, y) }.sum

  val input = scala.io.Source.fromFile("inputs/day09").mkString.split('\n').map(_.strip.map(_.asDigit).toArray)
  var grid = new Grid(input)

  println(solve(grid))
}
