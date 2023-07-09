import React from "react";
import {useDispatch} from "react-redux";
import {Link} from "react-router-dom";
import {logout} from "../../modules/auth/store/auth-slice";
import Menu from "../components-library/Menu";
import CartButton from "./CartButton";
import MainLogo from "./Logo";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faBagShopping, faLocationDot, faRightFromBracket, faUser} from "@fortawesome/free-solid-svg-icons";
import {useAuth} from "../../modules/auth/hooks/useAuth";

const Navbar: React.FC = () => {
    const auth = useAuth();
    const dispatch = useDispatch();

    return <div className="navbar rounded-box bg-base-100">
        <MainLogo width={100}/>

        <div className="ml-auto">
            <CartButton/>
            <div className="flex-none gap-2">
                <div className="dropdown dropdown-end z-50">
                    <Menu.Dropdown label={<div className="avatar placeholder h-full p-1">
                        <div className="bg-neutral-focus text-neutral-content rounded-full">
                            <span className="text-xl"> <FontAwesomeIcon icon={faUser}/> </span>
                        </div>
                    </div>}>
                        <Menu.Item>
                            <Link to={`/user/${auth.user?.user?.userID}`}>
                                <FontAwesomeIcon icon={faUser}/> Principale
                            </Link>
                        </Menu.Item>
                        <Menu.Item>
                            <Link to={`/user/${auth.user?.user?.userID}/address`}>
                                <FontAwesomeIcon icon={faLocationDot}/>I miei indirizzi
                            </Link>
                        </Menu.Item>
                        <Menu.Item>
                            <Link to={`/user/${auth.user?.user?.userID}/orders`}>
                                <FontAwesomeIcon icon={faBagShopping}/>I miei ordini
                            </Link>
                        </Menu.Item>
                        <Menu.Item>
                            <a onClick={() => dispatch(logout())}>
                                <FontAwesomeIcon icon={faRightFromBracket}/> Esci
                            </a>
                        </Menu.Item>
                    </Menu.Dropdown>
                </div>
            </div>
        </div>
    </div>;
};

export default Navbar;