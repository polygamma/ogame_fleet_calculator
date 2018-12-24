# ogame_fleet_calculator
Calculator for nearly optimal fleet composition

- https://golang.org/dl/
- change variables in `simulator.go` if needed, lines 14-21
- set players in `func main()`, e.g.

```
polygammaOne := player{
    "polygamma", 1, 2, 3,
    14, 15, 16,
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
```

This player has the name `polygamma` with coords of the planet `[1:2:3]`.
`14,15,16` techs and `15,13,9` drives.
Has 500 battleships AND wants to use at least 250 of them, 100 bomber, 100 destroyer and 100 battlecruiser.
He also has 1000 rocket launchers, 500 light lasers, 20 gauss cannons and 10 plasma turrets.
He also has 2kk metal, 1kk crystal and 1.5kk deuterium and you can get 50 percent loot of him.
Notice: This example does not make sense as is, since setting to use at least 250 of the 500 battleships will only be considered for attackers, but it explains the syntax.

- run the program with `go run simulator.go`
