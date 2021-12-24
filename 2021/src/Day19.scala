import scala.collection.immutable

object Day19 extends App {

  case class RelBeacon(x: Int, y: Int, z: Int)

  case class Transformation(rule: String) {
    def roll(b: RelBeacon): RelBeacon = RelBeacon(b.x, b.z, -b.y)
    def turn(b: RelBeacon): RelBeacon = RelBeacon(-b.y, b.x, b.z)

    def apply(b: RelBeacon): RelBeacon = {
      rule.foldLeft(b)((beacon, c) =>
        c match {
          case 'R' => roll(beacon)
          case 'T' => turn(beacon)
        }
      )
    }
  }

  def transformations(): Array[Transformation] = {
    var allTransformations: Set[Transformation] = Set.empty
    val rule = new StringBuilder()

    for (cycle <- 0 to 1) {
      for (step <- 0 to 2) {
        rule.append("R")
        allTransformations += Transformation(rule.toString())

        for (i <- 0 to 2) {
          rule.append("T")
          allTransformations += Transformation(rule.toString())
        }
      }

      rule.append("RTR")
    }

    allTransformations.toArray
  }

  def dist(first: RelBeacon, second: RelBeacon): Int =
    Math.abs(first.x - second.x) + Math.abs(first.y - second.y) + Math.abs(first.z - second.z)

  def distFrom(source: RelBeacon, others: Array[RelBeacon]): Array[Int] =
    others.filter(_ != source).map(dist(source, _)).sorted

  def toRelBeacons(input: Array[String]): Set[RelBeacon] =
    input.tail.map(_.split(',')).map { case Array(x, y, z) => RelBeacon(x.toInt, y.toInt, z.toInt) }.toSet

  def allOrigins(bS1: RelBeacon, bS2: RelBeacon): Array[(Int, Int, Int)] = Array(
    (bS1.x - bS2.x, bS1.y - bS2.y, bS1.z - bS2.z)
  )

  def step(input: (Array[Set[RelBeacon]], Array[Set[RelBeacon]])): (Array[Set[RelBeacon]], Array[Set[RelBeacon]]) = {
    println("==============================")
    val beaconsByScanner = input._1
    val scannersByScanner = input._2

    if (beaconsByScanner.length == 1)
      return (beaconsByScanner, scannersByScanner)

    var beaconDistances: Array[(Int, RelBeacon, Array[Int])] =
      beaconsByScanner.zipWithIndex.flatMap((beacons: Set[RelBeacon], i: Int) => {
        beacons.map((b) => (i, b, distFrom(b, beaconsByScanner(i).toArray)))
      })

    var matchedBeaconPairs = beaconDistances
      .combinations(2)
      .map { case Array(b1, b2) => (b1, b2) }
      .filter { case (b1, b2) => b1._1 < b2._1 }
      .map { case (b1, b2) => (b1._1, b1._2, b2._1, b2._2, b1._3.intersect(b2._3)) }
      // .filter { case (s1, b1, s2, b2, intersection) => intersection.length >= 12 }
      .toArray
      .sortBy(-_._5.length)

    // println(matchedBeaconPairs.map(t => (t._1, t._2, t._3, t._4, t._5.length, t._5.mkString(","))).mkString("\n"))
    val bestScannerPairs: Array[(Int, Int, Int)] = matchedBeaconPairs
      .groupBy(_._5.length)
      .flatMap { case (l, pairs) =>
        pairs.map(p => (l, p._1, p._3)).groupBy(identity).view.mapValues(_.size).toMap.filter(_._2 >= 2)
      }
      .toArray
      .map(_._1)
      .sortBy(-_._1)

    println("OTHER PROPOSAL: " + bestScannerPairs.head)
    val (s1: Int, s2: Int) = (bestScannerPairs.head._2, bestScannerPairs.head._3)

    var s1BeaconDistances = beaconsByScanner(s1).map((s) => (s, distFrom(s, beaconsByScanner(s1).toArray)))
    var s2BeaconDistances = beaconsByScanner(s2).map((s) => (s, distFrom(s, beaconsByScanner(s2).toArray)))

    var allMatches = s1BeaconDistances
      .flatMap { case (b, distances) =>
        s2BeaconDistances.map { case (b2, distances2) => (b, b2, distances.intersect(distances2).length) }
      }
      .toList
      .filter(_._3 > 0)
      .sortBy(-_._3)

    // println(allMatches)

    // Find the S2 coordinate relative to S1
    val bestMatches = allMatches.take(2)
    val goodOrigins: Array[(Transformation, RelBeacon)] = transformations()
      .filter {
        case (transformation) => {
          val (b1S1, b2S1) = (bestMatches(0)._1, bestMatches(1)._1)
          val b1S2 = transformation.apply(bestMatches(0)._2)
          val b2S2 = transformation.apply(bestMatches(1)._2)

          // println(s"$b1S1 $b1S2 " + dist(b1S1, b1S2) + " and " + s"$b2S1 $b2S2 " + dist(b2S1, b2S2))
          dist(b1S1, b1S2) == dist(b2S1, b2S2)
        }
      }
      .flatMap((transformation) => {
        println("take all origins")
        val (b1S1, b2S1) = (bestMatches(0)._1, bestMatches(1)._1)
        val b1S2 = transformation.apply(bestMatches(0)._2)
        val b2S2 = transformation.apply(bestMatches(1)._2)
        val matchingDistance: (RelBeacon, RelBeacon, RelBeacon) => Boolean = (bS1, bS2, hypothesis) => {
          dist(bS1, hypothesis) == dist(bS2, RelBeacon(0, 0, 0))
        }

        // println(allOrigins(b1S1, b1S2).mkString(",") + " || " + allOrigins(b2S1, b2S2).mkString(","))
        allOrigins(b1S1, b1S2)
          .intersect(allOrigins(b2S1, b2S2))
          .filter {
            case (x, y, z) => {
              matchingDistance(b1S1, b1S2, RelBeacon(x, y, z))
              && matchingDistance (b2S1, b2S2, RelBeacon(x, y, z))
            }
          }
          .headOption
          .map { case (x, y, z) => (transformation, RelBeacon(x, y, z)) }
      })

    println("GOOD ORIGINS: " + goodOrigins.mkString(", "))
    val (transformation, relOrigin) = goodOrigins.head

    val allTranslatedS2Beacons = beaconsByScanner(s2).map {
      case RelBeacon(x, y, z) => {
        val transformedBeacon = transformation.apply(RelBeacon(x, y, z))

        RelBeacon(
          transformedBeacon.x + relOrigin.x,
          transformedBeacon.y + relOrigin.y,
          transformedBeacon.z + relOrigin.z
        )
      }
    }

    val allTranslatedS2Scanners = Array(relOrigin) ++ scannersByScanner(s2).map {
      case RelBeacon(x, y, z) => {
        val transformedBeacon = transformation.apply(RelBeacon(x, y, z))

        RelBeacon(
          transformedBeacon.x + relOrigin.x,
          transformedBeacon.y + relOrigin.y,
          transformedBeacon.z + relOrigin.z
        )
      }
    }

    println("REL ORIGIN " + s2 + " " + relOrigin)
    val mergedBeacons = beaconsByScanner.zipWithIndex.flatMap { (bs, i) =>
      {
        if (i == s2) {
          None
        } else if (i == s1) {
          Some(bs ++ allTranslatedS2Beacons)
        } else {
          Some(bs)
        }
      }
    }

    val mergedScanners = scannersByScanner.zipWithIndex.flatMap { (ss, i) =>
      {
        if (i == s2) {
          None
        } else if (i == s1) {
          Some(ss ++ allTranslatedS2Scanners)
        } else {
          Some(ss)
        }
      }
    }

    println("How many scanners left: " + mergedBeacons.size)
    (mergedBeacons, mergedScanners)
  }

  def solve(input: Array[Set[RelBeacon]]): Int = {
    var scanners = Int.MaxValue
    var scannersByScanner: Array[Set[RelBeacon]] = Array.fill(input.length)(Set.empty)
    var beaconsByScanner = input

    while (scanners > 1) {
      scanners = beaconsByScanner.length
      val (newBeacons, newScannersByScanner) = step((beaconsByScanner, scannersByScanner))

      beaconsByScanner = newBeacons
      scannersByScanner = newScannersByScanner
    }

    println(scannersByScanner(0).toArray.combinations(2).map { case Array(a, b) => dist(a, b) }.max)
    println(beaconsByScanner(0).size)
    scanners
  }

  var beaconsByScanner = scala.io.Source
    .fromFile("inputs/day19")
    .mkString
    .split("\n\n")
    .map(_.strip.split("\n"))
    .map(toRelBeacons(_))

  println(solve(beaconsByScanner))
}
