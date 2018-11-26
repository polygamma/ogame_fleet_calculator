# ogame_fleet_calculator
Calculator for nearly optimal fleet composition

- https://golang.org/dl/
- change variables in `simulator.gò` if needed, lines 14-22
- set players in `func main()`, e.g.

```
onePlayer := player{
    "nameOfPlayer", 2, 3, 4, 15, 15, 15,
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
}
```

This player has the name `nameOfPlayer`, coords of the planet `[2:3:4]`, `15,15,15` techs, has no ships but a bit of defense.

```
anotherPlayer := player{
    "evenBetterName", 1, 1, 1, 14, 14, 14,
    []playerShip{
        {&knownShips[0], 0, 0},    // Kleiner Transporter
        {&knownShips[1], 0, 0},    // Großer Transporter
        {&knownShips[2], 0, 0},    // Leichter Jäger
        {&knownShips[3], 0, 0},    // Schwerer Jäger
        {&knownShips[4], 0, 0},    // Kreuzer
        {&knownShips[5], 500, 0},  // Schlachtschiff
        {&knownShips[6], 0, 0},    // Kolonieschiff
        {&knownShips[7], 0, 0},    // Recycler
        {&knownShips[8], 0, 0},    // Spionagesonde
        {&knownShips[9], 100, 0},  // Bomber
        {&knownShips[11], 100, 0}, // Zerstörer
        {&knownShips[12], 0, 0},   // Todesstern
        {&knownShips[13], 100, 50}, // Schlachtkreuzer
    },
    []playerDefense{},
}
```

Player with name `evenBetterName` at `[1:1:1]` with `14,14,14` techs.
Has 500 battleships, 100 bomber, 100 destroyer, 100 battlecruiser AND wants to use AT LEAST 50 of them.

- run the program with `go run simulator.go`
