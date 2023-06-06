import {faBagShopping, faLocationDot, faUser} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import {Link} from "react-router-dom";
import Menu from "../../../common/components-library/Menu";

const UserMenu: React.FC = () => {
    return <Menu>
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
    </Menu>;
};

export default UserMenu;