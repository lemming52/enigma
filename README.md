# Enigma

```
import "github.com/lemming52/enigma"
```

An exercise of a golang package, enigma provides a functioning replication of the enigma machine, and can be configured using a subset of the available rotors. The configured machine can then be used to encode alphanumeric strings, with an ugly output.

The machine functions identically to the physical machine, and can be configured using the provided rotors and reflectors. In theory the number of rotors can be extended beyond the four rotor limit of the physical machine, but the stepping motion, akin to the physical machine is limited to the first 3 rotors.


## Usage

Instantiate a machine with specific rotor configurations, from right to left on the physical machine:
```
em, err := enigma.New([]RotorConfiguration{
    {
        name: RotorIII,
        position: 5,
        ringSetting: 6
    },{
        name: RotorII,
        position: 23,
        ringSetting: 1
    },{
        name: RotorI,
        position: 0,
        ringSetting: 0
    }
}, ReflectorB, "AZ BC XT")
```

using the available rotors. The plugboard can be up to 10 pairs if letters and the positions and ring settings start from 0.

To encode a string
```
cipher := em.Encode("AAAAA")
>>> JTUJZ
```

### Available Components

Rotors | Reflectors
------ | ----------
RotorI | ReflectorA
RotorII | ReflectorB
RotorIII | ReflectorC
RotorIV | ReflectorBThin
RotorV | ReflectorCThing
RotorVI |
RotorVII |
RotorVIII |
RotorBeta |
RotorGamma

## Limitations

* Available characters are only alphanumeric, no punctuation. Numeric characters and whitespace are preserved.
* The output is not pretty-printed into the 4 block characters as seen in authentic messages
* None of the enigma practice is included, starting and ending messages with same string and so on
* Probably very fragile
* Minimal reporting, would be nice to turn on a setting to show complete traversal
* Any kind of interface, be it visual or cli
* Package organisation needs work
* Component list is incomplete
* Ability to generate either a random enigma machine, or a prexisting model

### Sources

Naturally the function here is derivative, predominantly from the [rotor details page](https://en.wikipedia.org/wiki/Enigma_rotor_details) of wikipedia, with the assistance of two existing enigma emulators:

* The Cryptii online tool: https://cryptii.com/pipes/enigma-machine
* Leelar Thaophialuang / Piotte13's enigma simulator: https://piotte13.github.io/enigma-cipher/, this features terrific visualisation of the rotor traversal which helped debug
