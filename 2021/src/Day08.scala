object Day08 extends App {
  def countUniqueDigits(segments: Array[String]): Int = {
    segments.map(_.length).filter(Set(2, 4, 3, 7).contains(_)).length
  }

  def solve(segmentLines: Array[(Array[String], Array[String])]): Int = {
    segmentLines.map(_._2).map(countUniqueDigits).sum
  }

  val segmentLine = (parts: Array[String]) => (parts(0).strip.split(' '), parts(1).strip.split(' '))
  val segments = scala.io.Source.fromFile("inputs/day08").mkString.split('\n').map(line => segmentLine(line.split('|')))

  println(solve(segments))
}
