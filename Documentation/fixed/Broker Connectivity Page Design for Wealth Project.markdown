# Broker Connectivity Page Design for Wealth Project

## Purpose
The **Broker Connectivity Page** enables users to link third-party brokers (CEX and traditional) for trading futures, commodities, forex, options, ETFs, stocks, and bonds, unifying access within Wealth Project. It’s inspired by **TradingView**’s broker integration panel and **GitHub**’s API settings, providing a secure, streamlined interface.

**Goals**:
- Connect CEX brokers (e.g., Binance, Bybit) and traditional brokers (e.g., Interactive Brokers, OANDA).
- Manage API keys for broker access (encrypted, Vault).
- Display connected brokers’ status and trading capabilities.
- Integrate with side panel for balance/news updates.

## Functionality
- **Broker List**:
  - Connected brokers: Binance (active), Interactive Brokers (pending).
  - Details: Broker name, status, markets (crypto, stocks), API key (partial).
  - Supported Brokers: Binance, Bybit, OKX, Bitfinex, Deribit, Interactive Brokers, OANDA, TD Ameritrade.
- **Connect Broker Form**:
  - Select broker (dropdown).
  - Input API key, secret (encrypted).
  - Verify connection (OAuth or API test).
- **Broker Status**:
  - Active, Pending, Disconnected.
  - Markets: Crypto, Futures, Stocks, etc.
- **Side Panel**:
  - Balance Updates: Broker connected (+0.1 WTH).
  - News: Broker updates (“Binance Adds Options”).
  - Social Alerts: Broker news (@WealthProject).

### Interactions
- **Broker List**:
  - Broker Card → Modal (details, disconnect).
  - Status → Toggle (connect/disconnect).
- **Connect Broker Form**:
  - Submit → Verification modal (success/error).
- **Transitions**:
  - Broker Card → Pop-up modal (0.2s).
  - Submit → Fade verification (0.2s).

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Broker Icon (Yellow), Add Broker (Yellow Button)]
[Left Sidebar: Same as Main]
[Right Side Panel: 250px Wide, Black #1E222D]
  [Balance Updates: Card (+0.1 WTH, Broker Connected)]
  [News: Card (“Binance Adds Options”)]
  [Social Alerts: Card (@WealthProject: “New Broker”)]
[Main Content: 750px Wide]
  [Header: "Broker Connectivity" (20px, White), Connected Brokers: 3]
  [Broker List: Cards (Binance: Active, Crypto/Futures; Interactive Brokers: Pending), 300x150px]
  [Connect Form: Dropdown (Broker: Binance), Inputs (API Key, Secret), Submit]
  [Status: Table (Broker, Status, Markets), 600x200px, Gray #2A2E39]
[Bottom Panel: Alerts (Broker Linked), 750x80px]
[Mobile: Stacked Cards, Compact Form]
```

## Notes
- TradingView-style broker cards, GitHub-style API management.
- Unifies CEX/traditional broker access, enhancing multi-market trading.