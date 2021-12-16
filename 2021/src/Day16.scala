import scala.collection.immutable.Nil.:::

object Day16 extends App {
  val toBinInput: (String) => String = _.flatMap(toBin)

  val toInt: (String) => Int = word => Integer.parseInt(word, 2)
  val toBin: (Char) => String = (c) =>
    String
      .format(
        "%4s",
        c match {
          case c if c <= '9' => (c - '0').toBinaryString
          case c             => (c - 'A' + 10).toBinaryString
        }
      )
      .replace(' ', '0')

  val version: (String) => Int = (i) => toInt(i.take(3))
  val typeId: (String) => Int = (i) => toInt(i.drop(3).take(3))

  // For version 4 (literal value)
  val literalValue: (String) => (Int, Int) = (i) => {
    val groups = i.drop(6).grouped(5).filter(_.length == 5).toList
    val goodGroups = groups.takeWhile(_.head == '1') ++ groups.dropWhile(_.head == '1').take(1).toList

    // val value = toInt(goodGroups.flatMap(_.drop(1)).mkString)
    val nonPaddedLength = 6 + goodGroups.length * 5

    // val rem = (nonPaddedLength + 4) % 4
    // val length = if (rem == 0) nonPaddedLength else nonPaddedLength + 4 - rem

    (0, nonPaddedLength)
  }

  // For version != 4 (operator)
  val lengthTypeId: (String) => Int = (i) => toInt(i.drop(6).take(1))
  val subpackets: (String, Int) => (String, Int) = (i, lengthTypeId) => {
    val marker = lengthTypeId match {
      case 0 => toInt(i.drop(7).take(15)) // LENGTH OF SUBPACKETS
      case 1 => toInt(i.drop(7).take(11)) // NUMBER OF SUBPACKETS
    }
    (i.drop(7).drop(if (lengthTypeId == 0) 15 else 11), marker)
  }

  def solve(input: String): Int = {
    1
    // literalValue(toBinInput(input))
  }

  def packetValue(input: String): (Int, Int, Int) = {
    typeId(input) match {
      case 4 => {
        // println(s"v${version(input)} literal value: " + input)
        val (a, b) = literalValue(input)
        (a, b, version(input))
      }
      case _ => {
        val lti = lengthTypeId(input)
        val (sps, marker) = subpackets(input, lti)
        println(s"v${version(input)} operator $lti $marker:\t" + input)
        lti match {
          case 0 => {
            var sum = 0
            var beginning = 0
            var totalV = 0

            while (marker - beginning >= 11) {
              println("--- now solving with beginning: " + beginning + " " + sps)
              val (a, b, partialV) = packetValue(sps.drop(beginning))
              println("case0: " + b)
              sum += a
              beginning += b
              totalV += partialV
            }
            (sum, marker + 7 + (if (lti == 0) 15 else 11), totalV + version(input))
          }
          case 1 => {
            var sum = 0
            var beginning = 0
            var totalV = 0

            for (k <- 1 to marker) {
              println("--- now solving with beginning: " + beginning + " " + sps)
              val (a, b, partialV) = packetValue(sps.drop(beginning))
              println("case1: " + b)
              sum += a
              beginning += b
              totalV += partialV
            }
            (sum, beginning + 7 + (if (lti == 0) 15 else 11), totalV + version(input))
          }
        }
      }
    }
  }

  val input = scala.io.Source.fromFile("inputs/day16").mkString
  println(packetValue(toBinInput(input))._3)
}
