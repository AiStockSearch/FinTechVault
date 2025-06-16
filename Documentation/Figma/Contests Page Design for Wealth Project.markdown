# Contests Page Design for Wealth Project

## Purpose
The **Contests Page** is a dedicated hub for participating in Wealth Project’s index creation competitions, viewing leaderboards, and earning WTH rewards. It’s distinct from the Competitions Screen (integrated in dashboard), offering a standalone, community-focused experience inspired by **TradingView**’s trading competitions and **GitHub**’s hackathon pages, driving engagement for traders and creators.

**Goals**:
- Display ongoing and past contests with details (rewards, duration).
- Enable participation: Create indices, submit entries.
- Show leaderboards with rankings and scores.
- Integrate with side panel for contest updates and airdrop promotion.

### Functionality
- **Contest List**:
  - Ongoing: “Index Challenge” (1000 WTH, ends 2025-06-30).
  - Past: “Q1 2025 Contest” (500 WTH, 5th place).
  - Details: Reward, duration, rules, participants.
- **Participation Form**:
  - Create index: Name, tokens (30,000+), weights.
  - Submit entry: Confirm index, pay fee (0.1 WTH).
- **Leaderboard**:
  - Top 10 users: Rank, username, score (e.g., index return).
  - Filter: Contest, Timeframe.
- **Side Panel**:
  - Balance Updates: Contest entry (+0.1 WTH).
  - News: Contest updates (“New Challenge Live”).
  - Social Alerts: Contest rankings (@WealthProject).

### Interactions
- **Contest List**:
  - Contest Card → Modal (details, join).
  - Filter → Update list.
- **Participation Form**:
  - Submit → Confirmation modal.
- **Leaderboard**:
  - Filter → Update rankings.
  - User Card → Profile modal.
- **Transitions**:
  - Contest Card → Pop-up modal (0.2s).
  - Submit → Fade confirmation (0.2s).

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Contests Icon (Yellow), Filter (Contest Dropdown)]
[Left Sidebar: Same as Main]
[Right Side Panel: 250px Wide, Black #1E222D]
  [Balance Updates: Card (+0.1 WTH, Contest Entry)]
  [News: Card (“New Challenge Live”)]
  [Social Alerts: Card (@WealthProject: “Top Rank”)]
[Main Content: 750px Wide]
  [Header: "Contests" (20px, White), Active Contests: 5]
  [Contest List: Cards (“Index Challenge”, 1000 WTH, 2025-06-30), 300x150px]
  [Participation Form: Inputs (Name), Dropdown (Tokens), Sliders (Weights), Submit]
  [Leaderboard: Table (Rank, User, Score), 600x400px]
[Bottom Panel: Alerts (Entry Submitted), 750x80px]
[Mobile: Stacked Cards, Compact Table]
```

## Notes
- TradingView-style leaderboard, GitHub-style contest cards.
- Enhances community engagement and competition.