import scala.collection.MapView

object Day08 extends App {
  def countUniqueDigits(segments: Array[String]): Int = {
    segments.map(_.length).filter(Set(2, 4, 3, 7).contains(_)).length
  }

  def groupBars(values: Array[String]): Map[Int, Set[Char]] =
    values
      .map(_.groupBy(_.toChar).view.mapValues(_.size).toList)
      .reduce(_ ++ _)
      .groupBy(_._1)
      .map({ case (k, v) =>
        (k -> v.map(_._2).sum)
      })
      .groupBy(_._2)
      .view
      .mapValues(_.keys.toSet)
      .toMap

  val addMatching = (digit: Int, bars: Array[Char], digitsByCode: Map[String, Int]) =>
    digitsByCode + (bars.mkString.sorted -> digit)

  def solveSegmentLine(segmentLine: ((Array[String], Array[String]))): Int = {
    var digitsByCode = Map.empty[String, Int]
    var codesByDigit = Map.empty[Int, String]

    // Fill unique digits
    segmentLine._1
      .foreach(s =>
        s.length match {
          case 2 => { digitsByCode += (s.sorted -> 1); codesByDigit += (1 -> s.sorted) }
          case 3 => { digitsByCode += (s.sorted -> 7); codesByDigit += (7 -> s.sorted) }
          case 4 => { digitsByCode += (s.sorted -> 4); codesByDigit += (4 -> s.sorted) }
          case 7 => { digitsByCode += (s.sorted -> 8); codesByDigit += (8 -> s.sorted) }
          case _ => ()
        }
      )

    // Deducting other digits: 0, 2, 3, 5, 6, 9
    val charByFrequency = groupBars(segmentLine._1)

    var barA = codesByDigit(7).toSet.removedAll(codesByDigit(1)).head
    var barB = charByFrequency(6).head
    var barC = charByFrequency(8).removedAll(Set(barA)).head
    var barD = codesByDigit(4).toSet.removedAll(codesByDigit(1)).removedAll(Set(barB)).head
    var barE = charByFrequency(4).head
    var barF = charByFrequency(9).head
    var barG = charByFrequency(7).removedAll(Set(barD)).head

    digitsByCode = addMatching(0, Array(barA, barB, barC, barE, barF, barG), digitsByCode)
    digitsByCode = addMatching(2, Array(barA, barC, barD, barE, barG), digitsByCode)
    digitsByCode = addMatching(3, Array(barA, barC, barD, barF, barG), digitsByCode)
    digitsByCode = addMatching(5, Array(barA, barB, barD, barF, barG), digitsByCode)
    digitsByCode = addMatching(6, Array(barA, barB, barD, barE, barF, barG), digitsByCode)
    digitsByCode = addMatching(9, Array(barA, barB, barC, barD, barF, barG), digitsByCode)

    segmentLine._2.map(segment => digitsByCode(segment.sorted)).mkString.toInt
  }

  def solve(segmentLines: Array[(Array[String], Array[String])]): Int = {
    segmentLines.map(_._2).map(countUniqueDigits).sum
  }

  def solve2(segmentLines: Array[(Array[String], Array[String])]): Int = {
    segmentLines.map(solveSegmentLine).sum
  }

  val segmentLine = (parts: Array[String]) => (parts(0).strip.split(' '), parts(1).strip.split(' '))
  val segments = scala.io.Source.fromFile("inputs/day08").mkString.split('\n').map(line => segmentLine(line.split('|')))

  println(solve2(segments))
}
