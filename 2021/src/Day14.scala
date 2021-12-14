object Day14 extends App {
  def solve(template: String, rules: Map[String, String]): Int = {
    val finalTemplate =
      (1 to 10).foldLeft(template)((t, _) =>
        t.zip(t.drop(1)).map((c, d) => "" + c + rules("" + c + d)).mkString + t.last
      )

    val frequencyMap = finalTemplate.groupBy(identity).map(c => (c._1 -> c._2.length)).toMap
    frequencyMap.maxBy(_._2)._2 - frequencyMap.minBy(_._2)._2
  }

  val input = scala.io.Source
    .fromFile("inputs/day14")
    .mkString
    .split("\n\n")

  val template = input(0).strip
  val rules = input(1).strip.split('\n').map(_.split(" -> ")).map { case Array(a, b) => (a -> b) }.toMap

  println(solve(template, rules))
}
