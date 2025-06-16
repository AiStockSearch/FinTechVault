import React from "react";
import "./firstLogin.css";

const FirstLogin: React.FC = () =>
{
    return (
        <section
            aria-label="Hero section"
            className="first-login-section">
            <div className="first-login-content">
                <h1 className="firstlogin_title">
                    The Network for
                    <br />
                    Programmable Liquidity
                </h1>
                <p className="firstlogin_description">
                    Leading the next generation of DeFi through the tokenization
                    <br />
                    of liquidity positions while ensuring seamless integration into
                    <br />
                    the Mitosis Ecosystem.
                </p>
            </div>

            <div className="first-login-cta">
                <a
                    href="https://docs.mitosis.org/docs/learn/litepaper"
                    target="_blank"
                    rel="noopener noreferrer"
                    className="first-login-link"
                >
                    More about Programmable Liquidity &rarr;
                </a>
            </div>
        </section>
    );
};

export default FirstLogin;
