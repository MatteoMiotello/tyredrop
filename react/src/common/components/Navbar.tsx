import React from "react";
import {useDispatch} from "react-redux";
import {Link} from "react-router-dom";
import {logout} from "../../modules/auth/store/auth-slice";
import Field from "../components-library/Input";
import Menu from "../components-library/Menu";
import CartButton from "./CartButton";
import MainLogo from "./Logo";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faBagShopping, faLocationDot, faRightFromBracket, faSearch, faUser} from "@fortawesome/free-solid-svg-icons";
import {useAuth} from "../../modules/auth/hooks/useAuth";

const Navbar: React.FC = () => {
    const auth = useAuth();
    const dispatch = useDispatch();

    return <div className="navbar bg-base-100">
        <MainLogo width={100}/>
        <div className="input-group">
            <Field.FormControl className="m-auto px-4 md:w-1/2 w-full">
                <Field.InputGroup size="sm">
                    <Field.Input type="text" name="search" placeholder="Ricerca facile" size="sm"/>
                    <span>
                        <FontAwesomeIcon icon={faSearch}/>
                    </span>
                </Field.InputGroup>
            </Field.FormControl>
        </div>
        <CartButton/>
        <div className="flex-none gap-2">
            <div className="dropdown dropdown-end z-50">
                <Menu.Dropdown label={ <div className="avatar placeholder h-full p-1">
                    <div className="bg-neutral-focus text-neutral-content rounded-full">
                        <span className="text-xl"> <FontAwesomeIcon icon={faUser}/> </span>
                    </div>
                </div> }>
                    <Menu.Item>
                        <Link to="/user">
                            <FontAwesomeIcon icon={faUser}/> Principale
                        </Link>
                    </Menu.Item>
                    <Menu.Item>
                        <Link to="/user/address">
                            <FontAwesomeIcon icon={faLocationDot}/>I miei indirizzi
                        </Link>
                    </Menu.Item>
                    <Menu.Item>
                        <Link to="/user/orders">
                            <FontAwesomeIcon icon={faBagShopping}/>I miei ordini
                        </Link>
                    </Menu.Item>
                    <Menu.Item>
                        <a onClick={() => dispatch( logout() )}>
                            <FontAwesomeIcon icon={faRightFromBracket}/> Esci
                        </a>
                    </Menu.Item>
                </Menu.Dropdown>
            </div>
        </div>
    </div>;
};

export default Navbar;