import scala.collection.mutable
import scala.collection.mutable.ListBuffer

object Day12 extends App {
  class Graph() {
    private val adj = mutable.Map.empty[String, mutable.Set[String]]

    def this(vertices: Set[String]) = {
      this()
      vertices.foreach((v) => adj += (v -> mutable.Set.empty))
    }

    def addEdge(a: String, b: String): Unit = {
      adj(a) += b
      adj(b) += a
    }

    def isVisitableOnce(v: String) = v.filterNot(_.isLower).isEmpty

    def count(s: String, visited: mutable.Map[String, Int], currentList: ListBuffer[String]): Int = {
      if (s == "end") {
        1
      } else {
        if (isVisitableOnce(s))
          visited(s) = Math.min(visited.getOrElseUpdate(s, 0) + 1, 2)

        var allPaths = 0
        for (t <- adj(s)) {
          if (
            t != "start" && (!visited
              .contains(t) || visited(t) == 0 || (visited(t) == 1 && currentList
              .filter(o => visited.getOrElseUpdate(o, 0) == 2)
              .isEmpty))
          ) {
            currentList.append(t)
            allPaths += count(t, visited, currentList)
            currentList.remove(currentList.lastIndexOf(t))
          }
        }

        if (visited.contains(s))
          visited(s) = Math.max(visited(s) - 1, 0)

        allPaths
      }
    }
  }

  def toGraph(adj: Array[(String, String)]): Graph = {
    val graph = new Graph(adj.flatMap { case (a, b) => Array(a, b) }.toSet)
    adj.foreach { case (a, b) => graph.addEdge(a, b) }
    graph
  }

  def solve(graph: Graph): Int = {
    graph.count("start", mutable.Map.empty, ListBuffer("start"))
  }

  val input = scala.io.Source
    .fromFile("inputs/day12")
    .mkString
    .split('\n')
    .map(_.strip.split("-"))
    .map { case Array(a, b) => (a, b) }

  val graph = toGraph(input)
  println(solve(graph))
}
