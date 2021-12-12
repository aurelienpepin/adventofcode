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

    def isVisitableOnce(v: String) = Set("start").contains(v) || v.filterNot(_.isLower).isEmpty

    def bfs(source: String): Int = {
      val visited = mutable.Set.empty[String]
      val pathsTo = mutable.Map((source -> 1))

      val queue = mutable.Queue.empty[String]
      queue.enqueue(source)

      while (!queue.isEmpty) {
        println(queue)
        val v = queue.dequeue()
        pathsTo(v) = pathsTo.getOrElseUpdate(v, 0) + 1

        if (isVisitableOnce(v))
          visited.add(v)

        for (u <- adj(v)) {
          if (!visited.contains(u))
            queue.enqueue(u)
        }
      }

      pathsTo("end")
    }

    def count(s: String, visited: mutable.Set[String], currentList: ListBuffer[String]): Unit = {
      if (s == "end") {
        println(currentList.mkString(","))
      } else {
        if (isVisitableOnce(s))
          visited.add(s)

        for (t <- adj(s)) {
          if (!visited.contains(t)) {
            currentList.append(t)
            count(t, visited, currentList)
            currentList.remove(currentList.lastIndexOf(t))
          }
        }

        visited.remove(s)
      }

      ()
    }
  }

  def toGraph(adj: Array[(String, String)]): Graph = {
    val graph = new Graph(adj.flatMap { case (a, b) => Array(a, b) }.toSet)
    adj.foreach { case (a, b) => graph.addEdge(a, b) }
    graph
  }

  def solve(graph: Graph): Int = {
    graph.count("start", mutable.Set.empty, ListBuffer("start"))
    1
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
