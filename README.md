# Valid Palindrome

> NeetCode 150 · Two Pointers · Blind 75 — converging pointers that skip junk characters in place; no filtered copy, no reversal. Go, stdlib only, tested.

## Problem

Ignoring everything except letters and digits, and ignoring case, does the string read the same forwards and backwards? ("A man, a plan, a canal: Panama" → yes.)

## Approach

The allocate-happy version builds a cleaned lowercase copy and compares it with its reverse — O(n) extra memory for a question that needs none. The two-pointer version walks inwards from both ends, each pointer skipping non-alphanumerics *as it goes*, comparing case-folded bytes when both rest on real characters. The inner skip loops keep `l < r` as their guard so the pointers can't cross mid-skip. `"0P"` is the famous trap this problem carries: digit-letter comparisons must fail even though '0' and 'P' differ by exactly 32, the same gap as a case fold — folding only letters, not digits, is the fix.

**Complexity:** O(n) time, O(1) space. (Byte-level; for full Unicode, switch to rune iteration and `unicode` classification.)

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
