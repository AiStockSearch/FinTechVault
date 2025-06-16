# Supported Markets for Wealth Project

## Overview
Wealth Project will support a diverse range of financial markets to position itself as a universal Web3 financial platform, competing with TradingView and CryptoIndex. The following markets will be integrated into the platform’s **Screener**, **DeFi Trading**, **Portfolio**, and **Indices** screens, with real-time data via **Chainlink oracles** and connectivity to **third-party brokers** (CEX and traditional) for unified access.

### Markets
1. **Futures**:
   - **Description**: Contracts to buy/sell assets (crypto, commodities, stocks) at a future date/price.
   - **Examples**: Bitcoin Futures (CME), Ethereum Futures, S&P 500 Futures.
   - **Integration**: Tradable via CEX brokers (e.g., Bybit, Binance), displayed in Screener/Portfolio.
   - **Data**: Price, volume, open interest, expiry dates.
2. **Commodities**:
   - **Description**: Physical assets like gold, oil, agricultural products.
   - **Examples**: Gold (XAU), Crude Oil (WTI), Wheat.
   - **Integration**: Tokenized commodities on DeFi (e.g., Pax Gold), CEX trading (e.g., OKX).
   - **Data**: Spot price, futures price, volatility.
3. **Forex**:
   - **Description**: Currency pairs for foreign exchange trading.
   - **Examples**: EUR/USD, USD/JPY, GBP/USD.
   - **Integration**: Forex brokers (e.g., OANDA, IG), DeFi forex pools (e.g., Curve).
   - **Data**: Bid/ask, spread, daily change.
4. **Options**:
   - **Description**: Contracts for buying/selling assets at a set price by a date.
   - **Examples**: Bitcoin Call/Put Options, Tesla Stock Options.
   - **Integration**: CEX brokers (e.g., Deribit), DeFi options (e.g., Opyn).
   - **Data**: Strike price, expiry, premium, implied volatility.
5. **Crypto Market**:
   - **Description**: Major cryptocurrencies (e.g., BTC, ETH).
   - **Examples**: Bitcoin, Ethereum, BNB.
   - **Integration**: DeFi (Uniswap, Aave), CEX (Binance, Bybit), wallet support.
   - **Data**: Price, market cap, volume, Chainlink oracles.
6. **Altcoins**:
   - **Description**: Non-major cryptocurrencies (30,000+ tokens).
   - **Examples**: SOL, MATIC, LINK, DOGE.
   - **Integration**: DEX (0x Protocol), CEX, wallet support.
   - **Data**: Price, volume, volatility, RSI.
7. **ETFs**:
   - **Description**: Exchange-traded funds for traditional assets.
   - **Examples**: SPY (S&P 500 ETF), QQQ (Nasdaq ETF).
   - **Integration**: Traditional brokers (e.g., Interactive Brokers), tokenized ETFs on DeFi.
   - **Data**: NAV, expense ratio, performance.
8. **Crypto ETFs**:
   - **Description**: ETFs tracking crypto assets or indices.
   - **Examples**: BITO (Bitcoin Strategy ETF), Custom Wealth Project Crypto ETF.
   - **Integration**: CEX (e.g., Bitfinex), DeFi (tokenized), Indices Screen for custom creation.
   - **Data**: NAV, underlying assets, performance.
9. **Stocks (Various Countries)**:
   - **Description**: Equities from global markets.
   - **Examples**: USA (AAPL, TSLA), Japan (SONY), Germany (SAP).
   - **Integration**: Traditional brokers (e.g., TD Ameritrade), tokenized stocks on DeFi.
   - **Data**: Price, P/E ratio, dividend yield.
10. **Bonds**:
    - **Description**: Fixed-income securities (government, corporate).
    - **Examples**: US Treasury Bonds, Corporate Bonds (Apple).
    - **Integration**: Tokenized bonds on DeFi (e.g., Aave), traditional brokers.
    - **Data**: Yield, maturity, coupon rate.

### Integration Details
- **Screener**: Filter all markets by price, volume, volatility, RSI, market cap.
- **DeFi Trading**: Trade crypto, altcoins, tokenized commodities/ETFs/stocks/bonds on DEX (Uniswap, Curve), futures/options on CEX (Bybit, Deribit).
- **Portfolio**: Track all markets with risk metrics (VaR, Sharpe Ratio, MACD).
- **Indices**: Create custom indices combining any market (e.g., 30% BTC, 20% AAPL, 50% Gold).
- **Broker Connectivity**: Unified API for CEX (Binance, Bybit) and traditional brokers (Interactive Brokers, OANDA) via a new **Broker Connectivity** page.
- **Data Sources**:
  - Crypto/Altcoins: Chainlink oracles, DEX APIs (0x Protocol).
  - Futures/Options: CEX APIs (Bybit, Deribit).
  - Commodities/Forex/ETFs/Stocks/Bonds: Broker APIs, tokenized data on DeFi.
- **XML Updates**: Add `<MarketData>` to relevant services (Screener, Portfolio, Indices):
  ```xml
  <MarketData type="futures">
      <Asset symbol="BTC_FUT" priceUSD="30000" volume="1000000" expiry="2025-12-31"/>
  </MarketData>
  ```

## Notes
- Supports 10 markets, ensuring Wealth Project is a universal platform.
- Integrates with DeFi (DEX, tokenized assets) and CEX/traditional brokers.
- Enhances Indices Screen for multi-market index creation.