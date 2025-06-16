# Updated Indices Screen Design for Wealth Project

## Purpose
The **Indices Screen** is the hub for creating, trading, and managing custom indices, now expanded to include all supported markets (futures, commodities, forex, options, crypto, altcoins, ETFs, crypto ETFs, stocks, bonds). It allows users to build personalized ETF-like portfolios, inspired by **TradingView**’s charting and **GitHub**’s structured project views, with a focus on flexibility and analytics.

**Goals**:
- Enable index creation with assets from all markets (e.g., 30% BTC, 20% AAPL, 50% Gold).
- Support custom ETF portfolio creation (tokenized crypto ETFs).
- Display index performance with TradingView-style charts and metrics.
- Integrate fees (0.1% entry/exit, 0.05% holder share) and WTH rewards.
- Connect to third-party brokers for trading non-crypto assets.
- Integrate with side panel for balance/news updates.

## Functionality
- **Index List**:
  - Custom indices: “Wealth Index” ($1000, 2% return).
  - ETF portfolios: “Crypto ETF” (BTC, ETH, $1M).
  - Details: Components (e.g., 30% BTC, 20% AAPL), fees, performance.
- **Create Index Form**:
  - Market selector: Futures, Commodities, Forex, Options, Crypto, Altcoins, ETFs, Stocks, Bonds.
  - Asset selector: 30,000+ tokens (DEX), stocks (AAPL), bonds (US Treasury).
  - Weights: Sliders (e.g., BTC 30%, Gold 50%).
  - Fees: Entry (0.1%), Exit (0.1%), Holder Share (0.05%).
  - ETF Option: Tokenize as crypto ETF (Polygon).
- **Performance Chart**:
  - Line graph (index return, 1D/1W/1M), metrics (Sharpe Ratio, VaR).
  - Toggle indicators: RSI, MACD, Bollinger Bands.
  - Oracle data: Chainlink prices (BTC $30,000, AAPL $150).
- **Broker Connectivity**:
  - Link to brokers (Bybit for futures, Interactive Brokers for stocks).
  - API key status (active, encrypted).
- **Quick Actions**:
  - Trade Index, Add to Portfolio, Share (X, Telegram).
- **Side Panel**:
  - Balance Updates: Index created (+0.1 WTH).
  - News: Market updates (“Gold Hits $2000”).
  - Social Alerts: Index shares (@WealthProject).

## Interactions
- **Index List**:
  - Index Card → Modal (details, trade).
  - ETF Card → Modal (details, API access).
- **Create Index Form**:
  - Select Market → Update asset list.
  - Submit → Confirmation modal (fees, ETF option).
- **Performance Chart**:
  - Toggle Indicators → Update chart.
  - Zoom → Adjust timeframe.
- **Broker Connectivity**:
  - Link Broker → Broker Connectivity page.
  - API Key → Copy, revoke.
- **Quick Actions**:
  - Trade → DeFi Trading page.
  - Share → Copy link, open social app.
- **Transitions**:
  - Index Card → Pop-up modal (0.2s).
  - Submit Form → Fade confirmation (0.2s).
  - Chart Toggle → Smooth update (0.3s).

## Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Indices Icon (Yellow), Create Index (Yellow Button), Broker Status (Blue Link)]
[Left Sidebar: Home, Wallet, DeFi, Portfolio, Indices (Yellow), Publications, Analytics, Tokenomics, DAO, Competitions, ETF, Notes, Watchlists, Settings]
[Right Side Panel: 250px Wide, Black #1E222D]
  [Balance Updates: Card (+0.1 WTH, Index Created)]
  [News: Card (“Gold Hits $2000”)]
  [Social Alerts: Card (@WealthProject: “New Index”)]
[Main Content: 750px Wide]
  [Header: "Indices" (20px, White), Total Value: $1,000,000 (24px, Yellow)]
  [Index List: Cards (“Wealth Index”: $1000, 2%; “Crypto ETF”: $1M), 300x150px, Gray #2A2E39]
  [Details: Table (BTC: 30%, AAPL: 20%, Gold: 50%), Fees (Entry: 0.1%), 600x200px]
  [Create Form: Dropdown (Market: Crypto, Stocks), Token Selector (BTC, AAPL), Sliders (Weights), Checkbox (ETF), Submit]
  [Chart: Line Graph (Return, 600x300px, Yellow/Blue), Toggles (RSI, MACD), Oracle Data (BTC: $30,000)]
  [Broker Connectivity: Card (Bybit: Active, Interactive Brokers: Link), API Key, 300x150px]
  [Quick Actions: Buttons (Trade, Add to Portfolio, Share), 150x80px]
[Bottom Panel: Alerts (Index Traded), 750x80px]
[Mobile: Stacked Cards, Dropdown for Market, Smaller Chart (300x150px)]
```

## Notes
- TradingView-style chart, GitHub-style structured form.
- Supports all markets for index/ETF creation, with broker integration.
- Enhances flexibility for traders and institutional users.