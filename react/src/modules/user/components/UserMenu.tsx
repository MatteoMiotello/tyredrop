import {faBagShopping, faLocationDot, faUser} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import {Link, useParams} from "react-router-dom";
import Menu from "../../../common/components-library/Menu";

const UserMenu: React.FC = () => {
    const params = useParams<{id: string}>();

    return <Menu>
        <Menu.Item>
            <Link to={`/user/${params.id}`}>
                <FontAwesomeIcon icon={faUser}/> Principale
            </Link>
        </Menu.Item>
        <Menu.Item>
            <Link to={`/user/${params.id}/address`}>
                <FontAwesomeIcon icon={faLocationDot}/>I miei indirizzi
            </Link>
        </Menu.Item>
        <Menu.Item>
            <Link to={`/user/${params.id}/orders`}>
                <FontAwesomeIcon icon={faBagShopping}/>I miei ordini
            </Link>
        </Menu.Item>
    </Menu>;
};

export default UserMenu;