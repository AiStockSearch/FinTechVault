import React from "react";
import "./secondLogin.css";

const SecondLogin: React.FC = () =>
{
    return (
        <section>
            <div className="second-login">
                <div className="second-login-inner">
                    <h1>Liquidity Stack,<br />
                        Reimagined.</h1>
                    <div className="content">
                        <p>
                            <strong>Supplying</strong> liquidity <span>is</span>
                            <strong>fundamental</strong> <span>to</span> DeFi,{" "}
                            <span>determining a</span> protocol's fate{" "}
                            <span>from the start.</span>
                        </p>
                        <p>
                            <span>Yet, the pursuit of</span> early-stage liquidity{" "}
                            <span>often involves</span> undisclosed contracts,
                        </p>
                        <p>
                            <span>fueling</span> <strong>information asymmetry</strong>
                            <span>and eroding</span> DeFi's core ethos.
                        </p>
                        <p>
                            <span>Mitosis</span> <strong>restores balance</strong>{" "}
                            <span>to</span>
                            fractured liquidity stack.
                        </p>
                    </div>
                </div>
                <div className="manifesto_button">
                    <a href="/manifesto" className="manifesto-link">
                        <span>Read Manifesto</span>
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            fill="none"
                            viewBox="0 0 16 16"
                            width="16"
                            height="16"
                        >
                            <path stroke="currentColor" d="m9 3 5 5m0 0-5 5m5-5H2"></path>
                        </svg>
                    </a>
                </div>
            </div>
        </section>
    );
};

export default SecondLogin;
