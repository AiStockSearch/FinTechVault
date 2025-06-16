import React from "react";
import "./threeLogin.css";

const translations = {
    title: {
        line1: "Be a Liquidity",
        line2: "Provider on Mitosis."
    },
    content: {
        line1: "Any Chain.",
        line2: "More Control.",
        line3: "More Rewards."
    },
    button: {
        text: "Provide Liquidity"
    },
    info: {
        assets: "5 Assets",
        chains: "9 Chains",
        tvl: "$48M TVL"
    }
};

const assets = [
    "/assets/third/assets/1.svg",
    "/assets/third/assets/2.svg",
    "/assets/third/assets/3.svg",
    "/assets/third/assets/4.svg",
    "/assets/third/assets/5.svg"
];

const chains = [
    "/assets/third/chains/1.svg",
    "/assets/third/chains/2.svg",
    "/assets/third/chains/3.svg",
    "/assets/third/chains/4.svg",
    "/assets/third/chains/5.svg",
    "/assets/third/chains/6.svg",
    "/assets/third/chains/7.svg",
    "/assets/third/chains/8.svg",
    "/assets/third/chains/9.svg"
];

const ThreeLogin = () =>
{
    return (
        <section className="three-login">
            <div className="three-login-inner">
                <h1>
                    <span className="block">{translations.title.line1}</span>
                    <span className="block">{translations.title.line2}</span>
                </h1>
                <div className="content">
                    <p>
                        <span className="block">{translations.content.line1}</span>
                        <span className="block">{translations.content.line2}</span>
                        <span className="block">{translations.content.line3}</span>
                    </p>
                </div>
                <div className="manifesto_button">
                    <a
                        href="https://app.mitosis.org"
                        target="_blank"
                        rel="noreferrer noopener"
                        className="manifesto-link"
                    >
                        <span>{translations.button.text}</span>
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 16" width="16" height="16">
                            <path stroke="currentColor" d="m9 3 5 5m0 0-5 5m5-5H2"></path>
                        </svg>
                    </a>
                </div>
            </div>
            <div className="three-login-assets">
                {assets.map( ( src, idx ) => (
                    <div className="asset-icon" key={src}>
                        <img src={src} alt={`asset-${ idx + 1 }`} width={28} height={28} />
                    </div>
                ) )}
                <div className="assets-info">{translations.info.assets}</div>
            </div>
            <div className="three-login-chains">
                {chains.map( ( src, idx ) => (
                    <div className="chain-icon" key={src}>
                        <img src={src} alt={`chain-${ idx + 1 }`} width={28} height={28} />
                    </div>
                ) )}
                <div className="chains-info">{translations.info.chains}</div>
                <div className="tvl-info">{translations.info.tvl}</div>
            </div>
        </section>
    );
};

export default ThreeLogin;