import React from "react";
import "./header.css";

const HeaderLogin: React.FC = () =>
{
    return (
        <nav>
            <div className="header_position relative flex w-full gap-10 justify-start md:justify-start md:gap-6 md:w-full sm:z-[9999]">
                <div
                    className="overflow-hidden sc-9232a9d6-1 jjcprg"
                    style={{ opacity: 1, filter: "blur(0px)", transform: "none" }}
                >
                    <a className="hover:opacity-70" href="/">
                        <div className="relative flex h-fit w-fit">
                            <div
                                className="relative flex h-fit w-fit rounded-full z-[1]"
                                style={{ width: "94px", height: "20px" }}
                            >
                                <img
                                    alt="logo image"
                                    loading="eager"
                                    width="94"
                                    height="10"
                                    decoding="async"
                                    data-nimg="1"
                                    style={{ color: "transparent" }}
                                    src="/src/assets/nav/logo.png"
                                />
                            </div>
                        </div>
                    </a>
                </div>
                <div className="header_block__row">
                    <div className="header_block__link">
                        <a href="/manifesto">
                            <span>
                                Manifesto
                            </span>
                        </a>
                    </div>
                    <div className="header_block__link">
                        <a href="/ecosystem">
                            <span>
                                Ecosystem
                            </span>
                        </a>
                    </div>
                    <div className="header_block__link">
                        <a href="/learn">
                            <span>
                                Learn
                            </span>
                        </a>
                    </div>
                    <div className="header_block__link">
                        <a href="/Community">
                            <span>
                                Community
                            </span>
                        </a>
                    </div>
                </div>
            </div>
            <div className="header_button">
                <a
                    href="https://app.mitosis.org"
                    target="_blank"
                    rel="noreferrer noopener"
                    className="sc-9232a9d6-4 ddENye"
                >
                    Launch App
                </a>
            </div>
        </nav>
    );
};

export default HeaderLogin;
