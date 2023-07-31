import logo from "../../assets/logo-transparent.png";
import React from "react";

type LogoProps = {
    className?: string
    width?: number
}
const Logo: React.FC<LogoProps> = ( props: LogoProps ) => {
    return <img src={logo} width={props.width} alt={"Logo"} className={ props.className }/>;
};

export default Logo;