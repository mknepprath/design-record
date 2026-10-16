---
kind: spec
number: 0002
title: Game of Life design
status: accepted
owner: Michael Knepprath
date: 2026-10-16
projects: [game-of-life]
evidence:
  - evidence/observations/2026-10-16-game-of-life-test-run.md
---

> Moved from `DESIGN.md` in `gol-rpg` at commit `d9ec957` (2026-08-16) on 2026-10-16. The text below is unchanged. Section numbers are unchanged because gol-rpg source files cite them.

# Game of Life — Design Spec

A turn-based roguelike where your balance sheet is your character sheet. Old-school
Final Fantasy grammar: menus, upkeep, status effects, a slow-charging win condition,
and encounters that can end a run.

Fully simulated. No real financial data enters this system.

---

## 1. The core loop

One **turn = one month**. A **run** = 120 turns (10 years) unless you die first.

```
INCOME      paycheck lands, income-affecting statuses tick
UPKEEP      rent/food/insurance premiums/debt minimums drain automatically
ACTION      you allocate what's left  ← the only place you make decisions
ENCOUNTER   draw from the event deck; insurance blocks, shield absorbs
INTEREST    debts compound
MARKET      investments roll
SETTLE      credit score updates, HP regen, sweep leftovers, check win/loss
```

The player only ever touches ACTION. Everything else is the world happening to them.
That framing is the point: most of a financial month is not a choice.

---

## 1.5 Two vocabularies

**Internally the model is financial. Externally the game is fantasy.** The split
is load-bearing: every parameter in here has to stay auditable against real-world
data (§10 requires sources), and it cannot be if the code talks about wards and
burdens.

The rule is: **ids are internal truth, `name` fields are external presentation.**
The engine keys off `'health'`, `'emergency'`, `'student'` forever. Only display
strings change. This document, the code, and every comment use the financial
vocabulary.

| Internal (model, code, this doc) | External (what the player reads) |
|---|---|
| HP / solvency | **Vitality** |
| Emergency fund | **Shield** |
| Income | **Mana** |
| Investments / portfolio | **Ventures** |
| Credit score | **Standing** |
| Credit tiers | Exalted · Honoured · Recognised · Wary · Shunned |
| Insurance policies | **Wards** |
| Health / auto / renters / disability cover | Ward of Vitality · of the Road · of Hearth · of Fortune |
| Universal cover (Norvang) | The Crown's Mercy / Sustenance |
| Debts | **Burdens** |
| Student loan | Scholar's Oath |
| Credit card / emergency credit | Usurer's Line / Usurer's Mark |
| Hospital payment plan | Infirmary Tithe |
| Index / bonds / single stock / crypto | Merchant Consortium · Crown Bonds · A Single House · Alchemists' Speculation |
| Bankruptcy discharge | **Jubilee** |
| Terminal insolvency | **Ruin** |
| Financial independence | **Free of the Wheel** |
| Survived to end of run | **Endured** |

**Events keep their real names.** "Surgery and Recovery", "Dead Battery",
"Apartment Fire". The friction between an ordinary medical bill and a character
sheet that calls it a Burden is the entire point of the game; renaming the
surgery to something with a dragon in it would throw that away.

**No acronyms in the interface.** Playtesting showed `SP`/`CRD`/`INS`/`DBT` are
opaque to anyone who has not read this document. Labels are words.

## 2. Resources

### HP — Solvency (0–100)

**Not dollars.** HP is how close you are to the wheels coming off. You lose HP only
when an expense lands that your cash and shield could not cover.

```
hpDamage = ceil(unabsorbed / netMonthlyIncome * 10)
```

One full month of income you couldn't cover = 10 HP. Ten months = dead. The formula
scales with income, which is correct: a $2,000 bill is a catastrophe at $2,500/mo and
an annoyance at $12,000/mo.

**HP 0 = bankruptcy = run over.**

Regen: `+2 HP/turn` while cash flow is positive *and* shield ≥ one month of expenses.
Deliberately slow. You cannot heal your way out of a spiral; you have to fix the
spiral.

### SP — Emergency Fund (shield)

Dollars. Absorbs damage before HP. Does not regenerate on its own — it is only ever
refilled by you choosing to refill it (or by the end-of-turn sweep, §5).

This is the single most load-bearing stat in the game and the tutorial should make you
feel its absence in the first ten turns.

### MP — Income (regen)

Your paycheck is mana. It arrives each turn, gets partly consumed by upkeep, and the
remainder is your **action budget**. Income is not a stat you can spend up front —
it's a rate. Raising the rate (§7) is the main long-game lever.

### Damage absorption order

```
cash → shield → HP + debt
```

Cash first (you'd really pay from checking), shield second, and only the remainder
converts to damage.

### §2a — a bill you can't pay becomes debt

An expense you cannot cover does not evaporate into HP. **35% of the notional HP
damage lands as immediate shock, capped at 35 HP, and the entire shortfall
becomes `Emergency Credit`** at 22.9% ± your credit tier's modifier — always
Poison.

The cap matters: a catastrophe must not one-shot you from full health. It should
leave you alive and carrying six figures of compounding debt. **The debt is what
kills you, not the bill.** That is what makes prevention worth paying for, and it
is why credit standing (§6) has teeth — poor standing prices your worst moment at
36.9% instead of 18.9%.

---

## 3. Debt — damage over time

Each debt is a **status effect**, not a number. Its APR determines which status:

| APR      | Status   | Feel                                              |
|----------|----------|---------------------------------------------------|
| ≥ 18%    | Poison   | Compounds visibly every turn. Fix it now.         |
| 8–18%    | Bleed    | Real drag, survivable, worth a plan.              |
| < 8%     | Slow     | A debuff, not an emergency. Paying it off early is often a *mistake* vs. investing. |

Monthly interest accrues on the balance: `balance *= (1 + apr/12)`.

Minimum payments are pulled during UPKEEP. **A missed minimum is the worst single
event in the game**: −40 credit score, a `DELINQUENT` status, and the balance keeps
compounding. Missing three in a row triggers default.

The Poison/Slow split is the whole lesson. A player who reflexively kills their 3%
student loan while carrying 24% card debt should lose runs.

## 4. Insurance — conditional block with upkeep

Each policy costs MP every single turn and **negates an event category down to a
deductible**.

| Policy     | Premium/mo | Blocks    | Deductible |
|------------|-----------:|-----------|-----------:|
| Health     |       $250 | MEDICAL   |     $1,500 |
| Auto       |        $95 | VEHICLE   |       $500 |
| Renters    |       $110 | PROPERTY  |       $250 |
| Disability | 2.4% of income | INCOME | 40% passthrough |

Premiums are priced at roughly **1.25× expected claims** — the real-world loss
ratio. Insurance is therefore EV-negative by construction, as it is in life. Its
only value is variance reduction near the ruin boundary.

Coverage starts **one month after purchase**, so you cannot buy a policy the month
you need it.

Disability is priced as a share of income because its payout scales with salary;
a flat premium would make it free money once you're earning well.

The trap is intentional and should be tempting: dropping health insurance frees
$250/mo forever, and in most runs it's free money. The rest end the run. This is
the cleanest fat-tail lesson in the design — never nerf it into safety.

**The success criterion is not "insurance is good."** It is: *insurance loses you
money in the median run and saves you in the bottom decile.* Check both columns in
`npm run sweep` — median and p10 — before changing any number here.

### 4.1 The criterion is met again, narrowly — and one number still disagrees

It was not met for a long time, and this section used to record the failure.
With the stability ladder (§8.2) and Wards counting toward `covered`, buying
cover now moves the bottom decile the right way:

| | p10 | median | terminal ruin |
|---|---:|---:|---:|
| with Wards | −$30,682 | $118,869 | 7.5% |
| no Wards | −$35,676 | $120,259 | 5.0% |

Wards cost about $1,400 of median and buy about $5,000 of bottom decile. That
is the shape §4 asks for.

**But the ruin column runs the other way**: carrying Wards makes you *more*
likely to be ruined outright, because the premiums are themselves what pushes
thin runs under. Net worth in the bad tail improves while the odds of the very
worst outcome get worse. Both are true, and the second is not obviously wrong —
paying for cover you cannot afford is a real way to go under. It is recorded
here rather than tuned away, and `npm run analysis` prints both.

### Cascading riders

Tier-3 events carry a second effect in a **different category**, so a single
catastrophe cannot be defended by a single policy. A surgery bills you (MEDICAL)
*and* stops you working (INCOME) — health covers the first, disability the second.
Real ruin is compound: sick → can't work → can't pay → the debt spirals.

## 5. Investments — scaling, volatile, delayed

Monthly return per asset:

```
return = mean + beta * marketFactor + idiosyncratic
```

- `marketFactor` — one shared roll per turn, `Normal(0, 3.5%)`. Systematic risk.
- `idiosyncratic` — per-asset roll, `Normal(0, assetSigma)`. Specific risk.

| Asset        | mean/mo | beta | sigma |
|--------------|--------:|-----:|------:|
| Index fund   |   0.70% | 1.00 |  0.8% |
| Bond fund    |   0.35% | 0.25 |  0.6% |
| Single stock |   1.00% | 1.20 | 11.0% |
| Crypto       |   1.40% | 1.60 | 25.0% |

**Diversification kills `idiosyncratic`, not `beta * marketFactor`.** Spreading across
four holdings smooths the specific risk and does nothing about a crash. That is the
real behavior of markets and the game should teach exactly that, not "diversify and
you're safe."

Liquidating early costs a **10% penalty** plus whatever the current drawdown is — so
investments are a bad emergency fund, which is the point of having a shield.

## 6. Credit score — faction standing (300–850)

| Change                        | Δ    |
|-------------------------------|-----:|
| All minimums paid this turn   |   +2 |
| Missed minimum                |  −40 |
| Revolving utilization > 30%   |  −5/turn |
| Revolving utilization > 70%   | −12/turn |
| Debt fully retired            |  +15 |

Tiers gate the APR you're offered on new debt:

| Tier      | Score | New-debt APR modifier |
|-----------|------:|----------------------:|
| Excellent |  750+ |                 −4.0% |
| Good      |  700+ |                 −1.5% |
| Fair      |  650+ |                  0.0% |
| Poor      |  580+ |                 +6.0% |
| Subprime  |  <580 |                +14.0% |

Standing you build slowly and lose instantly. Correct.

## 7. Classes and leveling

| Class            | Net income | Variance | Perk                                    |
|------------------|-----------:|---------:|-----------------------------------------|
| Wage Warrior     |     $2,500 |     none | Steady. Nothing surprising, good or bad.|
| Salaryman Paladin|     $3,100 |     none | Health + disability premiums are free.  |
| Freelance Rogue  |     $3,400 |     high | Income swings ±45%. Crits and dry spells.|
| Startup Mage     |     $1,900 |      low | Holds equity that may pay out enormously.|

**Leveling** is the `train` action: spend cash and a turn's discretionary budget for a
chance at a permanent raise. This is the only way to change your income rate, and it
compounds harder than any investment — which is the correct lesson about early-career
finance.

## 8. Win / loss

- **Loss** — HP hits 0. Bankruptcy.
- **Win** — investments ≥ 25× annual expenses (the 4% rule). Financial independence.
- **Survive** — reach turn 120 solvent but short of FI. Scored, not won.

## 8.1 Antes

A run is eight escalating stages, each with a requirement and a deadline. Miss
one and the run ends. **Independence is the final ante, not an early exit** —
ending a run the moment FI is reached truncated the exponential exactly where it
became interesting (10y median $63k → 30y $811k, but p90 barely moved after year
20 because every compounding run left the sample early).

**The ante ladder is a game conceit, not a sourced parameter.** Tuning it is
legitimate design and §10 does not apply. What §10 still forbids is tuning it
*per regime* to equalize win rates.

Two rules learned by measuring where runs actually die:

1. **Requirements must escalate faster than compounding.** A linear ladder
   (6→12→18→25 years of expenses) against an exponential portfolio meant
   clearing year 10 was effectively winning: 37% of deaths fell in the first
   decade and the last thirty years were a victory lap.
2. **Never let the tutorial ante be the deadliest.** The first version killed
   16.3% at ante 1.

Diagnose with `scripts/` sweeps before changing a number — a healthy ladder
kills across every rung rather than at one wall.

## 8.2 What the game should reward — stability, not a lottery ticket

**The behaviour we incentivise is a stable quality of life, not a slim chance at
enormous wealth.** This is a standing directive for all future ante tuning and
should override the convenient thing.

The current ladder does not honour it. Antes 6–8 ask for 9, 16 and 25 years of
expenses held in Ventures — an accumulation ladder. It rewards whoever earns
most (Salaryman and Rogue clear it 80–99% of the time), punishes the Wage
Warrior, and quietly frames Meridian's enormous ceiling as the prize. That is
backwards: in Meridian, chasing the ceiling means carrying the variance, and the
variance is what ruins people.

The replacement should measure whether a life *holds*:

- runway sustained, not touched once — stability is a duration, not a snapshot
- every exposure actually covered by a Ward
- no Poison debt
- income covering committed costs with real margin
- none of it collapsing for years at a time

The requirement vocabulary for this already exists in `AnteRequirement`
(`covered`, `cashFlowPositive`, `all`, `sustained`) and `anteStreak` is tracked;
no ante uses them yet.

Expected consequences when it is built, all of them improvements:

1. High earners stop coasting — accumulation is trivial for them, sustained
   stability is not.
2. Regimes stop producing 0.5%-vs-99% win rates, because a stability bar is
   reachable in a poor economy where a wealth bar is not — **without**
   equalizing outcomes, which §10 still forbids.
3. Financial independence becomes an optional flourish rather than the win,
   which is a truer statement of what the model has been saying all along.

## 9. Design rules

1. **The player only acts in ACTION.** Resist adding interactive prompts elsewhere.
2. **No choice is strictly dominant.** If one allocation always wins, the numbers are
   wrong. This governs decisions *inside* a run — how you allocate a month's budget.
   It does **not** govern the world you start in; see rule 6.
3. **Fat tails stay fat.** Insurance and crashes must be able to end runs.
4. **Determinism.** Every run is a seed. Same seed + same actions = same run, always.
5. **The engine never prints.** Rendering is somebody else's job.
6. **Never equalize regimes.** If a system is genuinely higher-variance or harder to
   survive, the model says so. See §10.
7. **A warning you can ignore for free is not a warning.** See §9.1.

### 9.1 Halts: transitions speak first, conditions speak again

Travel stops when the plan stops making sense, and every stop tests a
*transition* — "did this just become true" — because testing the condition stops
the road every month and the road never moves.

That rule alone had a hole. Ignoring a halt is one click and costs nothing, so a
condition you shrugged off was silenced for the rest of the run: your Shield
breaks its floor in month 40, you carry on, and sixty months later nothing has
mentioned it again. The halt that matters most was the cheapest to click past.

So the two halts that describe a *persisting* condition re-arm — the Shield floor
after 12 months, a deadline you are not on pace for after 3, since its window is
only 6 wide. The transition still earns the first warning; re-arming only extends
it. A halt that has never fired cannot re-arm, or a run that *begins* under its
Shield floor is stopped on turn one by something it never crossed into.

The other five halts (a gate cleared, a bill your plan cannot answer, a new
Burden, a raise or a layoff) are events, not conditions. They happen once and
have nothing to repeat.

Two things follow for the UI. A halt must be drawn against the control that
would ignore it, not elsewhere on the screen — the reason and the button are one
idea. And a halt that hands the player a real choice (a turning point) takes the
screen to itself, because a "carry on" button beside an offer is a way to lose it
without deciding to.

## 10. Regimes

Regimes are the country you were born in: healthcare, education debt, safety net,
taxes, housing. They are **difficulty and context selection, not balanced
alternatives.** A regime is not something the character chooses, so rule 2 does not
apply to it.

**We do not equalize outcomes across regimes.** If one produces fewer ruined lives,
the sweep reports that, and we ship it. Flattening the difference to make every mode
"viable" would be a lie, and it would delete the only real result this simulation can
produce. Players who want the higher-variance world can choose it — that's a
challenge mode, and challenge modes are supposed to be harder.

Honesty runs in both directions, though. Each regime must carry its **real** costs,
not a curated subset:

- higher marginal tax wedges and lower take-home pay
- lower top-end and equity outcomes (the Startup Mage archetype is not equally real
  everywhere)
- waiting times and coverage gaps where they genuinely exist

Model those because they are true, never to force parity. Then report whatever the
sweep says.

**Parameters must be sourced.** Effective tax wedges, unemployment replacement
rates, student debt loads, and medical-bankruptcy incidence are all documented. Cite
the source next to the number in the regime table. Fantasy names give us freedom to
isolate one variable at a time; they are not a licence to invent the numbers. Record
each regime's real-world referent so the model stays auditable.

The failure mode to avoid is a game that flatters one system by omission. The
defence is that every parameter is written down with a source, and the sweep is run
across regimes and classes before anything ships.

### 10.1 Regimes differ in spread, not just average

The honest difference between regimes is **the shape of the outcome distribution**,
and that is also the part a single-player game hides by default. A player only ever
sees their own run, so inequality is invisible unless we deliberately show it.

Meridian should produce a genuinely **fat right tail** — equity exits and windfalls
that are absurd, not merely comfortable — sitting alongside a fat left tail of ruin.
Its story is winners with far more than they need and losers with nothing. A
compressed regime produces neither: fewer enormous exits, far fewer six-figure
holes, most runs landing somewhere reasonable. Higher floor, lower ceiling, and a
much narrower middle.

That means a regime is not tuned by its median. It is tuned by **p10, median, p90,
and the spread between them**, and a regime whose median improves while its p10
collapses has got worse, not better.

**Measuring p10 correctly.** p10 converges acceptably at a few thousand paired
runs — two independent subsamples of 2,000 agree within a few percent. If you
write a convergence check, **split the sample by seed before sorting**. Slicing
an already-sorted array compares the bottom half of the distribution against the
top half and will tell you, wrongly, that nothing converges.

## 11. Bankruptcy discharge

Debt used to compound forever: nothing discharged it, no ceiling bounded it, and
a run could end carrying negative seven figures that no real person would ever
carry. Insolvency is the mechanism that bounds losses in life, and it is now the
mechanism that bounds the left tail here.

**Discharge is a setback, not a game over.** When HP reaches 0 and relief is
available:

- unsecured debt is wiped — **student debt survives**, as it does in life
- savings above the exemption and most of the portfolio are liquidated
- credit drops to the floor, so future borrowing is priced at Subprime
- HP is restored to 40 and the run continues

Relief cannot be granted again for **96 turns** (the real Chapter 7 rule is eight
years), so it is effectively once per run. Reaching 0 HP again inside that window
is the true loss condition — `outcome: 'bankrupt'` now means *insolvent with no
relief available*.

Effect on the tail, 4,000 paired runs, Wage Warrior: p10 moved from **−$475,740
to −$287,804**. Classes that never become insolvent are unaffected, as they
should be.

### 11.1 Calibrating insolvency against reality

**The target.** ~534,000 US consumer bankruptcy filings in 2025 against ~132m
households is roughly **0.4% of households per year**, and about **1 in 10
Americans file at some point in a lifetime**. Over a 10-year run that implies
~2–4% for the general population, and perhaps 5–10% for a low-wage archetype
carrying student debt without employer cover.

The model once produced **58.7%** for Wage Warrior. Three separate errors, none
of them a tuning problem:

**1. Every shortfall became revolving credit.** A $150k uninsured surgery turned
into $150k on a card at 23% with a minimum of 3% of an unbounded balance — a
$4,500/mo demand on a $2,500/mo income, unpayable by construction, so default was
certain rather than risky. Fixed by splitting the shortfall by kind:

- **medical** → 0%, minimum is a slice of *income*, does not compound
- **revolving credit** → capped at a real credit limit set by standing
- **arrears** → whatever exceeds the limit, at 0%, in collections

Hospital debt genuinely does not behave like card debt, which is why medical
bills are the leading contributor to bankruptcy *and* do not produce runaway
compounding.

**2. Wages never grew while rents inflated.** Cumulative rent increases add
~+$690/mo over a run; Wage Warrior's slack was $438. Cash-flow collapse was
arithmetic, not misfortune. Fixed with a cost-of-living adjustment —
`wageGrowthAnnual` — set from **nominal** wage growth, since real wages have been
roughly flat (2000–2019: +1.4% over nineteen years).

Both sides are now sourced and the gap between them is deliberate: rents run
~4.7%/yr against 3.5% wage growth, because US rents grew ~5.4%/yr over 2015–2025
and +30.4% over 2019–2023 against +20.2% for wages. **Rent outpacing pay is the
real story of the era and the model must not flatten it.**

**3. Startup Mage was underwater at character creation** — $1,900 income against
$1,911 of committed costs. "Below market" is not "below subsistence."

**4. The forced-insolvency trigger fired on debts bankruptcy cannot touch.** The
single largest cause of a forced filing was the *student loan* — a
non-dischargeable debt. Filing over it accomplishes nothing, which is why the
real road there is garnishment and income-driven repayment rather than court.
The trigger now requires a debt that is both **dischargeable** and **material**
(at least three months of income).

**5. Standing never recovered from a discharge.** Credit landed at exactly 300
and climbed +2/turn — 140 turns to escape Shunned, so a discharged character was
permanently paying 1.15× rent and 1.95× premiums, which made a second insolvency
close to inevitable. Recovery is now mean-reverting toward 700 at 4% of the gap
per turn: fast from the floor, unchanged near the top. Observed: **300 → 456 →
550 → 608 → 644** over four years, which is roughly how scores behave after a
Chapter 7 once the delinquencies are resolved by the discharge itself.

Current, 2,000 runs per class over 10 years:

| class | ever insolvent | ran out of relief |
|---|---:|---:|
| Wage Warrior | 14.1% | 5.3% |
| Salaryman Paladin | 1.2% | 0.3% |
| Freelance Rogue | 0.1% | 0.0% |
| Startup Mage | 12.9% | 5.5% |

Down from 58.7% for Wage Warrior. Plausible for the most exposed archetypes
against a 5–10% target, and repeat insolvency inside the cooldown is now
uncommon rather than routine.

To make it visible in play, the end-of-run summary should place the run **within its
regime's population** — a percentile, and the p10/median/p90 of that regime and
class — so the player sees whether they landed among the excess winners, the ruined,
or the reasonable middle. The distributions come from the sweep, precomputed per
regime and class. Without that readout the game can only say "you did fine," which
is exactly the omission §10 warns about.
