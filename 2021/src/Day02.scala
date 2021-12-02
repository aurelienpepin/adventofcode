object Day02 extends App {
  def solve(movements: Array[Movement]): Int = {
    val result = movements.foldLeft((0, 0))((tuple, m) => m.move(tuple))
    result._1 * result._2
  }

  def solve2(movements: Array[Movement]): Int = {
    val result = movements.foldLeft((0, 0, 0))((tuple, m) => m.move(tuple))
    result._1 * result._2
  }

  val input = scala.io.Source.fromFile("inputs/day02")
  var lines = (try input.mkString
  finally input.close())
    .split('\n')
    .map(_.split(' '))
    .map { case Array(movement, value) => (movement, value.toInt) }
    .map({
      case ("forward", v) => Forward(v)
      case ("up", v)      => Up(v)
      case ("down", v)    => Down(v)
    })

  // println(solve(lines))
  println(solve2(lines))
}

trait Movement {
  def move(data: (Int, Int)): (Int, Int)
  def move(data: (Int, Int, Int)): (Int, Int, Int)
}

case class Forward(x: Int) extends Movement {
  override def move(data: (Int, Int)): (Int, Int) = (data._1 + x, data._2)
  override def move(data: (Int, Int, Int)): (Int, Int, Int) =
    (data._1 + x, data._2 + (data._3 * x), data._3)
}

case class Up(y: Int) extends Movement {
  override def move(data: (Int, Int)): (Int, Int) = (data._1, data._2 - y)
  override def move(data: (Int, Int, Int)): (Int, Int, Int) =
    (data._1, data._2, data._3 - y)
}

case class Down(y: Int) extends Movement {
  override def move(data: (Int, Int)): (Int, Int) = (data._1, data._2 + y)
  override def move(data: (Int, Int, Int)): (Int, Int, Int) =
    (data._1, data._2, data._3 + y)
}
