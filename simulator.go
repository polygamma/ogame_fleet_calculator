package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

var debrisPercent = 50
var defenseIntoDebris = false
var maxGala = 6
var maxSunSystem = 499
var isDonutUniverse = true
var speedFactor = 3
var howOftenForAverage = 50
var toChangePercent = 10

func main() {
	debug.SetGCPercent(-1)
	rand.Seed(time.Now().UTC().UnixNano())

	polygammaOne := player{
		"polygamma", 1, 1, 1,
		14, 14, 14,
		15, 13, 9,
		[]playerShip{
			{&knownShips[0], 0, 0},     // Kleiner Transporter
			{&knownShips[1], 0, 0},     // Großer Transporter
			{&knownShips[2], 0, 0},     // Leichter Jäger
			{&knownShips[3], 0, 0},     // Schwerer Jäger
			{&knownShips[4], 0, 0},     // Kreuzer
			{&knownShips[5], 500, 250}, // Schlachtschiff
			{&knownShips[6], 0, 0},     // Kolonieschiff
			{&knownShips[7], 0, 0},     // Recycler
			{&knownShips[8], 0, 0},     // Spionagesonde
			{&knownShips[9], 100, 0},   // Bomber
			{&knownShips[11], 100, 0},  // Zerstörer
			{&knownShips[12], 0, 0},    // Todesstern
			{&knownShips[13], 100, 0},  // Schlachtkreuzer
		},
		[]playerDefense{},
		0, 0, 0, 0,
	}
	polygammaTwo := player{
		"polygamma", 2, 1, 1,
		14, 14, 14,
		15, 13, 9,
		[]playerShip{
			{&knownShips[0], 1500, 0}, // Kleiner Transporter
			{&knownShips[1], 250, 0},  // Großer Transporter
			{&knownShips[2], 1100, 0}, // Leichter Jäger
			{&knownShips[3], 0, 0},    // Schwerer Jäger
			{&knownShips[4], 0, 0},    // Kreuzer
			{&knownShips[5], 500, 0},  // Schlachtschiff
			{&knownShips[6], 0, 0},    // Kolonieschiff
			{&knownShips[7], 0, 0},    // Recycler
			{&knownShips[8], 0, 0},    // Spionagesonde
			{&knownShips[9], 100, 0},  // Bomber
			{&knownShips[11], 100, 0}, // Zerstörer
			{&knownShips[12], 0, 0},   // Todesstern
			{&knownShips[13], 100, 0}, // Schlachtkreuzer
		},
		[]playerDefense{},
		0, 0, 0, 0,
	}

	wazz := player{
		"wazz", 2, 2, 2,
		15, 15, 15,
		15, 13, 9,
		[]playerShip{
			{&knownShips[0], 0, 0},  // Kleiner Transporter
			{&knownShips[1], 0, 0},  // Großer Transporter
			{&knownShips[2], 0, 0},  // Leichter Jäger
			{&knownShips[3], 0, 0},  // Schwerer Jäger
			{&knownShips[4], 0, 0},  // Kreuzer
			{&knownShips[5], 0, 0},  // Schlachtschiff
			{&knownShips[6], 0, 0},  // Kolonieschiff
			{&knownShips[7], 0, 0},  // Recycler
			{&knownShips[8], 0, 0},  // Spionagesonde
			{&knownShips[9], 0, 0},  // Bomber
			{&knownShips[10], 0, 0}, // Solarsatellit
			{&knownShips[11], 0, 0}, // Zerstörer
			{&knownShips[12], 0, 0}, // Todesstern
			{&knownShips[13], 0, 0}, // Schlachtkreuzer
		},
		[]playerDefense{
			{&knownDefenses[0], 1000}, // Raketenwerfer
			{&knownDefenses[1], 500},  // Leichtes Lasergeschütz
			{&knownDefenses[2], 0},    // Schweres Lasergeschütz
			{&knownDefenses[3], 20},   // Gaußkanone
			{&knownDefenses[4], 0},    // Ionengeschütz
			{&knownDefenses[5], 10},   // Plasmawerfer
			{&knownDefenses[6], 0},    // Kleine Schildkuppel
			{&knownDefenses[7], 0},    // Große Schildkuppel
		},
		2000000, 1000000, 1500000, 50,
	}

	fight := createFight([]player{polygammaOne, polygammaTwo}, []player{wazz})
	fight.optimizeAttackers()
	fight.printFight()
}

type stationaryDataBase struct {
	name                      string
	metal, crystal, deuterium int
	shield, attack            float64
	rapidFireTableIndex       int
}

func (obj *stationaryDataBase) getName() string {
	return obj.name
}

func (obj *stationaryDataBase) getMetalCost() int {
	return obj.metal
}

func (obj *stationaryDataBase) getCrystalCost() int {
	return obj.crystal
}

func (obj *stationaryDataBase) getDeuteriumCost() int {
	return obj.deuterium
}

func (obj *stationaryDataBase) getBaseHull() float64 {
	return math.Round((float64(obj.getMetalCost()) + float64(obj.getCrystalCost())) / 10.0)
}

func (obj *stationaryDataBase) getBaseShield() float64 {
	return obj.shield
}

func (obj *stationaryDataBase) getBaseAttack() float64 {
	return obj.attack
}

func (obj *stationaryDataBase) getRapidFireTableIndex() int {
	return obj.rapidFireTableIndex
}

type movingDataBase struct {
	stationaryDataBase
	speed, consumption float64
}

func (obj *movingDataBase) getBaseSpeed() float64 {
	return obj.speed
}

func (obj *movingDataBase) getBaseConsumption() float64 {
	return obj.consumption
}

type shipBase struct {
	movingDataBase
	rapidFireTable []int
	capacity       int
	driveType      int
}

func (obj *shipBase) getRapidFireTable() []int {
	return obj.rapidFireTable
}

func (obj *shipBase) getCapacity() int {
	return obj.capacity
}

func (obj *shipBase) getDriveType() int {
	return obj.driveType
}

func (obj *shipBase) isShip() bool {
	return true
}

type defenseBase struct {
	stationaryDataBase
}

func (obj *defenseBase) isShip() bool {
	return false
}

func (obj *defenseBase) getRapidFireTable() []int {
	return []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func (obj *defenseBase) getCapacity() int {
	return 0
}

func (obj *defenseBase) getDriveType() int {
	return 0
}

type playerShip struct {
	ship                *shipBase
	count, minimumCount int
}

func (obj *playerShip) getObject() objectForFighting {
	return obj.ship
}

func (obj *playerShip) getShip() *shipBase {
	return obj.ship
}

func (obj *playerShip) getCount() int {
	return obj.count
}

func (obj *playerShip) getCountPointer() *int {
	return &obj.count
}

func (obj *playerShip) getMinimumCount() int {
	return obj.minimumCount
}

type playerDefense struct {
	defense *defenseBase
	count   int
}

func (obj *playerDefense) getObject() objectForFighting {
	return obj.defense
}

func (obj *playerDefense) getCount() int {
	return obj.count
}

type player struct {
	name                                                                 string
	galaxy, system, planet, weaponResearch, shieldResearch, hullResearch int
	combustionDrive, impulseDrive, hyperspaceDrive                       int
	ships                                                                []playerShip
	defenses                                                             []playerDefense
	metal, crystal, deuterium, lootPercent                               int
}

func (pl *player) getDistanceTo(galaxy, system, planet int) int {
	if pl.galaxy != galaxy {
		absVal := math.Abs(float64(pl.galaxy) - float64(galaxy))
		if isDonutUniverse {
			return 20000 * int(math.Min(absVal, float64(maxGala)-absVal))
		} else {
			return 20000 * int(absVal)
		}
	} else if pl.system != system {
		absVal := math.Abs(float64(pl.system) - float64(system))
		if isDonutUniverse {
			return 2700 + 95*int(math.Min(absVal, float64(maxSunSystem)-absVal))
		} else {
			return 2700 + 95*int(absVal)
		}
	} else if pl.planet != planet {
		return 1000 + 5*int(math.Abs(float64(pl.planet)-float64(planet)))
	} else {
		return 5
	}
}

func (obj *shipBase) getShipSpeed(player *player) int {
	driveType := obj.getDriveType()
	baseSpeed := obj.getBaseSpeed()

	if obj.getName() == "Kleiner Transporter" && player.impulseDrive <= 4 {
		baseSpeed = 5000
		driveType = 0
	} else if obj.getName() == "Bomber" && player.hyperspaceDrive <= 7 {
		baseSpeed = 4000
		driveType = 1
	} else if obj.getName() == "Recycler" {
		if player.hyperspaceDrive >= 15 {
			baseSpeed = 6000
			driveType = 2
		} else if player.impulseDrive >= 17 {
			baseSpeed = 4000
			driveType = 1
		}
	}

	driveFactor := 0.0
	if driveType == 0 {
		driveFactor = 0.001 * math.Round((1.0+10.0*float64(player.combustionDrive)/100.0)*1000.0)
	} else if driveType == 1 {
		driveFactor = 0.001 * math.Round((1.0+20.0*float64(player.impulseDrive)/100.0)*1000.0)
	} else if driveType == 2 {
		driveFactor = 0.001 * math.Round((1.0+30.0*float64(player.hyperspaceDrive)/100.0)*1000.0)
	}

	return int(math.Round(driveFactor * baseSpeed))
}

func getFlightTime(distance, minSpeed int) int {
	return int(math.Round((3500.0*math.Sqrt(float64(distance)*10.0/float64(minSpeed)) + 10.0) / float64(speedFactor)))
}

func (obj *shipBase) getFuel(distance, duration, count int, player *player) int {
	dummyVal := 35000.0 / float64(duration*speedFactor-10) * math.Sqrt(float64(distance)*10.0/float64(obj.getShipSpeed(player)))

	return int(math.Round(obj.getBaseConsumption() * float64(count) * float64(distance) / 35000.0 * ((dummyVal / 10.0) + 1.0) * ((dummyVal / 10.0) + 1.0)))
}

type objectForFighting interface {
	getName() string
	getMetalCost() int
	getCrystalCost() int
	getDeuteriumCost() int
	getBaseHull() float64
	getBaseShield() float64
	getBaseAttack() float64
	getRapidFireTableIndex() int
	getRapidFireTable() []int
	getCapacity() int
	getDriveType() int
	isShip() bool
}

type objectInFight interface {
	objectForFighting
	getCurrentHull() float64
	getMaxHull() float64
	setCurrentHull(toSet float64)
	getCurrentShield() float64
	getMaxShield() float64
	setCurrentShield(toSet float64)
	getAttack() float64
	getIsExploded() bool
	setIsExploded(toSet bool)
	tryToExplode() bool
}

type concreteFightObject struct {
	objectForFighting
	maxHull, currentHull, maxShield, currentShield, attack float64
	isExploded                                             bool
}

func (obj *concreteFightObject) getCurrentHull() float64 {
	return obj.currentHull
}

func (obj *concreteFightObject) getMaxHull() float64 {
	return obj.maxHull
}

func (obj *concreteFightObject) setCurrentHull(toSet float64) {
	obj.currentHull = 0.001 * math.Round(1000.0*toSet)
}

func (obj *concreteFightObject) getCurrentShield() float64 {
	return obj.currentShield
}

func (obj *concreteFightObject) getMaxShield() float64 {
	return obj.maxShield
}

func (obj *concreteFightObject) setCurrentShield(toSet float64) {
	obj.currentShield = 0.001 * math.Round(1000.0*toSet)
}

func (obj *concreteFightObject) getAttack() float64 {
	return obj.attack
}

func (obj *concreteFightObject) getIsExploded() bool {
	return obj.isExploded
}

func (obj *concreteFightObject) setIsExploded(toSet bool) {
	obj.isExploded = toSet
}

func (obj *concreteFightObject) tryToExplode() bool {
	hullRatio := obj.getCurrentHull() / obj.getMaxHull()
	if !(hullRatio <= 0.7) {
		return false
	}

	return rand.Float64() < (1.0 - hullRatio)
}

type fight struct {
	attackers, defenders                       []player
	attackerObjectsAlive, defenderObjectsAlive []objectInFight
	flightCosts, attackersLost, totalCapacity  int
}

func doRound(attackerObjs, defenderObjs []objectInFight, defenderAlive int) {
	for _, attackerObj := range attackerObjs {
		allowedToShoot := true
		for allowedToShoot {
			allowedToShoot = false
			targetObj := defenderObjs[rand.Intn(defenderAlive)]

			if attackerObj.getAttack() < targetObj.getCurrentShield() {
				percentDealt := int(math.Floor(100.0 * attackerObj.getAttack() / targetObj.getMaxShield()))
				targetObj.setCurrentShield(targetObj.getCurrentShield() - float64(percentDealt)*0.01*targetObj.getMaxShield())
			} else {
				targetObj.setCurrentShield(targetObj.getCurrentShield() - attackerObj.getAttack())
				targetObj.setCurrentHull(targetObj.getCurrentHull() + targetObj.getCurrentShield())
				targetObj.setCurrentShield(0.0)
			}

			if !targetObj.getIsExploded() && targetObj.tryToExplode() {
				targetObj.setIsExploded(true)
			}

			rapidValue := float64(attackerObj.getRapidFireTable()[targetObj.getRapidFireTableIndex()])
			if rapidValue > 0 && rand.Float64() < (1.0-1.0/rapidValue) {
				allowedToShoot = true
			}
		}
	}
}

func (fi *fight) hasAttackerWon() bool {
	return len(fi.attackerObjectsAlive) > 0 && len(fi.defenderObjectsAlive) == 0
}

func (fi *fight) attackerTotalLost() int {
	defender := fi.defenders[0]
	capacityLeft := fi.totalCapacity
	metalLoot, crystalLoot, deuteriumLoot := 0, 0, 0
	possibleMetalLoot := int(float64(defender.lootPercent) / 100.0 * float64(defender.metal))
	possibleCrystalLoot := int(float64(defender.lootPercent) / 100.0 * float64(defender.crystal))
	possibleDeuteriumLoot := int(float64(defender.lootPercent) / 100.0 * float64(defender.deuterium))
	if fi.hasAttackerWon() {
		// http://www.owiki.de/index.php?title=Beute
		// step 1
		metalLootFirst := int(math.Min(float64(capacityLeft)/3.0, float64(possibleMetalLoot)))
		capacityLeft -= metalLootFirst
		metalLoot += metalLootFirst
		possibleMetalLoot -= metalLootFirst
		// step 2
		crystalLootFirst := int(math.Min(float64(capacityLeft)/2.0, float64(possibleCrystalLoot)))
		capacityLeft -= crystalLootFirst
		crystalLoot += crystalLootFirst
		possibleCrystalLoot -= crystalLootFirst
		// step 3
		deuteriumLootFirst := int(math.Min(float64(capacityLeft), float64(possibleDeuteriumLoot)))
		capacityLeft -= deuteriumLootFirst
		deuteriumLoot += deuteriumLootFirst
		possibleDeuteriumLoot -= deuteriumLootFirst
		// step 4
		if capacityLeft > 0 {
			metalLootSecond := int(math.Min(float64(capacityLeft)/2.0, float64(possibleMetalLoot)))
			capacityLeft -= metalLootSecond
			metalLoot += metalLootSecond
			possibleMetalLoot -= metalLootSecond
			// step 5
			crystalLootSecond := int(math.Min(float64(capacityLeft), float64(possibleCrystalLoot)))
			capacityLeft -= crystalLootSecond
			crystalLoot += crystalLootSecond
			possibleCrystalLoot -= crystalLootSecond
		}
	}

	return 3*fi.flightCosts + fi.attackersLost - 3*deuteriumLoot - int(1.5*float64(crystalLoot)) - metalLoot
}

func (fi *fight) goLower(toChange *int, lowestValue int, bestLosses float64) float64 {
	currentValue := *toChange
	stepSize := int(math.Max(1, float64(currentValue-lowestValue)*float64(toChangePercent)/100.0))
	currentValue -= stepSize

	for currentValue >= lowestValue {
		oldValue := *toChange
		*toChange = currentValue

		lossesSum := fi.doFightsAndReturnLosses()
		if lossesSum < bestLosses {
			bestLosses = lossesSum
			fmt.Println("Current losses through going down: ", bestLosses)
		} else {
			*toChange = oldValue
		}

		currentValue -= stepSize
	}

	return bestLosses
}

func (fi *fight) goHigher(toChange *int, highestValue int, bestLosses float64) float64 {
	currentValue := *toChange
	stepSize := int(math.Max(1, float64(highestValue-currentValue)*float64(toChangePercent)/100.0))
	currentValue += stepSize

	for currentValue <= highestValue {
		oldValue := *toChange
		*toChange = currentValue

		lossesSum := fi.doFightsAndReturnLosses()
		if lossesSum < bestLosses {
			bestLosses = lossesSum
			fmt.Println("Current losses through going up: ", bestLosses)
		} else {
			*toChange = oldValue
		}

		currentValue += stepSize
	}

	return bestLosses
}

func (fi *fight) optimizeAttackers() {
	metaCurrentLosses, metaBestLosses := fi.doFightsAndReturnLosses(), math.MaxFloat64

	for metaCurrentLosses < metaBestLosses {
		metaBestLosses = metaCurrentLosses

		highestShipCounts := make([]int, 0)
		runtime.GC()
		for i := range fi.attackers {
			realI := len(fi.attackers) - 1 - i
			for _, ship := range fi.attackers[realI].ships {
				highestShipCounts = append(highestShipCounts, ship.getCount())
			}
		}

		currentLosses, bestLosses := metaBestLosses, math.MaxFloat64

		for currentLosses < bestLosses {
			bestLosses = currentLosses

			for i := range fi.attackers {
				realI := len(fi.attackers) - 1 - i
				for k := range fi.attackers[realI].ships {
					currentLosses = fi.goLower(fi.attackers[realI].ships[k].getCountPointer(), fi.attackers[realI].ships[k].getMinimumCount(), currentLosses)
				}
			}

			counter := 0
			for i := range fi.attackers {
				realI := len(fi.attackers) - 1 - i
				for k := range fi.attackers[realI].ships {
					currentLosses = fi.goHigher(fi.attackers[realI].ships[k].getCountPointer(), highestShipCounts[counter], currentLosses)
					counter++
				}
			}
		}

		counter := 0
		metaCurrentLosses = bestLosses
		for i := range fi.attackers {
			realI := len(fi.attackers) - 1 - i
			for k := range fi.attackers[realI].ships {
				metaCurrentLosses = fi.goHigher(fi.attackers[realI].ships[k].getCountPointer(), highestShipCounts[counter], metaCurrentLosses)
				counter++
			}
		}
	}
}

func (fi *fight) doFightsAndReturnLosses() float64 {
	lossesSum := 0.0
	for i := 0; i < howOftenForAverage; i++ {
		fi.setupFight()
		fi.doFight()
		lossesSum += float64(fi.attackerTotalLost())
	}
	lossesSum /= float64(howOftenForAverage)
	return lossesSum
}

func (fi *fight) doFight() {
	for i := 0; i < 6; i++ {
		attackerAlive, defenderAlive := len(fi.attackerObjectsAlive), len(fi.defenderObjectsAlive)

		if attackerAlive == 0 || defenderAlive == 0 {
			break
		}

		for _, obj := range fi.attackerObjectsAlive {
			obj.setCurrentShield(obj.getMaxShield())
		}
		for _, obj := range fi.defenderObjectsAlive {
			obj.setCurrentShield(obj.getMaxShield())
		}

		doRound(fi.attackerObjectsAlive, fi.defenderObjectsAlive, defenderAlive)
		doRound(fi.defenderObjectsAlive, fi.attackerObjectsAlive, attackerAlive)

		index := 0
		for _, obj := range fi.attackerObjectsAlive {
			if !obj.getIsExploded() {
				fi.attackerObjectsAlive[index] = obj
				index++
				continue
			}

			fi.totalCapacity -= obj.getCapacity()
			fi.attackersLost += 3*obj.getDeuteriumCost() + int((1.0-float64(debrisPercent)/100.0)*float64(float64(obj.getMetalCost())+1.5*float64(obj.getCrystalCost())))
		}
		for delIndex := index; delIndex < len(fi.attackerObjectsAlive); delIndex++ {
			fi.attackerObjectsAlive[delIndex] = nil
		}
		fi.attackerObjectsAlive = fi.attackerObjectsAlive[:index]

		index = 0
		for _, obj := range fi.defenderObjectsAlive {
			if !obj.getIsExploded() {
				fi.defenderObjectsAlive[index] = obj
				index++
				continue
			}

			if obj.isShip() || defenseIntoDebris {
				fi.attackersLost -= int((float64(debrisPercent) / 100.0) * float64(float64(obj.getMetalCost())+1.5*float64(obj.getCrystalCost())))
			}
		}
		for delIndex := index; delIndex < len(fi.defenderObjectsAlive); delIndex++ {
			fi.defenderObjectsAlive[delIndex] = nil
		}
		fi.defenderObjectsAlive = fi.defenderObjectsAlive[:index]
		runtime.GC()
	}
}

func getNameCoordStrings(players []player) []string {
	returnStrings := make([]string, 0)
	for _, player := range players {
		returnStrings = append(returnStrings, player.name+" ["+strconv.Itoa(player.galaxy)+":"+strconv.Itoa(player.system)+":"+strconv.Itoa(player.planet)+"]")
	}
	return returnStrings
}

func getShipsNamesAndCount(shipsOrDefenses []objectInFight) []string {
	if len(shipsOrDefenses) == 0 {
		return []string{}
	}

	returnStrings := make([]string, 0)

	name := ""
	count := 0
	for _, obj := range shipsOrDefenses {
		if count == 0 {
			count = 1
			name = obj.getName()
		} else if name == obj.getName() {
			count += 1
		} else {
			returnStrings = append(returnStrings, name+": "+strconv.Itoa(count))
			count = 1
			name = obj.getName()
		}
	}
	returnStrings = append(returnStrings, name+": "+strconv.Itoa(count))

	return returnStrings
}

func (fi *fight) printFight() {
	fi.setupFight()

	attackerStrings := getNameCoordStrings(fi.attackers)
	defenderStrings := getNameCoordStrings(fi.defenders)
	fmt.Println("\nAttacker(s): " + strings.Join(attackerStrings, ", "))
	fmt.Println("Defender(s): " + strings.Join(defenderStrings, ", "))
	fmt.Println("\n--- Before the fight ---")

	attackerObjsAlive := getShipsNamesAndCount(fi.attackerObjectsAlive)
	defenderObjsAlive := getShipsNamesAndCount(fi.defenderObjectsAlive)

	fmt.Println("Attacker(s): " + strings.Join(attackerObjsAlive, ", "))
	fmt.Println("Defender(s): " + strings.Join(defenderObjsAlive, ", "))

	fi.doFight()
	fmt.Println("\n--- After the fight ---")
	attackerObjsAlive = getShipsNamesAndCount(fi.attackerObjectsAlive)
	defenderObjsAlive = getShipsNamesAndCount(fi.defenderObjectsAlive)

	fmt.Println("Attacker(s): " + strings.Join(attackerObjsAlive, ", "))
	fmt.Println("Defender(s): " + strings.Join(defenderObjsAlive, ", "))
	fmt.Println("\n--- Summation ---")
	fmt.Println("Flight costs in deuterium: " + strconv.Itoa(fi.flightCosts))
	fmt.Println("Losses all in all in MSE: " + strconv.Itoa(fi.attackerTotalLost()))
}

func (fi *fight) setupFight() {
	for delIndex := 0; delIndex < len(fi.attackerObjectsAlive); delIndex++ {
		fi.attackerObjectsAlive[delIndex] = nil
	}
	for delIndex := 0; delIndex < len(fi.defenderObjectsAlive); delIndex++ {
		fi.defenderObjectsAlive[delIndex] = nil
	}
	fi.attackerObjectsAlive = fi.attackerObjectsAlive[:0]
	fi.defenderObjectsAlive = fi.defenderObjectsAlive[:0]
	runtime.GC()
	fi.flightCosts = 0
	fi.attackersLost = 0
	fi.totalCapacity = 0

	for _, attacker := range fi.attackers {
		minSpeed := math.MaxInt32
		for _, fightObject := range attacker.ships {
			fi.totalCapacity += fightObject.getCount() * fightObject.getShip().getCapacity()
			if fightObject.getCount() > 0 {
				currSpeed := fightObject.getShip().getShipSpeed(&attacker)
				if currSpeed < minSpeed {
					minSpeed = currSpeed
				}
			}
			maxHull, maxShield, attack := 0.001*math.Round(1000.0*fightObject.getObject().getBaseHull()*(1.0+0.1*float64(attacker.hullResearch))),
				0.001*math.Round(1000.0*fightObject.getObject().getBaseShield()*(1.0+0.1*float64(attacker.shieldResearch))),
				0.001*math.Round(1000.0*fightObject.getObject().getBaseAttack()*(1.0+0.1*float64(attacker.weaponResearch)))

			for i := 0; i < fightObject.getCount(); i++ {
				fi.attackerObjectsAlive = append(fi.attackerObjectsAlive, &concreteFightObject{
					fightObject.getObject(),
					maxHull, maxHull,
					maxShield, maxShield,
					attack,
					false,
				})
			}
		}
		distance := attacker.getDistanceTo(fi.defenders[0].galaxy, fi.defenders[0].system, fi.defenders[0].planet)
		flightTime := getFlightTime(distance, minSpeed)
		for _, fightObject := range attacker.ships {
			if fightObject.getCount() > 0 {
				fi.flightCosts += fightObject.getShip().getFuel(distance, flightTime, fightObject.getCount(), &attacker)
			}
		}
	}

	for _, defender := range fi.defenders {
		for _, fightObject := range defender.ships {
			maxHull, maxShield, attack := 0.001*math.Round(1000.0*fightObject.getObject().getBaseHull()*(1.0+0.1*float64(defender.hullResearch))),
				0.001*math.Round(1000.0*fightObject.getObject().getBaseShield()*(1.0+0.1*float64(defender.shieldResearch))),
				0.001*math.Round(1000.0*fightObject.getObject().getBaseAttack()*(1.0+0.1*float64(defender.weaponResearch)))

			for i := 0; i < fightObject.getCount(); i++ {
				fi.defenderObjectsAlive = append(fi.defenderObjectsAlive, &concreteFightObject{
					fightObject.getObject(),
					maxHull, maxHull,
					maxShield, maxShield,
					attack,
					false,
				})
			}
		}
		for _, fightObject := range defender.defenses {
			maxHull, maxShield, attack := 0.001*math.Round(1000.0*fightObject.getObject().getBaseHull()*(1.0+0.1*float64(defender.hullResearch))),
				0.001*math.Round(1000.0*fightObject.getObject().getBaseShield()*(1.0+0.1*float64(defender.shieldResearch))),
				0.001*math.Round(1000.0*fightObject.getObject().getBaseAttack()*(1.0+0.1*float64(defender.weaponResearch)))

			for i := 0; i < fightObject.getCount(); i++ {
				fi.defenderObjectsAlive = append(fi.defenderObjectsAlive, &concreteFightObject{
					fightObject.getObject(),
					maxHull, maxHull,
					maxShield, maxShield,
					attack,
					false,
				})
			}
		}
	}
}

func createFight(attackers, defenders []player) fight {
	toReturnFight := fight{
		attackers, defenders,
		nil, nil, 0, 0, 0,
	}
	numberOfShipsAttacker, numberOfShipsDefender := 0, 0
	for _, attacker := range attackers {
		for _, fightObject := range attacker.ships {
			numberOfShipsAttacker += fightObject.getCount()
		}
	}
	for _, defender := range defenders {
		for _, fightObject := range defender.ships {
			numberOfShipsDefender += fightObject.getCount()
		}
		for _, fightObject := range defender.defenses {
			numberOfShipsDefender += fightObject.getCount()
		}
	}

	toReturnFight.attackerObjectsAlive = make([]objectInFight, 0, numberOfShipsAttacker)

	toReturnFight.defenderObjectsAlive = make([]objectInFight, 0, numberOfShipsDefender)

	return toReturnFight
}

var knownShips = []shipBase{
	{
		movingDataBase{
			stationaryDataBase{
				"Kleiner Transporter", 2000, 2000, 0, 10, 5, 0,
			},
			10000, 20,
		},
		[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		5000, 1,
	},
	{
		movingDataBase{
			stationaryDataBase{
				"Großer Transporter", 6000, 6000, 0, 25, 5, 1,
			},
			7500, 50,
		},
		[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		25000, 0,
	},
	{
		movingDataBase{
			stationaryDataBase{
				"Leichter Jäger", 3000, 1000, 0, 10, 50, 2,
			},
			12500, 20,
		},
		[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		50, 0,
	},
	{
		movingDataBase{
			stationaryDataBase{
				"Schwerer Jäger", 6000, 4000, 0, 25, 150, 3,
			},
			10000, 75,
		},
		[]int{
			3, 0, 0, 0, 0, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		100, 1,
	},
	{
		movingDataBase{
			stationaryDataBase{
				"Kreuzer", 20000, 7000, 2000, 50, 400, 4,
			},
			15000, 300,
		},
		[]int{
			0, 0, 6, 0, 0, 0, 0, 0, 5, 0, 5, 0, 0, 0, 10, 0, 0, 0, 0, 0, 0, 0,
		},
		800, 1,
	},
	{
		movingDataBase{
			stationaryDataBase{
				"Schlachtschiff", 45000, 15000, 0, 200, 1000, 5,
			},
			10000, 500,
		},
		[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		1500, 2,
	},
	{
		movingDataBase{
			stationaryDataBase{
				"Kolonieschiff", 10000, 20000, 10000, 100, 50, 6,
			},
			2500, 1000,
		},
		[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		7500, 1,
	},
	{
		movingDataBase{
			stationaryDataBase{
				"Recycler", 10000, 6000, 2000, 10, 1, 7,
			},
			2000, 300,
		},
		[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		20000, 0,
	},
	{
		movingDataBase{
			stationaryDataBase{
				"Spionagesonde", 0, 1000, 0, 0.01, 0.01, 8,
			},
			100000000, 1,
		},
		[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		0, 0,
	},
	{
		movingDataBase{
			stationaryDataBase{
				"Bomber", 50000, 25000, 15000, 500, 1000, 9,
			},
			5000, 1000,
		},
		[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 5, 0, 0, 0, 20, 20, 10, 0, 10, 0, 0, 0,
		},
		500, 2,
	},
	{
		movingDataBase{
			stationaryDataBase{
				"Solarsatellit", 0, 2000, 500, 1, 1, 10,
			},
			0, 0,
		},
		[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		0, 0,
	},
	{
		movingDataBase{
			stationaryDataBase{
				"Zerstörer", 60000, 50000, 15000, 500, 2000, 11,
			},
			5000, 1000,
		},
		[]int{
			0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 5, 0, 0, 2, 0, 10, 0, 0, 0, 0, 0, 0,
		},
		2000, 2,
	},
	{
		movingDataBase{
			stationaryDataBase{
				"Todesstern", 5000000, 4000000, 1000000, 50000, 200000, 12,
			},
			100, 1,
		},
		[]int{
			250, 250, 200, 100, 33, 30, 250, 250, 1250, 25, 1250, 5, 0, 15, 200, 200, 100, 50, 100, 0, 0, 0,
		},
		1000000, 2,
	},
	{
		movingDataBase{
			stationaryDataBase{
				"Schlachtkreuzer", 30000, 40000, 15000, 400, 700, 13,
			},
			10000, 250,
		},
		[]int{
			3, 3, 0, 4, 4, 7, 0, 0, 5, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		750, 2,
	},
}

var knownDefenses = []defenseBase{
	{
		stationaryDataBase{
			"Raketenwerfer", 2000, 0, 0, 20, 80, 14,
		},
	},
	{
		stationaryDataBase{
			"Leichtes Lasergeschütz", 1500, 500, 0, 25, 100, 15,
		},
	},
	{
		stationaryDataBase{
			"Schweres Lasergeschütz", 6000, 2000, 0, 100, 250, 16,
		},
	},
	{
		stationaryDataBase{
			"Gaußkanone", 20000, 15000, 2000, 200, 1100, 17,
		},
	},
	{
		stationaryDataBase{
			"Ionengeschütz", 2000, 6000, 0, 500, 150, 18,
		},
	},
	{
		stationaryDataBase{
			"Plasmawerfer", 50000, 50000, 30000, 300, 3000, 19,
		},
	},
	{
		stationaryDataBase{
			"Kleine Schildkuppel", 10000, 10000, 0, 2000, 1, 20,
		},
	},
	{
		stationaryDataBase{
			"Große Schildkuppel", 50000, 50000, 0, 10000, 1, 21,
		},
	},
}
