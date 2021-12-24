import scala.collection.mutable

object Day20 extends App {
  val toInt: (String) => Int = word => Integer.parseInt(word, 2)

  case class Dimensions(minX: Int, maxX: Int, minY: Int, maxY: Int) {
    def extend(): Dimensions = {
      Dimensions(minX - 1, maxX + 1, minY - 1, maxY + 1)
    }
  }

  def enhance(x: Int, y: Int, image: mutable.Map[(Int, Int), Char], deadPixel: Char): Char = {
    val message = Array(
      (x - 1, y - 1),
      (x - 1, y),
      (x - 1, y + 1),
      (x, y - 1),
      (x, y),
      (x, y + 1),
      (x + 1, y - 1),
      (x + 1, y),
      (x + 1, y + 1)
    ).map { case (a, b) =>
      image.getOrElse((a, b), deadPixel)
    }.map(p => if (p == '.') '0' else '1')
      .mkString

    code(toInt(message))
  }

  def step(
      image: mutable.Map[(Int, Int), Char],
      dim: Dimensions,
      deadPixel: Char
  ): (mutable.Map[(Int, Int), Char], Dimensions, Char) = {
    var newImage = mutable.Map.empty[(Int, Int), Char]

    for (x <- dim.minX - 1 to dim.maxX + 1) {
      for (y <- dim.minY - 1 to dim.maxY + 1) {
        newImage += ((x, y) -> enhance(x, y, image, deadPixel))
      }
    }

    val newDeadPixel = code(toInt((deadPixel.toString * 9).map(p => if (p == '.') '0' else '1').mkString))
    (newImage, dim.extend(), newDeadPixel)
  }

  def solve(code: String, input: Array[String]): Int = {
    val image = mutable.Map.empty[(Int, Int), Char]

    (0 to input.length - 1).foreach(x => (0 to input(x).length - 1).foreach(y => image += ((x, y) -> input(x)(y))))
    val (newImage, dims, deadPixel) = step(image, Dimensions(0, input.length - 1, 0, input(0).length - 1), '.')
    val (newImage2, _, _) = step(newImage, dims, deadPixel)
    newImage2.values.filter(_ == '#').size
  }

  val Array(code, input) = scala.io.Source
    .fromFile("inputs/day20")
    .mkString
    .split("\n\n")
    .map(_.strip)

  val image = input.split("\n")

  println(solve(code, image))
}
