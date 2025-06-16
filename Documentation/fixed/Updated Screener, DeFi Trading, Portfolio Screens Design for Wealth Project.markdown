# Updated Screener, DeFi Trading, Portfolio Screens Design for Wealth Project

## Screener and Watchlists Screen

### Purpose
The **Screener and Watchlists Screen** is updated to support all markets (futures, commodities, forex, options, crypto, altcoins, ETFs, stocks, bonds), enabling users to discover and track assets with advanced filters and watchlists.

### Functionality
- **Screener**:
  - Filter by: Price, Volume, Volatility, RSI, Market Cap, Market Type (Crypto, Stocks).
  - Sort by: Price change (1h, 24h), volume, P/E ratio (stocks), yield (bonds).
  - Results: SOL ($200), AAPL ($150), Gold ($2000), US Treasury (2% yield).
- **Watchlists**:
  - Multiple lists: “DeFi”, “Stocks”, “Bonds”.
  - Triggers: Price, RSI, MACD, futures expiry.
- **Side Panel**:
  - Balance Updates: Token added (+100 MATIC).
  - News: Market updates (“Gold Surges”).
  - Social Alerts: Trends (@WealthProject).

### Interactions
- Filter → Update results.
- Token Card → Modal (details, trade).
- Watchlist Selector → Switch list.

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Watchlists Icon (Yellow), Search, Filter (Market Dropdown)]
[Left Sidebar: Same as Main]
[Right Side Panel: Balance Updates, News, Social Alerts]
[Main Content: 750px Wide]
  [Header: "Screener & Watchlists" (20px, White), Active Watchlist: DeFi]
  [Screener: Filters (Price, Volume, Market: Crypto), Sort, 600x100px]
  [Results: Table (SOL: $200, AAPL: $150, Gold: $2000), 600x400px]
  [Watchlists: Tabs (DeFi, Stocks), List (SOL: $200), 600x300px]
[Bottom Panel: Alerts (Price Trigger), 750x80px]
[Mobile: Dropdown for Tabs, Compact Table]
```

## DeFi Trading Screen

### Purpose
The **DeFi Trading Screen** is updated to include trading for tokenized assets (commodities, ETFs, stocks, bonds) and options, alongside crypto and altcoins.

### Functionality
- **Trade Form**:
  - Markets: Crypto (SOL/USDT), Tokenized Stocks (AAPL), Bonds (US Treasury).
  - Order types: Limit, Market, Options (Call/Put).
  - Slippage, gas settings.
- **Token Search**: 30,000+ tokens, tokenized assets.
- **Active Orders**: SOL limit ($200), AAPL call ($150).
- **Side Panel**:
  - Balance Updates: Trade executed (+0.1 WTH).
  - News: Market updates.
  - Social Alerts: Trade trends (@WealthProject).

### Interactions
- Token Search → Autocomplete.
- Trade Form → Submit → Modal.
- Order Card → Cancel/edit.

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: DeFi Icon (Yellow), Market Dropdown (Crypto, Stocks)]
[Left Sidebar: Same as Main]
[Right Side Panel: Balance Updates, News, Social Alerts]
[Main Content: 750px Wide]
  [Header: "DeFi Trading" (20px, White), Market: Crypto]
  [Trade Form: Dropdowns (SOL, AAPL), Inputs, Submit, 400x300px]
  [Token Search: Input (30,000+ Assets), Autocomplete]
  [Active Orders: Table (SOL $200, AAPL Call), 600x200px]
[Bottom Panel: Alerts (Order Placed), 750x80px]
[Mobile: Stacked Panels, Compact Table]
```

## Portfolio Screen

### Purpose
The **Portfolio Screen** is updated to track all markets with risk metrics and auto-rebalancing across diverse assets.

### Functionality
- **Overview**:
  - Portfolios: Short-term ($2800, SOL, AAPL), Mid-term ($3000, Gold), Long-term ($4000, Bonds).
  - Chart: Performance, VaR, RSI, MACD.
- **Asset Breakdown**:
  - SOL (37%), AAPL (20%), Gold (30%), pie chart.
- **Auto-Rebalance**: Toggle, frequency, market weights.
- **Side Panel**:
  - Balance Updates: Rebalanced (+0.1 WTH).
  - News: Portfolio updates.
  - Social Alerts: Portfolio shares (@WealthProject).

### Interactions
- Timeframe Selector → Update chart.
- Asset Card → Modal (edit).
- Toggle → Save rebalance.

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Portfolio Icon (Yellow), Timeframe Dropdown]
[Left Sidebar: Same as Main]
[Right Side Panel: Balance Updates, News, Social Alerts]
[Main Content: 750px Wide]
  [Header: "Portfolio" (20px, White), Total Value: $9,800]
  [Overview: Cards (Short-term: $2800, SOL, AAPL), 300x100px]
  [Chart: Line Graph (600x300px), Toggles (VaR, RSI)]
  [Asset Breakdown: Pie Chart (SOL 37%, AAPL 20%), 300x300px]
  [Auto-Rebalance: Toggle, Dropdown, Save]
[Bottom Panel: Alerts (Rebalanced), 750x80px]
[Mobile: Stacked Cards, Smaller Chart]
```

## Notes
- All screens updated to support new markets (futures, stocks, bonds, etc.).
- TradingView-style charts, GitHub-style structured data views.
- Enhances multi-market trading and portfolio management.