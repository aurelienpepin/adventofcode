import scala.collection.mutable.Stack

object Day18 extends App {
  sealed trait Operation
  case object OP_EXPLODE extends Operation
  case object OP_SPLIT extends Operation
  case object NOTHING extends Operation

  class Node(var v: Option[Int], var left: Node, var right: Node, var depth: Int) {
    def isPair(): Boolean = v.isEmpty

    def magnitude(): BigInt = {
      if (v.isDefined) {
        v.get
      } else {
        3 * left.magnitude() + 2 * right.magnitude()
      }
    }

    override def toString: String = if (isPair()) "[" + this.left + "," + this.right + "]" else v.get.toString
  }

  def getParent(root: Node, node: Node): Option[(Node, Direction)] = {
    if (root == node || root == null) {
      None
    } else if (root.left == node) {
      Some(root, Left())
    } else if (root.right == node) {
      Some(root, Right())
    } else {
      val fromLeft = getParent(root.left, node)
      if (fromLeft.isDefined)
        fromLeft
      else
        getParent(root.right, node)
    }
  }

  sealed trait Direction
  case class Left() extends Direction
  case class Right() extends Direction

  def shouldExplode(n: Node): Boolean = n.isPair() && n.depth >= 4
  def shouldSplit(n: Node): Boolean = n.v.getOrElse(Int.MinValue) >= 10

  def inorder(root: Node): Boolean = {
    println("~~ [New inorder] ~~")
    var forcedOperation = inorderFirstOperation(root)
    if (forcedOperation == NOTHING)
      return false

    val stack = Stack.empty[Node]
    var current: Node = root

    var prevRegular: Node = null
    var prevPrevRegular: Node = null
    var forwardForRight: Option[(Int, Node)] = None

    while (current != null || !stack.isEmpty) {
      while (current != null) {
        stack.push(current)
        current = current.left
      }

      current = stack.pop()

      // PROCESS NODE
      // println(">>>> " + current.depth + " " + current + " " + forwardForRight) // + "\t" + stack)
      // parent = getParent(root, current)
      if (!forwardForRight.isEmpty && !current.isPair()) { //} && forwardForRight.get._2 != getParent(root, current).get._1) {
        println("@ forwardForRight " + forwardForRight.get._2)
        println("@ Parent vs current: " + getParent(root, current) + " " + current)
        current.v = Some(current.v.get + forwardForRight.get._1)
        return true
      } else if (forcedOperation == OP_EXPLODE && forwardForRight.isEmpty && shouldExplode(current)) {
        println("* shouldExplode")
        println("* prevPrevRegular " + prevPrevRegular)
        println("* Parent vs current: " + getParent(root, current) + " " + current)
        if (prevPrevRegular != null) {
          prevPrevRegular.v = Some(prevPrevRegular.v.get + current.left.v.get)
        }

        val parent = getParent(root, current).get
        val newValue = leaf(0, current.depth)
        forwardForRight = Some((current.right.v.get, parent._1))

        parent._2 match {
          case Left()  => parent._1.left = newValue
          case Right() => parent._1.right = newValue
        }
        current = newValue
      } else if (forcedOperation == OP_SPLIT && forwardForRight.isEmpty && shouldSplit(current)) {
        println("§ shouldSplit")
        val number = current.v.get
        val newPair =
          node(leaf(number / 2, current.depth + 1), leaf(number - number / 2, current.depth + 1), current.depth)

        val parent = getParent(root, current).get
        parent._2 match {
          case Left()  => parent._1.left = newPair
          case Right() => parent._1.right = newPair
        }
        return true
      }

      if (!current.isPair()) {
        prevPrevRegular = prevRegular
        prevRegular = current
      }

      current = current.right
    }

    return !forwardForRight.isEmpty
  }

  def inorderFirstOperation(root: Node): Operation = {
    val stack = Stack.empty[Node]
    var current: Node = root

    var atLeastOneExplode = false
    var atLeastOneSplit = false

    while (current != null || !stack.isEmpty) {
      while (current != null) {
        stack.push(current)
        current = current.left
      }

      current = stack.pop()

      // PROCESS
      if (shouldExplode(current))
        atLeastOneExplode = true
      if (shouldSplit(current))
        atLeastOneSplit = true

      current = current.right
    }

    if (atLeastOneExplode)
      OP_EXPLODE
    else if (atLeastOneSplit)
      OP_SPLIT
    else
      NOTHING
  }

  def node(l: Node, r: Node, d: Int): Node = new Node(None, l, r, d)
  def leaf(v: Int, d: Int): Node = new Node(Some(v), null, null, d)

  def toTree(input: String, depth: Int): Node = {
    if (input forall Character.isDigit) {
      Node(Some(input.toInt), null, null, depth)
    } else {
      var c: Int = 1
      var count = if (input(c) == '[') 1 else 0

      while (count > 0) {
        c += 1
        if (input(c) == '[')
          count += 1
        else if (input(c) == ']')
          count -= 1
      }

      val firstPart = input.substring(1, c + 1)
      val secondPart = input.substring(c + 2, input.length - 1)
      Node(None, toTree(firstPart, depth + 1), toTree(secondPart, depth + 1), depth)
    }
  }

  def addTrees(a: Node, b: Node): Node = {
    // Increase depths in a
    def increaseDepth(n: Node): Unit = {
      if (n != null) {
        n.depth += 1
        increaseDepth(n.left)
        increaseDepth(n.right)
      }
    }

    increaseDepth(a)
    increaseDepth(b)
    new Node(None, a, b, 0)
  }

  def solve(input: Array[String]): BigInt = {
    var tree: Node = toTree(input(0), 0)

    input.tail
      .foldLeft(tree)((reducedTree, line) => {
        println("-- NEW TREE --")
        var newTree = toTree(line, 0)
        var treeToReduce = addTrees(reducedTree, newTree)
        while (inorder(treeToReduce)) {
          println(treeToReduce)
        }

        treeToReduce
      })
      .magnitude()
  }

  def solve2(input: Array[String]): BigInt = {
    input
      .combinations(2)
      .toArray
      .map { case Array(x, y) => (x, y) }
      .map((a: String, b: String) => {
        val treeToReduce = addTrees(toTree(a, 0), toTree(b, 0))
        while (inorder(treeToReduce)) {}

        treeToReduce.magnitude()
      })
      .max
  }

  val input = scala.io.Source.fromFile("inputs/day18").mkString.split('\n')

  // println(solve(input))
  println(solve2(input))
}
