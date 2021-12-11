import scala.collection.mutable

object Day11 extends App {
  class Grid(var grid: Array[Array[Int]]) {
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
      Array(
        (x - 1, y),
        (x + 1, y),
        (x, y + 1),
        (x, y - 1),
        (x - 1, y - 1),
        (x - 1, y + 1),
        (x + 1, y - 1),
        (x + 1, y + 1)
      ).map { case (x, y) => get(x, y) }.flatten
    }

    def mkStep(): (Int, Boolean) = {
      val flashables = mutable.Queue.empty[(Int, Int)]
      val haveFlashed = mutable.Set.empty[(Int, Int)]

      for (i <- 0 until grid.length) {
        for (j <- 0 until grid(i).length) {
          grid(i)(j) += 1
          if (grid(i)(j) == 10) {
            flashables.enqueue((i, j))
          }
        }
      }

      while (!flashables.isEmpty) {
        val (i, j) = flashables.dequeue
        val alreadyFlashed = haveFlashed.contains((i, j))

        // Flash
        grid(i)(j) = 0
        haveFlashed.add((i, j))

        if (!alreadyFlashed) {
          for ((x, y) <- neighbours((i, j))) {
            if (grid(x)(y) >= 9) {
              if (!haveFlashed.contains((x, y)))
                flashables.enqueue((x, y))
            } else {
              if (!haveFlashed.contains((x, y)))
                grid(x)(y) += 1 //= Math.min(grid(x)(y) + 1, 9)
            }
          }
        }
      }

      (haveFlashed.size, allPoints().filterNot(_ == 0).isEmpty)
    }

    override def toString(): String = {
      grid.map(_.mkString("")).mkString("\n")
    }
  }

  def solve(grid: Grid, steps: Int): Int = {
    var flashes = 0
    for (i <- 1 to steps) {
      flashes += grid.mkStep()._1
    }
    flashes
  }

  val input = scala.io.Source.fromFile("inputs/day11").mkString.split('\n').map(_.strip.map(_.asDigit).toArray)
  var grid = new Grid(input)

  println(solve(grid, 100))
}
