<p align="center" style="font-size:32pt;font-style:bold">
    <img src="docs/images/logo.png" width=400>
</p>
<p align="center" style="font-size:12pt;font-style:bold">
    <b>OpenNox</b> is an open-source community collaboration project extending the Nox engine. 
</p>
<p align="center">
    <a href="https://github.com/noxworld-dev/opennox/actions"><img alt="OpenNox Build Status (dev)" src="https://github.com/noxworld-dev/opennox/actions/workflows/build-and-release.yml/badge.svg"></a>
    <a href="https://www.gnu.org/licenses/gpl-3.0.en.html"><img alt="OpenNox license" src="https://img.shields.io/github/license/noxworld-dev/opennox?style=flat"></a>
    <br>
    <a href="https://www.patreon.com/opennox"><img alt="OpenNox on Patreon" src="https://img.shields.io/badge/patreon-Support%20us-blue?logo=patreon&logoColor=white&style=flat"></a>
    <a href="https://discord.gg/HgDUeXhAyW"><img alt="OpenNox on Discord" src="https://img.shields.io/badge/discord-OpenNox-blue?logo=discord&logoColor=white&style=flat"></a>
    <a href="https://matrix.to/#/#opennox:nwca.xyz"><img alt="OpenNox on Matrix" src="https://img.shields.io/badge/matrix-%23opennox-blue?logo=matrix&logoColor=white&style=flat"></a>
</p>

## Features

OpenNox supports all vanilla Nox features. You should be able to complete the campaign and play online with OpenNox.
If something doesn't work, please [open an issue](https://github.com/noxworld-dev/opennox/issues/new/choose).

For a list of new features see [this page](https://noxworld-dev.github.io/opennox-docs/opennox/features/index.html).

## Cooperative Campaign

This branch adds cooperative play for the original single-player campaigns. The Conjurer, Warrior and Wizard story maps can be played through together online, with the same chapters, quests, cutscenes and progression as solo play.

**How to play**

1. Every player needs this build of OpenNox and a copy of Nox (e.g. the GOG release).
2. One player picks **Multiplayer → Host Game**, sets **Game Type** to **Coop**, picks a campaign map (such as `con01a`, `war01a` or `wiz01a`) and presses **Go**.
3. Other players join through the normal **Join Game** browser.

**How it works**

- Every player creates their own character and class. Warriors learn abilities by leveling, while Wizards and Conjurers learn spells from spellbooks, exactly like single-player.
- Campaign progress is kept per character, separate from the multiplayer character's normal all-powers setup. Each character's coop stats, spells, abilities and gear are saved to `save/coop/<name>.plr` at every chapter transition and when they leave, and restored when they join or the next chapter loads — so players can drop out and rejoin mid-campaign without losing progress.
- Loot is instanced: each player picks up their own copy of items found in the world, which then disappears only for them, so nobody misses the gear or spellbooks the campaign expects them to have.
- Shops are instanced too: every player sees their own stock at shopkeepers, and a purchase only removes it from their own inventory list, never from someone else's.
- Players who join mid-campaign spawn near the host wherever the party currently is on the map, instead of back at the map's starting point.
- Experience is shared evenly: every player is awarded the same XP the game would grant in single-player, so the party levels at the pace the campaign was designed for.
- Monsters scale with the number of players to keep fights challenging.
- Chapter exits wait for the whole party: everyone alive must reach the exit before the next chapter loads. Players who die respawn and are revived on the next map.
- When a cutscene starts, the party is gathered around the player who triggered it and input is frozen until the scene ends.

**Difficulty settings**

Monster strength can be tuned in `opennox.yml` (created next to the game data after first launch). Each setting is a multiplier, and the "per player" variants add that amount for every player beyond the first:

```yaml
game:
  coop:
    enemy_health: 1.0             # base health multiplier
    enemy_health_per_player: 1.0  # extra health per additional player
    enemy_damage: 1.0             # base damage multiplier
    enemy_damage_per_player: 0.5  # extra damage per additional player
    enemy_speed: 1.0              # base speed multiplier
    enemy_speed_per_player: 0.0   # extra speed per additional player
```

With the defaults, two players face monsters with 2× health and 1.5× damage, three players 3× health and 2× damage, and so on. Raise or lower the numbers to taste.

## Download OpenNox

<a href="https://github.com/noxworld-dev/opennox/releases"><img alt="OpenNox releases" src="https://img.shields.io/github/downloads/noxworld-dev/opennox/total?style=flat&label=releases"></a>
<a href="https://snapcraft.io/opennox"><img alt="OpenNox Snap package" src="https://img.shields.io/badge/snap-Install-green?logo=snapcraft&logoColor=white&style=flat"></a>

### Release
All release builds are made from the `dev` branch. Recent OpenNox releases can be found [here](<https://github.com/noxworld-dev/opennox/releases>).

Linux releases are also available in `stable` channel of our [Snap package](https://snapcraft.io/opennox).

### Nightly
On each commit, an automated build of the `dev` branch is uploaded.
These builds contain all the latest merged features, but are not yet considered stable for release.
These builds are to help provide an insight to what the next release will contain and should only be used for active playtesting purposes **only**.

Linux nightly builds are also available in `edge` channel of our [Snap package](https://snapcraft.io/opennox).

## Build OpenNox
**NOTE: This section is only for people who wish to build the source code locally.**

### Linux
- [Linux](./docs/build-linux.md)
  
### Windows
- [Windows](./docs/build-windows.md)
- [Windows (on Linux)](./docs/build-windows-on-linux.md)

## Contributing
Read [CONTRIBUTING](CONTRIBUTING.md)!

## Legal

<a href="https://www.gnu.org/licenses/gpl-3.0.en.html"><img alt="OpenNox license" src="https://img.shields.io/github/license/noxworld-dev/opennox?style=flat"></a>

This project (OpenNox) is an unofficial community collaboration project for preservation, modding and compatibility purposes.
This project has no direct affiliation with Electronic Arts Inc. and/or the "Nox" brand. "Nox" is an Electronic Arts Inc. brand. All Rights Reserved.

No assets, texts, artwork or other media from the original game(s) is included in this project.
We do not condone piracy in any way, shape or form and encourage users to legally own the original game.

The video game "Nox" is copyright © 2000 Westwood Studios. All Rights Reserved.
Westwood Studios is a trademark or registered trademark of Electronic Arts in the U.S. and/or other countries. All rights reserved.

If not specified otherwise, the source code provided in this repository is licenced under the [GNU General Public License version 3](<https://www.gnu.org/licenses/gpl-3.0.html>). Please see the accompanying LICENSE file.

OpenNox logo created by [@CCHyper](https://github.com/CCHyper) under [CC0 license](https://creativecommons.org/share-your-work/public-domain/cc0/).

OpenNox project additionally follows [C&C Remastered Modding guideline](https://www.ea.com/games/command-and-conquer/command-and-conquer-remastered/modding-faq). All changes to the project MUST follow these rules.
