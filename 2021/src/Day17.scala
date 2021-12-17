import scala.util.control.Breaks._

object Day17 extends App {
  def allValidVx0(minX: Int, maxX: Int): Array[Int] = {
    (0 to 100).dropWhile((vx0) => minX > vx0 * (vx0 + 1) / 2).takeWhile((vx0) => vx0 <= maxX).toArray
  }

  def x(n: Int, vx0: Int): Double = {
    0 - (1 / 2.0) * (n - 1) * n * Math.signum(vx0) + n * vx0
  }

  def y(n: Int, vy0: Int): Double = {
    0 - (1 / 2.0) * n * (n - 2 * vy0 - 1)
  }

  val minTargetX = 29
  val maxTargetX = 73
  val minTargetY = -248
  val maxTargetY = -194

  var highestY = Double.MinValue

  for (vx0 <- allValidVx0(minTargetX, maxTargetX)) {
    // println(vx0)
    for (vy0 <- 0 to 1000) {
      // println(vx0 + " " + vy0)
      var x = 0
      var y = 0
      var vx = vx0
      var vy = vy0
      var currentHighestY = Int.MinValue

      breakable {

        while (true) {
          x = x + vx
          y = y + vy
          vx = vx - Math.signum(vx).toInt
          vy = vy - 1

          currentHighestY = Math.max(currentHighestY, y)

          if (y < minTargetY || x > maxTargetX)
            break

          if (
            minTargetX <= x && x <= maxTargetX
            && minTargetY <= y && y <= maxTargetY
            && currentHighestY > highestY
          ) {
            highestY = currentHighestY
            println(vx0 + " " + vy0 + " " + highestY)
          }
        }
      }
    }
  }

  println(highestY)
}
