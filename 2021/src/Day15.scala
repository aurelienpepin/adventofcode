import scala.collection.mutable

object Day15 extends App {
  case class OutEdge(p: (Int, Int), d: Int)
  case class DistTo(p: (Int, Int), d: Int) extends Ordered[DistTo] {
    import scala.math.Ordered.orderingToOrdered
    override def compare(that: DistTo): Int = (that.d, that.p) compare (d, p)
  }

  class Graph() {
    private val adj = mutable.Map.empty[(Int, Int), mutable.Set[OutEdge]]

    def this(vertices: Set[(Int, Int)]) = {
      this()
      vertices.foreach((v) => adj += (v -> mutable.Set.empty))
    }

    def addEdge(a: (Int, Int), b: (Int, Int), dist: Int): Unit = {
      adj(a) += OutEdge(b, dist)
    }

    def dijkstra(source: (Int, Int)): Int = {
      val dists = mutable.Map(adj.keys.map((_, Int.MaxValue)).toSeq: _*)
      dists(source) = 0

      val pq = new mutable.PriorityQueue[DistTo]()
      pq.enqueue(DistTo(source, 0))

      while (!pq.isEmpty) {
        val top = pq.dequeue

        if (top.d <= dists(top.p)) {
          for (oe <- adj(top.p)) {
            if (dists(top.p) + oe.d < dists(oe.p)) {
              dists(oe.p) = dists(top.p) + oe.d
              pq.enqueue(DistTo(oe.p, dists(oe.p)))
            }
          }
        }
      }

      dists(adj.keys.maxBy((x, y) => x * y))
    }
  }

  def toGraph(input: Array[Array[Int]]): Graph = {
    val get: ((Int, Int)) => Option[(Int, Int)] = { case (x, y) =>
      if (x < 0 || x == input.length || y < 0 || y == input(x).length) Option.empty else Some((x, y))
    }

    val points: Set[(Int, Int)] = (0 until input.length)
      .map((i) => (i, (0 until input(i).length).toArray))
      .flatMap({ case (v, r) => r.map((v, _)) })
      .toSet

    val graph = new Graph(points)
    points.foreach((i, j) =>
      Array((i - 1, j), (i + 1, j), (i, j - 1), (i, j + 1)).flatMap(get).foreach { case (x, y) =>
        graph.addEdge((i, j), (x, y), input(x)(y))
      }
    )

    graph
  }

  val input = scala.io.Source.fromFile("inputs/day15").mkString.split('\n').map(_.strip.map(_.asDigit).toArray)
  val graph = toGraph(input)

  println(graph.dijkstra((0, 0)))
}
