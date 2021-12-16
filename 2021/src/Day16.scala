import scala.collection.immutable.Nil.:::

object Day16 extends App {
  val toInt: (String) => Int = word => Integer.parseInt(word, 2)
  val toBigInt: (String) => BigInt = word => BigInt(word, 2)

  val toBinInput: (String) => String = _.flatMap(toBin)
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
  val literalValue: (String) => (BigInt, Int) = (i) => {
    val groups = i.drop(6).grouped(5).filter(_.length == 5).toList
    val goodGroups = groups.takeWhile(_.head == '1') ++ groups.dropWhile(_.head == '1').take(1).toList

    val value = toBigInt(goodGroups.flatMap(_.drop(1)).mkString)
    val nonPaddedLength = 6 + goodGroups.length * 5

    (value, nonPaddedLength)
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

  def apply(code: Int, values: Array[BigInt]): BigInt = {
    code match {
      case 0 => values.sum
      case 1 => values.product
      case 2 => values.min
      case 3 => values.max
      case 5 => if (values(0) > values(1)) 1 else 0
      case 6 => if (values(0) < values(1)) 1 else 0
      case 7 => if (values(0) == values(1)) 1 else 0
    }
  }

  def packetValue(input: String): (BigInt, Int, Int) = {
    typeId(input) match {
      case 4 => {
        val (a, b) = literalValue(input)
        (a, b, version(input))
      }
      case t => {
        val lti = lengthTypeId(input)
        val shift = 7 + (if (lti == 0) 15 else 11)
        val (sps, marker) = subpackets(input, lti)

        var values = Array.empty[BigInt]
        var beginning = 0
        var totalV = 0

        lti match {
          case 0 => {
            while (marker - beginning >= 11) {
              val (a, b, partialV) = packetValue(sps.drop(beginning))

              values = values.appended(a)
              beginning += b
              totalV += partialV
            }
            (apply(t, values), marker + shift, totalV + version(input))
          }
          case 1 => {
            for (k <- 1 to marker) {
              val (a, b, partialV) = packetValue(sps.drop(beginning))

              values = values.appended(a)
              beginning += b
              totalV += partialV
            }
            (apply(t, values), beginning + shift, totalV + version(input))
          }
        }
      }
    }
  }

  val input = scala.io.Source.fromFile("inputs/day16").mkString

  // println(packetValue(toBinInput(input))._3)
  println(packetValue(toBinInput(input))._1)
}
