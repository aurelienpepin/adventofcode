import scala.collection.mutable

object Day14 extends App {
  def solve(template: String, rules: Map[String, String]): Int = {
    val finalTemplate =
      (1 to 10).foldLeft(template)((t, _) =>
        t.zip(t.drop(1)).map((c, d) => "" + c + rules("" + c + d)).mkString + t.last
      )

    val frequencyMap = finalTemplate.groupBy(identity).map(c => (c._1 -> c._2.length)).toMap
    frequencyMap.maxBy(_._2)._2 - frequencyMap.minBy(_._2)._2
  }

  def solve2(template: String, rules: Map[String, String], steps: Int): BigInt = {
    var frequencyMap: mutable.Map[String, BigInt] = mutable.Map(rules.keys.map((_, BigInt(0))).toSeq: _*)
    template.zip(template.drop(1)).map((c, d) => frequencyMap("" + c + d) += 1)
    println(frequencyMap)

    for (_ <- 1 to steps) {
      var newFrequencyMap = frequencyMap.clone

      frequencyMap.foreach((leftPart, count) => {
        val newSymbol: String = rules(leftPart)(0).toString
        newFrequencyMap(leftPart) -= count

        newFrequencyMap(leftPart(0) + newSymbol) += count
        newFrequencyMap(newSymbol + leftPart(1)) += count
      })
      frequencyMap = newFrequencyMap
    }

    val finalMap = frequencyMap.groupBy(_._1.last).map((cm) => (cm._1, cm._2.values.sum + (if (template.head == cm._1) then 1 else 0)))
    finalMap.maxBy(_._2)._2 - finalMap.minBy(_._2)._2
  }

  val input = scala.io.Source
    .fromFile("inputs/day14")
    .mkString
    .split("\n\n")

  val template = input(0).strip
  val rules = input(1).strip.split('\n').map(_.split(" -> ")).map { case Array(a, b) => (a -> b) }.toMap

  // println(solve(template, rules))
  println(solve2(template, rules, 40))
}
