import scala.util.Try;

object Day04 extends App {
  class Board(
      var board: Map[Int, (Int, Int)],
      drawnInRows: Array[Int],
      drawnInColumns: Array[Int]
  ) {
    def this(numbers: Array[Array[Int]]) = {
      this(
        Map.from(
          numbers.zipWithIndex.flatMap(row =>
            row._1.zipWithIndex
              .map((_, row._2))
              .map((vi, j) => (vi._1 -> (vi._2, j)))
          )
        ),
        Array.fill(numbers.length) { 0 },
        Array.fill(numbers(0).length) { 0 }
      )
    }

    def draw(number: Int): Int = {
      val (newBoard, score) = board
        .get(number)
        .map(position => {
          drawnInRows(position._1) += 1
          drawnInColumns(position._2) += 1
          getScoreAfterDraw(number)
        })
        .map(s => (board.removed(number), s))
        .getOrElse((board, 0))

      board = newBoard
      score
    }

    def getScoreAfterDraw(number: Int): Int = {
      if (
        drawnInRows.exists(_ == drawnInColumns.length)
        || drawnInColumns.exists(_ == drawnInRows.length)
      ) {
        number * (board.keys.sum - number)
      } else {
        0
      }
    }
  }

  def solve(draws: Array[Int], boards: Array[Board]): Int = {
    draws.foldLeft(0)((bestScore, draw) => {
      if (bestScore > 0) {
        bestScore // keep the best score
      } else {
        boards.map(_.draw(draw)).max
      }
    })
  }

  def solve2(draws: Array[Int], boards: Array[Board]): Unit = {
    var remainingBoards = boards

    for (draw <- draws) {
      remainingBoards = remainingBoards
        .map(board => (board, board.draw(draw)))
        .flatMap((b, v) => {
          if (v > 0) {
            println(v)
            None
          } else {
            Some(b)
          }
        })
    }
  }

  val lines = scala.io.Source
    .fromFile("inputs/day04")
    .mkString
    .split("\n\n")

  val (draws, boards): (Array[Int], Array[Board]) =
    (
      lines.head.split(',').map(_.toInt),
      lines.tail
        .map(
          _.split('\n').map(_.split("\\s+").flatMap(v => Try(v.toInt).toOption))
        )
        .map(new Board(_))
    )

  println(solve2(draws, boards))
}
