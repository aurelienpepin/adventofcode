println(1)
println("Hello, " + "World!")

val x: Int = 1 - -1
println(x)

var y: Int = 1 + 2
y = 4 // works because y is var, not val
println(y * y)

println({
  var x: Int = 1 + 7
  x += 7
  println(x)
}) // result of the block := result of the last expression

/* FUNCTIONS */
val addOne = (x: Int) => x + 1
var noParam: (() => Int) = () => 50
println(addOne(2))
println(noParam())

/* METHODS */
def add(x: Int, y: Int): Int = x + y
println(add(1, 2))

def addThenMultiply(x: Int, y: Int)(mult: Int): Int = (x + y) * mult
val f = addThenMultiply(1, 5)
println(f(2))

def getSquareString(input: Double): String = {
  val square = input * input
  square.toString
}
println(getSquareString(2.5))