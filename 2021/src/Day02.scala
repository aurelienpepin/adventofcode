object Day02 extends App {
  def solve(movements: Array[Movement]): Int = {
    val result = movements.foldLeft((0, 0))((tuple, m) => m.move(tuple))
    result._1 * result._2
  }

  val input = scala.io.Source.fromFile("inputs/day02")
  var lines = (try input.mkString finally input.close())
    .split('\n')
    .map(_.split(' '))
    .map { case Array(movement, value) => (movement, value.toInt) }
    .map((movement, value) => (movement, value) match {
      case ("forward", _) => Forward(value)
      case ("up", _) => Up(value)
      case ("down", _) => Down(value)
    })

  println(solve(lines))
}

trait Movement {
  def move(data: (Int, Int)): (Int, Int)
}

case class Forward(x: Int) extends Movement {
  override def move(data: (Int, Int)): (Int, Int) = (data._1 + x, data._2)
}

case class Up(y: Int) extends Movement {
  override def move(data: (Int, Int)): (Int, Int) = (data._1, data._2 - y)
}

case class Down(y: Int) extends Movement {
  override def move(data: (Int, Int)): (Int, Int) = (data._1, data._2 + y)
}